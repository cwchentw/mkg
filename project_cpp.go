package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type CppProject struct {
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

func NewCppProject(param ProjectParam) *CppProject {
	p := new(CppProject)

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

func (p *CppProject) Prog() string {
	return p.prog
}

func (p *CppProject) Path() string {
	return p.path
}

func (p *CppProject) Config() string {
	return p.config
}

func (p *CppProject) Author() string {
	return p.author
}

func (p *CppProject) Brief() string {
	return p.brief
}

func (p *CppProject) Proj() ProjectType {
	return p.proj
}

func (p *CppProject) Layout() ProjectLayout {
	return p.layout
}

func (p *CppProject) License() License {
	return p.license
}

func (p *CppProject) Src() string {
	return p.src
}

func (p *CppProject) Include() string {
	return p.include
}

func (p *CppProject) Dist() string {
	return p.dist
}

func (p *CppProject) Test() string {
	return p.test
}

func (p *CppProject) Example() string {
	return p.example
}

func (p *CppProject) Create() {
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

func (p *CppProject) createGitignore() {
	path := filepath.Join(p.Path(), ".gitignore")
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tpl := fmt.Sprintf(gitignoreCpp)
	tmpl, err := template.New("gitignore").Parse(tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = tmpl.Execute(file, struct {
		DistDir string
	}{
		p.Dist(),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CppProject) createConfigAppFlat() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CXX

	CXX_STD

	CXXFLAGS_DEBUG

	CXXFLAGS_RELEASE

	TARGET

	CXX_FLAGS

	RM

	SEP

	PROGRAM

	OBJS

	RULE_APP_CXX

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

%s
%s`

	tpl := fmt.Sprintf(config,
		MakefilePlatform,
		makefile_cxx,
		MakefileCXXStandard,
		makefile_cxxflags_debug,
		makefile_cxxflags_release,
		MakefileTarget,
		makefile_cxxflags,
		MakefileRM,
		MakefileSep,
		MakefileProgram,
		makefileObjectCpp,
		makefileExternalLibraryCpp,
		makefileAppFlatCpp,
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

func (p *CppProject) createConfigLibFlat() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CXX

	CXX_STD

	CXXFLAGS_DEBUG

	CXXFLAGS_RELEASE

	TARGET

	CXX_FLAGS

	RM

	SEP

	LIBRARY

	OBJS

	RULE_LIB_CXX

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
		MakefilePlatform,
		makefile_cxx,
		MakefileCXXStandard,
		makefile_cxxflags_debug,
		makefile_cxxflags_release,
		MakefileTarget,
		makefile_cxxflags,
		MakefileRM,
		MakefileSep,
		makefileLibCpp,
		makefileObjCppLib,
		MakefileCExtLib,
		makefileLibFlatCxx,
		makefileLibClean)

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

func (p *CppProject) createConfigAppNested() {
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

	CXX_STD

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

%s
%s`

	tpl := fmt.Sprintf(config,
		MakefilePlatform,
		makefile_cxx,
		MakefileCXXStandard,
		makefile_cxxflags_debug,
		makefile_cxxflags_release,
		MakefileTarget,
		makefile_cxxflags,
		MakefileRM,
		MakefileSep,
		makefileProjectStructure,
		MakefileProgram,
		makefileObjectCpp,
		MakefileCExtLib,
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

func (p *CppProject) createConfigLibNested() {
	path := filepath.Join(p.Path(), p.Config())
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	PLATFORM

	CXX

	CXX_STD

	CXXFLAGS_DEBUG

	CXXFLAGS_RELEASE

	TARGET

	CXX_FLAGS

	RM

	SEP

	PROJECT_STRUCTURE

	LIBRARY

	OBJECTS

	EXTERNAL_LIBRARY

	RULE_LIB_CXX

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

%s
%s`

	tpl := fmt.Sprintf(config,
		MakefilePlatform,
		makefile_cxx,
		MakefileCXXStandard,
		makefile_cxxflags_debug,
		makefile_cxxflags_release,
		MakefileTarget,
		makefile_cxxflags,
		MakefileRM,
		MakefileSep,
		makefileProjectStructure,
		makefileLibCpp,
		makefileObjCppLib,
		MakefileCExtLib,
		makefileLibNested,
		makefileLibNestedClean)

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

func (p *CppProject) createConfigAppInternal() {
	pWindows := filepath.Join(p.Path(), p.Src(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Src(), "Makefile")

	p.createConfigAppInternalImpl(pWindows)
	p.createConfigAppInternalImpl(pUnix)
}

func (p *CppProject) createConfigAppInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s`

	var app string
	if filepath.Ext(path) == ".win" {
		app = makefileInternalAppCxxWin
	} else {
		app = makefileInternalAppCxx
	}

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

func (p *CppProject) createConfigLibInternal() {
	pWindows := filepath.Join(p.Path(), p.Src(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Src(), "Makefile")

	p.createConfigLibInternalImpl(pWindows)
	p.createConfigLibInternalImpl(pUnix)
}

func (p *CppProject) createConfigLibInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s`

	var lib string
	if filepath.Ext(path) == ".win" {
		lib = makefileInternalLibCxxWin
	} else {
		lib = makefileInternalLibCxx
	}

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

func (p *CppProject) createConfigLibTestInternal() {
	pWindows := filepath.Join(p.Path(), p.Test(), "Makefile.win")
	pUnix := filepath.Join(p.Path(), p.Test(), "Makefile")

	p.createConfigLibTestInternalImpl(pWindows)
	p.createConfigLibTestInternalImpl(pUnix)
}

func (p *CppProject) createConfigLibTestInternalImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* Makefile layout
	RULE_LIB_CXX

	RULE_RM
	*/
	const config = `%s
%s`

	var lib string
	var test string
	if filepath.Ext(path) == ".win" {
		lib = makefileInternalLibTestCxxWin
		test = makefileInternalLibTestCleanWin
	} else {
		lib = makefile_internal_lib_test_cxx
		test = makefile_internal_lib_test_clean
	}

	tmpl, err := template.New("internal").Parse(
		fmt.Sprintf(config, lib, test))
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

func (p *CppProject) createApp() {
	suffix := ".cpp"

	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(p.Path(), p.Src(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createAppImpl(path)
}

func (p *CppProject) createAppImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	_, err = file.WriteString(program_app_cpp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *CppProject) createAppTest() {
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

func (p *CppProject) createAppTestImpl(path string) {
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

		tmpl, err := template.New("test").Parse(test)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		err = tmpl.Execute(file, struct {
			Program string
			DistDir string
		}{
			p.Prog(),
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

func (p *CppProject) createHeader() {
	suffix := ".hpp"

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

func (p *CppProject) createHeaderImpl(path string) {
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

func (p *CppProject) createDef() {
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

func (p *CppProject) createDefImpl(path string) {
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

func (p *CppProject) createLib() {
	suffix := ".cpp"

	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(p.Path(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	} else {
		path = filepath.Join(
			p.Path(), p.Src(), fmt.Sprintf("%s%s", p.Prog(), suffix))
	}

	p.createLibImpl(path)
}

func (p *CppProject) createLibImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl, err := template.New("program").Parse(program_lib_cpp)
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

func (p *CppProject) createLibTest() {
	suffix := ".cpp"

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

func (p *CppProject) createLibTestImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	tmpl, err := template.New("test").Parse(program_lib_test_cxx)
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
