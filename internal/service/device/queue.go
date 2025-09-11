package device

import "context"

type Queue interface {
	Enqueue(ctx context.Context, p EnqueueParams) (*EnqueueResult, error)
	GetQueue(ctx context.Context, p GetQueueParams) ([]QueuedItem, error)
	FlushQueue(ctx context.Context, p FlushQueueParams) error
}

type EnqueueResult struct {
}

type EnqueueParams struct{}

type GetQueueParams struct{}

type FlushQueueParams struct{}

type QueuedItem struct{}
