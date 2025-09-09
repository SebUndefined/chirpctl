package config

import (
	"github.com/spf13/cobra"
)

func NewConfigCommand() *cobra.Command {
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
		newProfileCommand(),
	)
	return userCmd
}
