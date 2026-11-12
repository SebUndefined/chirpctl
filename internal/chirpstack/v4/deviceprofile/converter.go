package deviceprofile

import (
	"errors"
	"maps"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toDeviceProfileListItem(src *api.DeviceProfileListItem) (deviceprofile.ListItem, error) {
	if src == nil {
		return deviceprofile.ListItem{}, errors.New("device profile list item is nil")
	}
	return deviceprofile.ListItem{
		ID:                src.Id,
		CreatedAt:         shared.TsProtobufToPtrTime(src.GetCreatedAt()),
		UpdatedAt:         shared.TsProtobufToPtrTime(src.GetUpdatedAt()),
		Name:              src.GetName(),
		Region:            shared.ParseRegionSourceFromInt32(int32(src.GetRegion())),
		MacVersion:        shared.ParseMacVersionSourceFromInt32(int32(src.GetMacVersion())),
		RegParamsRevision: shared.ParseRegParamsRevisionSourceFromInt32(int32(src.GetRegParamsRevision())),
		SupportsOtaa:      src.GetSupportsOtaa(),
		SupportsClassB:    src.GetSupportsClassB(),
		SupportsClassC:    src.GetSupportsClassC(),
		VendorID:          src.GetVendorId(),
		VendorName:        src.GetVendorName(),
		DeviceID:          src.GetDeviceId(),
		DeviceName:        src.GetDeviceName(),
		FirmwareVersion:   src.GetFirmwareVersion(),
	}, nil
}

func toDeviceProfile(src *api.DeviceProfile, createdAt, updatedAt *timestamppb.Timestamp) (*deviceprofile.DeviceProfile, error) {
	if src == nil {
		return nil, errors.New("device profile response is nil")
	}
	res := &deviceprofile.DeviceProfile{
		ID:                               src.GetId(),
		CreatedAt:                        shared.TsProtobufToPtrTime(createdAt),
		UpdatedAt:                        shared.TsProtobufToPtrTime(updatedAt),
		TenantID:                         src.GetTenantId(),
		Name:                             src.GetName(),
		Description:                      src.GetDescription(),
		Region:                           shared.ParseRegionSourceFromInt32(int32(src.GetRegion())),
		MacVersion:                       shared.ParseMacVersionSourceFromInt32(int32(src.GetMacVersion())),
		RegParamsRevision:                shared.ParseRegParamsRevisionSourceFromInt32(int32(src.GetRegParamsRevision())),
		AdrAlgorithmID:                   src.GetAdrAlgorithmId(),
		PayloadCodecRuntime:              shared.ParseCodecRuntimeSourceFromInt32(int32(src.GetPayloadCodecRuntime())),
		PayloadCodecScript:               src.GetPayloadCodecScript(),
		FlushQueueOnActivate:             src.GetFlushQueueOnActivate(),
		UplinkInterval:                   src.GetUplinkInterval(),
		DeviceStatusReqInterval:          src.GetDeviceStatusReqInterval(),
		SupportsOtaa:                     src.GetSupportsOtaa(),
		SupportsClassB:                   src.GetSupportsClassB(),
		SupportsClassC:                   src.GetSupportsClassC(),
		ClassBTimeout:                    src.GetClassBTimeout(),
		ClassBPingSlotPeriodicity:        src.GetClassBPingSlotPeriodicity(),
		ClassBPingSlotDr:                 src.GetClassBPingSlotDr(),
		ClassBPingSlotFreq:               src.GetClassBPingSlotFreq(),
		ClassCTimeout:                    src.GetClassCTimeout(),
		AbpRx1Delay:                      src.GetAbpRx1Delay(),
		AbpRx1DrOffset:                   src.GetAbpRx1DrOffset(),
		AbpRx2Dr:                         src.GetAbpRx2Dr(),
		AbpRx2Freq:                       src.GetAbpRx2Freq(),
		Tags:                             maps.Clone(src.GetTags()),
		Measurements:                     nil,
		AutoDetectMeasurements:           src.GetAutoDetectMeasurements(),
		RegionConfigID:                   src.GetRegionConfigId(),
		IsRelay:                          src.GetIsRelay(),
		IsRelayEd:                        src.GetIsRelayEd(),
		RelayEdRelayOnly:                 src.GetRelayEdRelayOnly(),
		RelayEnabled:                     src.GetRelayEnabled(),
		RelayCadPeriodicity:              shared.ParseCadPeriodicitySourceFromInt32(int32(src.GetRelayCadPeriodicity())),
		RelayDefaultChannelIndex:         src.GetRelayDefaultChannelIndex(),
		RelaySecondChannelFreq:           src.GetRelaySecondChannelFreq(),
		RelaySecondChannelDr:             src.GetRelaySecondChannelDr(),
		RelaySecondChannelAckOffset:      shared.ParseSecondChAckOffsetSourceFromInt32(int32(src.GetRelaySecondChannelAckOffset())),
		RelayEdActivationMode:            shared.ParseRelayModeActivationSourceFromInt32(int32(src.GetRelayEdActivationMode())),
		RelayEdSmartEnableLevel:          src.GetRelayEdSmartEnableLevel(),
		RelayEdBackOff:                   src.GetRelayEdBackOff(),
		RelayEdUplinkLimitBucketSize:     src.GetRelayEdUplinkLimitBucketSize(),
		RelayEdUplinkLimitReloadRate:     src.GetRelayEdUplinkLimitReloadRate(),
		RelayJoinReqLimitReloadRate:      src.GetRelayJoinReqLimitReloadRate(),
		RelayNotifyLimitReloadRate:       src.GetRelayNotifyLimitReloadRate(),
		RelayGlobalUplinkLimitReloadRate: src.GetRelayGlobalUplinkLimitReloadRate(),
		RelayOverallLimitReloadRate:      src.GetRelayOverallLimitReloadRate(),
		RelayJoinReqLimitBucketSize:      src.GetRelayJoinReqLimitBucketSize(),
		RelayNotifyLimitBucketSize:       src.GetRelayNotifyLimitBucketSize(),
		RelayGlobalUplinkLimitBucketSize: src.GetRelayGlobalUplinkLimitBucketSize(),
		RelayOverallLimitBucketSize:      src.GetRelayOverallLimitBucketSize(),
		AllowRoaming:                     src.GetAllowRoaming(),
		Rx1Delay:                         src.GetRx1Delay(),
		AppLayerParams:                   nil,
		DeviceID:                         src.GetDeviceId(),
		FirmwareVersion:                  src.GetFirmwareVersion(),
		SupportedUplinkDataRates:         src.GetSupportedUplinkDataRates(),
		ClassBDownlinkOnly:               src.GetClassBDownlinkOnly(),
	}
	if src.Measurements != nil {
		for key, measurement := range src.Measurements {
			res.Measurements[key] = shared.Measurement{
				Name: measurement.Name,
				Kind: shared.ParseMeasurementKindSourceFromInt32(int32(measurement.Kind)),
			}
		}
	}
	if src.AppLayerParams != nil {
		res.AppLayerParams = &shared.AppLayerParams{
			Ts003Version: shared.ParseTs003VersionFromInt32(int32(src.AppLayerParams.Ts003Version)),
			Ts003FPort:   src.AppLayerParams.Ts003FPort,
			Ts004Version: shared.ParseTs004VersionFromInt32(int32(src.AppLayerParams.Ts004Version)),
			Ts004FPort:   src.AppLayerParams.Ts004FPort,
			Ts005Version: shared.ParseTs005VersionFromInt32(int32(src.AppLayerParams.Ts005Version)),
			Ts005FPort:   src.AppLayerParams.Ts005FPort,
		}
	}
	return res, nil
}
