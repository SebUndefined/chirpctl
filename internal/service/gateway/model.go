package gateway

import (
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

// Gateway represent a Device
type Gateway struct {
	GatewayID       string
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
	Gateway *Gateway
}

type GetParams struct {
	DevEUI string
}

type ListItem struct {
	GatewayID   string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	LastSeenAt  *time.Time
	Name        string
	Description string
	State       string
	Properties  map[string]string
	TenantID    string
}

type ListResult struct {
	TotalCount uint32
	Result     []*ListItem
}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	shared.ListTagsFilterRequest
	TenantID         string
	MulticastGroupID string
}

type DeleteParams struct {
	GatewayID string
}
