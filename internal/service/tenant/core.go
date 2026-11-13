package application

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateParams) (CreateResult, error)
	Get(ctx context.Context, in GetParams) (*Tenant, error)
	Update(ctx context.Context, in UpdateParams) error
	Delete(ctx context.Context, in DeleteParams) error
	List(ctx context.Context, in ListParams) (shared.ListResult[ListItem], error)
}

type Tenant struct {
	ID                  string            `json:"id" yaml:"id"`
	CreatedAt           *time.Time        `json:"createdAt" yaml:"createdAt"`
	UpdatedAt           *time.Time        `json:"updatedAt" yaml:"updatedAt"`
	Name                string            `json:"name" yaml:"name"`
	Description         string            `json:"description" yaml:"description"`
	CanHaveGateways     bool              `json:"canHaveGateways" yaml:"canHaveGateways"`
	MaxGatewayCount     uint32            `json:"maxGatewayCount" yaml:"maxGatewayCount"`
	MaxDeviceCount      uint32            `json:"maxDeviceCount" yaml:"maxDeviceCount"`
	PrivateGatewaysUp   bool              `json:"privateGatewaysUp" yaml:"privateGatewaysUp"`
	PrivateGatewaysDown bool              `json:"privateGatewaysDown" yaml:"privateGatewaysDown"`
	Tags                map[string]string `json:"tags" yaml:"tags"`
	DevAddrPrefixes     []string          `json:"devAddrPrefixes" yaml:"devAddrPrefixes"`
}

type CreateParams struct {
	*Tenant
}

type CreateResult struct {
	ID string `json:"id"`
}

type GetParams struct {
	ID string
}

type UpdateParams struct {
}

type DeleteParams struct{}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	UserID string
}

type ListItem struct {
	ID                  string     `json:"ID" yaml:"ID"`
	CreatedAt           *time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt" yaml:"updatedAt"`
	Name                string     `json:"name" yaml:"name"`
	CanHaveGateways     bool       `json:"canHaveGateways" yaml:"canHaveGateways"`
	PrivateGatewaysUp   bool       `json:"privateGatewaysUp" yaml:"privateGatewaysUp"`
	PrivateGatewaysDown bool       `json:"privateGatewaysDown" yaml:"privateGatewaysDown"`
	MaxGatewayCount     uint32     `json:"maxGatewayCount" yaml:"maxGatewayCount"`
	MaxDeviceCount      uint32     `json:"maxDeviceCount" yaml:"maxDeviceCount"`
}
