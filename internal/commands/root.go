package commands

import (
	"strings"
	"time"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/cmdutils"
	"github.com/sebundefined/chirpctl/internal/commands/application"
	"github.com/sebundefined/chirpctl/internal/commands/device"
	"github.com/sebundefined/chirpctl/internal/commands/gateway"
	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RootOptions struct {
	Profile    string
	Addr       string
	Insecure   bool
	CACert     string
	ServerName string
	Token      string
	Timeout    time.Duration
	Output     string
	Verbose    bool
	ApiVersion string // "v4"
	Transport  string // "grpc" | "rest" | "graphql"
}

func NewRootCmd(a *app.App) *cobra.Command {
	override := &a.Overrides

	cmd := &cobra.Command{
		Use:   "chirpctl",
		Short: "CLI ChipStark (gRPC)",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			// fusionner flags persistents et locaux du root
			merged := pflag.NewFlagSet("all", pflag.ContinueOnError)
			merged.AddFlagSet(cmd.InheritedFlags())
			merged.AddFlagSet(cmd.Flags())

			// resolve = profil YAML -> env -> flags
			res := options.Resolve(a.Cfg, merged, nil)

			// re-créer la factory
			if a.Factory != nil {
				_ = a.Factory.Close()
				a.Factory = nil
			}
			a.Factory = cmdutils.NewFactory(&cmdutils.ResolveConfig{
				Addr:       res.Addr,
				Insecure:   res.Insecure,
				CACert:     res.CACert,
				ServerName: res.ServerName,
				Token:      res.Token,
				ApiVersion: strings.TrimPrefix(res.ApiVersion, "v"),
				Transport:  "grpc",
				Timeout:    res.Timeout,
			})
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&override.Profile, "profile", "", "Profile to use")
	cmd.PersistentFlags().StringVar(&override.Addr, "addr", "", "gRPC address (host:port)")
	cmd.PersistentFlags().BoolVar(&override.Insecure, "insecure", false, "Disabled TLS certificate verification")
	cmd.PersistentFlags().StringVar(&override.CACert, "cacert", "", "Path to CA certificate")
	cmd.PersistentFlags().StringVar(&override.ServerName, "server-name", "", "SNI / TLS name")
	cmd.PersistentFlags().StringVar(&override.Token, "token", "", "API token (Bearer)")
	cmd.PersistentFlags().DurationVar(&override.Timeout, "timeout", 10*time.Second, "Command timeout")
	cmd.PersistentFlags().StringVarP(&override.Output, "output", "o", "table", "table|json|yaml")
	cmd.PersistentFlags().BoolVarP(&override.Verbose, "verbose", "v", false, "Verbose logs")

	// Brancher les sous-commandes
	cmd.AddCommand(
		application.NewApplicationCommand(a),
		device.NewDeviceCommand(a),
		gateway.NewGatewayCommand(a),
		/*deviceprofile.NewDeviceProfileCommand(),
		deviceprofiletemplate.NewDeviceProfileTemplateCommand(),
		,
		multicast.NewMulticastCommand(),
		relay.NewRelayCommand(),
		tenant.NewTenantCommand(),
		user.NewUserCommand(),
		configCmd.NewConfigCommand(),*/
	)

	return cmd
}
