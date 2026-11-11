package clientfactory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

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
	s.Equal("grpc", cfg.Transport)
}

func (s *FactorySuite) TestExplicitTransport() {

}

func (s *FactorySuite) TestUnsupportedTransport() {

}

func (s *FactorySuite) TestGatewayCached() {

}

func (s *FactorySuite) TestDeviceCached() {

}

func (s *FactorySuite) TestReconfigure() {

}

func (s *FactorySuite) TestMissingCACert() {

}

func (s *FactorySuite) TestInvalidCACert() {

}

func (s *FactorySuite) TestClose() {

}
