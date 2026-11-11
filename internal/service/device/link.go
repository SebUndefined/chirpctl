package device

import "context"

type GetLinkMetricsParams struct {
}

type LinkMetrics struct {
}

type Link interface {
	GetLinkMetrics(ctx context.Context, p GetLinkMetricsParams) (*LinkMetrics, error)
}
