package application

import (
	"context"
	"fmt"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/application"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type core struct {
	client Client
}

func newCore(client Client) *core {
	return &core{client: client}
}

func (s *core) List(ctx context.Context, p application.ListParams) (shared.ListResult[application.ListItem], error) {
	req := &apiv4.ListApplicationsRequest{
		Limit:    p.GetSafeLimit(),
		Offset:   p.Offset,
		Search:   p.Search,
		TenantId: p.TenantID,
	}
	out := shared.NewListResult[application.ListItem]()

	resp, err := s.client.List(ctx, req)
	if err != nil {
		return out, err
	}

	out.TotalCount = resp.GetTotalCount()
	items := resp.GetResult()
	out = shared.NewListResultCap[application.ListItem](len(items))

	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		converted, err := toApplicationListItem(item)
		if err != nil {
			return out, fmt.Errorf("convert application list item: %w", err)
		}
		out.Add(converted)
	}
	return out, nil
}

func (s *core) Create(ctx context.Context, p application.CreateParams) (application.CreateResult, error) {
	req := &apiv4.CreateApplicationRequest{
		Application: &apiv4.Application{
			Name:        p.Name,
			Description: p.Description,
			TenantId:    p.TenantID,
			Tags:        p.Tags,
		},
	}
	res, err := s.client.Create(ctx, req)
	if err != nil {
		return application.CreateResult{}, err
	}
	return application.CreateResult{
		ID: res.GetId(),
	}, err
}

func (s *core) Get(ctx context.Context, p application.GetParams) (*application.Application, error) {
	resp, err := s.client.Get(ctx, &apiv4.GetApplicationRequest{Id: p.ID})
	if err != nil {
		return nil, err
	}
	return toApplication(resp)
}

func (s *core) Update(ctx context.Context, p application.UpdateParams) error {
	d, err := s.Get(ctx, application.GetParams{DevEUI: p.DevEUI})
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

func (s *core) Delete(ctx context.Context, p application.DeleteParams) error {
	//TODO implement me
	panic("implement me")
}
