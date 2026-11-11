package gateway

import (
	"context"
	"fmt"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type core struct {
	client Client
}

var _ gateway.Core = (*core)(nil)

func newCore(client Client) *core {
	return &core{client: client}
}

func (s *core) List(ctx context.Context, p gateway.ListParams) (shared.ListResult[gateway.ListItem], error) {
	req := &apiv4.ListGatewaysRequest{
		Limit:            p.GetSafeLimit(),
		Offset:           p.Offset,
		Search:           p.Search,
		TenantId:         p.TenantID,
		MulticastGroupId: p.MulticastGroupID,
		OrderBy:          1,
		OrderByDesc:      p.OrderByDesc,
	}
	out := shared.NewListResult[gateway.ListItem]()
	resp, err := s.client.List(ctx, req)
	if err != nil {
		return out, err
	}
	items := resp.GetResult()
	out = shared.NewListResultCap[gateway.ListItem](len(items))
	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		converted, err := toGatewayListItem(item)
		if err != nil {
			return out, fmt.Errorf("convert gateway list item: %w", err)
		}
		out.Add(converted)
	}
	return out, nil
}

func (s *core) Get(ctx context.Context, p gateway.GetParams) (*gateway.Gateway, error) {
	req := &apiv4.GetGatewayRequest{
		GatewayId: p.GatewayID,
	}
	get, err := s.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return toGateway(get)
}

func (s *core) Create(ctx context.Context, p gateway.CreateParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *core) Delete(ctx context.Context, p gateway.DeleteParams) error {
	_, err := s.client.Delete(ctx, &apiv4.DeleteGatewayRequest{
		GatewayId: p.GatewayID,
	})
	return err
}
