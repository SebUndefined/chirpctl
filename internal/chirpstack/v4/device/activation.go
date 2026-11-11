package device

import (
	"context"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/device"
)

type activation struct {
	c apiv4.DeviceServiceClient
}

func newActivation(c apiv4.DeviceServiceClient) *activation {
	return &activation{
		c: c,
	}
}

var _ device.Activation = (*activation)(nil)

func (a *activation) Activate(ctx context.Context, p device.ActivationParams) error {
	//TODO implement me
	panic("implement me")
}

func (a *activation) Deactivate(ctx context.Context, p device.ActivationParams) error {
	//TODO implement me
	panic("implement me")
}

func (a *activation) GetActivation(ctx context.Context, p device.GetActivationParams) (*device.ActivationInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activation) GetRandomDevAddr(ctx context.Context, p device.GetRandomDevAddrParams) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *activation) GetNextFCntDown(ctx context.Context, p device.GetNextFCntDownParams) (uint32, error) {
	//TODO implement me
	panic("implement me")
}
