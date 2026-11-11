package multicastgrp

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	Create(ctx context.Context, in *apiv4.CreateMulticastGroupRequest, opts ...grpc.CallOption) (*apiv4.CreateMulticastGroupResponse, error)
	Get(ctx context.Context, in *apiv4.GetMulticastGroupRequest, opts ...grpc.CallOption) (*apiv4.GetMulticastGroupResponse, error)
	Update(ctx context.Context, in *apiv4.UpdateMulticastGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *apiv4.DeleteMulticastGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *apiv4.ListMulticastGroupsRequest, opts ...grpc.CallOption) (*apiv4.ListMulticastGroupsResponse, error)
}

type grpcMulticastGrpAPI struct {
	c apiv4.MulticastGroupServiceClient
}

var _ Client = (*grpcMulticastGrpAPI)(nil)

func NewWithClient(c apiv4.MulticastGroupServiceClient) Client {
	return &grpcMulticastGrpAPI{c: c}
}

func (g *grpcMulticastGrpAPI) Create(ctx context.Context, in *apiv4.CreateMulticastGroupRequest, opts ...grpc.CallOption) (*apiv4.CreateMulticastGroupResponse, error) {
	return g.c.Create(ctx, in, opts...)
}

func (g *grpcMulticastGrpAPI) Get(ctx context.Context, in *apiv4.GetMulticastGroupRequest, opts ...grpc.CallOption) (*apiv4.GetMulticastGroupResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcMulticastGrpAPI) Update(ctx context.Context, in *apiv4.UpdateMulticastGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcMulticastGrpAPI) Delete(ctx context.Context, in *apiv4.DeleteMulticastGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcMulticastGrpAPI) List(ctx context.Context, in *apiv4.ListMulticastGroupsRequest, opts ...grpc.CallOption) (*apiv4.ListMulticastGroupsResponse, error) {
	return g.c.List(ctx, in, opts...)
}
