// Package main is the entry point for chezmoi, a dotfile manager.
package main

import (
	"fmt"
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// Use exit code 2 for usage errors vs 1 for general errors,
		// but keep it simple here since cmd.Execute handles most cases.
		fmt.Fprintf(os.Stderr, "chezmoi: %v\n", err)
		os.Exit(1)
	}
}
