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

	Proj     ProjectType
	Layout   ProjectLayout
	PLicense License

	Src     string
	Include string
	Dist    string
	Test    string
	Example string
}

func GetProject(pr *ParsingResult) IProject {
	param := ProjectParam{
		Program: pr.Prog(),
		Path:    pr.Path(),
		Config:  pr.Config(),

		Author: pr.Author(),
		Brief:  pr.Brief(),

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
		return NewCProject(param)
	} else if pr.Lang() == LANG_CPP {
		return NewCppProject(param)
	} else {
		panic("Unknown language")
	}
}
