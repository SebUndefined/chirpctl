package gateway

import (
	"context"
	"errors"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type metrics struct {
	client Client
}

func newMetrics(client Client) *metrics {
	return &metrics{client: client}
}

var _ shared.Metrics[gateway.GetGatewayMetricsParams, gateway.GetGatewayMetricsResult] = (*metrics)(nil)

func (m *metrics) GetMetrics(ctx context.Context, in gateway.GetGatewayMetricsParams) (gateway.GetGatewayMetricsResult, error) {
	if in.Start == nil || in.End == nil {
		return gateway.GetGatewayMetricsResult{}, errors.New("start and end must be set")
	}
	req := api.GetGatewayMetricsRequest{
		GatewayId:   in.GatewayID,
		Start:       timestamppb.New(*in.Start),
		End:         timestamppb.New(*in.End),
		Aggregation: 0,
	}
	getMetrics, err := m.client.GetMetrics(ctx, &req)
	if err != nil {
		return gateway.GetGatewayMetricsResult{}, err
	}
	return toGatewayMetrics(getMetrics)
}
