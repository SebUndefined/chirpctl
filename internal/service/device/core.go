package device

import (
	"context"
	"fmt"
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
	Margin              int32   `json:"margin" yaml:"margin"`
	ExternalPowerSource bool    `json:"externalPowerSource" yaml:"externalPowerSource"`
	BatteryLevel        float32 `json:"batteryLevel" yaml:"batteryLevel"`
}

// Device represent a Device
type Device struct {
	DevEUI          string            `json:"devEUI" yaml:"devEUI"`
	CreatedAt       *time.Time        `json:"createdAt" yaml:"createdAt"`
	UpdatedAt       *time.Time        `json:"updatedAt" yaml:"updatedAt"`
	LastSeenAt      *time.Time        `json:"lastSeenAt" yaml:"lastSeenAt"`
	Name            string            `json:"name" yaml:"name"`
	Description     string            `json:"description" yaml:"description"`
	ApplicationID   string            `json:"applicationID" yaml:"applicationID"`
	DeviceProfileID string            `json:"deviceProfileID" yaml:"deviceProfileID"`
	SkipFcntCheck   bool              `json:"skipFcntCheck" yaml:"skipFcntCheck"`
	IsDisabled      bool              `json:"isDisabled" yaml:"isDisabled"`
	Variables       map[string]string `json:"variables" yaml:"variables"`
	Tags            map[string]string `json:"tags" yaml:"tags"`
	JoinEUI         string            `json:"joinEUI" yaml:"joinEUI"`
}

func (d Device) String() string {
	return fmt.Sprintf("name: %s, devEUI:%s, Description: %s, ApplicationID: %s, DeviceProfileID: %s, SkipFcntCheck: %t, IsDisabled: %t, Variables: %s, Tags: %s JoinEUI: %s",
		d.Name, d.DevEUI, d.Description, d.DeviceProfileID, d.ApplicationID, d.SkipFcntCheck, d.IsDisabled, d.Variables, d.Tags, d.JoinEUI)
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
	DevEUI            string            `json:"devEUI" yaml:"devEUI"`
	CreatedAt         *time.Time        `json:"createdAt" yaml:"createdAt"`
	UpdatedAt         *time.Time        `json:"updatedAt" yaml:"updatedAt"`
	LastSeenAt        *time.Time        `json:"lastSeenAt" yaml:"lastSeenAt"`
	Name              string            `json:"name" yaml:"name"`
	Description       string            `json:"description" yaml:"description"`
	DeviceProfileID   string            `json:"deviceProfileID" yaml:"deviceProfileID"`
	DeviceProfileName string            `json:"deviceProfileName" yaml:"deviceProfileName"`
	DeviceStatus      *Status           `json:"deviceStatus" yaml:"deviceStatus"`
	Tags              map[string]string `json:"tags" yaml:"tags"`
}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	shared.ListTagsFilterRequest
	DeviceProfileID string
	ApplicationID   string
}

type DeleteParams struct{}
