package clientfactory

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	deviceV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/device"
	deviceprofileV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/deviceprofile"
	gatewayV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/gateway"
	tenantV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/tenant"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Factory interface {
	Tenant() (tenant.API, error)
	Device() (device.API, error)
	Gateway() (gateway.API, error)
	DeviceProfile() (deviceprofile.API, error)
	Reconfigure(apply func(*ResolveConfig)) error
	GetCurrentConfig() (ResolveConfig, error)
	Close() error
}

type ResolveConfig struct {
	Addr, CACert, ServerName, Token, ApiVersion, Transport string
	Insecure, TLSSkipVerify                                bool
	Timeout                                                time.Duration
	MTls                                                   config.MTls
}

type factory struct {
	cfg ResolveConfig

	connOnce, devOnce, gwOnce onceErr
	conn                      *grpc.ClientConn
	tenant                    tenant.API
	dev                       device.API
	gw                        gateway.API
	deviceProfile             deviceprofile.API
}

func NewFactory(cfg *ResolveConfig) (Factory, error) {
	if cfg == nil {
		return nil, errors.New("factory configuration is nil")
	}

	c := *cfg

	if c.Timeout == 0 {
		c.Timeout = 15 * time.Second
	}

	version := APIVersion(
		strings.TrimPrefix(
			strings.ToLower(strings.TrimSpace(c.ApiVersion)),
			"v",
		),
	)

	transport := Transport(strings.ToLower(strings.TrimSpace(c.Transport)))

	switch version {
	case APIVersionV4:
		if transport == "" {
			transport = TransportGRPC
		}

		if transport != TransportGRPC {
			return nil, fmt.Errorf(
				"unsupported transport %q for ChirpStack v4",
				transport,
			)
		}

	default:
		return nil, fmt.Errorf(
			"unsupported ChirpStack API version %q",
			c.ApiVersion,
		)
	}

	c.ApiVersion = string(version)
	c.Transport = string(transport)

	return &factory{cfg: c}, nil
}

func (f *factory) Reconfigure(apply func(*ResolveConfig)) error {
	apply(&f.cfg)
	_ = f.Close()
	f.connOnce.reset()
	f.devOnce.reset()
	f.gwOnce.reset()
	f.conn, f.dev, f.gw = nil, nil, nil
	return nil
}

func (f *factory) Tenant() (tenant.API, error) {
	if err := f.gwOnce.do(func() error {
		switch f.cfg.Transport {
		case "grpc":
			conn, err := f.grpcConn()
			if err != nil {
				return err
			}
			f.tenant = tenantV4.NewAdapter(conn)
		default:
			return errors.New("transport non supporté pour tenant")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return f.tenant, nil
}

func (f *factory) Device() (device.API, error) {
	if err := f.devOnce.do(func() error {
		switch f.cfg.Transport {
		case "grpc":
			conn, err := f.grpcConn()
			if err != nil {
				return err
			}
			f.dev = deviceV4.NewAdapter(conn)
		default:
			return errors.New("transport non supporté pour Device")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return f.dev, nil
}

func (f *factory) Gateway() (gateway.API, error) {
	if err := f.gwOnce.do(func() error {
		switch f.cfg.Transport {
		case "grpc":
			conn, err := f.grpcConn()
			if err != nil {
				return err
			}
			f.gw = gatewayV4.NewAdapter(conn)
		default:
			return errors.New("transport non supporté pour Gateway")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return f.gw, nil
}

func (f *factory) DeviceProfile() (deviceprofile.API, error) {
	if err := f.gwOnce.do(func() error {
		switch f.cfg.Transport {
		case "grpc":
			conn, err := f.grpcConn()
			if err != nil {
				return err
			}
			f.deviceProfile = deviceprofileV4.NewAdapter(conn)
		default:
			return errors.New("transport non supporté pour Device Profile")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return f.deviceProfile, nil
}

func (f *factory) grpcConn() (*grpc.ClientConn, error) {
	if err := f.connOnce.do(func() error {
		transportCreds, err := buildTransportCredentials(f.cfg)
		if err != nil {
			return fmt.Errorf("build gRPC transport credentials: %w", err)
		}

		conn, err := grpc.NewClient(
			f.cfg.Addr,
			grpc.WithTransportCredentials(transportCreds),
			grpc.WithPerRPCCredentials(bearerCreds{
				token:          strings.TrimSpace(f.cfg.Token),
				allowPlaintext: f.cfg.Insecure,
			}),
			grpc.WithUserAgent("chirpctl/1"),
		)
		if err != nil {
			return fmt.Errorf("create gRPC client: %w", err)
		}

		f.conn = conn
		return nil
	}); err != nil {
		return nil, err
	}

	return f.conn, nil
}
func buildTransportCredentials(cfg ResolveConfig) (credentials.TransportCredentials, error) {
	if cfg.Insecure {
		if cfg.TLSSkipVerify {
			return nil, errors.New(
				"tls-skip-verify cannot be used with plaintext gRPC",
			)
		}

		if cfg.MTls.Enabled {
			return nil, errors.New(
				"mTLS cannot be enabled with plaintext gRPC",
			)
		}

		if cfg.CACert != "" || cfg.ServerName != "" {
			return nil, errors.New(
				"TLS settings cannot be used with plaintext gRPC",
			)
		}

		return insecure.NewCredentials(), nil
	}
	serverName := cfg.ServerName

	if serverName == "" {
		host, _, err := net.SplitHostPort(cfg.Addr)
		if err != nil {
			return nil, fmt.Errorf("invalid gRPC address: %w", err)
		}
		serverName = host
	}
	// Simple tls config
	tlsCfg := &tls.Config{
		ServerName:         serverName,
		InsecureSkipVerify: cfg.TLSSkipVerify,
		MinVersion:         tls.VersionTLS12,
	}
	// Case custom
	if len(strings.TrimSpace(cfg.CACert)) > 0 {
		roots, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("load system certificate pool: %w", err)
		}

		caPEM, err := os.ReadFile(cfg.CACert)
		if err != nil {
			return nil, fmt.Errorf("read custom CA certificate: %w", err)
		}

		if ok := roots.AppendCertsFromPEM(caPEM); !ok {
			return nil, errors.New("no valid CA certificates found")
		}

		tlsCfg.RootCAs = roots
	}
	if cfg.MTls.Enabled {
		if cfg.MTls.ClientCert == "" || cfg.MTls.ClientKey == "" {
			return nil, errors.New(
				"mTLS requires a client certificate and private key",
			)
		}
		cert, err := tls.LoadX509KeyPair(cfg.MTls.ClientCert, cfg.MTls.ClientKey)
		if err != nil {
			return nil, err
		}

		tlsCfg.Certificates = []tls.Certificate{cert}
	}
	return credentials.NewTLS(tlsCfg), nil
}

func (f *factory) GetCurrentConfig() (ResolveConfig, error) {
	return f.cfg, nil
}

func (f *factory) Close() error {
	if f.conn != nil {
		_ = f.conn.Close()
	}
	return nil
}

// --- petit util pour once + error ---
type onceErr struct {
	once bool
	err  error
}

func (o *onceErr) do(fn func() error) error {
	if o.once {
		return o.err
	}
	o.err = fn()
	o.once = true
	return o.err
}
func (o *onceErr) reset() { o.once, o.err = false, nil }
