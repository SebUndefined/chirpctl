package deviceprofile

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/service/shared"
)

type Core interface {
	Create(ctx context.Context, in CreateDeviceProfileParams) (CreateDeviceProfileResult, error)
	Get(ctx context.Context, in GetDeviceProfileParams) (GetDeviceProfileResult, error)
	Update(ctx context.Context, in UpdateDeviceProfileParams) error
	Delete(ctx context.Context, in DeleteDeviceProfileParams) error
	List(ctx context.Context, in ListDeviceProfilesParams) (shared.ListResult[ListItem], error)
	ListAdrAlgorithms(ctx context.Context) (ListDeviceProfileAdrAlgorithmsResult, error)
}

type CreateDeviceProfileParams struct{}

type GetDeviceProfileParams struct{}

type UpdateDeviceProfileParams struct{}

type DeleteDeviceProfileParams struct{}

type ListDeviceProfilesParams struct{}

type ListDeviceProfileAdrAlgorithmsResult struct {
}

type GetDeviceProfileResult struct{}

type ListItem struct{}

type CreateDeviceProfileResult struct{}
