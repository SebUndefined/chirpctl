package user

import (
	"context"
	"fmt"

	"github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/samber/lo"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"github.com/sebundefined/chirpctl/internal/service/user"
)

type core struct {
	client Client
}

var _ user.Core = (*core)(nil)

func newCore(client Client) *core {
	return &core{client: client}
}

func (c *core) Create(ctx context.Context, in user.CreateParams) (user.CreateResult, error) {
	userTenant, err := lo.MapErr(in.Tenants, func(item *user.Tenant, index int) (*api.UserTenant, error) {
		if item == nil {
			return nil, fmt.Errorf("tenant is nil")
		}
		return &api.UserTenant{
			TenantId:       item.ID,
			IsAdmin:        item.IsAdmin,
			IsDeviceAdmin:  item.IsDeviceAdmin,
			IsGatewayAdmin: item.IsGatewayAdmin,
		}, nil
	})
	if err != nil {
		return user.CreateResult{}, err
	}
	create, err := c.client.Create(ctx, &api.CreateUserRequest{
		User: &api.User{
			Id:       "",
			IsAdmin:  in.User.IsAdmin,
			IsActive: in.User.IsActive,
			Email:    in.User.Email,
			Note:     in.User.Note,
		},
		Password: in.Password,
		Tenants:  userTenant,
	})
	if err != nil {
		return user.CreateResult{}, err
	}
	return user.CreateResult{
		ID: create.Id,
	}, nil
}

func (c *core) Get(ctx context.Context, in user.GetParams) (user.User, error) {
	panic("implement me")
}

func (c *core) Update(ctx context.Context, in user.UpdateParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) Delete(ctx context.Context, in user.DeleteParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) List(ctx context.Context, in user.ListParams) (shared.ListResult[user.ListItem], error) {
	//TODO implement me
	panic("implement me")
}

func (c *core) UpdatePassword(ctx context.Context, in user.UpdatePasswordParams) error {
	//TODO implement me
	panic("implement me")
}
