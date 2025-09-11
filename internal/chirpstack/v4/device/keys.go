package device

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/device"
)

type keys struct {
}

var _ device.Keys = (*keys)(nil)

func (k *keys) GetKeys(ctx context.Context, p device.GetKeysParams) (*device.KeysModel, error) {
	//TODO implement me
	panic("implement me")
}

func (k *keys) UpdateKeys(ctx context.Context, p device.UpdateKeysParams) error {
	//TODO implement me
	panic("implement me")
}

func (k *keys) DeleteKeys(ctx context.Context, p device.DeleteKeysParams) error {
	//TODO implement me
	panic("implement me")
}

func (k *keys) FlushDevNonces(ctx context.Context, p device.FlushDevNoncesParams) error {
	//TODO implement me
	panic("implement me")
}

func (k *keys) CreateKeys(ctx context.Context, p device.CreateKeyParams) error {
	//TODO implement me
	panic("implement me")
}
