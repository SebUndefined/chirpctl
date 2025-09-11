package device

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/device"
)

type metrics struct {
}

var _ device.Metrics = (*metrics)(nil)

func (m metrics) GetMetrics(ctx context.Context, p device.GetMetricsParam) (*device.Metrics, error) {
	//TODO implement me
	panic("implement me")
}
