package device

import "context"

type Activation interface {
	Activate(ctx context.Context, p ActivationParams) error
	Deactivate(ctx context.Context, p ActivationParams) error
	GetActivation(ctx context.Context, p GetActivationParams) (*ActivationInfo, error)
	GetRandomDevAddr(ctx context.Context, p GetRandomDevAddrParams) (string, error)
	GetNextFCntDown(ctx context.Context, p GetNextFCntDownParams) (uint32, error)
}

type ActivationParams struct {
}

type GetActivationParams struct{}

type GetRandomDevAddrParams struct{}

type GetNextFCntDownParams struct{}

type ActivationInfo struct {
}
