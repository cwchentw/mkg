package main

import (
	"errors"
)

func langToString(lang Language) string {
	switch lang {
	case LANG_C:
		return "C"
	case LANG_CPP:
		return "C++"
	default:
		panic("Unknown language")
	}
}

func stringToLang(lang string) (Language, error) {
	switch lang {
	case "c":
		return LANG_C, nil
	case "cpp":
		return LANG_CPP, nil
	default:
		return LANG_C, errors.New("Invalid language")
	}
}

func projToString(proj ProjectType) string {
	switch proj {
	case PROJ_APP:
		return "application"
	case PROJ_LIB:
		return "library"
	default:
		panic("Unknown project type")
	}
}

func layoutToString(layout ProjectLayout) string {
	switch layout {
	case LAYOUT_NESTED:
		return "nested"
	case LAYOUT_FLAT:
		return "flat"
	default:
		panic("Unknown layout")
	}
}

func licenseToString(license License) string {
	switch license {
	case LICENSE_NONE:
		return "none"
	case LICENSE_APACHE2:
		return "apache2"
	case LICENSE_MIT:
		return "mit"
	case LICENSE_GPL3:
		return "gpl3"
	case LICENSE_BSD2:
		return "bsd2"
	case LICENSE_BSD3:
		return "bsd3"
	case LICENSE_EPL2:
		return "epl2"
	case LICENSE_AGPL3:
		return "agpl3"
	case LICENSE_GPL2:
		return "gpl2"
	case LICENSE_LGPL2:
		return "lgpl2"
	case LICENSE_LGPL3:
		return "lgpl3"
	case LICENSE_MPL2:
		return "mpl2"
	case LICENSE_UNLICENSE:
		return "unlicense"
	default:
		panic("Unknown license")
	}
}

func isValidFileName(name string) bool {
	// Modify it later.
	return name != ""
}

func isValidPath(path string) bool {
	// Modify it later.
	return path != ""
}
