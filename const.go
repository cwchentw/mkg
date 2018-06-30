package main

const PROGRAM = "mkg"
const VERSION = "0.1.1"

type Language int

const (
	LANG_C Language = iota
	LANG_CPP
)

type ProjectType int

const (
	PROJ_CONSOLE ProjectType = iota
	PROJ_LIBRARY
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

type ParsingEvent int

const (
	PARSING_EVENT_VERSION ParsingEvent = iota
	PARSING_EVENT_HELP
	PARSING_EVENT_LICENSES
	PARSING_EVENT_RUN
	PARSING_EVENT_ERROR
)
