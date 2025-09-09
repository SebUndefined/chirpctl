package app

import (
	"github.com/sebundefined/chirpctl/internal/cmdutils"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/logger"
	"github.com/sebundefined/chirpctl/internal/options"
)

// App porte l’état global *résolu au root* (pas via context)
type App struct {
	Logger    logger.Interface
	Cfg       *config.File
	Overrides options.Options
	Factory   cmdutils.Factory
}
