package device

import "context"

type Service interface {
	Get(ctx context.Context, p GetParams) (*Device, error)
	Create(ctx context.Context, p CreateParams) error
	List(ctx context.Context, p ListParams) (ListResult, error)
}
