package gateway

import (
	"context"
)

type DutyCycle interface {
	GetDutyCycleMetrics(ctx context.Context, in GetGatewayDutyCycleMetricsParams) (GetGatewayDutyCycleMetricsResult, error)
}

type GetGatewayDutyCycleMetricsResult struct {
}

type GetGatewayDutyCycleMetricsParams struct {
}
