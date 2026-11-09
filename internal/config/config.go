package config

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type MTls struct {
	Enabled    bool   `yaml:"enabled"`
	ClientCert string `yaml:"clientCert"`
	ClientKey  string `yaml:"clientKey"`
}

type Context struct {
	Addr       string `yaml:"addr"`
	Insecure   bool   `yaml:"insecure"`
	CACert     string `yaml:"caCert"`
	ServerName string `yaml:"serverName"`
	Token      string `yaml:"token"`
	TenantID   string `yaml:"tenantID"`
	Timeout    string `yaml:"timeout"` // "10s"
	MTls       MTls   `yaml:"mTls"`
	ApiVersion string `yaml:"apiVersion"`
}

func (c Context) Validate() error {
	if c.Addr == "" {
		return errors.New("addr is required")
	}
	if _, _, err := net.SplitHostPort(c.Addr); err != nil {
		return fmt.Errorf("invalid addr %q: %w", c.Addr, err)
	}
	if c.Timeout != "" {
		if _, err := time.ParseDuration(c.Timeout); err != nil {
			return fmt.Errorf("invalid timeout %q: %w", c.Timeout, err)
		}
	}
	if len(strings.TrimSpace(c.ApiVersion)) == 0 {
		return errors.New("apiVersion is required")
	}
	if c.MTls.Enabled {
		if c.MTls.ClientCert == "" || c.MTls.ClientKey == "" {
			return errors.New("mTLS enabled but client-cert or client-key missing")
		}
	}
	return nil
}

type UISettings struct {
	Output string `yaml:"output"` // "table-styled"|"table"|"json"|"yaml"|"csv" default is table-styled
}

type File struct {
	Version        int                `yaml:"version"`
	DefaultContext string             `yaml:"defaultContext"`
	Contexts       map[string]Context `yaml:"contexts"`
	UI             *UISettings        `yaml:"ui"`
}

func Load() (*File, string, error) {
	p, err := Path()
	if err != nil {
		return nil, "", err
	}

	b, err := os.ReadFile(p)
	if err != nil {
		return nil, "", err
	}

	var f File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return nil, "", err
	}
	if f.Contexts == nil {
		f.Contexts = map[string]Context{}
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
