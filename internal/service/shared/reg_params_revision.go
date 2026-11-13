package shared

import "strings"

type RegParamsRevision string

const (
	RegParamsRevisionUnknown   RegParamsRevision = "UNKNOWN"
	RegParamsRevisionA         RegParamsRevision = "A"
	RegParamsRevisionB         RegParamsRevision = "B"
	RegParamsRevisionRP002v100 RegParamsRevision = "RP002_1.0.0"
	RegParamsRevisionRP002v101 RegParamsRevision = "RP002_1.0.1"
	RegParamsRevisionRP002v102 RegParamsRevision = "RP002_1.0.2"
	RegParamsRevisionRP002v103 RegParamsRevision = "RP002_1.0.3"
	RegParamsRevisionRP002v104 RegParamsRevision = "RP002_1.0.4"
	RegParamsRevisionRP002v105 RegParamsRevision = "RP002_1.0.5"
)

func (r RegParamsRevision) String() string {
	if len(strings.TrimSpace(string(r))) == 0 {
		return RegParamsRevisionUnknown.String()
	}
	return string(r)
}

func ParseRegParamsRevisionSourceFromInt32(r int32) RegParamsRevision {
	switch r {
	case 0:
		return RegParamsRevisionA
	case 1:
		return RegParamsRevisionB
	case 2:
		return RegParamsRevisionRP002v100
	case 3:
		return RegParamsRevisionRP002v101
	case 4:
		return RegParamsRevisionRP002v102
	case 5:
		return RegParamsRevisionRP002v103
	case 6:
		return RegParamsRevisionRP002v104
	case 7:
		return RegParamsRevisionRP002v105
	default:
		return RegParamsRevisionUnknown
	}
}
