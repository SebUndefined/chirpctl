package shared

import "strings"

type Ts003Version string

const (
	Ts003VersionNotUnknown     Ts003Version = "UNKNOWN"
	Ts003VersionNotImplemented Ts003Version = "NOT_IMPLEMENTED"
	Ts003VersionV100           Ts003Version = "TS003_1.0.0"
	Ts003VersionV200           Ts003Version = "TS003_2.0.0"
)

func (c Ts003Version) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return Ts003VersionNotUnknown.String()
	}
	return string(c)
}

func ParseTs003VersionFromInt32(r int32) Ts003Version {
	switch r {
	case 0:
		return Ts003VersionNotImplemented
	case 1:
		return Ts003VersionV100
	case 2:
		return Ts003VersionV200
	default:
		return Ts003VersionNotUnknown
	}
}

type Ts004Version string

const (
	Ts004VersionNotUnknown     Ts004Version = "UNKNOWN"
	Ts004VersionNotImplemented Ts004Version = "NOT_IMPLEMENTED"
	Ts004VersionV100           Ts004Version = "TS004_1.0.0"
	Ts004VersionV200           Ts004Version = "TS004_2.0.0"
)

func (c Ts004Version) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return Ts004VersionNotUnknown.String()
	}
	return string(c)
}

func ParseTs004VersionFromInt32(r int32) Ts004Version {
	switch r {
	case 0:
		return Ts004VersionNotImplemented
	case 1:
		return Ts004VersionV100
	case 2:
		return Ts004VersionV200
	default:
		return Ts004VersionNotUnknown
	}
}

type Ts005Version string

const (
	Ts005VersionNotUnknown     Ts005Version = "UNKNOWN"
	Ts005VersionNotImplemented Ts005Version = "NOT_IMPLEMENTED"
	Ts005VersionV100           Ts005Version = "TS005_1.0.0"
	Ts005VersionV200           Ts005Version = "TS005_2.0.0"
)

func (c Ts005Version) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return Ts005VersionNotUnknown.String()
	}
	return string(c)
}

func ParseTs005VersionFromInt32(r int32) Ts005Version {
	switch r {
	case 0:
		return Ts005VersionNotImplemented
	case 1:
		return Ts005VersionV100
	case 2:
		return Ts005VersionV200
	default:
		return Ts005VersionNotUnknown
	}
}

type AppLayerParams struct {
	Ts003Version Ts003Version `json:"ts003Version" yaml:"ts003Version"`
	Ts003FPort   uint32       `json:"ts003FPort" yaml:"ts003FPort"`
	Ts004Version Ts004Version `json:"ts004Version" yaml:"ts004Version"`
	Ts004FPort   uint32       `json:"ts004FPort" yaml:"ts004FPort"`
	Ts005Version Ts005Version `json:"ts005Version" yaml:"ts005Version"`
	Ts005FPort   uint32       `json:"ts005FPort" yaml:"ts005FPort"`
}
