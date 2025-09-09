package chipctl

import (
	"context"
	"time"

	"github.com/sebundefined/chirpctl/internal/commands/contextx"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/logger"
	"github.com/sebundefined/chirpctl/internal/options"
	configSvc "github.com/sebundefined/chirpctl/internal/service/config"
	"github.com/sebundefined/chirpctl/internal/service/config/fsimpl"
	"github.com/spf13/cobra"
)

type closerKeyType struct{}

var closerKey closerKeyType

func AttachRuntime(cmd *cobra.Command, l logger.Interface) {
	origPre := cmd.PersistentPreRunE
	origPost := cmd.PersistentPostRunE

	cmd.PersistentPreRunE = func(c *cobra.Command, args []string) error {
		if origPre != nil {
			if err := origPre(c, args); err != nil {
				return err
			}
		}

		// 1) config + resolve
		cfg, _, err := config.Load()
		if err != nil {
			return err
		}
		l.Debug("Config loaded", logger.Any("config", cfg))
		opts := options.Resolve(cfg, c.Flags(), nil)
		l.Debug("Opts resolved", logger.Obj("opts", opts))
		// 2) timeout + options dans le ctx
		if opts.Timeout <= 0 {
			opts.Timeout = 10 * time.Second
		}
		ctx, cancel := context.WithTimeout(c.Context(), opts.Timeout)
		ctx = contextx.WithValue[logger.Interface](ctx, l)
		ctx = contextx.WithValue[options.Options](ctx, opts)
		cfgSvc := fsimpl.New()
		ctx = contextx.WithValue[configSvc.Service](ctx, cfgSvc)
		// 3) runtime (dial/provider) ou pas
		if shouldSkipRuntime(c) {
			// Pas de dial, mais on veut quand même cancel() en PostRun
			c.SetContext(context.WithValue(ctx, closerKey, func() error {
				cancel()
				return nil
			}))
			return nil
		}
		return nil
	}

	cmd.PersistentPostRunE = func(c *cobra.Command, args []string) error {
		if origPost != nil {
			if err := origPost(c, args); err != nil {
				return err
			}
		}
		if v := c.Context().Value(closerKey); v != nil {
			if cf, ok := v.(func() error); ok && cf != nil {
				return cf()
			}
		}
		return nil
	}
}

func shouldSkipRuntime(cmd *cobra.Command) bool {
	if cmd.Annotations != nil && cmd.Annotations["no-runtime"] == "true" {
		return true
	}
	// Sécurité: on check aussi quelques chemins courants
	switch cmd.Name() {
	case "help", "completion", "version":
		return true
	}
	// Exemple: le groupe "config" n'a pas besoin de dial
	if p := cmd; p != nil {
		for ; p.HasParent(); p = p.Parent() {
			if p.Name() == "config" {
				return true
			}
		}
	}
	return false
}
