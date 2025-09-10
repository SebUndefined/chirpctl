package device

import (
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers devices.
This device is a tool to generate the needed files
to quickly create a Cobra device.`,
		Run: func(cmd *cobra.Command, args []string) {
			//l := do.MustInvokeNamed[logger.Interface]("logger")
		},
	}
	return userCmd
}
