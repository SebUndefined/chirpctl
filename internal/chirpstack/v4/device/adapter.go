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
	svc := apiv4.NewDeviceServiceClient(conn)
	s := apiv4.NewDeviceProfileServiceClient()
	return &adapter{
		core:       newCore(svc),
		keys:       newKeys(svc),
		activation: newActivation(svc),
		queue:      newQueue(svc),
		metrics:    newMetrics(svc),
		link:       newLink(svc),
	}
}
