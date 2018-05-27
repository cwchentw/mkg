package main

import (
	"errors"
	"path/filepath"
)

func (r *ParsingResult) ParseArgument(args []string) (ParsingEvent, error) {
	size := len(args)
	setProg := false
	setPath := false
	for i := 1; i < size; i++ {
		switch args[i] {
		case "-v", "--version":
			return PARSING_EVENT_VERSION, nil
		case "-h", "--help":
			return PARSING_EVENT_HELP, nil
		case "--licenses":
			return PARSING_EVENT_LICENSES, nil
		case "--custom":
			return PARSING_EVENT_ERROR, errors.New("--custom should be the first argument")
		case "-f", "--force":
			r.SetForced(true)
		case "-p", "--program":
			r.SetProg(args[i])
			setProg = true
		case "-c", "-C":
			r.SetLang(LANG_C)
		case "-cpp", "-cxx":
			r.SetLang(LANG_CPP)
		case "--console":
			r.SetProj(PROJ_CONSOLE)
		case "--library":
			r.SetProj(PROJ_LIBRARY)
		case "--nested":
			r.SetLayout(LAYOUT_NESTED)
		case "--flat":
			r.SetLayout(LAYOUT_FLAT)
		case "-a", "--author":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid author")
			}

			r.SetAuthor(args[i+1])
			i++
		case "-b", "--brief":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid description")
			}

			r.SetBrief(args[i+1])
			i++
		case "-l", "--license":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid license")
			}

			l, err := stringToLicense(args[i+1])
			if err != nil {
				return PARSING_EVENT_ERROR, err
			}

			r.SetLicense(l)
			i++
		case "-s", "--source":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid source directory name")
			}

			if r.IsNested() {
				r.SetSrc(args[i+1])
			}

			i++
		case "-i", "--include":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid include directory name")
			}

			if r.IsNested() {
				r.SetInclude(args[i+1])
			}

			i++
		case "-d", "--dist":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid dist directory name")
			}

			if r.IsNested() {
				r.SetDist(args[i+1])
			}

			i++
		case "-t", "--test":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid test directory name")
			}

			if r.IsNested() {
				r.SetTest(args[i+1])
			}

			i++
		case "-e", "--example":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid example directory name")
			}

			if r.IsNested() {
				r.SetExample(args[i+1])
			}

			i++
		default:
			if !isValidPath(args[i]) {
				return PARSING_EVENT_ERROR, errors.New("Invalid path")
			}

			r.SetPath(args[i])

			if !setProg {
				r.SetProg(filepath.Base(r.Path()))
			}

			setPath = true
		}

		if setPath {
			break
		}
	}

	return PARSING_EVENT_RUN, nil
}
