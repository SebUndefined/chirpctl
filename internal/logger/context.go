package logger

import "context"

type loggerKey struct{}

// With injecte un Service de config dans le contexte
func With(ctx context.Context, l Interface) context.Context {
	return context.WithValue(ctx, loggerKey{}, l)
}

// From extrait le Service de config du contexte
func From(ctx context.Context) Interface {
	v, _ := ctx.Value(loggerKey{}).(Interface)
	return v
}
