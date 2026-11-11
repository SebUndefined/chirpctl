package application

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateApplicationParams) (CreateApplicationResult, error)
	Get(ctx context.Context, in GetApplicationParams) (Application, error)
	Update(ctx context.Context, in UpdateApplicationParams) error
	Delete(ctx context.Context, in DeleteApplicationParams) error
	List(ctx context.Context, in ListApplicationsParams) (shared.ListResult[Application], error)
}

type Application struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	TenantID    string
	Tags        map[string]string
}

type CreateApplicationParams struct {
	Application Application
}

type CreateApplicationResult struct {
	ID string
}

type GetApplicationParams struct {
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

type ListApplicationsParams struct{}
