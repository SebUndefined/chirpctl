package shared

import "strings"

type Aggregation string

const (
	AggregationUnknown Aggregation = "UNKNOWN"
	AggregationHour    Aggregation = "HOUR"
	AggregationDay     Aggregation = "DAY"
	AggregationMinute  Aggregation = "MINUTE"
)

func (a Aggregation) String() string {
	if len(strings.TrimSpace(string(a))) == 0 {
		return AggregationUnknown.String()
	}
	return string(a)
}

func ParseAggregationFromInt32(r int32) Aggregation {
	switch r {
	case 0:
		return AggregationUnknown
	case 1:
		return AggregationHour
	case 2:
		return AggregationDay
	default:
		return AggregationMinute
	}
}
