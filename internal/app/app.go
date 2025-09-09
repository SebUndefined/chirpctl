package app

import (
	"github.com/sebundefined/chirpctl/internal/cmdutils"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/logger"
	"github.com/sebundefined/chirpctl/internal/options"
	configSvc "github.com/sebundefined/chirpctl/internal/service/config"
)

// App porte l’état global *résolu au root* (pas via context)
type App struct {
	Logger    logger.Interface
	Cfg       *config.File
	CfgSvc    configSvc.Service
	Overrides options.Options
	Factory   cmdutils.Factory
}
