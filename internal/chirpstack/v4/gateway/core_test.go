package gateway

import (
	"testing"

	gatewaymocks "github.com/sebundefined/chirpctl/internal/chirpstack/v4/gateway/mocks"
	"github.com/stretchr/testify/suite"
)

type GatewayCoreSuite struct {
	suite.Suite

	client *gatewaymocks.ClientMock
	core   *core
}

func TestGatewayCoreSuite(t *testing.T) {
	suite.Run(t, new(GatewayCoreSuite))
}

func (s *GatewayCoreSuite) SetupTest() {
	s.client = gatewaymocks.NewClientMock(s.T())
	s.core = newCore(s.client)
}
