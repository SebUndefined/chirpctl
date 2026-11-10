package clientfactory

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"os"
	"strings"
	"time"

	deviceV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/device"
	gatewayV4 "github.com/sebundefined/chirpctl/internal/chirpstack/v4/gateway"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Factory interface {
	Device() (device.API, error)
	Gateway() (gateway.API, error)
	Reconfigure(apply func(*ResolveConfig)) error
	GetCurrentConfig() (ResolveConfig, error)
	Close() error
}

type ResolveConfig struct {
	Addr, CACert, ServerName, Token, ApiVersion, Transport string
	Insecure                                               bool
	Timeout                                                time.Duration
}

type factory struct {
	cfg ResolveConfig

	connOnce, devOnce, gwOnce onceErr
	conn                      *grpc.ClientConn
	dev                       device.API
	gw                        gateway.API
}

func NewFactory(cfg *ResolveConfig) Factory {
	c := *cfg
	if c.Timeout == 0 {
		c.Timeout = 15 * time.Second
	}
	if c.Transport == "" {
		if c.ApiVersion == "v4" {
			c.Transport = "grpc"
		} else {
			c.Transport = "rest"
		}
	}
	return &factory{cfg: c}
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

func (f *factory) grpcConn() (*grpc.ClientConn, error) {
	if err := f.connOnce.do(func() error {
		var dialOpts []grpc.DialOption
		if f.cfg.Insecure {
			dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		} else {
			tlsCfg := &tls.Config{ServerName: f.cfg.ServerName}
			if tlsCfg.ServerName == "" {
				if host, _, err := net.SplitHostPort(f.cfg.Addr); err == nil {
					tlsCfg.ServerName = host
				}
			}
			if f.cfg.CACert != "" {
				ca, err := os.ReadFile(f.cfg.CACert)
				if err != nil {
					return err
				}
				cp := x509.NewCertPool()
				cp.AppendCertsFromPEM(ca)
				tlsCfg.RootCAs = cp
			}
			dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
		}
		/*unary := grpc.WithUnaryInterceptor(func(
			ctx context.Context,
			method string,
			req, reply interface{},
			cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker,
			callOpts ...grpc.CallOption,
		) error {
			md := metadata.New(nil)
			if f.cfg.Token != "" {
				md.Append("authorization", "Bearer "+f.cfg.Token)
			}
			md.Append("user-agent", "chipctl/1")
			return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, callOpts...)
		})*/
		// Auth sur toutes les RPC
		//dialOpts = append(dialOpts, unary)
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(bearerCreds{
			token:    strings.TrimSpace(f.cfg.Token),
			insecure: f.cfg.Insecure,
		}))

		conn, err := grpc.NewClient(f.cfg.Addr, dialOpts...)
		if err != nil {
			return err
		}
		f.conn = conn
		return nil
	}); err != nil {
		return nil, err
	}
	return f.conn, nil
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
