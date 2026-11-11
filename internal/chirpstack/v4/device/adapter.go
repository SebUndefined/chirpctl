package device

import (
	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
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
	client := NewWithClient(apiv4.NewDeviceServiceClient(conn))
	return &adapter{
		core:       newCore(client),
		keys:       newKeys(client),
		activation: newActivation(client),
		queue:      newQueue(client),
		metrics:    newMetrics(client),
		link:       newLink(client),
	}
}
