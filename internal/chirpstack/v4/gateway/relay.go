package gateway

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type relay struct{}

var _ gateway.Relay = (*relay)(nil)

func (r *relay) GetRelayGateway(ctx context.Context, in gateway.GetRelayGatewayParams) (gateway.GetRelayGatewayResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *relay) ListRelayGateways(ctx context.Context, in gateway.ListRelayGatewaysParams) (shared.ListResult[gateway.ListRelayItem], error) {
	//TODO implement me
	panic("implement me")
}

func (r *relay) UpdateRelayGateway(ctx context.Context, in gateway.UpdateRelayGatewayParams) error {
	//TODO implement me
	panic("implement me")
}

func (r *relay) DeleteRelayGateway(ctx context.Context, in gateway.DeleteRelayGatewayParams) error {
	//TODO implement me
	panic("implement me")
}
