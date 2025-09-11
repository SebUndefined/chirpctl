package gateway

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	svcgateway "github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"google.golang.org/grpc"
)

type core struct{ c apiv4.GatewayServiceClient }

func NewGatewayService(conn *grpc.ClientConn) svcgateway.Service {
	return &core{c: apiv4.NewGatewayServiceClient(conn)}
}

func (s *core) List(ctx context.Context, p svcgateway.ListParams) (shared.ListResult[svcgateway.ListItem], error) {
	req := &apiv4.ListGatewaysRequest{
		Limit:            p.GetSafeLimit(),
		Offset:           p.Offset,
		Search:           p.Search,
		TenantId:         p.TenantID,
		MulticastGroupId: p.MulticastGroupID,
		OrderBy:          1,
		OrderByDesc:      p.OrderByDesc,
	}
	out := shared.NewListResult[svcgateway.ListItem]()
	resp, err := s.c.List(ctx, req)
	if err != nil {
		return out, err
	}
	items := resp.GetResult()
	out = shared.NewListResultCap[svcgateway.ListItem](len(items))
	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		out.Add(svcgateway.ListItem{
			GatewayID:   item.GatewayId,
			CreatedAt:   shared.TsProtobufToPtrTime(item.CreatedAt),
			UpdatedAt:   shared.TsProtobufToPtrTime(item.UpdatedAt),
			LastSeenAt:  shared.TsProtobufToPtrTime(item.LastSeenAt),
			Name:        item.Name,
			Description: item.Description,
			State:       svcgateway.ParseStateFromInt32(int32(item.State)),
			Properties:  item.Properties,
			TenantID:    item.TenantId,
		})
	}
	return out, nil
}

func (s *core) Get(ctx context.Context, p svcgateway.GetParams) (*svcgateway.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (s *core) Create(ctx context.Context, p svcgateway.CreateParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *core) Delete(ctx context.Context, p svcgateway.DeleteParams) error {
	_, err := s.c.Delete(ctx, &apiv4.DeleteGatewayRequest{
		GatewayId: p.GatewayID,
	})
	return err
}
