package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

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
	if isNoAuthor(pr.License()) {
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

func isNoAuthor(license License) bool {
	return license == LICENSE_AGPL3 || license == LICENSE_GPL2 ||
		license == LICENSE_GPL3 || license == LICENSE_MPL2
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

func createProjStruct(pr *ParsingResult) {
	// Create source dir
	path := filepath.Join(pr.Path(), pr.Src())
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create .gitkeep in source dir
	path = filepath.Join(path, ".gitkeep")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create include dir
	path = filepath.Join(pr.Path(), pr.Include())
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create .gitkeep in include dir
	path = filepath.Join(path, ".gitkeep")
	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create dist dir
	path = filepath.Join(pr.Path(), pr.Dist())
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create .gitkeep in dist dir
	path = filepath.Join(path, ".gitkeep")
	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create test dir
	path = filepath.Join(pr.Path(), pr.Test())
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create .gitkeep in test dir
	path = filepath.Join(path, ".gitkeep")
	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create example dir
	path = filepath.Join(pr.Path(), pr.Example())
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create .gitkeep in example dir
	path = filepath.Join(path, ".gitkeep")
	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
