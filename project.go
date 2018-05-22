package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func CreateProject(pr *ParsingResult) {
	_, err := os.Stat(pr.Path())
	if !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr,
			fmt.Sprintf("File or directory %s exists", pr.Path()))
		os.Exit(1)
	}

	err = os.MkdirAll(pr.Path(), os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if pr.License() != LICENSE_NONE {
		pathLicense := filepath.Join(pr.Path(), "LICENSE")
		fileLicense, err := os.Create(pathLicense)
		defer fileLicense.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		template := getTemplate(pr.License())
		now := time.Now()
		// Fix it later.
		_, err = fileLicense.WriteString(
			fmt.Sprintf(template, fmt.Sprint(now.Year()), "Michael Chen"))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}

func getTemplate(license License) string {
	switch license {
	case LICENSE_NONE:
		return ""
	case LICENSE_APACHE2:
		return template_apache2
	case LICENSE_MIT:
		return template_mit
	default:
		panic("Unknown license")
	}
}
