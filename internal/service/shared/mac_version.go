package shared

import "strings"

type MacVersion string

const (
	MacVersionUnknown    MacVersion = "UNKNOWN"
	MacVersionLorawan100 MacVersion = "LORAWAN_1.0.0"
	MacVersionLorawan101 MacVersion = "LORAWAN_1.0.1"
	MacVersionLorawan102 MacVersion = "LORAWAN_1.0.2"
	MacVersionLorawan103 MacVersion = "LORAWAN_1.0.3"
	MacVersionLorawan104 MacVersion = "LORAWAN_1.0.4"
	MacVersionLorawan110 MacVersion = "LORAWAN_1.1.0"
)

func (r MacVersion) String() string {
	if len(strings.TrimSpace(string(r))) == 0 {
		return MacVersionUnknown.String()
	}
	return string(r)
}

func ParseMacVersionSourceFromInt32(r int32) MacVersion {
	switch r {
	case 0:
		return MacVersionLorawan100
	case 1:
		return MacVersionLorawan101
	case 2:
		return MacVersionLorawan102
	case 3:
		return MacVersionLorawan103
	case 4:
		return MacVersionLorawan104
	case 5:
		return MacVersionLorawan110
	default:
		return MacVersionUnknown
	}
}
