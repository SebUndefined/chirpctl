package shared

import (
	"context"
)

type Metrics[T any, R any] interface {
	GetMetrics(ctx context.Context, in T) (R, error)
}
