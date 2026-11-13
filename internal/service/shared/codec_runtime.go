package shared

import "strings"

type CodecRuntime string

const (
	CodecRuntimeUnknown    CodecRuntime = "UNKNOWN"
	CodecRuntimeNone       CodecRuntime = "NONE"
	CodecRuntimeCayenneLPP CodecRuntime = "CAYENNE_LPP"
	CodecRuntimeJS         CodecRuntime = "JS"
)

func (c CodecRuntime) String() string {
	if len(strings.TrimSpace(string(c))) == 0 {
		return CodecRuntimeUnknown.String()
	}
	return string(c)
}

func ParseCodecRuntimeSourceFromInt32(c int32) CodecRuntime {
	switch c {
	case 0:
		return CodecRuntimeNone
	case 1:
		return CodecRuntimeCayenneLPP
	case 2:
		return CodecRuntimeJS
	default:
		return CodecRuntimeUnknown
	}
}
