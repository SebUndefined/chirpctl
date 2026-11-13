package shared

import "strings"

type MeasurementKind string

const (
	MeasurementKindUnknown  MeasurementKind = "UNKNOWN"
	MeasurementKindCounter  MeasurementKind = "COUNTER"
	MeasurementKindAbsolute MeasurementKind = "ABSOLUTE"
	MeasurementKindGauge    MeasurementKind = "GAUGE"
	MeasurementKindString   MeasurementKind = "STRING"
)

func (r MeasurementKind) String() string {
	if len(strings.TrimSpace(string(r))) == 0 {
		return MeasurementKindUnknown.String()
	}
	return string(r)
}

func ParseMeasurementKindSourceFromInt32(r int32) MeasurementKind {
	switch r {
	case 1:
		return MeasurementKindCounter
	case 2:
		return MeasurementKindAbsolute
	case 3:
		return MeasurementKindGauge
	case 4:
		return MeasurementKindString
	default:
		return MeasurementKindUnknown
	}
}

type Measurement struct {
	Name string          `json:"name" yaml:"name"`
	Kind MeasurementKind `json:"kind" yaml:"kind"`
}
