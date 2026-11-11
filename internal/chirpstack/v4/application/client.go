package application

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client interface {
	// Create creates the given application.
	Create(ctx context.Context, in *apiv4.CreateApplicationRequest, opts ...grpc.CallOption) (*apiv4.CreateApplicationResponse, error)
	// Get the application for the given ID.
	Get(ctx context.Context, in *apiv4.GetApplicationRequest, opts ...grpc.CallOption) (*apiv4.GetApplicationResponse, error)
	// Update updates the given application.
	Update(ctx context.Context, in *apiv4.UpdateApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the application for the given ID.
	Delete(ctx context.Context, in *apiv4.DeleteApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the list of applications.
	List(ctx context.Context, in *apiv4.ListApplicationsRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationsResponse, error)
	// List all configured integrations.
	ListIntegrations(ctx context.Context, in *apiv4.ListIntegrationsRequest, opts ...grpc.CallOption) (*apiv4.ListIntegrationsResponse, error)
	// Create HTTP integration.
	CreateHttpIntegration(ctx context.Context, in *apiv4.CreateHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get the configured HTTP integration.
	GetHttpIntegration(ctx context.Context, in *apiv4.GetHttpIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetHttpIntegrationResponse, error)
	// Update the HTTP integration.
	UpdateHttpIntegration(ctx context.Context, in *apiv4.UpdateHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete the HTTP integration.
	DeleteHttpIntegration(ctx context.Context, in *apiv4.DeleteHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create InfluxDb integration.
	CreateInfluxDbIntegration(ctx context.Context, in *apiv4.CreateInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get InfluxDb integration.
	GetInfluxDbIntegration(ctx context.Context, in *apiv4.GetInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetInfluxDbIntegrationResponse, error)
	// Update InfluxDb integration.
	UpdateInfluxDbIntegration(ctx context.Context, in *apiv4.UpdateInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete InfluxDb integration.
	DeleteInfluxDbIntegration(ctx context.Context, in *apiv4.DeleteInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create ThingsBoard integration.
	CreateThingsBoardIntegration(ctx context.Context, in *apiv4.CreateThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get ThingsBoard integration.
	GetThingsBoardIntegration(ctx context.Context, in *apiv4.GetThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetThingsBoardIntegrationResponse, error)
	// Update ThingsBoard integration.
	UpdateThingsBoardIntegration(ctx context.Context, in *apiv4.UpdateThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete ThingsBoard integration.
	DeleteThingsBoardIntegration(ctx context.Context, in *apiv4.DeleteThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create myDevices integration.
	CreateMyDevicesIntegration(ctx context.Context, in *apiv4.CreateMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get myDevices integration.
	GetMyDevicesIntegration(ctx context.Context, in *apiv4.GetMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetMyDevicesIntegrationResponse, error)
	// Update myDevices integration.
	UpdateMyDevicesIntegration(ctx context.Context, in *apiv4.UpdateMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete myDevices integration.
	DeleteMyDevicesIntegration(ctx context.Context, in *apiv4.DeleteMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create GCP Pub/Sub integration.
	CreateGcpPubSubIntegration(ctx context.Context, in *apiv4.CreateGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get GCP Pub/Sub integration.
	GetGcpPubSubIntegration(ctx context.Context, in *apiv4.GetGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetGcpPubSubIntegrationResponse, error)
	// Update GCP Pub/Sub integration.
	UpdateGcpPubSubIntegration(ctx context.Context, in *apiv4.UpdateGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete GCP Pub/Sub integration.
	DeleteGcpPubSubIntegration(ctx context.Context, in *apiv4.DeleteGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create AWS SNS integration.
	CreateAwsSnsIntegration(ctx context.Context, in *apiv4.CreateAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get AWS SNS integration.
	GetAwsSnsIntegration(ctx context.Context, in *apiv4.GetAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetAwsSnsIntegrationResponse, error)
	// Update AWS SNS integration.
	UpdateAwsSnsIntegration(ctx context.Context, in *apiv4.UpdateAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete AWS SNS integration.
	DeleteAwsSnsIntegration(ctx context.Context, in *apiv4.DeleteAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create Azure Service-Bus integration.
	CreateAzureServiceBusIntegration(ctx context.Context, in *apiv4.CreateAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get Azure Service-Bus integration.
	GetAzureServiceBusIntegration(ctx context.Context, in *apiv4.GetAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetAzureServiceBusIntegrationResponse, error)
	// Update Azure Service-Bus integration.
	UpdateAzureServiceBusIntegration(ctx context.Context, in *apiv4.UpdateAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete Azure Service-Bus integration.
	DeleteAzureServiceBusIntegration(ctx context.Context, in *apiv4.DeleteAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create Pilot Things integration.
	CreatePilotThingsIntegration(ctx context.Context, in *apiv4.CreatePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get Pilot Things integration.
	GetPilotThingsIntegration(ctx context.Context, in *apiv4.GetPilotThingsIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetPilotThingsIntegrationResponse, error)
	// Update Pilot Things integration.
	UpdatePilotThingsIntegration(ctx context.Context, in *apiv4.UpdatePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete Pilot Things integration.
	DeletePilotThingsIntegration(ctx context.Context, in *apiv4.DeletePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Create IFTTT integration.
	CreateIftttIntegration(ctx context.Context, in *apiv4.CreateIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get IFTTT integration.
	GetIftttIntegration(ctx context.Context, in *apiv4.GetIftttIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetIftttIntegrationResponse, error)
	// Update IFTTT integration.
	UpdateIftttIntegration(ctx context.Context, in *apiv4.UpdateIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete IFTTT integration.
	DeleteIftttIntegration(ctx context.Context, in *apiv4.DeleteIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Generates application ID specific client-certificate.
	GenerateMqttIntegrationClientCertificate(ctx context.Context, in *apiv4.GenerateMqttIntegrationClientCertificateRequest, opts ...grpc.CallOption) (*apiv4.GenerateMqttIntegrationClientCertificateResponse, error)
	// List device-profiles used within the given application.
	ListDeviceProfiles(ctx context.Context, in *apiv4.ListApplicationDeviceProfilesRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationDeviceProfilesResponse, error)
	// List device tags used within the given application.
	ListDeviceTags(ctx context.Context, in *apiv4.ListApplicationDeviceTagsRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationDeviceTagsResponse, error)
}

type grpcApplicationAPI struct {
	c apiv4.ApplicationServiceClient
}

// assert compile-time
var _ Client = (*grpcApplicationAPI)(nil)

func NewWithClient(c apiv4.ApplicationServiceClient) Client { return &grpcApplicationAPI{c: c} }

func (g *grpcApplicationAPI) Get(ctx context.Context, in *apiv4.GetApplicationRequest, opts ...grpc.CallOption) (*apiv4.GetApplicationResponse, error) {
	return g.c.Get(ctx, in, opts...)
}

func (g *grpcApplicationAPI) Update(ctx context.Context, in *apiv4.UpdateApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Update(ctx, in, opts...)
}

func (g *grpcApplicationAPI) Delete(ctx context.Context, in *apiv4.DeleteApplicationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.Delete(ctx, in, opts...)
}

func (g *grpcApplicationAPI) List(ctx context.Context, in *apiv4.ListApplicationsRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationsResponse, error) {
	return g.c.List(ctx, in, opts...)
}

func (g *grpcApplicationAPI) ListIntegrations(ctx context.Context, in *apiv4.ListIntegrationsRequest, opts ...grpc.CallOption) (*apiv4.ListIntegrationsResponse, error) {
	return g.c.ListIntegrations(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateHttpIntegration(ctx context.Context, in *apiv4.CreateHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateHttpIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetHttpIntegration(ctx context.Context, in *apiv4.GetHttpIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetHttpIntegrationResponse, error) {
	return g.c.GetHttpIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateHttpIntegration(ctx context.Context, in *apiv4.UpdateHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateHttpIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteHttpIntegration(ctx context.Context, in *apiv4.DeleteHttpIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteHttpIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateInfluxDbIntegration(ctx context.Context, in *apiv4.CreateInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateInfluxDbIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetInfluxDbIntegration(ctx context.Context, in *apiv4.GetInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetInfluxDbIntegrationResponse, error) {
	return g.c.GetInfluxDbIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateInfluxDbIntegration(ctx context.Context, in *apiv4.UpdateInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateInfluxDbIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteInfluxDbIntegration(ctx context.Context, in *apiv4.DeleteInfluxDbIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteInfluxDbIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateThingsBoardIntegration(ctx context.Context, in *apiv4.CreateThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateThingsBoardIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetThingsBoardIntegration(ctx context.Context, in *apiv4.GetThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetThingsBoardIntegrationResponse, error) {
	return g.c.GetThingsBoardIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateThingsBoardIntegration(ctx context.Context, in *apiv4.UpdateThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateThingsBoardIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteThingsBoardIntegration(ctx context.Context, in *apiv4.DeleteThingsBoardIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteThingsBoardIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateMyDevicesIntegration(ctx context.Context, in *apiv4.CreateMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateMyDevicesIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetMyDevicesIntegration(ctx context.Context, in *apiv4.GetMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetMyDevicesIntegrationResponse, error) {
	return g.c.GetMyDevicesIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateMyDevicesIntegration(ctx context.Context, in *apiv4.UpdateMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateMyDevicesIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteMyDevicesIntegration(ctx context.Context, in *apiv4.DeleteMyDevicesIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteMyDevicesIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateGcpPubSubIntegration(ctx context.Context, in *apiv4.CreateGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateGcpPubSubIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetGcpPubSubIntegration(ctx context.Context, in *apiv4.GetGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetGcpPubSubIntegrationResponse, error) {
	return g.c.GetGcpPubSubIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateGcpPubSubIntegration(ctx context.Context, in *apiv4.UpdateGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateGcpPubSubIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteGcpPubSubIntegration(ctx context.Context, in *apiv4.DeleteGcpPubSubIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteGcpPubSubIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateAwsSnsIntegration(ctx context.Context, in *apiv4.CreateAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateAwsSnsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetAwsSnsIntegration(ctx context.Context, in *apiv4.GetAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetAwsSnsIntegrationResponse, error) {
	return g.c.GetAwsSnsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateAwsSnsIntegration(ctx context.Context, in *apiv4.UpdateAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateAwsSnsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteAwsSnsIntegration(ctx context.Context, in *apiv4.DeleteAwsSnsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteAwsSnsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateAzureServiceBusIntegration(ctx context.Context, in *apiv4.CreateAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateAzureServiceBusIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetAzureServiceBusIntegration(ctx context.Context, in *apiv4.GetAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetAzureServiceBusIntegrationResponse, error) {
	return g.c.GetAzureServiceBusIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateAzureServiceBusIntegration(ctx context.Context, in *apiv4.UpdateAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateAzureServiceBusIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteAzureServiceBusIntegration(ctx context.Context, in *apiv4.DeleteAzureServiceBusIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteAzureServiceBusIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreatePilotThingsIntegration(ctx context.Context, in *apiv4.CreatePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreatePilotThingsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetPilotThingsIntegration(ctx context.Context, in *apiv4.GetPilotThingsIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetPilotThingsIntegrationResponse, error) {
	return g.c.GetPilotThingsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdatePilotThingsIntegration(ctx context.Context, in *apiv4.UpdatePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdatePilotThingsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeletePilotThingsIntegration(ctx context.Context, in *apiv4.DeletePilotThingsIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeletePilotThingsIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) CreateIftttIntegration(ctx context.Context, in *apiv4.CreateIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.CreateIftttIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GetIftttIntegration(ctx context.Context, in *apiv4.GetIftttIntegrationRequest, opts ...grpc.CallOption) (*apiv4.GetIftttIntegrationResponse, error) {
	return g.c.GetIftttIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) UpdateIftttIntegration(ctx context.Context, in *apiv4.UpdateIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.UpdateIftttIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) DeleteIftttIntegration(ctx context.Context, in *apiv4.DeleteIftttIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return g.c.DeleteIftttIntegration(ctx, in, opts...)
}

func (g *grpcApplicationAPI) GenerateMqttIntegrationClientCertificate(ctx context.Context, in *apiv4.GenerateMqttIntegrationClientCertificateRequest, opts ...grpc.CallOption) (*apiv4.GenerateMqttIntegrationClientCertificateResponse, error) {
	return g.c.GenerateMqttIntegrationClientCertificate(ctx, in, opts...)
}

func (g *grpcApplicationAPI) ListDeviceProfiles(ctx context.Context, in *apiv4.ListApplicationDeviceProfilesRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationDeviceProfilesResponse, error) {
	return g.c.ListDeviceProfiles(ctx, in, opts...)
}

func (g *grpcApplicationAPI) ListDeviceTags(ctx context.Context, in *apiv4.ListApplicationDeviceTagsRequest, opts ...grpc.CallOption) (*apiv4.ListApplicationDeviceTagsResponse, error) {
	return g.c.ListDeviceTags(ctx, in, opts...)
}

func (g *grpcApplicationAPI) Create(ctx context.Context, in *apiv4.CreateApplicationRequest, opts ...grpc.CallOption) (*apiv4.CreateApplicationResponse, error) {
	return g.c.Create(ctx, in, opts...)
}
