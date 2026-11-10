package gateway

import (
	"context"
	"testing"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	gatewaymocks "github.com/sebundefined/chirpctl/internal/chirpstack/v4/gateway/mocks"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	testGatewayID = "0016c001f12a7d36"
	testTenantID  = "5409bf8b-777c-4ec7-b3da-c6276899b499"
)

type GatewayCoreSuite struct {
	suite.Suite

	client *gatewaymocks.ClientMock
	core   *core
}

func TestGatewayCoreSuite(t *testing.T) {
	suite.Run(t, new(GatewayCoreSuite))
}

func (s *GatewayCoreSuite) SetupTest() {
	s.client = gatewaymocks.NewClientMock(s.T())
	s.core = newCore(s.client)
}

func (s *GatewayCoreSuite) TestGetGatewayErrorGrpc() {
	expectedErr := status.Error(codes.Unavailable, "service unavailable")

	s.client.EXPECT().
		Get(mock.Anything, mock.Anything).
		Return(nil, expectedErr).
		Once()

	result, err := s.core.Get(
		context.Background(),
		gateway.GetParams{GatewayID: testGatewayID},
	)

	s.ErrorIs(err, expectedErr)
	s.Nil(result)
	s.Equal(codes.Unavailable, status.Code(err))
}

func (s *GatewayCoreSuite) TestGetGatewaySuccess() {
	gatewayID := "test-641A"
	s.client.EXPECT().
		Get(
			mock.Anything,
			mock.MatchedBy(func(req *apiv4.GetGatewayRequest) bool {
				return req.GetGatewayId() == gatewayID
			}),
		).
		Return(&apiv4.GetGatewayResponse{
			Gateway: &apiv4.Gateway{
				GatewayId: gatewayID,
				Name:      "test-gateway",
			},
		}, nil).
		Once()
	result, err := s.core.Get(
		context.Background(),
		gateway.GetParams{GatewayID: gatewayID},
	)
	s.Require().NoError(err)
	s.Require().NotNil(result)
	s.Equal(gatewayID, result.GatewayID)
	s.Equal("test-gateway", result.Name)
}

func (s *GatewayCoreSuite) TestGetGatewayEmptyResponse() {
	s.client.EXPECT().
		Get(mock.Anything, mock.Anything).
		Return(&apiv4.GetGatewayResponse{}, nil).
		Once()

	result, err := s.core.Get(
		context.Background(),
		gateway.GetParams{GatewayID: testGatewayID},
	)

	s.Require().Error(err)
	s.ErrorContains(err, "gateway response is empty")
	s.Nil(result)
}

func (s *GatewayCoreSuite) TestListGatewaysRequestSuccess() {
	params := gateway.ListParams{
		TenantID: testTenantID,

		Search:      "smaep",
		Offset:      5,
		OrderByDesc: true,
	}
	s.client.EXPECT().
		List(
			mock.Anything,
			mock.MatchedBy(func(req *apiv4.ListGatewaysRequest) bool {
				return req != nil &&
					req.GetTenantId() == testTenantID &&
					req.GetSearch() == "smaep" &&
					req.GetOffset() == 5 &&
					req.GetOrderByDesc()
			}),
		).
		Return(&apiv4.ListGatewaysResponse{
			TotalCount: 1,
			Result: []*apiv4.GatewayListItem{
				{
					GatewayId: testGatewayID,
					Name:      "test-gateway",
					TenantId:  testTenantID,
				},
			},
		}, nil).
		Once()
	result, err := s.core.List(context.Background(), params)
	s.Require().NoError(err)
	s.Equal(uint32(1), result.TotalCount)
	s.Require().Len(result.Items, 1)

	s.Equal(testGatewayID, result.Items[0].GatewayID)
	s.Equal(testTenantID, result.Items[0].TenantID)
	s.Equal("test-gateway", result.Items[0].Name)
}

func (s *GatewayCoreSuite) TestDeleteGatewaySuccess() {
	s.client.EXPECT().
		Delete(
			mock.Anything,
			mock.MatchedBy(func(req *apiv4.DeleteGatewayRequest) bool {
				return req != nil && req.GetGatewayId() == testGatewayID
			}),
		).
		Return(&emptypb.Empty{}, nil).
		Once()

	err := s.core.Delete(
		context.Background(),
		gateway.DeleteParams{GatewayID: testGatewayID},
	)

	s.Require().NoError(err)
}

func (s *GatewayCoreSuite) TestDeleteGatewayErrorGrpc() {
	expectedErr := status.Error(codes.PermissionDenied, "permission denied")

	s.client.EXPECT().
		Delete(mock.Anything, mock.Anything).
		Return(nil, expectedErr).
		Once()

	err := s.core.Delete(
		context.Background(),
		gateway.DeleteParams{GatewayID: testGatewayID},
	)

	s.ErrorIs(err, expectedErr)
	s.Equal(codes.PermissionDenied, status.Code(err))
}
