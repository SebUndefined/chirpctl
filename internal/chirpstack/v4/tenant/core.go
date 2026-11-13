package tenant

import (
	"context"
	"fmt"

	apiv4 "github.com/chirpstack/chirpstack/api/go/v4/api"
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
)

type core struct {
	client Client
}

var _ tenant.Core = (*core)(nil)

func newCore(client Client) *core {
	return &core{client: client}
}

func (c *core) Create(ctx context.Context, in tenant.CreateParams) (tenant.CreateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (c *core) Get(ctx context.Context, in tenant.GetParams) (*tenant.Tenant, error) {
	req := &apiv4.GetTenantRequest{
		Id: in.ID,
	}
	get, err := c.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return toTenant(get.Tenant, get.CreatedAt, get.UpdatedAt)
}

func (c *core) Update(ctx context.Context, in tenant.UpdateParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) Delete(ctx context.Context, in tenant.DeleteParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) List(ctx context.Context, in tenant.ListParams) (shared.ListResult[tenant.ListItem], error) {
	req := &apiv4.ListTenantsRequest{
		Limit:  in.GetSafeLimit(),
		Offset: in.Offset,
		Search: in.Search,
		UserId: in.UserID,
	}
	out := shared.NewListResult[tenant.ListItem]()
	resp, err := c.client.List(ctx, req)
	if err != nil {
		return out, err
	}
	items := resp.GetResult()
	out = shared.NewListResultCap[tenant.ListItem](len(items))
	out.TotalCount = resp.GetTotalCount()
	for _, item := range items {
		converted, err := toTenantListItem(item)
		if err != nil {
			return out, fmt.Errorf("convert tenant list item: %w", err)
		}
		out.Add(converted)
	}
	return out, nil
}

func (c *core) AddUser(ctx context.Context, in tenant.AddUserParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) GetUser(ctx context.Context, in tenant.GetUserParams) (*apiv4.GetTenantUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *core) UpdateUser(ctx context.Context, in tenant.UpdateUserParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) DeleteUser(ctx context.Context, in tenant.DeleteUserParams) error {
	//TODO implement me
	panic("implement me")
}

func (c *core) ListUsers(ctx context.Context, in tenant.ListUsersParams) (shared.ListResult[tenant.ListUserItem], error) {
	//TODO implement me
	panic("implement me")
}
