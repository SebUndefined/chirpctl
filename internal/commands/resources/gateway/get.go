package gateway

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/sebundefined/chirpctl/internal/validation"
	"github.com/spf13/cobra"
)

// NewGetCommand registers gateway retrieval operations.
func NewGetCommand(a *app.App) *cobra.Command {
	var params gateway.ListParams
	var allTenants bool
	cmd := &cobra.Command{
		Use:     "gateways [GATEWAY_ID]",
		Aliases: []string{"gateway"},
		Short:   "Display gateways or a specific gateway",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				// Use the context default unless the tenant was explicitly selected.
				if !cmd.Flags().Changed("tenant-id") {
					params.TenantID = a.ResolvedOptions.TenantID
				}

				// Explicitly request all tenants, overriding the context default.
				if allTenants {
					params.TenantID = ""
				}

				// Validate the effective tenant filter when present.
				if params.TenantID != "" {
					if err := validation.ValidateUUID(params.TenantID); err != nil {
						return fmt.Errorf("invalid tenant ID: %w", err)
					}
				}

				return listGateways(cmd, a, params)
			}

			return getGateway(cmd, a, args[0])
		},
	}

	cmd.Flags().StringVarP(
		&params.TenantID, "tenant-id", "t", "",
		"Filter gateways by tenant UUID",
	)
	// Allow global administrators to bypass the default tenant filter.
	cmd.Flags().BoolVar(
		&allTenants,
		"all-tenants",
		false,
		"List gateways across all tenants (requires global admin permissions)",
	)

	// Prevent conflicting tenant selection flags.
	cmd.MarkFlagsMutuallyExclusive("tenant-id", "all-tenants")
	cmdutils.AttachPaginationFlags(cmd, &params.ListPaginationRequest)
	cmdutils.AttachOrderByFlags(cmd, &params.ListOrderByRequest)

	return cmd
}

func getGateway(cmd *cobra.Command, a *app.App, gatewayID string) error {
	if err := validation.ValidateHexLen(gatewayID, 8); err != nil {
		return fmt.Errorf("invalid gateway ID: %w", err)
	}
	svc, err := a.Factory.Gateway()
	if err != nil {
		return err
	}
	result, err := svc.Get(cmd.Context(), gateway.GetParams{
		GatewayID: gatewayID,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("get gateway: empty response")
	}

	return a.Printer.Print(result, buildGatewayDetailTable(result))
}

// listGateways retrieves a page of gateways and prints the result.
func listGateways(cmd *cobra.Command, a *app.App, params gateway.ListParams) error {
	svc, err := a.Factory.Gateway()
	if err != nil {
		return fmt.Errorf("initialize gateway service: %w", err)
	}

	result, err := svc.List(cmd.Context(), params)
	if err != nil {
		return fmt.Errorf("list gateways: %w", err)
	}

	return a.Printer.Print(result, buildGatewayListTable(result.Items))
}
