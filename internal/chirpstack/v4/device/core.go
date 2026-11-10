package device

import (
	"context"
	"fmt"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type core struct {
	client Client
}

func newCore(client Client) *core {
	return &core{client: client}
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

	resp, err := s.client.List(ctx, req)
	if err != nil {
		return out, err
	}

	out.TotalCount = resp.GetTotalCount()
	items := resp.GetResult()
	out = shared.NewListResultCap[device.ListItem](len(items))

	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		converted, err := toDeviceListItem(item)
		if err != nil {
			return out, fmt.Errorf("convert gateway list item: %w", err)
		}
		out.Add(converted)
	}
	return out, nil
}

func (s *core) Get(ctx context.Context, p device.GetParams) (*device.Device, error) {
	resp, err := s.client.Get(ctx, &apiv4.GetDeviceRequest{DevEui: p.DevEUI})
	if err != nil {
		return nil, err
	}
	return toDevice(resp)
}

func (s *core) Create(ctx context.Context, p device.CreateParams) (string, error) {
	req := &apiv4.CreateDeviceRequest{
		Device: &apiv4.Device{
			DevEui:          p.Device.DevEUI,
			JoinEui:         p.Device.JoinEUI,
			Name:            p.Device.Name,
			Description:     p.Device.Description,
			ApplicationId:   p.Device.ApplicationID,
			DeviceProfileId: p.Device.DeviceProfileID,
			SkipFcntCheck:   p.Device.SkipFcntCheck,
			IsDisabled:      p.Device.IsDisabled,
			Variables:       p.Device.Variables,
			Tags:            p.Device.Tags,
		},
	}
	_, err := s.client.Create(ctx, req)
	return req.Device.DevEui, err
}

func (s *core) Update(ctx context.Context, p device.UpdateParams) error {
	d, err := s.Get(ctx, device.GetParams{DevEUI: p.DevEUI})
	if err != nil {
		return err
	}
	request := apiv4.UpdateDeviceRequest{
		Device: &apiv4.Device{
			DevEui:          d.DevEUI,
			Name:            d.Name,
			Description:     d.Description,
			ApplicationId:   d.ApplicationID,
			DeviceProfileId: d.DeviceProfileID,
			SkipFcntCheck:   d.SkipFcntCheck,
			Variables:       d.Variables,
			Tags:            d.Tags,
			JoinEui:         d.JoinEUI,
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
		request.Device.Variables = shared.UpdateMap(p.Patch.Variables, d.Variables)
	}
	if p.Patch.Tags != nil {
		request.Device.Tags = shared.UpdateMap(p.Patch.Tags, d.Tags)
	}
	_, err = s.client.Update(ctx, &request)
	return err
}

func (s *core) Delete(ctx context.Context, p device.DeleteParams) error {
	//TODO implement me
	panic("implement me")
}
