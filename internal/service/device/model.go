package device

import (
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

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
