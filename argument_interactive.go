package main

import (
	"bufio"
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

	if lang != "" {
		l, err := stringToLang(lang)
		if err != nil {
			return err
		}

		r.SetLang(l)
	}

	proj, err := prompt(
		fmt.Sprintf("Project type (app/lib) [%s]: ", projToString(r.Proj())))
	if err != nil {
		return err
	}

	if proj != "" {
		p, err := stringToProj(proj)
		if err != nil {
			return err
		}

		r.SetProj(p)
	}

	fmt.Println("")

	printLicenses()

	fmt.Println("")

	cert, err := prompt(
		fmt.Sprintf("Project license [%s]: ", licenseToString(r.License())))
	if err != nil {
		return err
	}

	if cert != "" {
		c, err := stringToLicense(cert)
		if err != nil {
			return err
		}

		r.SetLicense(c)
	}

	return nil
}

func (r *ParsingResult) Run() error {
	err := r.RunWithDefaults()
	if err != nil {
		return err
	}

	layout, err := prompt(
		fmt.Sprintf("Project layout (nested/flat) [%s]: ", layoutToString(r.Layout())))
	if err != nil {
		return err
	}

	if layout != "" {
		l, err := stringToLayout(layout)
		if err != nil {
			return err
		}

		r.SetLayout(l)
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
