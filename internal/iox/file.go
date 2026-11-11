package iox

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sebundefined/chirpctl/utils"
)

// OpenReader open a file
func OpenReader(file string) (io.Reader, func() error, error) {
	if file == "" || file == "-" {
		// stdin
		return os.Stdin, func() error { return nil }, nil
	}

	// chemin de base
	p, err := utils.ExpandPath(file)
	if err != nil {
		return nil, nil, err
	}

	// en debug, préfixe uniquement si chemin relatif
	if utils.IsDevModeEnabled() && !filepath.IsAbs(p) {
		root, err := utils.DevRoot()
		if err != nil {
			return nil, nil, err
		}
		p = filepath.Join(root, p)
	}

	f, err := os.Open(p)
	if err != nil {
		return nil, nil, fmt.Errorf("open %q: %w", p, err)
	}
	return f, f.Close, nil
}

type WriterOption func(*writerOptions)

type writerOptions struct {
	noClobber bool
}

func WithNoClobber() WriterOption {
	return func(o *writerOptions) {
		o.noClobber = true
	}
}

// OpenWriter ouvre un fichier en écriture ou retourne os.Stdout si file == "-".
// Retourne aussi un closer pour être defer.
func OpenWriter(file string, opts ...WriterOption) (io.Writer, func() error, error) {
	if file == "" || file == "-" {
		return os.Stdout, func() error { return nil }, nil
	}
	// chemin de base
	p, err := utils.ExpandPath(file)
	if err != nil {
		return nil, nil, err
	}
	// en debug, préfixe uniquement si chemin relatif
	if utils.IsDevModeEnabled() && !filepath.IsAbs(p) {
		root, err := utils.DevRoot()
		if err != nil {
			return nil, nil, err
		}
		p = filepath.Join(root, p)
	}
	cfg := writerOptions{}
	for _, opt := range opts {
		opt(&cfg)
	}

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC

	if cfg.noClobber {
		flags = os.O_CREATE | os.O_EXCL | os.O_WRONLY
	}

	f, err := os.OpenFile(p, flags, 0600)
	if err != nil {
		return nil, nil, fmt.Errorf("open output %q: %w", p, err)
	}

	return f, f.Close, nil
}
