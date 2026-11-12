package deviceprofile

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateDeviceProfileParams) (CreateDeviceProfileResult, error)
	Get(ctx context.Context, in GetDeviceProfileParams) (*DeviceProfile, error)
	Update(ctx context.Context, in UpdateDeviceProfileParams) error
	Delete(ctx context.Context, in DeleteDeviceProfileParams) error
	List(ctx context.Context, in ListParams) (shared.ListResult[ListItem], error)
	ListAdrAlgorithms(ctx context.Context) (ListDeviceProfileAdrAlgorithmsResult, error)
}

type CreateDeviceProfileParams struct {
	DeviceProfile
}

type CreateDeviceProfileResult struct {
	ID string `json:"id" yaml:"id"`
}

type GetDeviceProfileParams struct {
	ID string
}

type UpdateDeviceProfileParams struct{}

type DeleteDeviceProfileParams struct{}

type ListParams struct {
	shared.ListPaginationRequest
	shared.ListOrderByRequest
	shared.ListTagsFilterRequest
	TenantID   string
	DeviceID   string
	GlobalOnly bool
	TenantOnly bool
}

type ListDeviceProfileAdrAlgorithmsResult struct {
}

type DeviceProfile struct {
	ID                               string                        `json:"id" yaml:"id"`
	CreatedAt                        *time.Time                    `json:"createdAt" yaml:"createdAt"`
	UpdatedAt                        *time.Time                    `json:"updatedAt" yaml:"updatedAt"`
	TenantID                         string                        `json:"tenantID" yaml:"tenantID"`
	Name                             string                        `json:"name" yaml:"name"`
	Description                      string                        `json:"description" yaml:"description"`
	Region                           shared.Region                 `json:"region" yaml:"region"`
	MacVersion                       shared.MacVersion             `json:"macVersion" yaml:"macVersion"`
	RegParamsRevision                shared.RegParamsRevision      `json:"regParamsRevision" yaml:"regParamsRevision"`
	AdrAlgorithmID                   string                        `json:"adrAlgorithmID" yaml:"adrAlgorithmID"`
	PayloadCodecRuntime              shared.CodecRuntime           `json:"payloadCodecRuntime" yaml:"payloadCodecRuntime"`
	PayloadCodecScript               string                        `json:"payloadCodecScript" yaml:"payloadCodecScript"`
	FlushQueueOnActivate             bool                          `json:"flushQueueOnActivate" yaml:"flushQueueOnActivate"`
	UplinkInterval                   uint32                        `json:"uplinkInterval" yaml:"uplinkInterval"`
	DeviceStatusReqInterval          uint32                        `json:"deviceStatusReqInterval" yaml:"deviceStatusReqInterval"`
	SupportsOtaa                     bool                          `json:"supportsOtaa" yaml:"supportsOtaa"`
	SupportsClassB                   bool                          `json:"supportsClassB" yaml:"supportsClassB"`
	SupportsClassC                   bool                          `json:"supportsClassC" yaml:"supportsClassC"`
	ClassBTimeout                    uint32                        `json:"classBTimeout" yaml:"classBTimeout"`
	ClassBPingSlotPeriodicity        uint32                        `json:"classBPingSlotPeriodicity" yaml:"classBPingSlotPeriodicity"`
	ClassBPingSlotDr                 uint32                        `json:"classBPingSlotDr" yaml:"classBPingSlotDr"`
	ClassBPingSlotFreq               uint32                        `json:"classBPingSlotFreq" yaml:"classBPingSlotFreq"`
	ClassCTimeout                    uint32                        `json:"classCTimeout" yaml:"classCTimeout"`
	AbpRx1Delay                      uint32                        `json:"abpRx1Delay" yaml:"abpRx1Delay"`
	AbpRx1DrOffset                   uint32                        `json:"abpRx1DrOffset" yaml:"abpRx1DrOffset"`
	AbpRx2Dr                         uint32                        `json:"abpRx2Dr" yaml:"abpRx2Dr"`
	AbpRx2Freq                       uint32                        `json:"abpRx2Freq" yaml:"abpRx2Freq"`
	Tags                             map[string]string             `json:"tags" yaml:"tags"`
	Measurements                     map[string]shared.Measurement `json:"measurements" yaml:"measurements"`
	AutoDetectMeasurements           bool                          `json:"autoDetectMeasurements" yaml:"autoDetectMeasurements"`
	RegionConfigID                   string                        `json:"regionConfigID" yaml:"regionConfigID"`
	IsRelay                          bool                          `json:"isRelay" yaml:"isRelay"`
	IsRelayEd                        bool                          `json:"isRelayEd" yaml:"isRelayEd"`
	RelayEdRelayOnly                 bool                          `json:"relayEdRelayOnly" yaml:"relayEdRelayOnly"`
	RelayEnabled                     bool                          `json:"relayEnabled" yaml:"relayEnabled"`
	RelayCadPeriodicity              shared.CadPeriodicity         `json:"relayCadPeriodicity" yaml:"relayCadPeriodicity"`
	RelayDefaultChannelIndex         uint32                        `json:"relayDefaultChannelIndex" yaml:"relayDefaultChannelIndex"`
	RelaySecondChannelFreq           uint32                        `json:"relaySecondChannelFreq" yaml:"relaySecondChannelFreq"`
	RelaySecondChannelDr             uint32                        `json:"relaySecondChannelDr" yaml:"relaySecondChannelDr"`
	RelaySecondChannelAckOffset      shared.SecondChAckOffset      `json:"relaySecondChannelAckOffset" yaml:"relaySecondChannelAckOffset"`
	RelayEdActivationMode            shared.RelayModeActivation    `json:"relayEdActivationMode" yaml:"relayEdActivationMode"`
	RelayEdSmartEnableLevel          uint32                        `json:"relayEdSmartEnableLevel" yaml:"relayEdSmartEnableLevel"`
	RelayEdBackOff                   uint32                        `json:"relayEdBackOff" yaml:"relayEdBackOff"`
	RelayEdUplinkLimitBucketSize     uint32                        `json:"relayEdUplinkBucketSize" yaml:"relayEdUplinkBucketSize"`
	RelayEdUplinkLimitReloadRate     uint32                        `json:"relayEdUplinkReloadRate" yaml:"relayEdUplinkReloadRate"`
	RelayJoinReqLimitReloadRate      uint32                        `json:"relayJoinReqLimitReloadRate" yaml:"relayJoinReqLimitReloadRate"`
	RelayNotifyLimitReloadRate       uint32                        `json:"relayNotifyLimitReloadRate" yaml:"relayNotifyLimitReloadRate"`
	RelayGlobalUplinkLimitReloadRate uint32                        `json:"relayGlobalUplinkLimitReloadRate" yaml:"relayGlobalUplinkLimitReloadRate"`
	RelayOverallLimitReloadRate      uint32                        `json:"relayOverallLimitReloadRate" yaml:"relayOverallLimitReloadRate"`
	RelayJoinReqLimitBucketSize      uint32                        `json:"relayJoinReqLimitBucketSize" yaml:"relayJoinReqLimitBucketSize"`
	RelayNotifyLimitBucketSize       uint32                        `json:"relayNotifyLimitBucketSize" yaml:"relayNotifyLimitBucketSize"`
	RelayGlobalUplinkLimitBucketSize uint32                        `json:"relayGlobalUplinkLimitBucketSize" yaml:"relayGlobalUplinkLimitBucketSize"`
	RelayOverallLimitBucketSize      uint32                        `json:"relayOverallLimitBucketSize" yaml:"relayOverallLimitBucketSize"`
	AllowRoaming                     bool                          `json:"allowRoaming" yaml:"allowRoaming"`
	Rx1Delay                         uint32                        `json:"rx1Delay" yaml:"rx1Delay"`
	AppLayerParams                   *shared.AppLayerParams        `json:"appLayerParams" yaml:"appLayerParams"`
	DeviceID                         string                        `json:"deviceID" yaml:"deviceID"`
	FirmwareVersion                  string                        `json:"firmwareVersion" yaml:"firmwareVersion"`
	SupportedUplinkDataRates         []uint32                      `json:"supportedUplinkDataRates" yaml:"supportedUplinkDataRates"`
	ClassBDownlinkOnly               bool                          `json:"classBDownlinkOnly" yaml:"classBDownlinkOnly"`
}

