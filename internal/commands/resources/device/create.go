package device

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/spf13/cobra"
)

func NewCreateCommand(a *app.App) *cobra.Command {
	// userCmd represents the user command
	var paramsCreateDevice device.CreateParams
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a device",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			svc, err := a.Factory.Device()
			if err != nil {
				return err
			}
			devEui, err := svc.Create(cmd.Context(), paramsCreateDevice)
			if err != nil {
				return err
			}
			fmt.Printf("Device created with DevEui: %s\n", devEui)
			return nil
		},
	}
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.ApplicationID,
		"app-id", "", "Application UUID",
	)
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.DevEUI,
		"dev-eui", "", "Device EUI",
	)
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.JoinEUI,
		"join-eui", "", "Join EUI",
	)
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.Name,
		"name", "", "Device name",
	)
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.Description,
		"description", "", "Device description",
	)
	cmd.Flags().StringVar(
		&paramsCreateDevice.Device.DeviceProfileID,
		"device-profile", "", "Device profile ID",
	)
	cmd.Flags().BoolVar(
		&paramsCreateDevice.Device.IsDisabled,
		"disabled", false, "Disable device",
	)
	cmd.Flags().BoolVar(
		&paramsCreateDevice.Device.SkipFcntCheck,
		"skipFcntCheck", false, "Disable frame-counter validation",
	)
	cmd.Flags().StringToStringVar(&paramsCreateDevice.Device.Variables, "vars",
		nil, "Device variables")
	cmd.Flags().StringToStringVar(&paramsCreateDevice.Device.Tags, "tags",
		nil, "Device tags")

	_ = cmd.MarkFlagRequired("app-id")
	_ = cmd.MarkFlagRequired("dev-eui")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("device-profile")
	return cmd
}
