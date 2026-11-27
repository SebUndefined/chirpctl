package application

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateParams) (CreateResult, error)
	Get(ctx context.Context, in GetParams) (Application, error)
	Update(ctx context.Context, in UpdateApplicationParams) error
	Delete(ctx context.Context, in DeleteApplicationParams) error
	List(ctx context.Context, in ListParams) (shared.ListResult[Application], error)
}

type Application struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string
	Description     string
	TenantID        string
	Tags            map[string]string
	MeasurementKeys []string
}

type CreateParams struct {
	Name        string
	Description string
	TenantID    string
	Tags        map[string]string
}

type CreateResult struct {
	ID string
}

type GetParams struct {
	ID string
}

type UpdateApplicationParams struct {
	ID    string
	Patch Patch
}

type Patch struct{}

type DeleteApplicationParams struct {
	ID string
}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	shared.ListTagsFilterRequest
	TenantID string
}

type ListItem struct {
	ID          string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Name        string
	Description string
}
