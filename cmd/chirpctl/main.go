package main

import (
	"os"

	"github.com/sebundefined/chirpctl/internal/chirpctl"
)

func main() {
	code := chirpctl.Run()
	os.Exit(int(code))
}
