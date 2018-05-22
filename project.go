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

	createLicense(pr)
	createREADME(pr)
	createGitignore(pr)
}

func createLicense(pr *ParsingResult) {
	if pr.License() == LICENSE_NONE {
		return
	}

	path := filepath.Join(pr.Path(), "LICENSE")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	template := getTemplate(pr.License())
	now := time.Now()
	if pr.License() == LICENSE_GPL3 {
		_, err = file.WriteString(template)
	} else {
		_, err = file.WriteString(
			fmt.Sprintf(template, fmt.Sprint(now.Year()), pr.Author()))
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createREADME(pr *ParsingResult) {
	path := filepath.Join(pr.Path(), "README.md")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	now := time.Now()
	_, err = file.WriteString(
		fmt.Sprintf(template_readme,
			pr.Prog(), pr.Brief(), fmt.Sprint(now.Year()), pr.Author()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createGitignore(pr *ParsingResult) {
	path := filepath.Join(pr.Path(), ".gitignore")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var ignore string
	switch pr.Lang() {
	case LANG_C:
		ignore = gitignore_c
	case LANG_CPP:
		ignore = gitignore_cpp
	default:
		panic("Unknown language")
	}

	_, err = file.WriteString(ignore)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
