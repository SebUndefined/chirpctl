package device

import (
	"github.com/sebundefined/chirpctl/internal/service/device"
	"google.golang.org/grpc"
)

type adapter struct {
	*core
	*keys
	*activation
	*queue
	*metrics
	*link
}

var _ device.API = (*adapter)(nil)

func NewAdapter(conn *grpc.ClientConn) device.API {
	return &adapter{
		core: newCore(conn),
	}
}
