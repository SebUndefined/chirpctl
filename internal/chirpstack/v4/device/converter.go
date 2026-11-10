package device

import (
	"errors"
	"maps"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/service/shared"
)

// toDeviceListItem converts a ChirpStack v4 list item into the internal model.
func toDeviceListItem(src *api.DeviceListItem) (device.ListItem, error) {
	if src == nil {
		return device.ListItem{}, errors.New("gateway list item is nil")
	}

	return device.ListItem{
		DevEUI:            src.GetDevEui(),
		CreatedAt:         shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:         shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		LastSeenAt:        shared.TsProtobufToPtrTime(src.GetLastSeenAt()),
		Name:              src.GetName(),
		Description:       src.GetDescription(),
		DeviceProfileID:   src.GetDeviceProfileId(),
		DeviceProfileName: src.GetDeviceProfileName(),
		DeviceStatus:      toDeviceStatus(src.GetDeviceStatus()),
		Tags:              maps.Clone(src.Tags),
	}, nil
}

// toDevice converts a ChirpStack v4 response into the internal Device model.
func toDevice(resp *api.GetDeviceResponse) (*device.Device, error) {
	if resp == nil || resp.GetDevice() == nil {
		return nil, errors.New("device response is empty")
	}
	src := resp.GetDevice()
	return &device.Device{
		DevEUI:          src.GetDevEui(),
		CreatedAt:       shared.TsProtobufToPtrTime(resp.CreatedAt),
		UpdatedAt:       shared.TsProtobufToPtrTime(resp.UpdatedAt),
		LastSeenAt:      shared.TsProtobufToPtrTime(resp.LastSeenAt),
		Name:            src.GetName(),
		Description:     src.GetDescription(),
		ApplicationID:   src.GetApplicationId(),
		DeviceProfileID: src.GetDeviceProfileId(),
		SkipFcntCheck:   src.GetSkipFcntCheck(),
		IsDisabled:      src.GetIsDisabled(),
		Variables:       maps.Clone(src.GetVariables()),
		Tags:            maps.Clone(src.GetTags()),
		JoinEUI:         src.GetJoinEui(),
	}, nil
}

func toDeviceStatus(src *api.DeviceStatus) *device.Status {
	if src == nil {
		return nil
	}
	return &device.Status{
		Margin:              src.Margin,
		ExternalPowerSource: src.ExternalPowerSource,
		BatteryLevel:        src.BatteryLevel,
	}
}
