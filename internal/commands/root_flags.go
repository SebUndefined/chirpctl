package commands

import (
	"time"

	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/spf13/cobra"
)

func registerRootFlags(cmd *cobra.Command, opts *options.Options) {
	cmd.PersistentFlags().StringVar(&opts.Context, "context", "", "Context to use")
	cmd.PersistentFlags().StringVar(&opts.Addr, "addr", "", "gRPC address (host:port)")
	cmd.PersistentFlags().BoolVar(&opts.Insecure, "insecure", false, "Use an unencrypted gRPC connection (unsafe)")
	cmd.PersistentFlags().StringVar(&opts.ApiVersion, "api-version", "v4", "ChirpStack API version")
	cmd.PersistentFlags().StringVar(&opts.CACert, "cacert", "", "Path to CA certificate")
	cmd.PersistentFlags().StringVar(&opts.ServerName, "server-name", "", "SNI / TLS name")
	cmd.PersistentFlags().StringVar(&opts.Token, "token", "", "API token (Bearer)")
	cmd.PersistentFlags().DurationVar(&opts.Timeout, "timeout", 10*time.Second, "Command timeout")
	cmd.PersistentFlags().StringVarP(&opts.Output, "output", "o", "table-styled", "table-styled|table|json|yaml|csv")
	cmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "Enable verbose logging")
}
