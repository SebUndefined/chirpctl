package deviceprofile

import (
	"github.com/spf13/cobra"
)

func NewDeviceProfileCommand() *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "device-profile",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	userCmd.AddCommand(
		newListCommand(),
		newGetCommand(),
		newCreateCommand(),
		newUpdateCommand(),
		newDeleteCommand(),
	)
	return userCmd
}
