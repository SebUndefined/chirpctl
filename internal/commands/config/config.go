package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/config"
	configWizard "github.com/sebundefined/chirpctl/internal/tui/config"
	"github.com/spf13/cobra"
)

func newProfileCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "profile",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	cmd.AddCommand(
		newInitConfigCommand(a),
		newGetCurrentProfileCommand(a),
	)
	return cmd
}

func newGetCurrentProfileCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "view",
		Short: "A brief description of your command",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			/*ctx := cmd.Context()
			svc := configctx.From(ctx)
			name, prof, err := svc.CurrentProfile(ctx)
			if err != nil {
				return err
			}

			p := tea.NewProgram(NewProfileView(name, prof))
			res, err := p.RunNewContextWizard()
			return err
			a.CfgSvc.Init(cmd.Context())
			fmt.Println(a.Cfg)*/
			return nil
		},
	}

	return cmd
}

func newInitConfigCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Init the profile",
		Long:  `Guide you into the steps for creating a new profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			get, err := a.CfgSvc.Get(cmd.Context())
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}
			if get != nil && len(get.Contexts) > 0 {
				return errors.New("config Already Exists")
			}
			res, err := configWizard.RunNewContextWizard()
			if err != nil {
				return err
			}

			init, err := a.CfgSvc.Init(cmd.Context(), config.File{
				Version:        1,
				DefaultContext: res.Name,
				Contexts: map[string]config.Context{
					res.Name: {
						Addr:       res.Addr,
						Insecure:   res.Insecure,
						CACert:     res.CACert,
						ServerName: res.ServerName,
						Token:      res.Token,
						Timeout:    res.Timeout,
						MTls:       config.MTls{},
						ApiVersion: res.ApiVersion,
					},
				},
			})
			if err != nil {
				return err
			}
			fmt.Println(res)
			fmt.Println(get)
			fmt.Println(init)
			return nil
		},
	}

	return cmd
}
