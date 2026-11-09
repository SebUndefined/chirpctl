package chipctl

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/commands"
	"github.com/sebundefined/chirpctl/internal/service/config/fsimpl"
)

type ExitCode int

const (
	exitOK      ExitCode = 0
	exitError   ExitCode = 1
	exitCancel  ExitCode = 2
	exitAuth    ExitCode = 4
	exitPending ExitCode = 8
)

func Run() ExitCode {
	a := &app.App{
		CfgSvc: fsimpl.New(),
	}
	rootCmd := commands.NewRootCmd(a)
	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		return exitError
	}
	return exitOK
}
