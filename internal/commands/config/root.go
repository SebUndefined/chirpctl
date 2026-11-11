package config

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/spf13/cobra"
)

func NewConfigCommand(a *app.App) *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "config",
		Short: "Information about the Chirpctl configuration",
		//Long: ``,
	}
	userCmd.AddCommand(
		newGetContextsCommand(a),
		newCurrentContextCommand(a),
		newInitConfigCommand(a),
		newSetContextCommand(a),
		newDeleteContextsCommand(a),
	)
	return userCmd
}
