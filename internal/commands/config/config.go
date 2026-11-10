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
	"github.com/sebundefined/chirpctl/internal/printing"
	configWizard "github.com/sebundefined/chirpctl/internal/tui/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func newCurrentContextCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "current-context",
		Short: "Show the current context",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			name, ctx, err := a.CfgSvc.CurrentProfile(cmd.Context())
			if err != nil {
				return fmt.Errorf("get current context: %w", err)
			}

			view := ContextView{
				Name:       name,
				Current:    true,
				APIVersion: ctx.ApiVersion,
				Address:    ctx.Addr,
				TenantID:   ctx.TenantID,
				TLS:        !ctx.Insecure,
				HasToken:   ctx.Token != "",
			}

			return a.Printer.Print(
				view,
				buildCurrentContextTable(view),
			)
		},
	}

	return cmd
}

func newGetContextsCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-contexts",
		Short: "A brief description of your command",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgFile, err := a.CfgSvc.Get(cmd.Context())
			if err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			}

			views := buildContextViews(cfgFile)

			td := buildContextsTable(
				views,
				a.Printer.Format() == printing.FormatTableStyled,
				a.Printer.UseColor(),
			)

			return a.Printer.Print(views, td)
		},
	}

	return cmd
}

func newDeleteContextsCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete-context",
		Short: "A brief description of your command",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if _, err := a.CfgSvc.DeleteProfile(cmd.Context(), name); err != nil {
				return err
			}

			_, err := fmt.Fprintf(cmd.OutOrStdout(), "Context %q deleted\n", name)
			return err
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
			if cfgFile.Contexts == nil {
				cfgFile.Contexts = make(map[string]config.Context)
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
			cmdutils.SetIfString(cmd, "tenant-id", func(v string) { ctx.TenantID = v })
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
			if existed {
				_, err = a.CfgSvc.UpdateProfile(cmd.Context(), contextName, ctx)
			} else {
				_, err = a.CfgSvc.CreateProfile(cmd.Context(), contextName, ctx)
			}

			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("tenant-id", "", "Default tenant UUID for this context")
	return cmd
}

func newInitConfigCommand(a *app.App) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Init the profile",
		Long:  `Guide you into the steps for creating a new profile.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !term.IsTerminal(int(os.Stdin.Fd())) ||
				!term.IsTerminal(int(os.Stdout.Fd())) {
				return errors.New(
					"config init requires an interactive terminal; use config set-context for automation",
				)
			}
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

			_, err = a.CfgSvc.Init(cmd.Context(), config.File{
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
						TenantID:   res.TenantID,
					},
				},
			})
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
