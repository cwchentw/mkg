package main

import (
	"errors"
	"fmt"
)

type Language int

const (
	C Language = iota
	Cpp
)

type ProjectType int

const (
	Application ProjectType = iota
	Library
)

type ProjectLayout int

const (
	Nested ProjectLayout = iota
	Flat
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

	result.lang = C
	result.proj = Application
	result.layout = Nested

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
`, r.prog, r.path, langToString(r.lang), projToString(r.proj),
		layoutToString(r.layout), r.src, r.include, r.test, r.example)
}

func langToString(lang Language) string {
	switch lang {
	case C:
		return "C"
	case Cpp:
		return "C++"
	default:
		panic("Unknown language")
	}
}

func projToString(proj ProjectType) string {
	switch proj {
	case Application:
		return "application"
	case Library:
		return "library"
	default:
		panic("Unknown project type")
	}
}

func layoutToString(layout ProjectLayout) string {
	switch layout {
	case Nested:
		return "nested"
	case Flat:
		return "flat"
	default:
		panic("Unknown layout")
	}
}

func (r *ParsingResult) ParseArgument(args []string) *ParsingResult {
	return r
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
