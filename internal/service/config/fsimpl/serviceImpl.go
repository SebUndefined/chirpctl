package fsimpl

import (
	"context"
	"fmt"

	"github.com/sebundefined/chirpctl/internal/config"
	configsvc "github.com/sebundefined/chirpctl/internal/service/config"
)

type configService struct{}

func New() configsvc.Service { return &configService{} }

func (s *configService) Get(ctx context.Context) (config.File, error) {
	f, _, err := config.Load()
	return *f, err
}

func (s *configService) Init(ctx context.Context, f config.File) (config.File, error) {
	if f.Profiles == nil {
		f.Profiles = map[string]config.Profile{}
	}
	if err := config.Save(&f); err != nil {
		return f, err
	}
	return f, nil
}

func (s *configService) ListProfiles(ctx context.Context) (map[string]config.Profile, error) {
	f, _, err := config.Load()
	if err != nil {
		return nil, err
	}

	return f.Profiles, nil
}

func (s *configService) CreateProfile(ctx context.Context, name string, p config.Profile) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if f.Profiles == nil {
		f.Profiles = map[string]config.Profile{}
	}
	f.Profiles[name] = p
	if f.DefaultProfile == "" {
		f.DefaultProfile = name
	}
	return *f, config.Save(f)
}

func (s *configService) UpdateProfile(ctx context.Context, name string, p config.Profile) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if _, ok := f.Profiles[name]; !ok {
		return *f, fmt.Errorf("profil %q introuvable", name)
	}
	f.Profiles[name] = p
	return *f, config.Save(f)
}

func (s *configService) DeleteProfile(ctx context.Context, name string) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	delete(f.Profiles, name)
	if f.DefaultProfile == name {
		f.DefaultProfile = ""
	}
	return *f, config.Save(f)
}

func (s *configService) SetDefaultProfile(ctx context.Context, name string) (config.File, error) {
	f, _, err := config.Load()
	if err != nil {
		return config.File{}, err
	}
	if _, ok := f.Profiles[name]; !ok {
		return *f, fmt.Errorf("profil %q introuvable", name)
	}
	f.DefaultProfile = name
	return *f, config.Save(f)
}

func (s *configService) CurrentProfile(ctx context.Context) (string, config.Profile, error) {
	f, _, err := config.Load()
	if err != nil {
		return "", config.Profile{}, err
	}
	p, ok := f.Profiles[f.DefaultProfile]
	if !ok {
		return "", config.Profile{}, fmt.Errorf("aucun profil par défaut")
	}
	return f.DefaultProfile, p, nil
}
