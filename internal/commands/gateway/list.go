package gateway

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/sebundefined/chirpctl/internal/service/gateway"
	"github.com/spf13/cobra"
)

type ListGatewayOptions struct {
	TenantID         uuid.UUID
	MulticastGroupID uuid.UUID
}

func (o *ListGatewayOptions) Complete(_ *cobra.Command) error {
	// rien de spécial pour cet exemple
	return nil
}

func (o *ListGatewayOptions) Validate() error {
	// optionnel: vérifier formats, combinaisons invalides, etc.
	return nil
}

func (o *ListGatewayOptions) Run(parent context.Context) error {
	fmt.Printf("Listing gateways...\n")
	return nil
}

func newListCommand(a *app.App) *cobra.Command {
	var listParams gateway.ListParams
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List gateways",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {

			svc, err := a.Factory.Gateway()
			if err != nil {
				return err
			}
			list, err := svc.List(context.Background(), listParams)
			if err != nil {
				return err
			}
			format := printing.ParseFormat(a.Overrides.Output) // vient du flag/config
			noColor := os.Getenv("CHIRPCTL_NO_COLOR") != ""

			td := printing.TableData{
				Headers: printing.Row{"Gateway ID", "Name", "State", "Created At", "Updated At", "Last seen", "Tenant ID"},
				Rows:    make([]printing.Row, 0, len(list.Result)),
			}
			for _, d := range list.Result {
				stateCol := printing.StatusDot(printing.DotGray)
				switch d.State {
				case gateway.StateOffline:
					stateCol = printing.StatusDot(printing.DotRed)
				case gateway.StateOnline:
					stateCol = printing.StatusDot(printing.DotGreen)
				case gateway.StateNeverSeen:
					stateCol = printing.StatusDot(printing.DotYellow)
				default:
					stateCol = printing.StatusDot(printing.DotGray)

				}
				td.Rows = append(td.Rows, printing.Row{
					d.GatewayID,
					d.Name,
					stateCol,
					d.CreatedAt.Format("2006-01-02 15:04:05"),
					d.UpdatedAt.Format("2006-01-02 15:04:05"),
					d.LastSeenAt.Format("2006-01-02 15:04:05"),
					d.TenantID,
				})
			}
			return printing.RenderList(format, td, list, noColor)

		},
	}
	cmd.Flags().StringVarP(&listParams.TenantID, "tenant-id", "t", "", "Tenant UUID (required)")
	cmdutils.AttachPaginationFlags(cmd, &listParams.ListPaginationRequest)
	cmdutils.AttachOrderByFlags(cmd, &listParams.ListOrderByRequest)
	//cmdutils.AttachTagFlags(cmd, &listParams.Tags)
	_ = cmd.MarkFlagRequired("tenant-id")
	return cmd
}
