package main

import (
	"errors"
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
		return "app"
	case PROJ_LIB:
		return "lib"
	default:
		panic("Unknown project type")
	}
}

func stringToProj(proj string) (ProjectType, error) {
	switch proj {
	case "app":
		return PROJ_APP, nil
	case "lib":
		return PROJ_LIB, nil
	default:
		return PROJ_APP, errors.New("Invalid project type")
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

func stringToLicense(cert string) (License, error) {
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

func getTemplate(license License) string {
	switch license {
	case LICENSE_NONE:
		return ""
	case LICENSE_APACHE2:
		return license_apache2
	case LICENSE_MIT:
		return license_mit
	case LICENSE_BSD2:
		return license_bsd2
	case LICENSE_GPL2:
		return license_gpl2
	case LICENSE_GPL3:
		return license_gpl3
	case LICENSE_MPL2:
		return license_mpl2
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
