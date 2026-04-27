// Package main is the entry point for chezmoi, a dotfile manager.
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/twpayne/chezmoi/v2/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "chezmoi: %v\n", err)
		// Exit with code 2 for usage errors, 1 for all other errors.
		// See: https://www.gnu.org/software/bash/manual/html_node/Exit-Status.html
		var usageErr *cmd.UsageError
		if errors.As(err, &usageErr) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
