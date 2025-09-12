package user

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	Create(ctx context.Context, in *apiv4.CreateUserRequest, opts ...grpc.CallOption) (*apiv4.CreateUserResponse, error)
	Get(ctx context.Context, in *apiv4.GetUserRequest, opts ...grpc.CallOption) (*apiv4.GetUserResponse, error)
	Update(ctx context.Context, in *apiv4.UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *apiv4.DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *apiv4.ListUsersRequest, opts ...grpc.CallOption) (*apiv4.ListUsersResponse, error)
	UpdatePassword(ctx context.Context, in *apiv4.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type grpcUserAPI struct{ c apiv4.UserServiceClient }

func (g *grpcUserAPI) Create(ctx context.Context, in *apiv4.CreateUserRequest, opts ...grpc.CallOption) (*apiv4.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ Client = (*grpcUserAPI)(nil)

func NewWithClient(c apiv4.UserServiceClient) Client { return &grpcUserAPI{c: c} }

func (g *grpcUserAPI) Get(ctx context.Context, in *apiv4.GetUserRequest, opts ...grpc.CallOption) (*apiv4.GetUserResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcUserAPI) Update(ctx context.Context, in *apiv4.UpdateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcUserAPI) Delete(ctx context.Context, in *apiv4.DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcUserAPI) List(ctx context.Context, in *apiv4.ListUsersRequest, opts ...grpc.CallOption) (*apiv4.ListUsersResponse, error) {
	return g.c.List(ctx, in, opts...)
}

func (g *grpcUserAPI) UpdatePassword(ctx context.Context, in *apiv4.UpdateUserPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdatePassword(ctx, in, opts...)
}
