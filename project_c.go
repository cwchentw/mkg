package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type CProject struct {
	prog   string
	path   string
	config string

	author string
	brief  string

	proj   ProjectType
	layout ProjectLayout

	license License

	src     string
	include string
	dist    string
	test    string
	example string
}

func NewCProject(param ProjectParam) *CProject {
	p := new(CProject)

	p.prog = param.Program
	p.path = param.Path
	p.config = param.Config

	p.author = param.Author
	p.brief = param.Brief

	p.proj = param.Proj
	p.layout = param.Layout
	p.license = param.PLicense

	p.src = param.Src
	p.include = param.Include
	p.dist = param.Dist
	p.test = param.Test
	p.example = param.Example

	return p
}

func (p *CProject) Prog() string {
	return p.prog
}

func (p *CProject) Path() string {
	return p.path
}

func (p *CProject) Config() string {
	return p.config
}

func (p *CProject) Author() string {
	return p.author
}

func (p *CProject) Brief() string {
	return p.brief
}

func (p *CProject) Proj() ProjectType {
	return p.proj
}

func (p *CProject) Layout() ProjectLayout {
	return p.layout
}

func (p *CProject) License() License {
	return p.license
}

func (p *CProject) Src() string {
	return p.src
}

func (p *CProject) Include() string {
	return p.include
}

func (p *CProject) Dist() string {
	return p.dist
}

func (p *CProject) Test() string {
	return p.test
}

func (p *CProject) Example() string {
	return p.example
}

func (p *CProject) Create() {
	createLicense(p)
	createREADME(p)
	p.createGitignore()

	if p.Proj() == PROJ_CONSOLE {
		if p.Layout() == LAYOUT_FLAT {
			p.createConfigAppFlat()
			p.createApp()
			p.createAppTest()
		} else if p.Layout() == LAYOUT_NESTED {
			createProjStruct(p)
			p.createConfigAppNested()
			p.createConfigAppInternal()
			p.createApp()
			p.createAppTest()
		} else {
			panic("Unknown project layout")
		}
	} else if p.Proj() == PROJ_LIBRARY {
		if p.Layout() == LAYOUT_FLAT {
			p.createConfigLibFlat()
			p.createHeader()
			p.createDef()
			p.createLib()
			p.createLibTest()
		} else if p.Layout() == LAYOUT_NESTED {
			createProjStruct(p)
			p.createConfigLibNested()
			p.createConfigLibInternal()
			p.createConfigLibTestInternal()
			p.createHeader()
			p.createDef()
			p.createLib()
			p.createLibTest()
		} else {
			panic("Unknown project layout")
		}
	} else {
		panic("Unknown project type")
	}
}

