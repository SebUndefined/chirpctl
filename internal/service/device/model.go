package device

import "time"

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
	Device *Device
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

type ListResult struct {
	TotalCount uint32
	Result     []*ListItem
}

type ListParams struct {
	ApplicationID string
}
