package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/gateway"
)

type dutyCycle struct{}

var _ gateway.DutyCycle = (*dutyCycle)(nil)

func (d *dutyCycle) GetDutyCycleMetrics(ctx context.Context, in gateway.GetGatewayDutyCycleMetricsParams) (gateway.GetGatewayDutyCycleMetricsResult, error) {
	//TODO implement me
	panic("implement me")
}
