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

type License int

const (
	LICENSE_NONE License = iota
	LICENSE_APACHE2
	LICENSE_BSD2
	LICENSE_BSD3
	LICENSE_MIT
	LICENSE_GPL2
	LICENSE_GPL3
	LICENSE_AGPL3
	LICENSE_LGPL2
	LICENSE_LGPL3
	LICENSE_EPL2
	LICENSE_MPL2
	LICENSE_UNLICENSE
)
