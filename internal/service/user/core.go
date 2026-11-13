package user

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateParams) (CreateResult, error)
	Get(ctx context.Context, in GetParams) (User, error)
	Update(ctx context.Context, in UpdateParams) error
	Delete(ctx context.Context, in DeleteParams) error
	List(ctx context.Context, in ListParams) (shared.ListResult[ListItem], error)
	UpdatePassword(ctx context.Context, in UpdatePasswordParams) error
}

type CreateParams struct {
	User     *User
	Password string
	Tenants  []*Tenant
}

type Tenant struct {
	ID             string
	IsAdmin        bool
	IsDeviceAdmin  bool
	IsGatewayAdmin bool
}

type CreateResult struct {
	ID string `json:"id"`
}

type GetParams struct {
	ID string
}

type User struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" yaml:"updatedAt"`
	IsAdmin   bool       `json:"isAdmin" yaml:"isAdmin"`
	IsActive  bool       `json:"isActive" yaml:"isActive"`
	Email     string     `json:"email" yaml:"email"`
	Note      string     `json:"note" yaml:"note"`
}

type UpdateParams struct {
	User *User
}

type DeleteParams struct {
	ID string `json:"id" yaml:"id"`
}

type ListParams struct {
	shared.ListPaginationRequest
}

type ListItem struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" yaml:"updatedAt"`
	IsAdmin   bool       `json:"isAdmin" yaml:"isAdmin"`
	IsActive  bool       `json:"isActive" yaml:"isActive"`
	Email     string     `json:"email" yaml:"email"`
}

type UpdatePasswordParams struct {
	ID       string
	Password string
}
