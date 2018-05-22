package main

import (
	"errors"
	"path/filepath"
)

func (r *ParsingResult) ParseArgument(args []string) (ParsingEvent, error) {
	size := len(args)
	setProg := false
	for i := 1; i < size; i++ {
		switch args[i] {
		case "-v", "--version":
			return PARSING_EVENT_VERSION, nil
		case "-h", "--help":
			return PARSING_EVENT_HELP, nil
		case "--licenses":
			return PARSING_EVENT_LICENSES, nil
		case "-p", "--program":
			r.SetProg(args[i])
			setProg = true
		case "-c", "-C":
			r.SetLang(LANG_C)
		case "-cpp", "-cxx":
			r.SetLang(LANG_CPP)
		case "-app", "--application":
			r.SetProj(PROJ_APP)
		case "-lib", "--library":
			r.SetProj(PROJ_LIB)
		case "--nested":
			r.SetLayout(LAYOUT_NESTED)
		case "--flat":
			r.SetLayout(LAYOUT_FLAT)
		case "-l", "--license":
			if i+1 >= len(args) {
				return PARSING_EVENT_ERROR, errors.New("No valid license")
			}

			l, err := stringToLicense(args[i+1])
			if err != nil {
				return PARSING_EVENT_ERROR, err
			}

			r.SetLicense(l)
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
		}
	}

	return PARSING_EVENT_RUN, nil
}
