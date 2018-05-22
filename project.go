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

	now := time.Now()
	if pr.License() != LICENSE_NONE {
		pathLicense := filepath.Join(pr.Path(), "LICENSE")
		fileLicense, err := os.Create(pathLicense)
		defer fileLicense.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		template := getTemplate(pr.License())

		if pr.License() == LICENSE_GPL3 {
			_, err = fileLicense.WriteString(template)
		} else {
			_, err = fileLicense.WriteString(
				fmt.Sprintf(template, fmt.Sprint(now.Year()), pr.Author()))
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	pathREADME := filepath.Join(pr.Path(), "README.md")
	fileREADME, err := os.Create(pathREADME)
	defer fileREADME.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = fileREADME.WriteString(
		fmt.Sprintf(template_readme,
			pr.Prog(), pr.Brief(), fmt.Sprint(now.Year()), pr.Author()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
