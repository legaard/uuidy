package main

import (
	"os"
	"uuid/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
