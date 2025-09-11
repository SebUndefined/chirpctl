package v4

import (
	"context"
	"time"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/samber/lo"
	svcgateway "github.com/sebundefined/chirpctl/internal/service/gateway"
	"google.golang.org/grpc"
)

type gatewaySvc struct{ c apiv4.GatewayServiceClient }

func NewGatewayService(conn *grpc.ClientConn) svcgateway.Service {
	return &gatewaySvc{c: apiv4.NewGatewayServiceClient(conn)}
}

func (s *gatewaySvc) List(ctx context.Context, p svcgateway.ListParams) (svcgateway.ListResult, error) {
	list, err := s.c.List(ctx, &apiv4.ListGatewaysRequest{
		Limit:            p.GetSafeLimit(),
		Offset:           p.Offset,
		Search:           p.Search,
		TenantId:         p.TenantID,
		MulticastGroupId: p.MulticastGroupID,
		OrderBy:          1,
		OrderByDesc:      p.OrderByDesc,
	})
	if err != nil {
		return svcgateway.ListResult{}, err
	}
	res := svcgateway.ListResult{
		TotalCount: list.GetTotalCount(),
	}
	res.Result = lo.Map(list.Result, func(item *apiv4.GatewayListItem, _ int) *svcgateway.ListItem {
		return &svcgateway.ListItem{
			GatewayID: item.GatewayId,
			CreatedAt: lo.TernaryF(item.CreatedAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.CreatedAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			UpdatedAt: lo.TernaryF(item.UpdatedAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.UpdatedAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			LastSeenAt: lo.TernaryF(item.LastSeenAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.LastSeenAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			Name:        item.Name,
			Description: item.Description,
			State:       svcgateway.ParseStateFromInt32(int32(item.State)),
			Properties:  item.Properties,
			TenantID:    item.TenantId,
		}
	})
	return res, nil
}

func (s *gatewaySvc) Get(ctx context.Context, p svcgateway.GetParams) (*svcgateway.Gateway, error) {
	//TODO implement me
	panic("implement me")
}

func (s *gatewaySvc) Create(ctx context.Context, p svcgateway.CreateParams) error {
	//TODO implement me
	panic("implement me")
}

func (s *gatewaySvc) Delete(ctx context.Context, p svcgateway.DeleteParams) error {
	_, err := s.c.Delete(ctx, &apiv4.DeleteGatewayRequest{
		GatewayId: p.GatewayID,
	})
	return err
}
