package device

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/device"
)

type link struct{}

var _ device.Link = (*link)(nil)

func (l *link) GetLinkMetrics(ctx context.Context, p device.GetLinkMetricsParams) (*device.LinkMetrics, error) {
	//TODO implement me
	panic("implement me")
}
