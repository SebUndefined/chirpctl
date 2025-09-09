package main

import (
	"os"

	"github.com/sebundefined/chirpctl/internal/chipctl"
)

func main() {
	code := chipctl.Run()
	os.Exit(int(code))
}
