package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func (r *ParsingResult) RunWithDefaults() error {
	prog, err := prompt(fmt.Sprintf("Program name [%s]: ", r.Prog()))
	if err != nil {
		return err
	}

	if prog != "" {
		err = r.SetProg(prog)
		if err != nil {
			return err
		}

		err = r.SetPath(prog)
		if err != nil {
			return err
		}
	}

	path, err := prompt(fmt.Sprintf("Project path [%s]: ", r.Path()))
	if err != nil {
		return err
	}

	if path != "" {
		err = r.SetPath(path)
		if err != nil {
			return err
		}
	}

	lang, err := prompt(
		fmt.Sprintf("Project language (c/cpp) [%s]: ", langToString(r.Lang())))
	if err != nil {
		return err
	}

	if lang == "c" {
		r.SetLang(C)
	} else if lang == "cpp" {
		r.SetLang(Cpp)
	} else if lang == "" {
		// Do nothing.
	} else {
		return errors.New("Invalid language")
	}

	return nil
}

func (r *ParsingResult) Run() error {
	err := r.RunWithDefaults()
	if err != nil {
		return err
	}

	return nil
}

func prompt(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	out, err := reader.ReadString('\n')
	if err != nil {
		return out, err
	}

	if runtime.GOOS == "windows" {
		out = strings.TrimSuffix(out, "\r\n")
	} else {
		out = strings.TrimSuffix(out, "\n")
	}

	return out, err
}
