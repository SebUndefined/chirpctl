package tenant

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
	AddUser(ctx context.Context, in AddUserParams) error
	GetUser(ctx context.Context, in GetUserParams) (*User, error)
	UpdateUser(ctx context.Context, in UpdateUserParams) error
	DeleteUser(ctx context.Context, in DeleteUserParams) error
	ListUsers(ctx context.Context, in ListUsersParams) (shared.ListResult[ListUserItem], error)
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

type AddUserParams struct{}

type GetUserParams struct {
	ID       string
	TenantID string
}

type User struct {
	ID             string               `json:"id" yaml:"id"`
	CreatedAt      *time.Time           `json:"createdAt" yaml:"createdAt"`
	UpdatedAt      *time.Time           `json:"updatedAt" yaml:"updatedAt"`
	TenantID       string               `json:"tenantID" yaml:"tenantID"`
	IsAdmin        bool                 `json:"isAdmin" yaml:"isAdmin"`
	IsDeviceAdmin  bool                 `json:"isDeviceAdmin" yaml:"isDeviceAdmin"`
	IsGatewayAdmin bool                 `json:"isGatewayAdmin" yaml:"isGatewayAdmin"`
	Email          string               `json:"email" yaml:"email"`
	DeviceProfiles []*UserDeviceProfile `json:"deviceProfiles" yaml:"deviceProfiles"`
	Applications   []*UserApplication   `json:"applications" yaml:"applications"`
}

type UserApplication struct {
	ApplicationID string `json:"applicationID" yaml:"applicationID"`
	IsReadOnly    bool   `json:"isReadOnly" yaml:"isReadOnly"`
}

type UserDeviceProfile struct {
	DeviceProfileID string `json:"deviceProfileID" yaml:"deviceProfileID"`
}

type UpdateUserParams struct{}

type DeleteUserParams struct{}

type ListUsersParams struct {
	shared.ListPaginationRequest
	TenantID string
}

type ListUserItem struct {
	ID             string     `json:"id" yaml:"id"`
	CreatedAt      *time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt" yaml:"updatedAt"`
	Email          string     `json:"email" yaml:"email"`
	IsAdmin        bool       `json:"isAdmin" yaml:"isAdmin"`
	IsGatewayAdmin bool       `json:"isGatewayAdmin" yaml:"isGatewayAdmin"`
	IsDeviceAdmin  bool       `json:"isDeviceAdmin" yaml:"isDeviceAdmin"`
}
