package clientfactory

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/connectivity"
)

func testConfig() *ResolveConfig {
	return &ResolveConfig{
		Addr:       "127.0.0.1:65535",
		ApiVersion: "v4",
		Transport:  "grpc",
		Insecure:   true,
		Timeout:    10 * time.Second,
	}
}

type FactorySuite struct {
	suite.Suite
}

func TestFactorySuite(t *testing.T) {
	suite.Run(t, new(FactorySuite))
}

func (s *FactorySuite) TestDefaultConfig() {
	f, err := NewFactory(&ResolveConfig{ApiVersion: "v4"})
	s.Require().NoError(err)
	cfg, err := f.GetCurrentConfig()

	s.Require().NoError(err)
	s.Equal("grpc", cfg.Transport)
	s.Equal(15*time.Second, cfg.Timeout)
}

func (s *FactorySuite) TestNewFactoryNormalizedVersion() {
	f, err := NewFactory(&ResolveConfig{
		ApiVersion: "4",
	})
	s.Require().NoError(err)

	cfg, err := f.GetCurrentConfig()

	s.Require().NoError(err)
	s.Require().Equal("grpc", cfg.Transport)
}

func (s *FactorySuite) TestExplicitTransport() {
	cfg := testConfig()
	f, err := NewFactory(cfg)
	s.Require().NoError(err)
	current, err := f.GetCurrentConfig()
	s.Require().NoError(err)
	s.Require().Equal("grpc", current.Transport)
}

func (s *FactorySuite) TestUnsupportedTransport() {
	cfg := testConfig()
	cfg.Transport = "invalid"
	_, err := NewFactory(cfg)
	s.Error(err)
}

func (s *FactorySuite) TestGatewayCached() {
	f, err := NewFactory(testConfig())
	s.Require().NoError(err)
	s.T().Cleanup(func() { _ = f.Close() })

	first, err := f.Gateway()
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := f.Gateway()
	s.Require().NoError(err)

	s.Same(first, second)
}

func (s *FactorySuite) TestDeviceCached() {
	f, err := NewFactory(testConfig())
	s.Require().NoError(err)
	s.T().Cleanup(func() { _ = f.Close() })

	first, err := f.Device()
	s.Require().NoError(err)
	s.Require().NotNil(first)

	second, err := f.Device()
	s.Require().NoError(err)

	s.Same(first, second)
}

func (s *FactorySuite) TestReconfigure() {
	f, err := NewFactory(testConfig())
	s.Require().NoError(err)
	newAddr := "127.0.0.1:65539"
	err = f.Reconfigure(func(cfg *ResolveConfig) {
		cfg.Addr = newAddr
	})
	s.Require().NoError(err)
	cfg, err := f.GetCurrentConfig()
	s.Require().NoError(err)
	s.Require().Equal(newAddr, cfg.Addr)
}

func (s *FactorySuite) TestInvalidCACert() {
	cfg := testConfig()
	cfg.Insecure = false

	path := filepath.Join(s.T().TempDir(), "invalid-ca.pem")

	err := os.WriteFile(path, []byte("invalid certificate"), 0o600)
	s.Require().NoError(err)

	cfg.CACert = path

	f, err := NewFactory(cfg)
	s.Require().NoError(err)

	_, err = f.Gateway()
	s.Require().Error(err)
	s.ErrorContains(err, "no valid CA certificates found")
}

func (s *FactorySuite) TestMissingCACert() {
	cfg := testConfig()
	cfg.Insecure = false
	cfg.CACert = filepath.Join(s.T().TempDir(), "missing-ca.pem")

	f, err := NewFactory(cfg)
	s.Require().NoError(err)

	_, err = f.Gateway()
	s.Require().Error(err)
	s.ErrorIs(err, os.ErrNotExist)
}

func (s *FactorySuite) TestClose() {
	f, err := NewFactory(testConfig())
	s.Require().NoError(err)

	_, err = f.Gateway()
	s.Require().NoError(err)

	concrete := f.(*factory)
	s.Require().NotNil(concrete.conn)

	err = f.Close()
	s.Require().NoError(err)

	s.Equal(
		connectivity.Shutdown,
		concrete.conn.GetState(),
	)
}
