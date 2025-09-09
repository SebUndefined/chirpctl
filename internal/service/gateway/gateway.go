package gateway

import (
	"context"
	"time"
)

type Service interface {
	Get(ctx context.Context, p GetParams) (*Gateway, error)
	Create(ctx context.Context, p CreateParams) error
	List(ctx context.Context, p ListParams) (ListResult, error)
}

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
	ApplicationID string
}
