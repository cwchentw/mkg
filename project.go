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

	// Remove it later.
	if pr.Layout() == LAYOUT_NESTED || pr.Proj() == PROJ_LIB {
		panic("Unimplemented")
	}

	createLicense(pr)
	createREADME(pr)
	createGitignore(pr)

	if pr.Layout() == LAYOUT_FLAT && pr.Proj() == PROJ_APP {
		createConfigAppFlat(pr)
		createAppFlat(pr)
	} else {
		// Implement it later.
	}
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

func createConfigAppFlat(pr *ParsingResult) {
	path := filepath.Join(pr.Path(), pr.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CC or CXX

	CFLAGS_DEBUG or CXXFLAGS_DEBUG

	CFLAGS_RELEASE or CXXFLAGS_RELEASE

	TARGET

	CFLAGS or CXX_FLAGS

	RM

	PROGRAM

	RULE_APP_C or RULE_APP_CXX

	RULE_RM
	*/
	config := `%s
%s
%s
%s
%s
%s
%s
%s
%s

%s
%s`

	var template string
	if pr.Lang() == LANG_C {
		template = fmt.Sprintf(config,
			config_platform,
			config_cc,
			config_cflags_debug,
			config_cflags_release,
			config_target,
			config_cflags,
			config_rm,
			config_program,
			config_lib,
			config_app_flat_c,
			config_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			config_platform,
			config_cxx,
			config_cxxflags_debug,
			config_cxxflags_release,
			config_target,
			config_cxxflags,
			config_rm,
			config_program,
			config_lib,
			config_app_flat_cpp,
			config_clean)
	} else {
		panic("Unknown language")
	}

	var src string
	if pr.Lang() == LANG_C {
		src = "%.c"
	} else if pr.Lang() == LANG_CPP {
		src = "%.cpp"
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(
		fmt.Sprintf(template, pr.Prog(), pr.Prog(), "%.obj", src, "%.o", src))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createAppFlat(pr *ParsingResult) {
	var suffix string
	if pr.Lang() == LANG_C {
		suffix = ".c"
	} else if pr.Lang() == LANG_CPP {
		suffix = ".cpp"
	} else {
		panic("Unknown language")
	}

	path := filepath.Join(pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var program string
	if pr.Lang() == LANG_C {
		program = program_app_c
	} else if pr.Lang() == LANG_CPP {
		program = program_app_cpp
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
