package tenant

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	Create(ctx context.Context, in *apiv4.CreateTenantRequest, opts ...grpc.CallOption) (*apiv4.CreateTenantResponse, error)
	Get(ctx context.Context, in *apiv4.GetTenantRequest, opts ...grpc.CallOption) (*apiv4.GetTenantResponse, error)
	Update(ctx context.Context, in *apiv4.UpdateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *apiv4.DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *apiv4.ListTenantsRequest, opts ...grpc.CallOption) (*apiv4.ListTenantsResponse, error)
	AddUser(ctx context.Context, in *apiv4.AddTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUser(ctx context.Context, in *apiv4.GetTenantUserRequest, opts ...grpc.CallOption) (*apiv4.GetTenantUserResponse, error)
	UpdateUser(ctx context.Context, in *apiv4.UpdateTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteUser(ctx context.Context, in *apiv4.DeleteTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListUsers(ctx context.Context, in *apiv4.ListTenantUsersRequest, opts ...grpc.CallOption) (*apiv4.ListTenantUsersResponse, error)
}

type grpcTenantAPI struct{ c apiv4.TenantServiceClient }

var _ Client = (*grpcTenantAPI)(nil)

func NewWithClient(c apiv4.TenantServiceClient) Client { return &grpcTenantAPI{c: c} }

func (g *grpcTenantAPI) Create(ctx context.Context, in *apiv4.CreateTenantRequest, opts ...grpc.CallOption) (*apiv4.CreateTenantResponse, error) {
	return g.c.Create(ctx, in, opts...)
}

func (g *grpcTenantAPI) Get(ctx context.Context, in *apiv4.GetTenantRequest, opts ...grpc.CallOption) (*apiv4.GetTenantResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcTenantAPI) Update(ctx context.Context, in *apiv4.UpdateTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcTenantAPI) Delete(ctx context.Context, in *apiv4.DeleteTenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcTenantAPI) List(ctx context.Context, in *apiv4.ListTenantsRequest, opts ...grpc.CallOption) (*apiv4.ListTenantsResponse, error) {
	return g.c.List(ctx, in, opts...)
}

func (g *grpcTenantAPI) AddUser(ctx context.Context, in *apiv4.AddTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.AddUser(ctx, in, opts...)
}

func (g *grpcTenantAPI) GetUser(ctx context.Context, in *apiv4.GetTenantUserRequest, opts ...grpc.CallOption) (*apiv4.GetTenantUserResponse, error) {
	return g.c.GetUser(ctx, in, opts...)
}

func (g *grpcTenantAPI) UpdateUser(ctx context.Context, in *apiv4.UpdateTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateUser(ctx, in, opts...)
}

func (g *grpcTenantAPI) DeleteUser(ctx context.Context, in *apiv4.DeleteTenantUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteUser(ctx, in, opts...)
}

func (g *grpcTenantAPI) ListUsers(ctx context.Context, in *apiv4.ListTenantUsersRequest, opts ...grpc.CallOption) (*apiv4.ListTenantUsersResponse, error) {
	return g.c.ListUsers(ctx, in, opts...)
}
