package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MTls struct {
	Enabled    bool   `yaml:"enabled"`
	ClientCert string `yaml:"clientCert"`
	ClientKey  string `yaml:"clientKey"`
}

type Profile struct {
	Addr       string `yaml:"addr"`
	Insecure   bool   `yaml:"insecure"`
	CACert     string `yaml:"caCert"`
	ServerName string `yaml:"serverName"`
	Token      string `yaml:"token"`
	Timeout    string `yaml:"timeout"` // "10s"
	MTls       MTls   `yaml:"mTls"`
	ApiVersion string `yaml:"apiVersion"`
}

type File struct {
	Version        int                `yaml:"version"`
	DefaultProfile string             `yaml:"default_profile"`
	Profiles       map[string]Profile `yaml:"profiles"`
}

func Load() (*File, string, error) {
	p, err := Path()
	if err != nil {
		return nil, "", err
	}

	b, err := os.ReadFile(p)
	if errors.Is(err, os.ErrNotExist) {
		return &File{Version: 1, Profiles: map[string]Profile{}}, p, nil
	}
	if err != nil {
		return nil, "", err
	}

	var f File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return nil, "", err
	}
	if f.Profiles == nil {
		f.Profiles = map[string]Profile{}
	}
	return &f, p, nil
}

func Save(f *File) error {
	p, err := Path()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}

	tmp := p + ".tmp"
	b, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}
