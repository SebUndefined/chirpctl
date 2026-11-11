package gateway

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/sebundefined/chirpctl/internal/app"
	clientfactorymocks "github.com/sebundefined/chirpctl/internal/chirpstack/clientfactory/mocks"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	svcgatewaymocks "github.com/sebundefined/chirpctl/internal/service/gateway/mocks"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GatewayGetSuite struct {
	suite.Suite

	factory *clientfactorymocks.FactoryMock
	api     *svcgatewaymocks.APIMock

	app    *app.App
	output *bytes.Buffer
	cmd    *cobra.Command
}

func TestGatewayGetSuite(t *testing.T) {
	suite.Run(t, new(GatewayGetSuite))
}

func (s *GatewayGetSuite) SetupTest() {
	s.factory = clientfactorymocks.NewFactoryMock(s.T())
	s.api = svcgatewaymocks.NewAPIMock(s.T())
	s.output = new(bytes.Buffer)

	s.app = &app.App{
		Factory: s.factory,
		Printer: printing.NewPrinter(
			s.output,
			printing.Options{Format: printing.FormatTable},
		),
	}

	s.cmd = NewGetCommand(s.app)

	s.cmd = NewGetCommand(s.app)
	s.cmd.SetOut(s.output)
	s.cmd.SetErr(io.Discard)
	s.cmd.SilenceUsage = true
	s.cmd.SilenceErrors = true
}

func (s *GatewayGetSuite) TestGetGatewayFactoryError() {
	expectedErr := errors.New("factory unavailable")

	s.factory.EXPECT().
		Gateway().
		Return(nil, expectedErr).
		Once()

	s.cmd.SetArgs([]string{"0016c001f12a7d36"})

	err := s.cmd.ExecuteContext(context.Background())

	s.ErrorIs(err, expectedErr)
}

func (s *GatewayGetSuite) TestGetGatewaySuccess() {
	const gatewayID = "0016c001f12a7d36"

	s.factory.EXPECT().
		Gateway().
		Return(s.api, nil).
		Once()

	s.api.EXPECT().
		Get(mock.Anything, gateway.GetParams{
			GatewayID: gatewayID,
		}).
		Return(&gateway.Gateway{
			GatewayID: gatewayID,
			Name:      "test-gateway",
		}, nil).
		Once()

	s.cmd.SetArgs([]string{gatewayID})

	err := s.cmd.ExecuteContext(context.Background())

	s.Require().NoError(err)
	s.Contains(s.output.String(), "test-gateway")
	s.Contains(s.output.String(), gatewayID)
}

func (s *GatewayGetSuite) TestGetGatewayInvalidID() {
	s.cmd.SetArgs([]string{"invalid-id"})

	err := s.cmd.ExecuteContext(context.Background())

	s.Require().Error(err)
	s.ErrorContains(err, "invalid gateway ID")

	s.factory.AssertNotCalled(s.T(), "Gateway")
}

func (s *GatewayGetSuite) TestGetGatewayNilResult() {
	const gatewayID = "0016c001f12a7d36"

	s.factory.EXPECT().
		Gateway().
		Return(s.api, nil).
		Once()

	s.api.EXPECT().
		Get(mock.Anything, gateway.GetParams{
			GatewayID: gatewayID,
		}).
		Return((*gateway.Gateway)(nil), nil).
		Once()

	s.cmd.SetArgs([]string{gatewayID})

	err := s.cmd.ExecuteContext(context.Background())

	s.ErrorContains(err, "empty response")
	s.Empty(s.output.String())
}

// expectGatewayList configures the expected tenant filter and service response.
func (s *GatewayGetSuite) expectGatewayList(
	tenantID string,
	result shared.ListResult[gateway.ListItem],
	serviceErr error,
) {
	s.factory.EXPECT().
		Gateway().
		Return(s.api, nil).
		Once()

	s.api.EXPECT().
		List(
			mock.Anything,
			mock.MatchedBy(func(p gateway.ListParams) bool {
				return p.TenantID == tenantID
			}),
		).
		Return(result, serviceErr).
		Once()
}

func (s *GatewayGetSuite) TestListGatewaysDefaultTenant() {
	const tenantID = "5409bf8b-777c-4ec7-b3da-c6276899b499"

	s.app.ResolvedOptions.TenantID = tenantID

	s.expectGatewayList(
		tenantID,
		shared.ListResult[gateway.ListItem]{},
		nil,
	)

	s.cmd.SetArgs([]string{})

	err := s.cmd.ExecuteContext(context.Background())

	s.Require().NoError(err)
}

func (s *GatewayGetSuite) TestListGatewaysAllTenants() {
	s.app.ResolvedOptions.TenantID = "5409bf8b-777c-4ec7-b3da-c6276899b499"

	s.expectGatewayList(
		"",
		shared.ListResult[gateway.ListItem]{},
		nil,
	)

	s.cmd.SetArgs([]string{"--all-tenants"})

	err := s.cmd.ExecuteContext(context.Background())

	s.Require().NoError(err)
}
