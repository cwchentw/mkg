package main

import (
	"fmt"
	"os"
)

func main() {
	pr := NewParsingResult()

	ev := PARSING_FSS_EVENT_ERROR
	var err error

	if len(os.Args) <= 1 {
		// Run it interactively with sensible defaults.
		err = pr.RunWithDefaults()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else if len(os.Args) == 2 && os.Args[1] == "--custom" {
		// Run it interactively with more customization.
		err = pr.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		// Run it in batch mode.
		ev, err = pr.ParseArgument(os.Args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Trick to use ev
		fmt.Sprintln(ev)
	}

	switch ev {
	case PARSING_FSS_EVENT_VERSION:
		printVersion()
		os.Exit(0)
	case PARSING_FSS_EVENT_HELP:
		printHelp(os.Stdout)
		os.Exit(0)
	case PARSING_FSS_EVENT_LICENSES:
		printLicenses()
		os.Exit(0)
	}

	// Remove it later.
	/*
		fmt.Println("")

		fmt.Print(pr)

		fmt.Println("")
	*/

	p := GetProject(pr)
	CreateDir(pr)
	p.Create()
}
