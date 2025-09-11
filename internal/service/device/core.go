package device

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Get(ctx context.Context, p GetParams) (*Device, error)
	Create(ctx context.Context, p CreateParams) (string, error)
	List(ctx context.Context, p ListParams) (shared.ListResult[ListItem], error)
	Update(ctx context.Context, p UpdateParams) error
	Delete(ctx context.Context, p DeleteParams) error
}

// Device represent a Device
type Device struct {
	DevEUI          string
	Name            string
	Description     string
	ApplicationID   string
	DeviceProfileID string
	SkipFcntCheck   bool
	IsDisabled      bool
	Variables       map[string]string
	Tags            map[string]string
	JoinEui         string
}

type CreateParams struct {
	Device Device
}

type Patch struct {
	Name            *string
	Description     *string
	ApplicationID   *string
	DeviceProfileID *string
	SkipFcntCheck   *bool
	IsDisabled      *bool
	Variables       *shared.MapPatch
	Tags            *shared.MapPatch
	JoinEui         *string
}

type UpdateParams struct {
	DevEUI string
	Patch  Patch
}

type GetParams struct {
	DevEUI string
}

type ListItem struct {
	DevEUI            string
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	LastSeenAt        *time.Time
	Name              string
	Description       string
	DeviceProfileID   string
	DeviceProfileName string
	DeviceStatus      *string
	Tags              map[string]string
}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	shared.ListTagsFilterRequest
	DeviceProfileID string
	ApplicationID   string
}

type DeleteParams struct{}
