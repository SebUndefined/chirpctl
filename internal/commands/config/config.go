package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands/cmdutils"
	"github.com/sebundefined/chirpctl/internal/config"
	configWizard "github.com/sebundefined/chirpctl/internal/tui/config"
	"github.com/spf13/cobra"
)

func newProfileCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "config",
		Short: "Manage chirpctl configuration",
		Long:  ``,
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

func newSetContextCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "set-context",
		Short: "Set a context entry in the configuration file",
		Long:  `Specifying a name that already exists will merge new fields on top of existing values for those fields.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, err := a.CfgSvc.Get(cmd.Context())
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}
			if cfgFile == nil {
				cfgFile = &config.File{Version: 1, Contexts: map[string]config.Context{}}
			}
			contextName, _ := cmd.Flags().GetString("context")
			if contextName == "" {
				return errors.New("--context is required")
			}
			ctx, existed := cfgFile.Contexts[contextName] // zéro value si nouveau

			// 3) merge avec flags root (uniquement ceux passés)
			cmdutils.SetIfString(cmd, "addr", func(v string) { ctx.Addr = v })
			cmdutils.SetIfBool(cmd, "insecure", func(v bool) { ctx.Insecure = v })
			cmdutils.SetIfString(cmd, "cacert", func(v string) { ctx.CACert = v })
			cmdutils.SetIfString(cmd, "server-name", func(v string) {
				ctx.ServerName = v
				if !ctx.Insecure && ctx.ServerName == "" {
					if host, _, err := net.SplitHostPort(ctx.Addr); err == nil {
						ctx.ServerName = host
					}
				}
			})
			cmdutils.SetIfString(cmd, "token", func(v string) { ctx.Token = v })
			cmdutils.SetIfDuration(cmd, "timeout", func(d time.Duration) { ctx.Timeout = d.String() })
			cmdutils.SetIfString(cmd, "api-version", func(v string) { ctx.ApiVersion = v })
			if errValidation := ctx.Validate(); errValidation != nil {
				return errValidation
			}
			// 4) écrire et sauver
			cfgFile.Contexts[contextName] = ctx
			if len(strings.TrimSpace(cfgFile.DefaultContext)) == 0 {
				cfgFile.DefaultContext = contextName
			}
			if err := config.Save(cfgFile); err != nil {
				return err
			}

			if err := config.Save(cfgFile); err != nil {
				return err
			}
			if !existed {
				fmt.Printf("context %q created\n", contextName)
			} else {
				fmt.Printf("context %q updated\n", contextName)
			}
			if cfgFile.DefaultContext == contextName {
				fmt.Printf("default context: %s\n", contextName)
			}
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
