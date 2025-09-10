package commands

import (
	"strings"
	"time"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/cmdutils"
	"github.com/sebundefined/chirpctl/internal/commands/application"
	"github.com/sebundefined/chirpctl/internal/commands/config"
	"github.com/sebundefined/chirpctl/internal/commands/device"
	"github.com/sebundefined/chirpctl/internal/commands/gateway"
	"github.com/sebundefined/chirpctl/internal/commands/shared"
	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RootOptions struct {
	Context    string
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
			// 1) Les commandes `chipctl config ...` n’ont pas besoin de config chargée
			if isConfigCommand(cmd) {
				return nil
			}
			// fusionner flags persistents et locaux du root
			merged := pflag.NewFlagSet("all", pflag.ContinueOnError)
			merged.AddFlagSet(cmd.InheritedFlags())
			merged.AddFlagSet(cmd.Flags())
			file, err := a.CfgSvc.Get(cmd.Context())
			if err != nil {
				return err
			}
			a.Cfg = file
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

	cmd.PersistentFlags().StringVar(&override.Context, "context", "", "Context to user")
	cmd.PersistentFlags().StringVar(&override.Addr, "addr", "", "gRPC address (host:port)")
	cmd.PersistentFlags().BoolVar(&override.Insecure, "insecure", false, "Disabled TLS certificate verification")
	cmd.PersistentFlags().StringVar(&override.ApiVersion, "api-version", "v4", "Set the API version")
	cmd.PersistentFlags().StringVar(&override.CACert, "cacert", "", "Path to CA certificate")
	cmd.PersistentFlags().StringVar(&override.ServerName, "server-name", "", "SNI / TLS name")
	cmd.PersistentFlags().StringVar(&override.Token, "token", "", "API token (Bearer)")
	cmd.PersistentFlags().DurationVar(&override.Timeout, "timeout", 10*time.Second, "Command timeout")
	cmd.PersistentFlags().StringVarP(&override.Output, "output", "o", "table", "table|json|yaml")
	cmd.PersistentFlags().BoolVarP(&override.Verbose, "verbose", "v", false, "Verbose logs")

	cmd.AddCommand(
		config.NewConfigCommand(a),
		NewVerbFirstRouter(shared.Get.String(), map[shared.Resource][]string{
			shared.Device: {shared.Device.String(), shared.Get.String()},
			//Gateway: gateway.NewGatewayGetCmd(a),
		}),
		NewVerbFirstRouter(shared.List.String(), map[shared.Resource][]string{
			shared.Device:  {shared.Device.String(), shared.List.String()},
			shared.Gateway: {shared.Gateway.String(), shared.List.String()},
			//Gateway: gateway.NewGatewayGetCmd(a),
		}),
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

// isConfigCommand retourne vrai si la commande en cours est sous l’arbre "config"
func isConfigCommand(cmd *cobra.Command) bool {
	for c := cmd; c != nil; c = c.Parent() {
		if c.Name() == "config" {
			return true
		}
	}
	return false
}
