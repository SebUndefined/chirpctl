package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Relay interface {
	GetRelayGateway(ctx context.Context, in GetRelayGatewayParams) (GetRelayGatewayResult, error)
	ListRelayGateways(ctx context.Context, in ListRelayGatewaysParams) (shared.ListResult[ListRelayItem], error)
	UpdateRelayGateway(ctx context.Context, in UpdateRelayGatewayParams) error
	DeleteRelayGateway(ctx context.Context, in DeleteRelayGatewayParams) error
}

type GetRelayGatewayParams struct{}
type ListRelayGatewaysParams struct{}
type UpdateRelayGatewayParams struct{}
type DeleteRelayGatewayParams struct{}
type GetRelayGatewayResult struct{}
type ListRelayItem struct{}
