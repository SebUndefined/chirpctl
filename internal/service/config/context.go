package config

import "context"

type configKey struct{}

// With injecte un Service de config dans le contexte
func With(ctx context.Context, s Service) context.Context {
	return context.WithValue(ctx, configKey{}, s)
}

// From extrait le Service de config du contexte
func From(ctx context.Context) Service {
	v, _ := ctx.Value(configKey{}).(Service)
	return v
}
