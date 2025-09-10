package device

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Service interface {
	Get(ctx context.Context, p GetParams) (*Device, error)
	Create(ctx context.Context, p CreateParams) error
	List(ctx context.Context, p ListParams) (shared.ListResult[ListItem], error)
	Update(ctx context.Context, p UpdateParams) error
}
