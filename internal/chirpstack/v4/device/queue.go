package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
)

type queue struct {
	c apiv4.DeviceServiceClient
}

func newQueue(c apiv4.DeviceServiceClient) *queue {
	return &queue{c: c}
}

var _ device.Queue = (*queue)(nil)

func (q *queue) Enqueue(ctx context.Context, p device.EnqueueParams) (*device.EnqueueResult, error) {
	//TODO implement me
	panic("implement me")
}

func (q *queue) GetQueue(ctx context.Context, p device.GetQueueParams) ([]device.QueuedItem, error) {
	//TODO implement me
	panic("implement me")
}

func (q *queue) FlushQueue(ctx context.Context, p device.FlushQueueParams) error {
	//TODO implement me
	panic("implement me")
}
