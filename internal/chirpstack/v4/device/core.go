package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"google.golang.org/grpc"
)

type core struct {
	c apiv4.DeviceServiceClient
}

func newCore(conn *grpc.ClientConn) *core {
	return &core{c: apiv4.NewDeviceServiceClient(conn)}
}

func (s *core) List(ctx context.Context, p device.ListParams) (shared.ListResult[device.ListItem], error) {
	req := &apiv4.ListDevicesRequest{
		Limit:         p.GetSafeLimit(),
		Offset:        p.Offset,
		Search:        p.Search,
		ApplicationId: p.ApplicationID,
		//OrderBy:         p.OrderBy,
		OrderByDesc:     p.OrderByDesc,
		DeviceProfileId: p.DeviceProfileID,
		Tags:            p.Tags,
	}
	out := shared.NewListResult[device.ListItem]()

	resp, err := s.c.List(ctx, req)
	if err != nil {
		return out, err
	}

	out.TotalCount = resp.GetTotalCount()
	items := resp.GetResult()
	out = shared.NewListResultCap[device.ListItem](len(items))

	out.TotalCount = resp.GetTotalCount()
	for _, it := range items {
		out.Add(device.ListItem{
			DevEUI:            it.DevEui,
			CreatedAt:         shared.TsProtobufToPtrTime(it.CreatedAt),
			UpdatedAt:         shared.TsProtobufToPtrTime(it.UpdatedAt),
			LastSeenAt:        shared.TsProtobufToPtrTime(it.LastSeenAt),
			Name:              it.Name,
			Description:       it.Description,
			DeviceProfileID:   it.DeviceProfileId,
			DeviceProfileName: it.DeviceProfileName,
			DeviceStatus:      nil,
			Tags:              it.Tags,
		})
	}
	return out, nil
}

func (s *core) Get(ctx context.Context, p device.GetParams) (*device.Device, error) {
	resp, err := s.c.Get(ctx, &apiv4.GetDeviceRequest{DevEui: p.DevEUI})
	if err != nil {
		return nil, err
	}
	d := resp.GetDevice()
	return &device.Device{
		DevEUI:          d.GetDevEui(),
		Name:            d.GetName(),
		ApplicationID:   d.GetApplicationId(),
		DeviceProfileID: d.GetDeviceProfileId(),
	}, nil
}

func (s *core) Create(ctx context.Context, p device.CreateParams) (string, error) {
	req := &apiv4.CreateDeviceRequest{
		Device: &apiv4.Device{DevEui: p.Device.DevEUI, ApplicationId: p.Device.ApplicationID},
	}
	/*if p.Device.Name != nil {
		req.Device.Name = *p.Device.Name
	}
	if p.DeviceProfileID != nil {
		req.Device.DeviceProfileID = *p.DeviceProfileID
	}*/
	_, err := s.c.Create(ctx, req)
	return req.Device.DevEui, err
}

func (s *core) Update(ctx context.Context, p device.UpdateParams) error {
	device, err := s.Get(ctx, device.GetParams{DevEUI: p.DevEUI})
	if err != nil {
		return err
	}
	request := apiv4.UpdateDeviceRequest{
		Device: &apiv4.Device{
			DevEui:          device.DevEUI,
			Name:            device.Name,
			Description:     device.Description,
			ApplicationId:   device.ApplicationID,
			DeviceProfileId: device.DeviceProfileID,
			SkipFcntCheck:   device.SkipFcntCheck,
			Variables:       device.Variables,
			Tags:            device.Tags,
			JoinEui:         device.JoinEui,
		},
	}
	if p.Patch.Name != nil {
		request.Device.Name = *p.Patch.Name
	}
	if p.Patch.Description != nil {
		request.Device.Description = *p.Patch.Description
	}
	if p.Patch.ApplicationID != nil {
		request.Device.ApplicationId = *p.Patch.ApplicationID
	}
	if p.Patch.DeviceProfileID != nil {
		request.Device.DeviceProfileId = *p.Patch.DeviceProfileID
	}
	if p.Patch.SkipFcntCheck != nil {
		request.Device.SkipFcntCheck = *p.Patch.SkipFcntCheck
	}
	if p.Patch.Variables != nil {
		request.Device.Variables = shared.UpdateMap(p.Patch.Variables, device.Variables)
	}
	if p.Patch.Tags != nil {
		request.Device.Tags = shared.UpdateMap(p.Patch.Tags, device.Tags)
	}
	_, err = s.c.Update(ctx, &request)
	return err
}

func (s *core) Delete(ctx context.Context, p device.DeleteParams) error {
	//TODO implement me
	panic("implement me")
}
