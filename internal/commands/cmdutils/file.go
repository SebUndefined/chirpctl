package cmdutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sebundefined/chirpctl/utils"
)

// openReader ouvre un fichier ou retourne os.Stdin si file == "-".
// La fonction retourne aussi un "closer" à appeler dans un defer.
// Exemple :
//
//	r, closeFn, err := openReader(path)
//	if err != nil { return err }
//	defer closeFn()
//	... utiliser r ...
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

// OpenWriter ouvre un fichier en écriture ou retourne os.Stdout si file == "-".
// Retourne aussi un closer pour être defer.
func OpenWriter(file string) (io.Writer, func() error, error) {
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
	f, err := os.Create(p)
	if err != nil {
		return nil, nil, err
	}
	return f, f.Close, nil
}
