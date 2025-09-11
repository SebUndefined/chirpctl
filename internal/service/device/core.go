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

type Status struct {
	Margin              int32
	ExternalPowerSource bool
	BatteryLevel        float32
}

// Device represent a Device
type Device struct {
	DevEUI          string            `json:"devEUI" yaml:"devEUI"`
	Name            string            `json:"name" yaml:"name"`
	Description     string            `json:"description" yaml:"description"`
	ApplicationID   string            `json:"applicationID" yaml:"applicationID"`
	DeviceProfileID string            `json:"deviceProfileID" yaml:"deviceProfileID"`
	SkipFcntCheck   bool              `json:"skipFcntCheck" yaml:"skipFcntCheck"`
	IsDisabled      bool              `json:"isDisabled" yaml:"isDisabled"`
	Variables       map[string]string `json:"variables" yaml:"variables"`
	Tags            map[string]string `json:"tags" yaml:"tags"`
	JoinEui         string            `json:"joinEUI" yaml:"joinEUI"`
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
	DeviceStatus      *Status
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
