package device

import (
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/spf13/cobra"
)

func NewListCommand(a *app.App) *cobra.Command {
	var listOptions device.ListParams
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List the devices",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			svc, err := a.Factory.Device()
			if err != nil {
				return err
			}
			list, err := svc.List(cmd.Context(), device.ListParams{
				ApplicationID: listOptions.ApplicationID,
			})
			if err != nil {
				return err
			}
			format := printing.ParseFormat(a.Overrides.Output) // vient du flag/config
			noColor := os.Getenv("CHIRPCTL_NO_COLOR") != ""

			td := printing.TableData{
				Headers: printing.Row{"DevEUI", "Name", "Created At", "Updated At", "Last seen", "Margin", "Bat. Level", "Profile name"},
				Rows:    make([]printing.Row, 0, len(list.Items)),
			}
			for _, d := range list.Items {
				td.Rows = append(td.Rows, printing.Row{
					d.DevEUI,
					d.Name,
					d.CreatedAt.Format("2006-01-02 15:04:05"),
					d.UpdatedAt.Format("2006-01-02 15:04:05"),
					d.LastSeenAt.Format("2006-01-02 15:04:05"),
					lo.Ternary(d.DeviceStatus != nil, string(d.DeviceStatus.Margin), ""),
					lo.Ternary(d.DeviceStatus != nil, fmt.Sprint(d.DeviceStatus.BatteryLevel), ""),
					d.DeviceProfileName,
				})
			}
			return printing.RenderList(format, td, list, noColor)
		},
	}
	cmd.Flags().StringVarP(&listOptions.ApplicationID, "app-id", "a", "", "Application UUID")
	_ = cmd.MarkFlagRequired("app-id")
	return cmd
}
