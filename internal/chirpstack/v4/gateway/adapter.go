package gateway

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
	*relay
	*cert
	*dutyCycle
	*metrics
}

var _ gateway.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) gateway.API {
	client := NewWithClient(apiv4.NewGatewayServiceClient(conn))
	return &adapter{
		core:      newCore(client),
		relay:     newRelay(client),
		cert:      newCert(client),
		dutyCycle: newDutyCycle(client),
		metrics:   newMetrics(client),
	}
}
