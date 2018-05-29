package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
			makefile_platform,
			makefile_cc,
			makefile_cflags_debug,
			makefile_cflags_release,
			makefile_target,
			makefile_cflags,
			makefile_rm,
			makefile_sep,
			makefile_program,
			makefile_objects,
			makefile_external_library,
			makefile_app_flat_c,
			makefile_app_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			makefile_platform,
			makefile_cxx,
			makefile_cxxflags_debug,
			makefile_cxxflags_release,
			makefile_target,
			makefile_cxxflags,
			makefile_rm,
			makefile_sep,
			makefile_program,
			makefile_objects,
			makefile_external_library,
			makefile_app_flat_cpp,
			makefile_app_clean)
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
			makefile_platform,
			makefile_cc,
			makefile_cflags_debug,
			makefile_cflags_release,
			makefile_target,
			makefile_cflags,
			makefile_rm,
			makefile_sep,
			makefile_library,
			makefile_objects,
			makefile_external_library,
			makefile_lib_flat_c,
			makefile_lib_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			makefile_platform,
			makefile_cxx,
			makefile_cxxflags_debug,
			makefile_cxxflags_release,
			makefile_target,
			makefile_cxxflags,
			makefile_rm,
			makefile_sep,
			makefile_library,
			makefile_objects,
			makefile_external_library,
			makefile_lib_flat_cxx,
			makefile_lib_clean)
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

	PROJECT_STRUCTURE

	PROGRAM

	OBJECTS

	EXTERNAL_LIBRARY

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
			makefile_platform,
			makefile_cc,
			makefile_cflags_debug,
			makefile_cflags_release,
			makefile_target,
			makefile_cflags,
			makefile_rm,
			makefile_sep,
			makefile_project_structure,
			makefile_program,
			makefile_objects,
			makefile_external_library,
			makefile_app_nested,
			makefile_app_nested_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			makefile_platform,
			makefile_cxx,
			makefile_cxxflags_debug,
			makefile_cxxflags_release,
			makefile_target,
			makefile_cxxflags,
			makefile_rm,
			makefile_sep,
			makefile_project_structure,
			makefile_program,
			makefile_objects,
			makefile_external_library,
			makefile_app_nested,
			makefile_app_nested_clean)
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(
		fmt.Sprintf(template,
			pr.Src(), pr.Include(), pr.Dist(), pr.Test(), pr.Example(),
			pr.Prog(), pr.Prog()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createConfigLibNested(pr *ParsingResult) {
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

	PROJECT_STRUCTURE

	LIBRARY

	OBJECTS

	EXTERNAL_LIBRARY

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
			makefile_platform,
			makefile_cc,
			makefile_cflags_debug,
			makefile_cflags_release,
			makefile_target,
			makefile_cflags,
			makefile_rm,
			makefile_sep,
			makefile_project_structure,
			makefile_library,
			makefile_objects,
			makefile_external_library,
			makefile_lib_nested,
			makefile_lib_nested_clean)
	} else if pr.Lang() == LANG_CPP {
		panic("Unimplemented")
	} else {
		panic("Unknown language")
	}

	_, err = file.WriteString(
		fmt.Sprintf(template,
			pr.Src(), pr.Include(), pr.Dist(), pr.Test(), pr.Example(),
			pr.Prog()))
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
			makefile_internal_app_c,
			makefile_internal_clean)
	} else if pr.Lang() == LANG_CPP {
		template = fmt.Sprintf(config,
			makefile_internal_app_cxx,
			makefile_internal_clean)
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
		fmt.Sprintf(template, "%.o", src))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createConfigLibInternal(pr *ParsingResult) {
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
			makefile_internal_lib_c,
			makefile_internal_clean)
	} else if pr.Lang() == LANG_CPP {
		panic("Unimplemented")
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
		fmt.Sprintf(template, "%.o", src))
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
		path = filepath.Join(
			pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
	} else {
		path = filepath.Join(
			pr.Path(), pr.Include(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
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
		path = filepath.Join(
			pr.Path(), pr.Src(), fmt.Sprintf("%s%s", pr.Prog(), suffix))
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

func createAppTest(pr *ParsingResult) {
	var path string
	if pr.Layout() == LAYOUT_FLAT {
		path = filepath.Join(
			pr.Path(), fmt.Sprintf("%s%s", pr.Prog(), ".bash"))
	} else {
		path = filepath.Join(
			pr.Path(), pr.Test(), fmt.Sprintf("%s%s", pr.Prog(), ".bash"))
	}

	createAppTestImpl(pr, path)
}

func createAppTestImpl(pr *ParsingResult, path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var test string
	if pr.Layout() == LAYOUT_FLAT {
		test = fmt.Sprintf(program_app_test, pr.Prog())
	} else if pr.Layout() == LAYOUT_NESTED {
		test = fmt.Sprintf(program_app_test_nested, pr.Prog(), pr.Dist())
	} else {
		panic("Unknown project layout")
	}
	_, err = file.WriteString(test)
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
