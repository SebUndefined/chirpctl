package user

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/tenant"
	"github.com/sebundefined/chirpctl/internal/validation"
	"github.com/spf13/cobra"
)

func NewGetCommand(a *app.App) *cobra.Command {
	var params tenant.ListUsersParams
	var cmd = &cobra.Command{
		Use:     "users TENANT_ID [USER_ID]",
		Short:   "Display users or a specific user for a specific tenant",
		Aliases: []string{"user"},
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			tenantID := args[0]
			if len(args) == 1 {
				params.TenantID = tenantID
				return listUsers(cmd, a, params)
			}

			return getUser(cmd, a, tenantID, args[0])
		},
	}
	cmdutils.AttachPaginationFlags(cmd, &params.ListPaginationRequest)
	return cmd
}

func getUser(cmd *cobra.Command, a *app.App, tenantID, s string) error {
	if err := validation.ValidateUUID(s); err != nil {
		return fmt.Errorf("invalid tenant ID: %w", err)
	}
	svc, err := a.Factory.Tenant()
	if err != nil {
		return err
	}
	result, err := svc.GetUser(cmd.Context(), tenant.GetUserParams{
		ID:       s,
		TenantID: tenantID,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("get gateway: empty response")
	}

	return a.Printer.Print(result, buildTenantUserDetailTable(result,
		a.Printer.Format() == printing.FormatTableStyled, a.Printer.UseColor()))
}

func listUsers(cmd *cobra.Command, a *app.App, params tenant.ListUsersParams) error {
	svc, err := a.Factory.Tenant()
	if err != nil {
		return fmt.Errorf("initialize tenant service: %w", err)
	}
	result, err := svc.ListUsers(cmd.Context(), params)
	if err != nil {
		return fmt.Errorf("list tenant: %w", err)
	}
	return a.Printer.Print(result, buildTenantUsersListTable(result.Items,
		a.Printer.Format() == printing.FormatTableStyled,
		a.Printer.UseColor()))
}
