package deviceprofile

import (
	"context"
	"fmt"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type core struct {
	client Client
}

var _ deviceprofile.Core = (*core)(nil)

func newCore(client Client) *core {
	return &core{client: client}
}

func (c *core) Create(ctx context.Context, in deviceprofile.CreateDeviceProfileParams) (deviceprofile.CreateDeviceProfileResult, error) {
	req := &apiv4.CreateDeviceProfileRequest{
		DeviceProfile: &apiv4.DeviceProfile{
			Id:                               "",
			TenantId:                         "",
			Name:                             "",
			Description:                      "",
			Region:                           0,
			MacVersion:                       0,
			RegParamsRevision:                0,
			AdrAlgorithmId:                   "",
			PayloadCodecRuntime:              0,
			PayloadCodecScript:               "",
			FlushQueueOnActivate:             false,
			UplinkInterval:                   0,
			DeviceStatusReqInterval:          0,
			SupportsOtaa:                     false,
			SupportsClassB:                   false,
			SupportsClassC:                   false,
			ClassBTimeout:                    0,
			ClassBPingSlotPeriodicity:        0,
			ClassBPingSlotDr:                 0,
			ClassBPingSlotFreq:               0,
			ClassCTimeout:                    0,
			AbpRx1Delay:                      0,
			AbpRx1DrOffset:                   0,
			AbpRx2Dr:                         0,
			AbpRx2Freq:                       0,
			Tags:                             nil,
			Measurements:                     nil,
			AutoDetectMeasurements:           false,
			RegionConfigId:                   "",
			IsRelay:                          false,
			IsRelayEd:                        false,
			RelayEdRelayOnly:                 false,
			RelayEnabled:                     false,
			RelayCadPeriodicity:              0,
			RelayDefaultChannelIndex:         0,
			RelaySecondChannelFreq:           0,
			RelaySecondChannelDr:             0,
			RelaySecondChannelAckOffset:      0,
			RelayEdActivationMode:            0,
			RelayEdSmartEnableLevel:          0,
			RelayEdBackOff:                   0,
			RelayEdUplinkLimitBucketSize:     0,
			RelayEdUplinkLimitReloadRate:     0,
			RelayJoinReqLimitReloadRate:      0,
			RelayNotifyLimitReloadRate:       0,
			RelayGlobalUplinkLimitReloadRate: 0,
			RelayOverallLimitReloadRate:      0,
			RelayJoinReqLimitBucketSize:      0,
			RelayNotifyLimitBucketSize:       0,
			RelayGlobalUplinkLimitBucketSize: 0,
			RelayOverallLimitBucketSize:      0,
			AllowRoaming:                     false,
			Rx1Delay:                         0,
			AppLayerParams:                   nil,
			DeviceId:                         "",
			FirmwareVersion:                  "",
			SupportedUplinkDataRates:         nil,
			ClassBDownlinkOnly:               false,
		},
	}
	create, err := c.client.Create(ctx, req)
	if err != nil {
		return deviceprofile.CreateDeviceProfileResult{}, err
	}
	return deviceprofile.CreateDeviceProfileResult{
		ID: create.Id,
	}, nil
}

func (c *core) Get(ctx context.Context, in deviceprofile.GetParams) (*deviceprofile.DeviceProfile, error) {
	req := &apiv4.GetDeviceProfileRequest{
		Id: in.ID,
	}
	get, err := c.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return toDeviceProfile(get.DeviceProfile, get.CreatedAt, get.UpdatedAt)
}

func (c *core) Update(ctx context.Context, in deviceprofile.UpdateDeviceProfileParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) Delete(ctx context.Context, in deviceprofile.DeleteDeviceProfileParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) List(ctx context.Context, in deviceprofile.ListParams) (shared.ListResult[deviceprofile.ListItem], error) {
	req := &apiv4.ListDeviceProfilesRequest{
		Limit:      in.GetSafeLimit(),
		Offset:     in.Offset,
		Search:     in.Search,
		TenantId:   in.TenantID,
		DeviceId:   in.DeviceID,
		GlobalOnly: in.GlobalOnly,
		TenantOnly: in.TenantOnly,
	}
	out := shared.NewListResult[deviceprofile.ListItem]()
	resp, err := c.client.List(ctx, req)
	if err != nil {
		return out, err
	}
	items := resp.GetResult()
	out = shared.NewListResultCap[deviceprofile.ListItem](len(items))
	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		converted, err := toDeviceProfileListItem(item)
		if err != nil {
			return out, fmt.Errorf("convert device profile list item: %w", err)
		}
		out.Add(converted)
	}
	return out, nil
}

func (c *core) ListAdrAlgorithms(ctx context.Context) (deviceprofile.ListDeviceProfileAdrAlgorithmsResult, error) {
	//TODO implement me
	panic("implement me")
}
