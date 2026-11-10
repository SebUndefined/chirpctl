package commands

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/build"
	"github.com/sebundefined/chirpctl/internal/chirpstack/clientfactory"
	"github.com/sebundefined/chirpctl/internal/commands/config"
	"github.com/sebundefined/chirpctl/internal/commands/verbs"
	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/sebundefined/chirpctl/internal/printing"
	"github.com/spf13/cobra"
)

func NewRootCmd(a *app.App) *cobra.Command {
	override := &a.OverrideOptions

	cmd := &cobra.Command{
		Version: build.Version,
		Long: fmt.Sprintf(
			"CLI admin for Chirpstack.\n\nVersion: %s\nBuild date: %s",
			build.Version,
			build.Date,
		),
		Use:   "chirpctl",
		Short: "CLI ChipStark (gRPC)",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// Configuration commands do not require the remote runtime.
			if isConfigCommand(cmd) {
				return nil
			}

			// Load the configuration file.
			file, err := a.CfgSvc.Get(cmd.Context())
			if err != nil {
				return err
			}
			a.Cfg = file

			// Resolve settings: YAML -> environment -> CLI flags.
			resolved := options.Resolve(a.Cfg, cmd.Flags(), nil)
			a.ResolvedOptions = resolved
			// Validate the requested output format.
			format, err := printing.ParseFormat(resolved.Output)
			if err != nil {
				return err
			}

			// Configure the output renderer.
			a.Printer = printing.NewPrinter(
				cmd.OutOrStdout(),
				printing.Options{
					Format: format,
				},
			)
			// Configure the logging level.
			level := slog.LevelWarn
			if resolved.Verbose {
				level = slog.LevelDebug
			}

			// Write diagnostic logs to stderr.
			handler := slog.NewTextHandler(
				cmd.ErrOrStderr(),
				&slog.HandlerOptions{Level: level},
			)
			a.Logger = slog.New(handler)
			a.Logger.Debug(
				"ChirpStack connection configuration",
				"context", resolved.Context,
				"address", resolved.Addr,
				"api_version", resolved.ApiVersion,
				"tls_enabled", !resolved.Insecure,
				"token_configured", strings.TrimSpace(resolved.Token) != "",
			)
			// Keep your existing factory initialization here.
			// re-créer la factory
			if a.Factory != nil {
				_ = a.Factory.Close()
				a.Factory = nil
			}
			a.Logger.Debug("Runtime configuration resolved", "context", resolved.Context, "api_version", resolved.ApiVersion)
			a.Factory, err = clientfactory.NewFactory(&clientfactory.ResolveConfig{
				Addr:       resolved.Addr,
				Insecure:   resolved.Insecure,
				CACert:     resolved.CACert,
				ServerName: resolved.ServerName,
				Token:      resolved.Token,
				ApiVersion: strings.TrimPrefix(resolved.ApiVersion, "v"),
				Transport:  "grpc",
				Timeout:    resolved.Timeout,
			})
			if err != nil {
				return fmt.Errorf("initialize ChirpStack client factory: %w", err)
			}
			return nil
		},
	}
	cmd.SetVersionTemplate(
		fmt.Sprintf("chirpctl version %s (built %s)\n", build.Version, build.Date),
	)
	registerRootFlags(cmd, override)

	cmd.AddCommand(
		config.NewConfigCommand(a),
		verbs.NewGetCommand(a),
		verbs.NewCreateCommand(a),
		/*verbs.NewUpdateCommand(a),
		verbs.NewDeleteCommand(a),
		verbs.NewImportCommand(a),*/
	)

	return cmd
}

// isConfigCommand retourne vrai si la commande en cours est sous l’arbre "config"
func isConfigCommand(cmd *cobra.Command) bool {
	for c := cmd; c != nil; c = c.Parent() {
		if c.Name() == "config" {
			return true
		}
	}
	return false
}
