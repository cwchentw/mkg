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
			return PARSING_FSS_EVENT_VERSION, nil
		case "-h", "--help":
			return PARSING_FSS_EVENT_HELP, nil
		case "--licenses":
			return PARSING_FSS_EVENT_LICENSES, nil
		case "--standards":
			return PARSING_FSS_EVNET_STANDARDS, nil
		case "--custom":
			return PARSING_FSS_EVENT_ERROR, errors.New("--custom should be the first argument")
		case "-f", "--force":
			r.SetForced(true)
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
		case "-std", "--standard":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid standard")
			}

			if args[i+1] == "ansi" {
				if r.Lang() == LANG_C {
					r.SetStd(STD_C89)
				} else if r.Lang() == LANG_CPP {
					r.SetStd(STD_CXX98)
				} else {
					return PARSING_FSS_EVENT_ERROR, errors.New("Invalid standard")
				}

				i++
				continue
			}

			std, err := stringToStd(args[i+1])
			if err != nil {
				return PARSING_FSS_EVENT_ERROR, err
			}

			r.SetStd(std)
			i++
		case "-p", "--program":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid program")
			}

			r.SetProg(args[i+1])
			setProg = true
			i++
		case "-a", "--author":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid author")
			}

			r.SetAuthor(args[i+1])
			i++
		case "-b", "--brief":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid description")
			}

			r.SetBrief(args[i+1])
			i++
		case "-l", "--license":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid license")
			}

			l, err := reprToLicense(args[i+1])
			if err != nil {
				return PARSING_FSS_EVENT_ERROR, err
			}

			r.SetLicense(l)
			i++
		case "-s", "--source":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid source directory name")
			}

			if r.IsNested() {
				r.SetSrc(args[i+1])
			}

			i++
		case "-i", "--include":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid include directory name")
			}

			if r.IsNested() {
				r.SetInclude(args[i+1])
			}

			i++
		case "-d", "--dist":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid dist directory name")
			}

			if r.IsNested() {
				r.SetDist(args[i+1])
			}

			i++
		case "-t", "--test":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid test directory name")
			}

			if r.IsNested() {
				r.SetTest(args[i+1])
			}

			i++
		case "-e", "--example":
			if i+1 >= len(args) {
				return PARSING_FSS_EVENT_ERROR, errors.New("No valid example directory name")
			}

			if r.IsNested() {
				r.SetExample(args[i+1])
			}

			i++
		default:
			if !isValidPath(args[i]) {
				return PARSING_FSS_EVENT_ERROR, errors.New("Invalid path")
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

	// Auto-correct the language standard.
	if r.Lang() == LANG_CPP && r.Std() == STD_C99 {
		r.SetStd(STD_CXX11)
	}

	return PARSING_FSS_EVENT_RUN, nil
}
