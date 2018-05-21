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

	src     string
	include string
	test    string
	example string
}

func NewParsingResult() *ParsingResult {
	result := new(ParsingResult)

	result.prog = "hello"
	result.path = result.prog
	result.config = "Makefile"

	result.lang = LANG_C
	result.proj = PROJ_APP
	result.layout = LAYOUT_NESTED

	result.src = "src"
	result.include = "include"
	result.test = "test"
	result.example = "examples"

	return result
}

func (r *ParsingResult) String() string {
	return fmt.Sprintf(`Program name: %s
Project path: %s
Project language: %s
Project type: %s
Project layout: %s
Project source directory: %s
Project include directory: %s
Project test directory: %s
Project example directory: %s
`, r.Prog(), r.Path(), langToString(r.Lang()), projToString(r.Proj()),
		layoutToString(r.Layout()), r.Src(), r.Include(), r.Test(), r.Example())
}

func langToString(lang Language) string {
	switch lang {
	case LANG_C:
		return "LANG_C"
	case LANG_CPP:
		return "LANG_C++"
	default:
		panic("Unknown language")
	}
}

func projToString(proj ProjectType) string {
	switch proj {
	case PROJ_APP:
		return "application"
	case PROJ_LIB:
		return "library"
	default:
		panic("Unknown project type")
	}
}

func layoutToString(layout ProjectLayout) string {
	switch layout {
	case LAYOUT_NESTED:
		return "nested"
	case LAYOUT_FLAT:
		return "flat"
	default:
		panic("Unknown layout")
	}
}

func (r *ParsingResult) ParseArgument(args []string) error {
	return nil
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

func isValidFileName(name string) bool {
	// Modify it later.
	return name != ""
}

func isValidPath(path string) bool {
	// Modify it later.
	return path != ""
}
