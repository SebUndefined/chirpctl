package device

import (
	"errors"
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/sebundefined/chirpctl/internal/validation"
	"github.com/spf13/cobra"
)

func NewGetCommand(a *app.App) *cobra.Command {
	var params device.ListParams
	var showVariables bool
	// deviceCmd represents the device command
	var deviceCmd = &cobra.Command{
		Use:     "devices [DEVICE_ID]",
		Aliases: []string{"device"},
		Short:   "Display devices or a specific device",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return listDevices(cmd, a, params)
			}

			return getDevice(cmd, a, args[0], showVariables)
		},
	}
	deviceCmd.Flags().StringVarP(
		&params.ApplicationID, "application-id", "a", "",
		"Filter devices by application ID",
	)
	// Allow global administrators to bypass the default tenant filter.
	deviceCmd.Flags().BoolVar(
		&showVariables,
		"show-variables",
		false,
		"Show variables for the device",
	)
	cmdutils.AttachPaginationFlags(deviceCmd, &params.ListPaginationRequest)
	cmdutils.AttachOrderByFlags(deviceCmd, &params.ListOrderByRequest)
	return deviceCmd
}

func listDevices(cmd *cobra.Command, a *app.App, params device.ListParams) error {
	applicationID := params.ApplicationID
	if applicationID == "" {
		return errors.New("application ID is required (--application-id)")
	}
	if err := validation.ValidateUUID(applicationID); err != nil {
		return fmt.Errorf("invalid application ID: %w", err)
	}
	svc, err := a.Factory.Device()
	if err != nil {
		return fmt.Errorf("initialize device service: %w", err)
	}
	result, err := svc.List(cmd.Context(), params)
	if err != nil {
		return fmt.Errorf("list devices: %w", err)
	}
	return a.Printer.Print(result, buildDeviceListTable(result.Items,
		a.Printer.Format() == printing.FormatTableStyled,
		a.Printer.UseColor()))
}

func getDevice(cmd *cobra.Command, a *app.App, devEui string, showVariables bool) error {
	if err := validation.ValidateHexLen(devEui, 8); err != nil {
		return fmt.Errorf("invalid device ID: %w", err)
	}
	svc, err := a.Factory.Device()
	if err != nil {
		return fmt.Errorf("initialize device service: %w", err)
	}
	result, err := svc.Get(cmd.Context(), device.GetParams{
		DevEUI: devEui,
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("get device: empty response")
	}
	return a.Printer.Print(result, buildDeviceDetailTable(result,
		showVariables,
		a.Printer.Format() == printing.FormatTableStyled,
		a.Printer.UseColor()))
}
