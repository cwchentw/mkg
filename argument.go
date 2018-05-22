package main

import (
	"errors"
	"fmt"
)

type ParsingResult struct {
	prog   string
	path   string
	config string

	lang   Language
	proj   ProjectType
	layout ProjectLayout

	license License

	src     string
	include string
	test    string
	example string
}

func NewParsingResult() *ParsingResult {
	result := new(ParsingResult)

	result.prog = "myapp"
	result.path = result.prog
	result.config = "Makefile"

	result.lang = LANG_C
	result.proj = PROJ_APP
	result.layout = LAYOUT_NESTED

	result.license = LICENSE_NONE

	result.src = "src"
	result.include = "include"
	result.test = "test"
	result.example = "examples"

	return result
}

func (r *ParsingResult) String() string {
	out := fmt.Sprintf(`Program name: %s
Project path: %s
Project language: %s
Project type: %s
Project license: %s
Project layout: %s
`, r.Prog(), r.Path(), langToString(r.Lang()), projToString(r.Proj()),
		licenseToString(r.License()), layoutToString(r.Layout()),
	)

	if !r.IsNested() {
		return out
	}

	more := fmt.Sprintf(`Project source directory: %s
Project include directory: %s
Project test directory: %s
Project example directory: %s
`, r.Src(), r.Include(), r.Test(), r.Example())

	return fmt.Sprintf("%s%s", out, more)
}

func (r *ParsingResult) Prog() string {
	return r.prog
}

func (r *ParsingResult) SetProg(prog string) error {
	if !isValidFileName(prog) {
		return errors.New("Invalid program name")
	}

	r.prog = prog

	return nil
}

func (r *ParsingResult) Path() string {
	return r.path
}

func (r *ParsingResult) SetPath(path string) error {
	if !isValidPath(path) {
		return errors.New("Invalid project path")
	}

	r.path = path

	return nil
}

func (r *ParsingResult) Config() string {
	return r.config
}

func (r *ParsingResult) SetConfig(config string) error {
	if !isValidFileName(config) {
		return errors.New("Invalid config file name")
	}

	r.config = config

	return nil
}

func (r *ParsingResult) Lang() Language {
	return r.lang
}

func (r *ParsingResult) SetLang(lang Language) {
	r.lang = lang
}

func (r *ParsingResult) Proj() ProjectType {
	return r.proj
}

func (r *ParsingResult) SetProj(proj ProjectType) {
	r.proj = proj
}

func (r *ParsingResult) Layout() ProjectLayout {
	return r.layout
}

func (r *ParsingResult) SetLayout(layout ProjectLayout) {
	r.layout = layout
}

func (r *ParsingResult) IsNested() bool {
	return r.layout == LAYOUT_NESTED
}

func (r *ParsingResult) License() License {
	return r.license
}

func (r *ParsingResult) SetLicense(license License) {
	r.license = license
}

func (r *ParsingResult) Src() string {
	return r.src
}

func (r *ParsingResult) SetSrc(src string) error {
	if !isValidPath(src) {
		return errors.New("Invalid source path")
	}

	r.src = src

	return nil
}

func (r *ParsingResult) Include() string {
	return r.include
}

func (r *ParsingResult) SetInclude(include string) error {
	if !isValidPath(include) {
		return errors.New("Invalid include path")
	}

	r.include = include

	return nil
}

func (r *ParsingResult) Test() string {
	return r.test
}

func (r *ParsingResult) SetTest(test string) error {
	if !isValidPath(test) {
		return errors.New("Invalid test path")
	}

	r.test = test

	return nil
}

func (r *ParsingResult) Example() string {
	return r.example
}

func (r *ParsingResult) SetExample(ex string) error {
	if !isValidPath(ex) {
		return errors.New("Invalid example path")
	}

	r.example = ex

	return nil
}
