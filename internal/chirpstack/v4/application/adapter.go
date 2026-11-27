package application

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/application"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
}

var _ application.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) device.API {
	client := NewWithClient(apiv4.NewApplicationServiceClient(conn))
	return &adapter{
		core: newCore(client),
	}
}
