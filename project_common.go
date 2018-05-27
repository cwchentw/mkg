package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	SEP

	PROGRAM

	OBJS

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
			config_sep,
			config_program,
			config_objects,
			config_external_library,
			config_app_flat_c,
			config_app_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			config_platform,
			config_cxx,
			config_cxxflags_debug,
			config_cxxflags_release,
			config_target,
			config_cxxflags,
			config_rm,
			config_sep,
			config_program,
			config_objects,
			config_external_library,
			config_app_flat_cpp,
			config_app_clean)
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

func createConfigLibFlat(pr *ParsingResult) {
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

	SEP

	LIBRARY

	OBJS

	RULE_LIB_C or RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s
%s
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
			config_sep,
			config_library,
			config_objects,
			config_external_library,
			config_lib_flat_c,
			config_lib_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			config_platform,
			config_cxx,
			config_cxxflags_debug,
			config_cxxflags_release,
			config_target,
			config_cxxflags,
			config_rm,
			config_sep,
			config_library,
			config_objects,
			config_external_library,
			config_lib_flat_cxx,
			config_lib_clean)
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
		fmt.Sprintf(template, pr.Prog(), "%.obj", src, "%.o", src))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createConfigAppNested(pr *ParsingResult) {
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

	SEP

	PROGRAM

	OBJS

	PROJECT_STRUCTURE

	LIBRARY

	RULE_LIB_C or RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s
%s
%s
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
			config_sep,
			config_program,
			config_objects,
			config_project_structure,
			config_external_library,
			config_app_nested,
			config_app_nested_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			config_platform,
			config_cxx,
			config_cxxflags_debug,
			config_cxxflags_release,
			config_target,
			config_cxxflags,
			config_rm,
			config_sep,
			config_program,
			config_objects,
			config_project_structure,
			config_external_library,
			config_app_nested,
			config_app_nested_clean)
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(
		fmt.Sprintf(template, pr.Prog(), pr.Prog(),
			pr.Src(), pr.Include(), pr.Dist(), pr.Test(), pr.Example()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createConfigAppInternal(pr *ParsingResult) {
	path := filepath.Join(pr.Path(), pr.Src(), "Makefile")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	RULE_LIB_C or RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s`

	var template string
	if pr.Lang() == LANG_C {
		template = fmt.Sprintf(config,
			config_internal_app_c,
			config_internal_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			config_internal_app_cxx,
			config_internal_clean)
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
		fmt.Sprintf(template, "%.obj", src, "%.o", src))
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

func createApp(pr *ParsingResult) {
	var suffix string
	if pr.Lang() == LANG_C {
		suffix = ".c"
	} else if pr.Lang() == LANG_CPP {
		suffix = ".cpp"
	} else {
		panic("Unknown language")
	}

	var path string
	if pr.Layout() == LAYOUT_FLAT {
		path = filepath.Join(pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	} else {
		path = filepath.Join(pr.Path(), pr.Src(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	}

	createAppImpl(pr, path)
}

func createAppImpl(pr *ParsingResult, path string) {
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

func createHeader(pr *ParsingResult) {
	var suffix string
	if pr.Lang() == LANG_C {
		suffix = ".h"
	} else if pr.Lang() == LANG_CPP {
		suffix = ".hpp"
	} else {
		panic("Unknown language")
	}

	var path string

	if pr.Layout() == LAYOUT_FLAT {
		path = filepath.Join(pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	} else {
		panic("Unimplemented")
	}

	createHeaderImpl(pr, path)
}

func createHeaderImpl(pr *ParsingResult, path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	header := program_header
	progUpper := strings.ToUpper(pr.Prog())

	_, err = file.WriteString(
		fmt.Sprintf(header, progUpper, progUpper, progUpper))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createLib(pr *ParsingResult) {
	var suffix string
	if pr.Lang() == LANG_C {
		suffix = ".c"
	} else if pr.Lang() == LANG_CPP {
		suffix = ".cpp"
	} else {
		panic("Unknown language")
	}

	var path string
	if pr.Layout() == LAYOUT_FLAT {
		path = filepath.Join(pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	} else {
		panic("Unimplemented")
	}

	createLibImpl(pr, path)
}

func createLibImpl(pr *ParsingResult, path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var program string
	if pr.Lang() == LANG_C {
		program = program_lib_c
	} else if pr.Lang() == LANG_CPP {
		program = program_lib_cpp
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createTest(pr *ParsingResult) {
	var path string
	if pr.Layout() == LAYOUT_FLAT {
		path = filepath.Join(pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), ".bash"))
	} else {
		panic("Unimplemented")
	}

	createTestImpl(pr, path)
}

func createTestImpl(pr *ParsingResult, path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	template := program_app_test
	_, err = file.WriteString(
		fmt.Sprintf(template, pr.Prog()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = os.Chmod(path, 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
