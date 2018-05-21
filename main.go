package main

import (
	"fmt"
	"os"
)

func main() {
	pr := NewParsingResult()

	if len(os.Args) <= 1 {
		// Run it interactively with sensible defaults.
	} else if len(os.Args) == 2 && os.Args[1] == "--custom" {
		// Run it interactively with more customization.
	} else {
		// Run it in batch mode.
		_ = pr.ParseArgument(os.Args)
	}

	// Remove it later.
	fmt.Print(pr)
}
