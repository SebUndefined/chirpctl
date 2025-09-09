package device

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/service/device"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ApplicationID string
}

func newListCommand(a *app.App) *cobra.Command {
	var listOptions ListOptions
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List all users",
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
			tw := table.NewWriter()
			tw.SetOutputMirror(os.Stdout)
			tw.AppendHeader(table.Row{"DevEUI", "Name", "Created At", "Updated At", "Last seen", "Profile name"})
			for _, d := range list.Result {
				tw.AppendRow(table.Row{
					d.DevEUI,
					d.Name,
					d.CreatedAt.Format("2006-01-02 15:04:05"),
					d.UpdatedAt.Format("2006-01-02 15:04:05"),
					d.LastSeenAt.Format("2006-01-02 15:04:05"),
					d.DeviceProfileName,
				})
			}
			if len(list.Result) == 0 {
				tw.AppendRow(table.Row{"—", "—", "—", "—", "—", "—"})
			}
			tw.Render()
			return nil
		},
	}
	cmd.Flags().StringVarP(&listOptions.ApplicationID, "app-id", "a", "", "Application UUID")
	_ = cmd.MarkFlagRequired("app")
	return cmd
}
