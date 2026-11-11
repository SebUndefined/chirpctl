package verbs

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/resources/device"
	"github.com/spf13/cobra"
)

func NewImportCommand(a *app.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import resources",
	}

	cmd.AddCommand(
		device.NewImportCommand(a),
	)

	return cmd
}
