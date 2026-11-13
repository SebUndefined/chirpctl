package verbs

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/resources/device"
	"github.com/sebundefined/chirpctl/internal/commands/resources/deviceprofile"
	"github.com/sebundefined/chirpctl/internal/commands/resources/gateway"
	"github.com/sebundefined/chirpctl/internal/commands/resources/tenant"
	"github.com/spf13/cobra"
)

func NewGetCommand(a *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display one or more resources",
	}

	cmd.AddCommand(
		device.NewGetCommand(a),
		gateway.NewGetCommand(a),
		deviceprofile.NewGetCommand(a),
		tenant.NewGetCommand(a),
	)

	return cmd
}
