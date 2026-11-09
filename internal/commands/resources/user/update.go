package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newUpdateCommand() *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a user",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user update called")
		},
	}
	return userCmd
}
