package app

import (
	"log/slog"

	"github.com/sebundefined/chirpctl/internal/cmdutils"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/sebundefined/chirpctl/internal/printing"
	configSvc "github.com/sebundefined/chirpctl/internal/service/config"
)

// App porte l’état global *résolu au root* (pas via context)
type App struct {
	Logger          *slog.Logger
	Cfg             *config.File
	CfgSvc          configSvc.Service
	ResolvedOptions options.Options
	OverrideOptions options.Options
	Factory         cmdutils.Factory
	Printer         *printing.Printer
}
