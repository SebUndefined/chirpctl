package device

import "context"

type Keys interface {
	CreateKeys(ctx context.Context, p CreateKeyParams) error
	GetKeys(ctx context.Context, p GetKeysParams) (*KeysModel, error)
	UpdateKeys(ctx context.Context, p UpdateKeysParams) error
	DeleteKeys(ctx context.Context, p DeleteKeysParams) error
	FlushDevNonces(ctx context.Context, p FlushDevNoncesParams) error
}

type CreateKeyParams struct {
	DevEui    string
	NwkKey    string
	AppKey    string
	GenAppKey string
}

type KeysModel struct{}

type GetKeysParams struct{}

type UpdateKeysParams struct{}

type DeleteKeysParams struct{}

type FlushDevNoncesParams struct{}
