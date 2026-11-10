package gateway

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	// Create creates the given gateway.
	Create(ctx context.Context, in *apiv4.CreateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get returns the gateway for the given Gateway ID.
	Get(ctx context.Context, in *apiv4.GetGatewayRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayResponse, error)
	// Update updates the given gateway.
	Update(ctx context.Context, in *apiv4.UpdateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete deletes the gateway matching the given Gateway ID.
	Delete(ctx context.Context, in *apiv4.DeleteGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the list of gateways.
	List(ctx context.Context, in *apiv4.ListGatewaysRequest, opts ...grpc.CallOption) (*apiv4.ListGatewaysResponse, error)
	// Generate client-certificate for the gateway.
	GenerateClientCertificate(ctx context.Context, in *apiv4.GenerateGatewayClientCertificateRequest, opts ...grpc.CallOption) (*apiv4.GenerateGatewayClientCertificateResponse, error)
	// GetMetrics returns the gateway metrics.
	GetMetrics(ctx context.Context, in *apiv4.GetGatewayMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayMetricsResponse, error)
	// GetDutyCycleMetrics returns the duty-cycle metrics.
	// Note that only the last 2 hours of data are stored. Currently only per minute aggregation is available.
	GetDutyCycleMetrics(ctx context.Context, in *apiv4.GetGatewayDutyCycleMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayDutyCycleMetricsResponse, error)
	// Get the given Relay Gateway.
	GetRelayGateway(ctx context.Context, in *apiv4.GetRelayGatewayRequest, opts ...grpc.CallOption) (*apiv4.GetRelayGatewayResponse, error)
	// List the detected Relay Gateways.
	ListRelayGateways(ctx context.Context, in *apiv4.ListRelayGatewaysRequest, opts ...grpc.CallOption) (*apiv4.ListRelayGatewaysResponse, error)
	// Update the given Relay Gateway.
	UpdateRelayGateway(ctx context.Context, in *apiv4.UpdateRelayGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the given Relay Gateway.
	DeleteRelayGateway(ctx context.Context, in *apiv4.DeleteRelayGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type grpcGatewayAPI struct {
	c apiv4.GatewayServiceClient
}

var _ Client = (*grpcGatewayAPI)(nil)

func NewWithClient(c apiv4.GatewayServiceClient) Client { return &grpcGatewayAPI{c: c} }

func (g *grpcGatewayAPI) Create(ctx context.Context, in *apiv4.CreateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Create(ctx, in, opts...)
}

func (g *grpcGatewayAPI) Get(ctx context.Context, in *apiv4.GetGatewayRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcGatewayAPI) Update(ctx context.Context, in *apiv4.UpdateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcGatewayAPI) Delete(ctx context.Context, in *apiv4.DeleteGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcGatewayAPI) List(ctx context.Context, in *apiv4.ListGatewaysRequest, opts ...grpc.CallOption) (*apiv4.ListGatewaysResponse, error) {
	return g.c.List(ctx, in, opts...)
}

func (g *grpcGatewayAPI) GenerateClientCertificate(ctx context.Context, in *apiv4.GenerateGatewayClientCertificateRequest, opts ...grpc.CallOption) (*apiv4.GenerateGatewayClientCertificateResponse, error) {
	return g.c.GenerateClientCertificate(ctx, in, opts...)
}

func (g *grpcGatewayAPI) GetMetrics(ctx context.Context, in *apiv4.GetGatewayMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayMetricsResponse, error) {
	return g.c.GetMetrics(ctx, in, opts...)
}

func (g *grpcGatewayAPI) GetDutyCycleMetrics(ctx context.Context, in *apiv4.GetGatewayDutyCycleMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetGatewayDutyCycleMetricsResponse, error) {
	return g.c.GetDutyCycleMetrics(ctx, in, opts...)
}

func (g *grpcGatewayAPI) GetRelayGateway(ctx context.Context, in *apiv4.GetRelayGatewayRequest, opts ...grpc.CallOption) (*apiv4.GetRelayGatewayResponse, error) {
	return g.c.GetRelayGateway(ctx, in, opts...)
}

func (g *grpcGatewayAPI) ListRelayGateways(ctx context.Context, in *apiv4.ListRelayGatewaysRequest, opts ...grpc.CallOption) (*apiv4.ListRelayGatewaysResponse, error) {
	return g.c.ListRelayGateways(ctx, in, opts...)
}

func (g *grpcGatewayAPI) UpdateRelayGateway(ctx context.Context, in *apiv4.UpdateRelayGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateRelayGateway(ctx, in, opts...)
}

func (g *grpcGatewayAPI) DeleteRelayGateway(ctx context.Context, in *apiv4.DeleteRelayGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteRelayGateway(ctx, in, opts...)
}
