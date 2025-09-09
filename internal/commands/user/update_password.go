package user

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newUpdatePasswordCommand() *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "update-password",
		Short: "Update password user",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user update-password called")
		},
	}
	return userCmd
}
