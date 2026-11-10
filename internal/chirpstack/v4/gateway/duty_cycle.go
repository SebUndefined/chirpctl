package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/gateway"
)

type dutyCycle struct {
	client Client
}

var _ gateway.DutyCycle = (*dutyCycle)(nil)

func newDutyCycle(client Client) *dutyCycle {
	return &dutyCycle{client: client}
}

func (d *dutyCycle) GetDutyCycleMetrics(ctx context.Context, in gateway.GetGatewayDutyCycleMetricsParams) (gateway.GetGatewayDutyCycleMetricsResult, error) {
	//TODO implement me
	panic("implement me")
}
