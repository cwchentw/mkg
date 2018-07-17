package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

func CreateDir(pr *ParsingResult) {
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
}

func createLicense(pr IProject) {
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

func createREADME(pr IProject) {
	path := filepath.Join(pr.Path(), "README.md")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tpl := fmt.Sprintf(templateREADME)
	tmpl, err := template.New("readmd").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	now := time.Now()

	if pr.License() == LICENSE_NONE {
		err = tmpl.Execute(file, struct {
			Prog    string
			Author  string
			Brief   string
			Year    int
			License string
		}{
			pr.Prog(),
			pr.Author(),
			pr.Brief(),
			now.Year(),
			"",
		})
	} else {
		err = tmpl.Execute(file, struct {
			Prog    string
			Author  string
			Brief   string
			Year    int
			License string
		}{
			pr.Prog(),
			pr.Author(),
			pr.Brief(),
			now.Year(),
			licenseToString(pr.License()),
		})
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createProjStruct(pr IProject) {
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
