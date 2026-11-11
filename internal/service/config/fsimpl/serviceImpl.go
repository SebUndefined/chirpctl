package fsimpl

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/sebundefined/chirpctl/internal/config"
	configsvc "github.com/sebundefined/chirpctl/internal/service/config"
)

type configService struct{}

func New() configsvc.Service { return &configService{} }

func (s *configService) Get(ctx context.Context) (*config.File, error) {
	f, _, err := config.Load()
	return f, err
}

func (s *configService) Init(ctx context.Context, f config.File) (config.File, error) {
	if f.Contexts == nil {
		f.Contexts = map[string]config.Context{}
	}
	if err := f.Validate(); err != nil {
		return f, fmt.Errorf("invalid context %w", err)
	}
	if err := config.Save(&f); err != nil {
		return f, err
	}
	return f, nil
}

func (s *configService) ListProfiles(ctx context.Context) (map[string]config.Context, error) {
	f, _, err := config.Load()
	if err != nil {
		return nil, err
	}

	return f.Contexts, nil
}

func (s *configService) CreateProfile(ctx context.Context, name string, p config.Context) (config.File, error) {
	f, err := loadEditableConfig()
	if err != nil {
		return config.File{}, err
	}
	if f.Contexts == nil {
		f.Contexts = map[string]config.Context{}
	}
	if _, exists := f.Contexts[name]; exists {
		return *f, fmt.Errorf("context %q already exists", name)
	}

	if err := p.Validate(); err != nil {
		return *f, fmt.Errorf("invalid context %q: %w", name, err)
	}
	f.Contexts[name] = p
	if f.DefaultContext == "" {
		f.DefaultContext = name
	}
	return *f, config.Save(f)
}

func (s *configService) UpdateProfile(ctx context.Context, name string, p config.Context) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if _, ok := f.Contexts[name]; !ok {
		return *f, fmt.Errorf("profil %q introuvable", name)
	}
	if err := p.Validate(); err != nil {
		return *f, fmt.Errorf("invalid context %q: %w", name, err)
	}
	f.Contexts[name] = p
	return *f, config.Save(f)
}

func (s *configService) DeleteProfile(ctx context.Context, name string) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if _, exists := f.Contexts[name]; !exists {
		return *f, fmt.Errorf("context %q not found", name)
	}

	if f.DefaultContext == name {
		return *f, fmt.Errorf(
			"cannot delete the default context %q; switch context first",
			name,
		)
	}

	delete(f.Contexts, name)
	return *f, config.Save(f)
}

func (s *configService) SetDefaultProfile(ctx context.Context, name string) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if _, ok := f.Contexts[name]; !ok {
		return *f, fmt.Errorf("profil %q introuvable", name)
	}
	f.DefaultContext = name
	return *f, config.Save(f)
}

func (s *configService) CurrentProfile(ctx context.Context) (string, config.Context, error) {
	f, _, err := config.Load()
	if err != nil {
		return "", config.Context{}, err
	}
	p, ok := f.Contexts[f.DefaultContext]
	if !ok {
		return "", config.Context{}, fmt.Errorf("aucun profil par défaut")
	}
	return f.DefaultContext, p, nil
}

func loadEditableConfig() (*config.File, error) {
	f, _, err := config.Load()

	if errors.Is(err, os.ErrNotExist) {
		return &config.File{
			Version:  1,
			Contexts: make(map[string]config.Context),
		}, nil
	}

	if err != nil {
		return nil, err
	}

	if f == nil {
		return &config.File{
			Version:  1,
			Contexts: make(map[string]config.Context),
		}, nil
	}

	if f.Contexts == nil {
		f.Contexts = make(map[string]config.Context)
	}

	return f, nil
}
