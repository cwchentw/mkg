package main

type Language int

const (
	LANG_C Language = iota
	LANG_CPP
)

type ProjectType int

const (
	PROJ_APP ProjectType = iota
	PROJ_LIB
)

type ProjectLayout int

const (
	LAYOUT_NESTED ProjectLayout = iota
	LAYOUT_FLAT
)
