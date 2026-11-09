package gateway

import (
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/spf13/cobra"
)

func NewCreateCommand(a *app.App) *cobra.Command {
	// userCmd represents the user command
	var userCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("user create called")
		},
	}
	return userCmd
}
