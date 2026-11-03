package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	Level  Level
	Format string // "console" or "json". "console" is the default
}
type ConcreteLogger struct {
	logger *zerolog.Logger
}

func NewLoggerWithConfig(cfg Config) Interface {
	zerolog.SetGlobalLevel(zerolog.Level(cfg.Level))

	var base zerolog.Logger
	if cfg.Format == "json" {
		base = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		base = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			With().Timestamp().Logger()
	}
	return &ConcreteLogger{
		logger: &base,
	}
}

func NewConcreteLogger(level Level) Interface {
	return NewLoggerWithConfig(Config{Level: level})
}

func NewDisabledLogger() Interface { return NewConcreteLogger(LevelDisabled) }

func (c *ConcreteLogger) With(f ...Field) Interface {
	ctx := c.logger.With()
	child := ctx.EmbedObject(fieldObject{fields: f}).Logger()
	return &ConcreteLogger{logger: &child}
}

func (c *ConcreteLogger) Trace(msg string, f ...Field) { c.dispatch(c.logger.Trace(), msg, f...) }
func (c *ConcreteLogger) Debug(msg string, f ...Field) { c.dispatch(c.logger.Debug(), msg, f...) }
func (c *ConcreteLogger) Info(msg string, f ...Field)  { c.dispatch(c.logger.Info(), msg, f...) }
func (c *ConcreteLogger) Warn(msg string, f ...Field)  { c.dispatch(c.logger.Warn(), msg, f...) }
func (c *ConcreteLogger) Error(msg string, f ...Field) { c.dispatch(c.logger.Error(), msg, f...) }
func (c *ConcreteLogger) Fatal(msg string, f ...Field) { c.dispatch(c.logger.Fatal(), msg, f...) }

func (c *ConcreteLogger) dispatch(e *zerolog.Event, msg string, f ...Field) {
	for _, fn := range f {
		if fn != nil {
			fn(e)
		}
	}
	e.Msg(msg)
}
