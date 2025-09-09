package logger

import (
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Level int

const (
	LevelTrace Level = iota - 1
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
	LevelDisabled
)

func LevelFromString(level string) Level {
	switch strings.ToLower(level) {
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	case "panic":
		return LevelPanic
	default:
		return LevelDisabled
	}
}

type Field func(*zerolog.Event)

type fieldObject struct {
	fields []Field
}

func (o fieldObject) MarshalZerologObject(e *zerolog.Event) {
	for _, f := range o.fields {
		if f != nil {
			f(e)
		}
	}
}

func Str(k, v string) Field         { return func(e *zerolog.Event) { e.Str(k, v) } }
func Int(k string, v int) Field     { return func(e *zerolog.Event) { e.Int(k, v) } }
func Int64(k string, v int64) Field { return func(e *zerolog.Event) { e.Int64(k, v) } }
func Float64(k string, v float64) Field {
	return func(e *zerolog.Event) { e.Float64(k, v) }
}
func Bool(k string, v bool) Field { return func(e *zerolog.Event) { e.Bool(k, v) } }
func Time(k string, t time.Time) Field {
	return func(e *zerolog.Event) { e.Time(k, t) }
}
func Dur(k string, d time.Duration) Field {
	return func(e *zerolog.Event) { e.Dur(k, d) }
}
func Any(k string, v any) Field { return func(e *zerolog.Event) { e.Interface(k, v) } }

func Err(err error) Field       { return func(e *zerolog.Event) { e.Err(err) } }
func Obj(k string, v any) Field { return func(e *zerolog.Event) { e.Interface(k, v) } }

type Interface interface {
	// Trace logs a message with trace level and key/value pairs, if any.
	Trace(msg string, f ...Field)
	// Debug logs a message with debug level and key/value pairs, if any.
	Debug(msg string, f ...Field)
	// Info logs a message with info level and key/value pairs, if any.
	Info(msg string, f ...Field)
	// Warn logs a message with warning level and key/value pairs, if any.
	Warn(msg string, f ...Field)
	// Error logs a message with error level and key/value pairs, if any.
	Error(msg string, f ...Field)
	// Fatal logs a message with fatal level and key/value pairs, if any.
	Fatal(msg string, f ...Field)
	With(f ...Field) Interface
}
