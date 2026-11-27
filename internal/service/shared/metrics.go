package shared

import (
	"context"
	"strings"
	"time"
)

type Metrics[T any, R any] interface {
	GetMetrics(ctx context.Context, in T) (R, error)
}

type MetricKind string

const (
	MetricKindUnknown  MetricKind = "UNKNOWN"
	MetricKindCounter  MetricKind = "COUNTER"
	MetricKindAbsolute MetricKind = "ABSOLUTE"
	MetricKindGauge    MetricKind = "GAUGE"
)

func (a MetricKind) String() string {
	if len(strings.TrimSpace(string(a))) == 0 {
		return MetricKindUnknown.String()
	}
	return string(a)
}

func ParseMetricKindFromInt32(r int32) MetricKind {
	switch r {
	case 0:
		return MetricKindCounter
	case 1:
		return MetricKindAbsolute
	case 2:
		return MetricKindGauge
	default:
		return MetricKindUnknown
	}
}

type MetricDataset struct {
	Label string    `json:"label" yaml:"label"`
	Data  []float32 `json:"data" yaml:"data"`
}

type Metric struct {
	Name       string           `json:"name" yaml:"name"`
	Timestamps []*time.Time     `json:"timestamps" yaml:"timestamps"`
	Datasets   []*MetricDataset `json:"datasets" yaml:"datasets"`
	Kind       MetricKind       `json:"kind" yaml:"kind"`
}
