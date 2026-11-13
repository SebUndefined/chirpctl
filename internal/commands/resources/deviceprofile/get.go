package deviceprofile

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/deviceprofile"
	"github.com/sebundefined/chirpctl/internal/validation"
	"github.com/spf13/cobra"
)

func NewGetCommand(a *app.App) *cobra.Command {
	var params deviceprofile.ListParams
	var cmd = &cobra.Command{
		Use:     "device-profiles [DEVICE_PROFILE_ID]",
		Short:   "Display device profiles or a specific device profile",
		Aliases: []string{"device-profile"},
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return listDeviceProfiles(cmd, a, params)
			}

			return getDeviceProfile(cmd, a, args[0])
		},
	}
	cmd.Flags().StringVarP(
		&params.TenantID, "tenant-id", "t", "",
		"Filter gateways by tenant UUID",
	)
	cmdutils.AttachPaginationFlags(cmd, &params.ListPaginationRequest)
	cmdutils.AttachOrderByFlags(cmd, &params.ListOrderByRequest)
	return cmd
}

func getDeviceProfile(cmd *cobra.Command, a *app.App, s string) error {
	if err := validation.ValidateUUID(s); err != nil {
		return fmt.Errorf("invalid device profile ID: %w", err)
	}
	svc, err := a.Factory.DeviceProfile()
	if err != nil {
		return err
	}
	result, err := svc.Get(cmd.Context(), deviceprofile.GetParams{
		ID: s,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("get gateway: empty response")
	}

	return a.Printer.Print(result, buildDeviceProfileDetailTable(result,
		a.Printer.Format() == printing.FormatTableStyled, a.Printer.UseColor()))
}

func listDeviceProfiles(cmd *cobra.Command, a *app.App, params deviceprofile.ListParams) error {
	svc, err := a.Factory.DeviceProfile()
	if err != nil {
		return fmt.Errorf("initialize device profile service: %w", err)
	}
	result, err := svc.List(cmd.Context(), params)
	if err != nil {
		return fmt.Errorf("list device profiles: %w", err)
	}
	return a.Printer.Print(result, buildDeviceProfileListTable(result.Items,
		a.Printer.Format() == printing.FormatTableStyled,
		a.Printer.UseColor()))
}
