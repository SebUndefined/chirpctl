package contextx

import (
	"context"
)

type key[T any] struct{}

func WithValue[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, key[T]{}, v)
}

func Value[T any](ctx context.Context) (T, bool) {
	v, ok := ctx.Value(key[T]{}).(T)
	return v, ok
}
