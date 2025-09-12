package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	Create(ctx context.Context, in *apiv4.CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get returns the device for the given DevEUI.
	Get(ctx context.Context, in *apiv4.GetDeviceRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceResponse, error)
	// Update the given device.
	Update(ctx context.Context, in *apiv4.UpdateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the device with the given DevEUI.
	Delete(ctx context.Context, in *apiv4.DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the list of devices.
	List(ctx context.Context, in *apiv4.ListDevicesRequest, opts ...grpc.CallOption) (*apiv4.ListDevicesResponse, error)
	// Create the given device-keys.
	CreateKeys(ctx context.Context, in *apiv4.CreateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the device-keys for the given DevEUI.
	GetKeys(ctx context.Context, in *apiv4.GetDeviceKeysRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceKeysResponse, error)
	// Update the given device-keys.
	UpdateKeys(ctx context.Context, in *apiv4.UpdateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the device-keys for the given DevEUI.
	DeleteKeys(ctx context.Context, in *apiv4.DeleteDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// FlushDevNonces flushes the OTAA device nonces.
	FlushDevNonces(ctx context.Context, in *apiv4.FlushDevNoncesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Activate (re)activates the device with the given parameters (for ABP or for
	// importing OTAA activations).
	Activate(ctx context.Context, in *apiv4.ActivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deactivate de-activates the device.
	Deactivate(ctx context.Context, in *apiv4.DeactivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetActivation returns the current activation details of the device (OTAA or
	// ABP).
	GetActivation(ctx context.Context, in *apiv4.GetDeviceActivationRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceActivationResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into
	// account.
	GetRandomDevAddr(ctx context.Context, in *apiv4.GetRandomDevAddrRequest, opts ...grpc.CallOption) (*apiv4.GetRandomDevAddrResponse, error)
	// GetMetrics returns the device metrics.
	// Note that this requires a device-profile with codec and measurements
	// configured.
	GetMetrics(ctx context.Context, in *apiv4.GetDeviceMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceMetricsResponse, error)
	// GetLinkMetrics returns the device link metrics.
	// This includes uplinks, downlinks, RSSI, SNR, etc...
	GetLinkMetrics(ctx context.Context, in *apiv4.GetDeviceLinkMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceLinkMetricsResponse, error)
	// Enqueue adds the given item to the downlink queue.
	Enqueue(ctx context.Context, in *apiv4.EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*apiv4.EnqueueDeviceQueueItemResponse, error)
	// FlushQueue flushes the downlink device-queue.
	FlushQueue(ctx context.Context, in *apiv4.FlushDeviceQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetQueue returns the downlink device-queue.
	GetQueue(ctx context.Context, in *apiv4.GetDeviceQueueItemsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceQueueItemsResponse, error)
	// GetNextFCntDown returns the next FCntDown to use for enqueing encrypted
	// downlinks. The difference with the DeviceActivation f_cont_down is that
	// this method takes potential existing queue-items into account.
	GetNextFCntDown(ctx context.Context, in *apiv4.GetDeviceNextFCntDownRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceNextFCntDownResponse, error)
}

type grpcDeviceAPI struct{ c apiv4.DeviceServiceClient }

func NewWithClient(c apiv4.DeviceServiceClient) Client { return &grpcDeviceAPI{c: c} }

// assert compile-time
var _ Client = (*grpcDeviceAPI)(nil)

func (g *grpcDeviceAPI) Get(ctx context.Context, in *apiv4.GetDeviceRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Update(ctx context.Context, in *apiv4.UpdateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Delete(ctx context.Context, in *apiv4.DeleteDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcDeviceAPI) List(ctx context.Context, in *apiv4.ListDevicesRequest, opts ...grpc.CallOption) (*apiv4.ListDevicesResponse, error) {
	return g.c.List(ctx, in, opts...)
}

func (g *grpcDeviceAPI) CreateKeys(ctx context.Context, in *apiv4.CreateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateKeys(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetKeys(ctx context.Context, in *apiv4.GetDeviceKeysRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceKeysResponse, error) {
	return g.c.GetKeys(ctx, in, opts...)
}

func (g *grpcDeviceAPI) UpdateKeys(ctx context.Context, in *apiv4.UpdateDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateKeys(ctx, in, opts...)
}

func (g *grpcDeviceAPI) DeleteKeys(ctx context.Context, in *apiv4.DeleteDeviceKeysRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteKeys(ctx, in, opts...)
}

func (g *grpcDeviceAPI) FlushDevNonces(ctx context.Context, in *apiv4.FlushDevNoncesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.FlushDevNonces(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Activate(ctx context.Context, in *apiv4.ActivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Activate(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Deactivate(ctx context.Context, in *apiv4.DeactivateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Deactivate(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetActivation(ctx context.Context, in *apiv4.GetDeviceActivationRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceActivationResponse, error) {
	return g.c.GetActivation(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetRandomDevAddr(ctx context.Context, in *apiv4.GetRandomDevAddrRequest, opts ...grpc.CallOption) (*apiv4.GetRandomDevAddrResponse, error) {
	return g.c.GetRandomDevAddr(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetMetrics(ctx context.Context, in *apiv4.GetDeviceMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceMetricsResponse, error) {
	return g.c.GetMetrics(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetLinkMetrics(ctx context.Context, in *apiv4.GetDeviceLinkMetricsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceLinkMetricsResponse, error) {
	return g.c.GetLinkMetrics(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Enqueue(ctx context.Context, in *apiv4.EnqueueDeviceQueueItemRequest, opts ...grpc.CallOption) (*apiv4.EnqueueDeviceQueueItemResponse, error) {
	return g.c.Enqueue(ctx, in, opts...)
}

func (g *grpcDeviceAPI) FlushQueue(ctx context.Context, in *apiv4.FlushDeviceQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.FlushQueue(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetQueue(ctx context.Context, in *apiv4.GetDeviceQueueItemsRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceQueueItemsResponse, error) {
	return g.c.GetQueue(ctx, in, opts...)
}

func (g *grpcDeviceAPI) GetNextFCntDown(ctx context.Context, in *apiv4.GetDeviceNextFCntDownRequest, opts ...grpc.CallOption) (*apiv4.GetDeviceNextFCntDownResponse, error) {
	return g.c.GetNextFCntDown(ctx, in, opts...)
}

func (g *grpcDeviceAPI) Create(ctx context.Context, in *apiv4.CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Create(ctx, in, opts...)
}