type ListItem struct {
	ID                string                   `json:"id" yaml:"id"`
	CreatedAt         *time.Time               `json:"createdAt" yaml:"createdAt"`
	UpdatedAt         *time.Time               `json:"updatedAt" yaml:"updatedAt"`
	Name              string                   `json:"name" yaml:"name"`
	Region            shared.Region            `json:"region" yaml:"region"`
	MacVersion        shared.MacVersion        `json:"macVersion" yaml:"macVersion"`
	RegParamsRevision shared.RegParamsRevision `json:"regParamsRevision" yaml:"regParamsRevision"`
	SupportsOtaa      bool                     `json:"supportsOtaa" yaml:"supportsOtaa"`
	SupportsClassB    bool                     `json:"supportsClassB" yaml:"supportsClassB"`
	SupportsClassC    bool                     `json:"supportsClassC" yaml:"supportsClassC"`
	VendorID          string                   `json:"vendorID" yaml:"vendorID"`
	VendorName        string                   `json:"vendorName" yaml:"vendorName"`
	DeviceID          string                   `json:"deviceID" yaml:"deviceID"`
	DeviceName        string                   `json:"deviceName" yaml:"deviceName"`
	FirmwareVersion   string                   `json:"firmwareVersion" yaml:"firmwareVersion"`
}
