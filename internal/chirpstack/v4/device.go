package v4

import (
	"context"
	"time"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/samber/lo"
	svcdevice "github.com/sebundefined/chirpctl/internal/service/device"
	"google.golang.org/grpc"
)

type deviceSvc struct{ c apiv4.DeviceServiceClient }

func NewDeviceService(conn *grpc.ClientConn) svcdevice.Service {
	return &deviceSvc{c: apiv4.NewDeviceServiceClient(conn)}
}

func (s *deviceSvc) List(ctx context.Context, p svcdevice.ListParams) (svcdevice.ListResult, error) {
	list, err := s.c.List(ctx, &apiv4.ListDevicesRequest{
		Limit:            10,
		Offset:           0,
		Search:           "",
		ApplicationId:    p.ApplicationID,
		MulticastGroupId: "",
		OrderBy:          1,
		OrderByDesc:      false,
		Tags:             nil,
		DeviceProfileId:  "",
	})
	if err != nil {
		return svcdevice.ListResult{}, err
	}
	res := svcdevice.ListResult{
		TotalCount: list.GetTotalCount(),
	}
	res.Result = lo.Map(list.Result, func(item *apiv4.DeviceListItem, _ int) *svcdevice.ListItem {
		return &svcdevice.ListItem{
			DevEUI: item.DevEui,
			CreatedAt: lo.TernaryF(item.CreatedAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.CreatedAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			UpdatedAt: lo.TernaryF(item.UpdatedAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.UpdatedAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			LastSeenAt: lo.TernaryF(item.LastSeenAt != nil, func() *time.Time {
				return lo.ToPtr(time.Unix(item.LastSeenAt.Seconds, 0))
			}, func() *time.Time {
				return nil
			}),
			Name:              item.Name,
			Description:       item.Description,
			DeviceProfileID:   item.DeviceProfileId,
			DeviceProfileName: item.DeviceProfileName,
			DeviceStatus:      nil,
			Tags:              item.Tags,
		}
	})
	return res, nil
}

func (s *deviceSvc) Get(ctx context.Context, p svcdevice.GetParams) (*svcdevice.Device, error) {
	resp, err := s.c.Get(ctx, &apiv4.GetDeviceRequest{DevEui: p.DevEUI})
	if err != nil {
		return nil, err
	}
	d := resp.GetDevice()
	return &svcdevice.Device{
		DevEUI:          d.GetDevEui(),
		Name:            d.GetName(),
		ApplicationID:   d.GetApplicationId(),
		DeviceProfileID: d.GetDeviceProfileId(),
	}, nil
}

func (s *deviceSvc) Create(ctx context.Context, p svcdevice.CreateParams) error {
	req := &apiv4.CreateDeviceRequest{
		Device: &apiv4.Device{DevEui: p.Device.DevEUI, ApplicationId: p.Device.ApplicationID},
	}
	/*if p.Device.Name != nil {
		req.Device.Name = *p.Device.Name
	}
	if p.DeviceProfileID != nil {
		req.Device.DeviceProfileId = *p.DeviceProfileID
	}*/
	_, err := s.c.Create(ctx, req)
	return err
}
