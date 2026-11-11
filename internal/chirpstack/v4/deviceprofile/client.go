package deviceprofile

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	Create(ctx context.Context, in *apiv4.CreateDeviceProfileRequest, opts ...grpc.CallOption) (*apiv4.CreateDeviceProfileResponse, error)
	Get(ctx context.Context, in *apiv4.GetDeviceProfileRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceProfileResponse, error)
	Update(ctx context.Context, in *apiv4.UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *apiv4.DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *apiv4.ListDeviceProfilesRequest, opts ...grpc.CallOption) (*apiv4.ListDeviceProfilesResponse, error)
	ListAdrAlgorithms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*apiv4.ListDeviceProfileAdrAlgorithmsResponse, error)
}

type grpcDeviceProfileAPI struct {
	c apiv4.DeviceProfileServiceClient
}

var _ Client = (*grpcDeviceProfileAPI)(nil)

func NewWithClient(c apiv4.DeviceProfileServiceClient) Client { return &grpcDeviceProfileAPI{c: c} }

func (c *grpcDeviceProfileAPI) Create(ctx context.Context, in *apiv4.CreateDeviceProfileRequest, opts ...grpc.CallOption) (*apiv4.CreateDeviceProfileResponse, error) {
	return c.c.Create(ctx, in)
}

func (c *grpcDeviceProfileAPI) Get(ctx context.Context, in *apiv4.GetDeviceProfileRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceProfileResponse, error) {
	return c.c.Get(ctx, in)
}

func (c *grpcDeviceProfileAPI) Update(ctx context.Context, in *apiv4.UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.c.Update(ctx, in)
}

func (c *grpcDeviceProfileAPI) Delete(ctx context.Context, in *apiv4.DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return c.c.Delete(ctx, in)
}

func (c *grpcDeviceProfileAPI) List(ctx context.Context, in *apiv4.ListDeviceProfilesRequest, opts ...grpc.CallOption) (*apiv4.ListDeviceProfilesResponse, error) {
	return c.c.List(ctx, in)
}

func (c *grpcDeviceProfileAPI) ListAdrAlgorithms(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*apiv4.ListDeviceProfileAdrAlgorithmsResponse, error) {
	return c.c.ListAdrAlgorithms(ctx, in)
}
