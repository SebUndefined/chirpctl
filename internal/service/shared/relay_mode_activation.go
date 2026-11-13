package shared

import "strings"

type RelayModeActivation string

const (
	RelayModeActivationUnknown             RelayModeActivation = "UNKNOWN"
	RelayModeActivationDisableMode         RelayModeActivation = "DISABLE_MODE"
	RelayModeActivationEnableMode          RelayModeActivation = "ENABLE_MODE"
	RelayModeActivationDynamic             RelayModeActivation = "DYNAMIC"
	RelayModeActivationEndDeviceControlled RelayModeActivation = "END_DEVICE_CONTROLLED"
)

func (c RelayModeActivation) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return RelayModeActivationUnknown.String()
	}
	return string(c)
}

func ParseRelayModeActivationSourceFromInt32(r int32) RelayModeActivation {
	switch r {
	case 0:
		return RelayModeActivationDisableMode
	case 1:
		return RelayModeActivationEnableMode
	case 2:
		return RelayModeActivationDynamic
	case 3:
		return RelayModeActivationEndDeviceControlled
	default:
		return RelayModeActivationUnknown
	}
}
