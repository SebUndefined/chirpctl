package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

func IsDevModeEnabled() bool {
	debugValue, isDebugSet := os.LookupEnv("CHIRPCTL_DEV")

	if !isDebugSet {
		return false
	}
	switch strings.ToLower(debugValue) {
	case "false", "0", "no", "":
		return false
	default:
		return true
	}
}

func ExpandPath(p string) (string, error) {
	if p == "" {
		return p, nil
	}
	if strings.HasPrefix(p, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home: %w", err)
		}
		p = filepath.Join(home, strings.TrimPrefix(p, "~"))
	}
	return filepath.Clean(p), nil
}

func DevRoot() (string, error) {
	if root := os.Getenv("CHIRPCTL_DEV_ROOT"); strings.TrimSpace(root) != "" {
		return ExpandPath(root)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getwd: %w", err)
	}
	return filepath.Join(cwd, "chirpctl-local-dev"), nil
}

var TerminalSize = func(w interface{}) (int, int, error) {
	if f, isFile := w.(*os.File); isFile {
		return term.GetSize(int(f.Fd()))
	}

	return 0, 0, fmt.Errorf("%v is not a file", w)
}
