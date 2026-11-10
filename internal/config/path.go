package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sebundefined/chirpctl/utils"
)

func Dir() (string, error) {
	switch runtime.GOOS {
	case "windows":
		if os.Getenv("APPDATA") != "" {
			return os.Getenv("APPDATA"), nil
		}
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support", "chirpctl"), nil
	default:
		if d := os.Getenv("XDG_CONFIG_HOME"); d != "" {
			return filepath.Join(d, "chirpctl"), nil
		}
		home, _ := os.UserHomeDir()
		if home != "" {
			return filepath.Join(home, ".config", "chirpctl"), nil
		}
	}
	return "", errors.New("unable to determine the home directory")
}

// Path get the path of the config file
func Path() (string, error) {
	if v := os.Getenv("CHIPCTL_CONFIG"); v != "" {
		p := v
		if info, err := os.Stat(p); err == nil && info.IsDir() {
			return filepath.Join(p, "config.yaml"), nil
		}
		if !filepath.IsAbs(p) {
			if cwd, err := os.Getwd(); err == nil {
				p = filepath.Join(cwd, p)
			}
		}
		return p, nil
	}
	if utils.IsDevModeEnabled() {
		if cwd, err := os.Getwd(); err == nil {
			return filepath.Join(cwd, "chirpctl-local-dev", "config.yaml"), nil
		}
	}
	d, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, "config.yaml"), nil
}
