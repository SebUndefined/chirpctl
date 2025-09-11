package device

import "context"

type Metrics interface {
	GetMetrics(ctx context.Context, p GetMetricsParam) (*Metrics, error)
}

type GetMetricsParam struct{}
