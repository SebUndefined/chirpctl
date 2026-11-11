package config

import (
	"context"

	"github.com/sebundefined/chirpctl/internal/config"
)

type Service interface {
	Get(ctx context.Context) (*config.File, error)
	Init(ctx context.Context, f config.File) (config.File, error)
	ListProfiles(ctx context.Context) (map[string]config.Context, error)
	CreateProfile(ctx context.Context, name string, p config.Context) (config.File, error)
	DeleteProfile(ctx context.Context, name string) (config.File, error)
	UpdateProfile(ctx context.Context, name string, patch config.Context) (config.File, error)
	SetDefaultProfile(ctx context.Context, name string) (config.File, error)
	CurrentProfile(ctx context.Context) (string, config.Context, error)
}
