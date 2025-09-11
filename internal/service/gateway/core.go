package gateway

import (
	"context"
	"strings"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Service interface {
	Get(ctx context.Context, p GetParams) (*Gateway, error)
	Create(ctx context.Context, p CreateParams) error
	List(ctx context.Context, p ListParams) (ListResult, error)
	Delete(ctx context.Context, p DeleteParams) error
}

type State string

const (
	StateNeverSeen State = "NEVER_SEEN"
	StateOnline    State = "ONLINE"
	StateOffline   State = "OFFLINE"
	StateUnknown   State = "UNKNOWN"
)

func (s State) String() string {
	if len(strings.TrimSpace(string(s))) == 0 {
		return StateUnknown.String()
	}
	return string(s)
}

func ParseStateFromInt32(s int32) State {
	switch s {
	case 0:
		return StateNeverSeen
	case 1:
		return StateOnline
	case 2:
		return StateOffline
	default:
		return StateUnknown
	}
}

// Gateway represent a Gateway
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
	State       State
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
