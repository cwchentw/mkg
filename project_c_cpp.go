package main

import (
	"fmt"
	"os"
)

func CreateCorCppProject(pr *ParsingResult) {
	_, err := os.Stat(pr.Path())
	if !os.IsNotExist(err) {
		if pr.IsForced() {
			err = os.RemoveAll(pr.Path())
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintln(os.Stderr,
				fmt.Sprintf("File or directory %s exists", pr.Path()))
			os.Exit(1)
		}
	}

	err = os.MkdirAll(pr.Path(), os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	createLicense(pr)
	createREADME(pr)
	createGitignore(pr)

	if pr.Layout() == LAYOUT_FLAT && pr.Proj() == PROJ_CONSOLE {
		createConfigAppFlat(pr)
		createApp(pr)
		createAppTest(pr)
	} else if pr.Layout() == LAYOUT_FLAT && pr.Proj() == PROJ_LIBRARY {
		createConfigLibFlat(pr)
		createHeader(pr)
		createLib(pr)
	} else if pr.Layout() == LAYOUT_NESTED && pr.Proj() == PROJ_CONSOLE {
		createProjStruct(pr)
		createConfigAppNested(pr)
		createConfigAppInternal(pr)
		createApp(pr)
		createAppTest(pr)
	} else if pr.Layout() == LAYOUT_NESTED && pr.Proj() == PROJ_LIBRARY {
		createProjStruct(pr)
		createConfigLibNested(pr)
		createConfigLibInternal(pr)
		createHeader(pr)
		createLib(pr)
	}
}
