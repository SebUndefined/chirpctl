package config

import (
	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/spf13/cobra"
)

func NewConfigCommand(a *app.App) *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "config",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
