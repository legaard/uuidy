package main

import (
	"os"

	"github.com/legaard/uuidy/cmd"
)

var (
	// value is overridden during build by ldflags
	version = "dev"
)

func main() {
	if err := cmd.Execute(version); err != nil {
		os.Exit(-1)
	}
}