func (p *CProject) createGitignore() {
	path := filepath.Join(p.Path(), ".gitignore")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = file.WriteString(gitignoreC)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigAppFlat() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CC

	CFLAGS_DEBUG

	CFLAGS_RELEASE

	TARGET

	CFLAGS

	RM

	SEP

	PROGRAM

	OBJS

	RULE_APP_C

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

	tpl := fmt.Sprintf(config,
		makefilePlatform,
		makefile_cc,
		makefile_cflags_debug,
		makefile_cflags_release,
		makefileTarget,
		makefile_cflags,
		makefileRm,
		makefileSep,
		makefile_program,
		makefile_objects,
		makefile_external_library,
		makefileAppFlatC,
		makefileAppClean)

	tmpl, err := template.New("appFlat").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigLibFlat() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CC

	CFLAGS_DEBUG

	CFLAGS_RELEASE

	TARGET

	CFLAGS

	RM

	SEP

	LIBRARY

	OBJS

	RULE_LIB_C

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

	tpl := fmt.Sprintf(config,
		makefilePlatform,
		makefile_cc,
		makefile_cflags_debug,
		makefile_cflags_release,
		makefileTarget,
		makefile_cflags,
		makefileRm,
		makefileSep,
		makefile_library,
		makefileObjLib,
		makefile_external_library,
		makefileLibFlatC,
		makefileLibClean)

	tmpl, err := template.New("libFlat").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigAppNested() {
	path := filepath.Join(p.Path(), p.Config())
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

	tpl := fmt.Sprintf(config,
		makefilePlatform,
		makefile_cc,
		makefile_cflags_debug,
		makefile_cflags_release,
		makefileTarget,
		makefile_cflags,
		makefileRm,
		makefileSep,
		makefileProjectStructure,
		makefile_program,
		makefile_objects,
		makefile_external_library,
		makefileAppNested,
		makefileAppNestedClean)

	tmpl, err := template.New("appNested").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program    string
		SrcDir     string
		IncludeDir string
		DistDir    string
		TestDir    string
		ExampleDir string
	}{
		p.Prog(),
		p.Src(),
		p.Include(),
		p.Dist(),
		p.Test(),
		p.Example(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigLibNested() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CC

	CFLAGS_DEBUG

	CFLAGS_RELEASE

	TARGET

	CFLAGS

	RM

	SEP

	PROJECT_STRUCTURE

	LIBRARY

	OBJECTS

	EXTERNAL_LIBRARY

	RULE_LIB_C

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

	tpl := fmt.Sprintf(config,
		makefilePlatform,
		makefile_cc,
		makefile_cflags_debug,
		makefile_cflags_release,
		makefileTarget,
		makefile_cflags,
		makefileRm,
		makefileSep,
		makefileProjectStructure,
		makefile_library,
		makefileObjLib,
		makefile_external_library,
		makefileLibNested,
		makefileLibNestedClean)

	tmpl, err := template.New("libNested").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program    string
		SrcDir     string
		IncludeDir string
		DistDir    string
		TestDir    string
		ExampleDir string
	}{
		p.Prog(),
		p.Src(),
		p.Include(),
		p.Dist(),
		p.Test(),
		p.Example(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigAppInternal() {
	pWindows := filepath.Join(p.Path(), p.Src(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Src(), "Makefile")

	p.createConfigAppInternalImpl(pWindows)
	p.createConfigAppInternalImpl(pUnix)
}

func (p *CProject) createConfigAppInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var app string
	if filepath.Ext(path) == ".win" {
		app = makefileInternalAppCWin
	} else {
		app = makefileInternalAppC
	}

	/* Makefile layout
	RULE_APP_C

	RULE_RM
	*/
	const config = `%s
%s`

	tmpl, err := template.New("internal").Parse(
		fmt.Sprintf(config, app,
			makefileInternalClean))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigLibInternal() {
	pWindows := filepath.Join(p.Path(), p.Src(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Src(), "Makefile")

	p.createConfigLibInternalImpl(pWindows)
	p.createConfigLibInternalImpl(pUnix)
}

func (p *CProject) createConfigLibInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var lib string
	if filepath.Ext(path) == ".win" {
		lib = makefileInternalLibCWin
	} else {
		lib = makefileInternalLibC
	}

	/* Makefile layout
	RULE_LIB_C

	RULE_RM
	*/
	const config = `%s
%s`

	tmpl, err := template.New("internal").Parse(
		fmt.Sprintf(config, lib,
			makefileInternalClean))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createConfigLibTestInternal() {
	pWindows := filepath.Join(p.Path(), p.Test(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Test(), "Makefile")

	p.createConfigLibTestInternalImpl(pWindows)
	p.createConfigLibTestInternalImpl(pUnix)
}

func (p *CProject) createConfigLibTestInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	RULE_LIB_C

	RULE_RM
	*/
	const config = `%s
%s`

	var lib string
	var clean string

	if filepath.Ext(path) == ".win" {
		lib = makefileInternalLibTestCWin
		clean = makefileInternalLibTestCleanWin
	} else {
		lib = makefile_internal_lib_test_c
		clean = makefile_internal_lib_test_clean
	}

	tmpl, err := template.New("internal").Parse(
		fmt.Sprintf(config, lib, clean))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createApp() {
	suffix := ".c"

	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(p.Path(), p.Src(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createAppImpl(path)
}

func (p *CProject) createAppImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	program := program_app_c

	_, err = file.WriteString(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createHeader() {
	suffix := ".h"

	var path string

	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(
			p.Path(), p.Include(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createHeaderImpl(path)
}

func (p *CProject) createHeaderImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	header := program_header
	progUpper := strings.ToUpper(p.Prog())

	tmpl, err := template.New("header").Parse(header)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		progUpper,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createDef() {
	suffix := ".def"

	var path string

	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(
			p.Path(), p.Src(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createDefImpl(path)
}

func (p *CProject) createDefImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl, err := template.New("def").Parse(programLibDef)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createLib() {
	suffix := ".c"

	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(
			p.Path(), p.Src(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createLibImpl(path)
}

func (p *CProject) createLibImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl, err := template.New("program").Parse(program_lib_c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createAppTest() {

	var pWindows string
	var pUnix string
	if p.Layout() == LAYOUT_FLAT {
		pWindows = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s", p.Prog(), ".vbs"))
		pUnix = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s", p.Prog(), ".bash"))
	} else {
		pWindows = filepath.Join(
			p.Path(), p.Test(), fmt.Sprintf("%s%s", p.Prog(), ".vbs"))
		pUnix = filepath.Join(
			p.Path(), p.Test(), fmt.Sprintf("%s%s", p.Prog(), ".bash"))
	}

	p.createAppTestImpl(pWindows)
	p.createAppTestImpl(pUnix)
}

func (p *CProject) createAppTestImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if p.Layout() == LAYOUT_FLAT {
		var test string
		if filepath.Ext(path) == ".vbs" {
			test = programAppTestWin
		} else {
			test = programAppTest
		}

		tmpl, err := template.New("test").Parse(test)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		err = tmpl.Execute(file, struct {
			Program string
		}{
			p.Prog(),
		})
	} else if p.Layout() == LAYOUT_NESTED {
		var test string
		if filepath.Ext(path) == ".vbs" {
			test = programAppTestNestedWin
		} else {
			test = programAppTestNested
		}

		var prog string
		if filepath.Ext(path) == ".vbs" {
			prog = fmt.Sprintf("%s%s", p.Prog(), ".exe")
		} else {
			prog = p.Prog()
		}

		tmpl, err := template.New("test").Parse(test)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = tmpl.Execute(file, struct {
			Program string
			DistDir string
		}{
			prog,
			p.Dist(),
		})
	} else {
		panic("Unknown project layout")
	}

	err = os.Chmod(path, 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CProject) createLibTest() {
	suffix := ".c"

	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s%s", p.Prog(), "_test", suffix))
	} else {
		path = filepath.Join(
			p.Path(), p.Test(), fmt.Sprintf("%s%s%s", p.Prog(), "_test", suffix))
	}

	p.createLibTestImpl(path)
}

func (p *CProject) createLibTestImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl, err := template.New("test").Parse(program_lib_test_c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		Program string
	}{
		p.Prog(),
	})
}
