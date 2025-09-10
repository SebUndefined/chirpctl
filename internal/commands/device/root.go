package device

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/shared"
	"github.com/spf13/cobra"
)

func NewDeviceCommand(a *app.App) *cobra.Command {
	var userCmd = &cobra.Command{
		Use:   shared.Device.String(),
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	userCmd.AddCommand(
		NewListCommand(a),
		NewGetCommand(a),
		NewCreateCommand(),
		NewUpdateCommand(),
		NewDeleteCommand(),
	)
	return userCmd
}
