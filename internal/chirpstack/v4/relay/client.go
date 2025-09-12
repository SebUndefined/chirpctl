package relay

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	List(ctx context.Context, in *apiv4.ListRelaysRequest, opts ...grpc.CallOption) (*apiv4.ListRelaysResponse, error)
	AddDevice(ctx context.Context, in *apiv4.AddRelayDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveDevice(ctx context.Context, in *apiv4.RemoveRelayDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListDevices(ctx context.Context, in *apiv4.ListRelayDevicesRequest, opts ...grpc.CallOption) (*apiv4.ListRelayDevicesResponse, error)
}

type grpcRelayAPI struct{ c apiv4.RelayServiceClient }

var _ Client = (*grpcRelayAPI)(nil)

func NewWithClient(c apiv4.RelayServiceClient) Client { return &grpcRelayAPI{c: c} }

func (g *grpcRelayAPI) List(ctx context.Context, in *apiv4.ListRelaysRequest, opts ...grpc.CallOption) (*apiv4.ListRelaysResponse, error) {
	return g.c.List(ctx, in)
}

func (g *grpcRelayAPI) AddDevice(ctx context.Context, in *apiv4.AddRelayDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.AddDevice(ctx, in)
}

func (g *grpcRelayAPI) RemoveDevice(ctx context.Context, in *apiv4.RemoveRelayDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.RemoveDevice(ctx, in)
}

func (g *grpcRelayAPI) ListDevices(ctx context.Context, in *apiv4.ListRelayDevicesRequest, opts ...grpc.CallOption) (*apiv4.ListRelayDevicesResponse, error) {
	return g.c.ListDevices(ctx, in)
}
