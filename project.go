package main

type IProject interface {
	Prog() string
	Path() string
	Config() string

	Author() string
	Brief() string

	Proj() ProjectType
	Layout() ProjectLayout
	License() License

	Src() string
	Include() string
	Dist() string
	Test() string
	Example() string

	Create()
}

type ProjectParam struct {
	Program string
	Path    string
	Config  string

	Author string
	Brief  string

	Std      Standard
	Proj     ProjectType
	Layout   ProjectLayout
	PLicense License

	Src     string
	Include string
	Dist    string
	Test    string
	Example string
}

func NewProject(pr *ParsingResult) IProject {
	param := ProjectParam{
		Program: pr.Prog(),
		Path:    pr.Path(),
		Config:  pr.Config(),

		Author: pr.Author(),
		Brief:  pr.Brief(),

		Std:      pr.Std(),
		Proj:     pr.Proj(),
		Layout:   pr.Layout(),
		PLicense: pr.License(),

		Src:     pr.Src(),
		Include: pr.Include(),
		Dist:    pr.Dist(),
		Test:    pr.Test(),
		Example: pr.Example(),
	}

	if pr.Lang() == LANG_C {
		if !IsValidCStd(param.Std) {
			panic("Invalid C Standard")
		}

		return NewCProject(param)
	} else if pr.Lang() == LANG_CPP {
		if !IsValidCXXStd(param.Std) {
			panic("Invalid C++ Standard")
		}

		return NewCppProject(param)
	} else {
		panic("Unknown language")
	}
}
