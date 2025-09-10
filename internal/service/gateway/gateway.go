package gateway

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, p GetParams) (*Gateway, error)
	Create(ctx context.Context, p CreateParams) error
	List(ctx context.Context, p ListParams) (ListResult, error)
	Delete(ctx context.Context, p DeleteParams) error
}
