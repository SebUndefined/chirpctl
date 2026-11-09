package verbs

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/resources/device"
	"github.com/sebundefined/chirpctl/internal/commands/resources/gateway"
	"github.com/spf13/cobra"
)

func NewCreateCommand(a *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a resources",
	}

	cmd.AddCommand(
		device.NewCreateCommand(a),
		gateway.NewCreateCommand(a),
	)

	return cmd
}
