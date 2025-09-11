package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
)

type link struct {
	c apiv4.DeviceServiceClient
}

func newLink(c apiv4.DeviceServiceClient) *link {
	return &link{c: c}
}

var _ device.Link = (*link)(nil)

func (l *link) GetLinkMetrics(ctx context.Context, p device.GetLinkMetricsParams) (*device.LinkMetrics, error) {
	//TODO implement me
	panic("implement me")
}
