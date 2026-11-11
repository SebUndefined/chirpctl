package device

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/device"
)

type queue struct {
	client Client
}

func newQueue(client Client) *queue {
	return &queue{client: client}
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
