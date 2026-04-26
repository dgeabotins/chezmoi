// Package main is the entry point for chezmoi, a dotfile manager.
package main

import (
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
