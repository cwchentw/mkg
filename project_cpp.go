package main

import (
	"fmt"
	"os"
	"path/filepath"
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
		// Implement it later.
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

	_, err = file.WriteString(gitignore_cpp)
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
%s`

	tpl := fmt.Sprintf(config,
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

	CXX

	CXXFLAGS_DEBUG

	CXXFLAGS_RELEASE

	TARGET

	CXX_FLAGS

	RM

	SEP

	PROJECT_STRUCTURE

	PROGRAM

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
%s`

	tpl := fmt.Sprintf(config,
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

	tpl := fmt.Sprintf(config,
		makefile_platform,
		makefile_cxx,
		makefile_cxxflags_debug,
		makefile_cxxflags_release,
		makefile_target,
		makefile_cxxflags,
		makefile_rm,
		makefile_sep,
		makefile_project_structure,
		makefile_library,
		makefile_objects,
		makefile_external_library,
		makefile_lib_nested,
		makefile_lib_nested_clean)

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
	var path string
	if p.Layout() == LAYOUT_FLAT {
		path = filepath.Join(
			p.Path(), fmt.Sprintf("%s%s", p.Prog(), ".bash"))
	} else {
		path = filepath.Join(
			p.Path(), p.Test(), fmt.Sprintf("%s%s", p.Prog(), ".bash"))
	}

	p.createAppTestImpl(path)
}

func (p *CppProject) createAppTestImpl(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if p.Layout() == LAYOUT_FLAT {
		tmpl, err := template.New("test").Parse(program_app_test)
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
		tmpl, err := template.New("test").Parse(program_app_test_nested)
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
