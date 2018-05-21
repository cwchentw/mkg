package main

import (
	"fmt"
	"os"
)

func main() {
	pr := NewParsingResult()

	if len(os.Args) <= 1 {
		// Run it interactively with sensible defaults.
		err := pr.RunWithDefaults()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else if len(os.Args) == 2 && os.Args[1] == "--custom" {
		// Run it interactively with more customization.
		err := pr.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		// Run it in batch mode.
		err := pr.ParseArgument(os.Args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	// Remove it later.
	fmt.Print(pr)
}
