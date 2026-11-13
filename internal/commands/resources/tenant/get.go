package tenant

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
	var params tenant.ListParams
	var cmd = &cobra.Command{
		Use:     "tenants [TENANT_ID]",
		Short:   "Display tenants or a specific tenant",
		Aliases: []string{"tenant"},
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return listTenants(cmd, a, params)
			}

			return getTenant(cmd, a, args[0])
		},
	}
	cmdutils.AttachPaginationFlags(cmd, &params.ListPaginationRequest)
	cmdutils.AttachOrderByFlags(cmd, &params.ListOrderByRequest)
	return cmd
}

func getTenant(cmd *cobra.Command, a *app.App, s string) error {
	if err := validation.ValidateUUID(s); err != nil {
		return fmt.Errorf("invalid tenant ID: %w", err)
	}
	svc, err := a.Factory.Tenant()
	if err != nil {
		return err
	}
	result, err := svc.Get(cmd.Context(), tenant.GetParams{
		ID: s,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("get gateway: empty response")
	}

	return a.Printer.Print(result, buildTenantDetailTable(result,
		a.Printer.Format() == printing.FormatTableStyled, a.Printer.UseColor()))
}

func listTenants(cmd *cobra.Command, a *app.App, params tenant.ListParams) error {
	svc, err := a.Factory.Tenant()
	if err != nil {
		return fmt.Errorf("initialize tenant service: %w", err)
	}
	result, err := svc.List(cmd.Context(), params)
	if err != nil {
		return fmt.Errorf("list tenant: %w", err)
	}
	return a.Printer.Print(result, buildTenantListTable(result.Items,
		a.Printer.Format() == printing.FormatTableStyled,
		a.Printer.UseColor()))
}
