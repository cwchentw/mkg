package main

import (
	"errors"
	"strings"
)

func langToString(lang Language) string {
	switch lang {
	case LANG_C:
		return "c"
	case LANG_CPP:
		return "cpp"
	default:
		panic("Unknown language")
	}
}

func stdToString(std Standard) string {
	switch std {
	case STD_C89:
		return "c89"
	case STD_C99:
		return "c99"
	case STD_C11:
		return "c11"
	case STD_C17:
		return "c17"
	case STD_C_GNU89:
		return "gnu89"
	case STD_C_GNU99:
		return "gnu99"
	case STD_C_GNU11:
		return "gnu11"
	case STD_C_GNU17:
		return "gnu17"
	case STD_CXX98:
		return "c++98"
	case STD_CXX11:
		return "c++11"
	case STD_CXX14:
		return "c++14"
	case STD_CXX17:
		return "c++17"
	case STD_CXX_GNU98:
		return "gnu++98"
	case STD_CXX_GNU11:
		return "gnu++11"
	case STD_CXX_GNU14:
		return "gnu++14"
	case STD_CXX_GNU17:
		return "gnu++17"
	default:
		panic("Unknown standard")
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

func stringToStd(std string) (Standard, error) {
	switch std {
	case "c89", "c90":
		return STD_C89, nil
	case "c99":
		return STD_C99, nil
	case "c11":
		return STD_C11, nil
	case "c17", "c18":
		return STD_C17, nil
	case "gnu89", "gnu90":
		return STD_C_GNU89, nil
	case "gnu99":
		return STD_C_GNU99, nil
	case "gnu11":
		return STD_C_GNU11, nil
	case "gnu17", "gnu18":
		return STD_C_GNU17, nil
	case "c++98", "c++03":
		return STD_CXX98, nil
	case "c++11":
		return STD_CXX11, nil
	case "c++14":
		return STD_CXX14, nil
	case "c++17":
		return STD_CXX17, nil
	case "gnu++98", "gnu++03":
		return STD_CXX_GNU98, nil
	case "gnu++11":
		return STD_CXX_GNU11, nil
	case "gnu++14":
		return STD_CXX_GNU14, nil
	case "gnu++17":
		return STD_CXX_GNU17, nil
	default:
		return STD_C99, errors.New("Invalid standard")
	}
}

func stdToStringWin(std Standard) string {
	switch std {
	case STD_CXX98:
		fallthrough
	case STD_CXX11:
		fallthrough
	case STD_CXX14:
		fallthrough
	case STD_CXX_GNU98:
		fallthrough
	case STD_CXX_GNU11:
		fallthrough
	case STD_CXX_GNU14:
		return "c++14"
	case STD_CXX17:
		fallthrough
	case STD_CXX_GNU17:
		return "c++17"
	}

	panic("Invalid C++ standard for Visual C++")
}

func projToString(proj ProjectType) string {
	switch proj {
	case PROJ_CONSOLE:
		return "console"
	case PROJ_LIBRARY:
		return "library"
	default:
		panic("Unknown project type")
	}
}

func stringToProj(proj string) (ProjectType, error) {
	switch proj {
	case "console":
		return PROJ_CONSOLE, nil
	case "library":
		return PROJ_LIBRARY, nil
	default:
		return PROJ_CONSOLE, errors.New("Invalid project type")
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

func stringToLayout(layout string) (ProjectLayout, error) {
	switch layout {
	case "nested":
		return LAYOUT_NESTED, nil
	case "flat":
		return LAYOUT_FLAT, nil
	default:
		return LAYOUT_NESTED, errors.New("Invalid project layout")
	}
}

func licenseToRepr(license License) string {
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

func reprToLicense(cert string) (License, error) {
	switch cert {
	case "none":
		return LICENSE_NONE, nil
	case "apache2":
		return LICENSE_APACHE2, nil
	case "mit":
		return LICENSE_MIT, nil
	case "gpl3":
		return LICENSE_GPL3, nil
	case "bsd2":
		return LICENSE_BSD2, nil
	case "bsd3":
		return LICENSE_BSD3, nil
	case "epl2":
		return LICENSE_EPL2, nil
	case "agpl3":
		return LICENSE_AGPL3, nil
	case "gpl2":
		return LICENSE_GPL2, nil
	case "lgpl2":
		return LICENSE_LGPL2, nil
	case "lgpl3":
		return LICENSE_LGPL3, nil
	case "mpl2":
		return LICENSE_MPL2, nil
	case "unlicense":
		return LICENSE_UNLICENSE, nil
	default:
		return LICENSE_MIT, errors.New("Invalid license")
	}
}

func licenseToString(license License) string {
	switch license {
	case LICENSE_NONE:
		return "None"
	case LICENSE_APACHE2:
		return "Apache 2.0"
	case LICENSE_MIT:
		return "MIT"
	case LICENSE_GPL3:
		return "GPL 3.0"
	case LICENSE_BSD2:
		return "BSD 2.0"
	case LICENSE_BSD3:
		return "BSD 3.0"
	case LICENSE_EPL2:
		return "EPL 2"
	case LICENSE_AGPL3:
		return "AGPL 3.0"
	case LICENSE_GPL2:
		return "GPL 2.1"
	case LICENSE_LGPL2:
		return "LGPL 2.1"
	case LICENSE_LGPL3:
		return "LGPL 3.0"
	case LICENSE_MPL2:
		return "MPL 2.0"
	case LICENSE_UNLICENSE:
		return "Unlicense"
	default:
		panic("Unknown license")
	}
}

func getTemplate(license License) string {
	switch license {
	case LICENSE_NONE:
		return ""
	case LICENSE_APACHE2:
		return LicenseApache2
	case LICENSE_MIT:
		return LicenseMIT
	case LICENSE_BSD2:
		return LicenseBSD2
	case LICENSE_BSD3:
		return LicenseBSD3
	case LICENSE_EPL2:
		return LicenseEPL2
	case LICENSE_GPL2:
		return LicenseGPL2
	case LICENSE_GPL3:
		return LicenseGPL3
	case LICENSE_AGPL3:
		return LicenseAGPL3
	case LICENSE_LGPL2:
		return LicenseLGPL2
	case LICENSE_LGPL3:
		return LicenseLGPL3
	case LICENSE_MPL2:
		return LicenseMPL2
	case LICENSE_UNLICENSE:
		return LicenseUnlicense
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
	if path == "" {
		return false
	}

	ss := strings.Split(path, "")

	if ss[0] == "-" {
		return false
	}

	return true
}

func IsValidCStd(std Standard) bool {
	switch std {
	case STD_C89:
		fallthrough
	case STD_C99:
		fallthrough
	case STD_C11:
		fallthrough
	case STD_C17:
		fallthrough
	case STD_C_GNU89:
		fallthrough
	case STD_C_GNU99:
		fallthrough
	case STD_C_GNU11:
		fallthrough
	case STD_C_GNU17:
		return true
	}

	return false
}

func IsValidCXXStd(std Standard) bool {
	switch std {
	case STD_CXX98:
		fallthrough
	case STD_CXX11:
		fallthrough
	case STD_CXX14:
		fallthrough
	case STD_CXX17:
		fallthrough
	case STD_CXX_GNU98:
		fallthrough
	case STD_CXX_GNU11:
		fallthrough
	case STD_CXX_GNU14:
		fallthrough
	case STD_CXX_GNU17:
		return true
	}

	return false
}
