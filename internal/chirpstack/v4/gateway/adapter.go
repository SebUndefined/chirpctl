package gateway

import (
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
	return &adapter{
		core: newCore(conn),
	}
}
