package chipctl

import (
	"context"
	"fmt"

	"github.com/sebundefined/chirpctl/internal/app"
	"github.com/sebundefined/chirpctl/internal/build"
	"github.com/sebundefined/chirpctl/internal/commands"
	"github.com/sebundefined/chirpctl/internal/config"
	"github.com/sebundefined/chirpctl/internal/options"
	"github.com/sebundefined/chirpctl/utils"
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
	buildDate := build.Date
	buildVersion := build.Version
	hasDebug, _ := utils.IsDebugEnabled()
	fmt.Println("Chirspack CLI Version:", buildVersion)
	fmt.Println("Chirspack CLI Build Date:", buildDate)
	fmt.Println("Chirspack CLI has debug:", hasDebug)

	cfg, _, err := config.Load()
	if err != nil {
		return 1
	}
	fmt.Println("Chirspack CLI Config:", cfg)
	a := &app.App{
		Cfg: cfg,
		Overrides: options.Options{
			// valeurs par défaut des flags (root les exposera)
			Output:     "table",
			ApiVersion: "v4",
			Timeout:    10_000_000_000, // 10s
		},
	}
	rootCmd := commands.NewRootCmd(a)
	//ctx := context.Background()
	//cmd := commands.NewRootCmd(cfg)
	//cmdFactory := cmdutils.NewFactory()
	err = rootCmd.ExecuteContext(context.Background())
	if err != nil {
		return exitError
	}
	return exitOK
}

/*func Run() int {
	// Si tu veux un contexte racine custom (cancel sur SIGINT, etc.)
	ctx := context.Background()
	l := logger.NewLoggerWithConfig(logger.Config{
		Level:  logger.LevelDebug,
		Format: "console",
	})
	l.Debug("Starting chipctlapp")
	// 1) config + resolve
	cfg, _, err := config.Load()
	if err != nil {
		return 1
	}

	cmd := commands.NewRootCmd(cfg)
	//AttachRuntime(cmd, l)
	if _, err := cmd.ExecuteContextC(ctx); err != nil {
		// mapping propre des codes de sortie
		if errors.Is(err, context.Canceled) {
			return 130 // Ctrl-C
		}
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}*/
