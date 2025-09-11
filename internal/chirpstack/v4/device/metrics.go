package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
)

type metrics struct {
	c apiv4.DeviceServiceClient
}

func newMetrics(c apiv4.DeviceServiceClient) *metrics {
	return &metrics{c: c}
}

var _ device.Metrics = (*metrics)(nil)

func (m metrics) GetMetrics(ctx context.Context, p device.GetMetricsParam) (*device.Metrics, error) {
	//TODO implement me
	panic("implement me")
}
