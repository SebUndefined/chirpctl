package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type metrics struct {
	client Client
}

func newMetrics(client Client) *metrics {
	return &metrics{client: client}
}

var _ shared.Metrics[gateway.GetGatewayMetricsParams, gateway.GetGatewayMetricsResult] = (*metrics)(nil)

func (m *metrics) GetMetrics(ctx context.Context, in gateway.GetGatewayMetricsParams) (gateway.GetGatewayMetricsResult, error) {
	//TODO implement me
	panic("implement me")
}
