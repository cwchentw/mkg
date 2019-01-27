package main

import (
	"fmt"
	"io"
)

func printVersion() {
	fmt.Println(VERSION)
}

func printHelp(stream io.Writer) {
	fmt.Fprintf(stream, `Usage: %s [option] path/to/project

Options:
    -p _prog_           Set program name as _prog_
    --program _prog_
    -a _author_         Set project author as _author_
    --author _author_
    -b _brief_          Set project description as _brief_
    --brief _brief_
    -o _makefile_       Set config file to _makefile_, default to Makefile
    --output _makefile_
    -l _license_        Choose a open-source _license_
    --license _license_
    --licenses          Show available licenses

    -c, -C              Generate a C project (default)
    -cpp, -cxx          Generate a C++ project
    -std _std_          Set the standard version of the project
    --standard _std_
    --standards         Show available standards
    --console           Generate a console project (default)
    --library           Generate a library project
    --nested            Generate a nested project (default)
    --flat              Generate a flat project

    -s _dir_            Make a custom source directory at _dir_
    --source _dir_
    -i _dir_            Make a custom include directory at _dir_
    --include _dir_
    -d _dir             Make a cust dist directory at _dir_
    --dist _dist_
    -t _dir_            Make a custom test directory at _dir_
    --test _dir_
    -e _dir_            Make a custom example directory at _dir_
    --example _dir_
    
    Custom directories only make effects in nested projects.

    --custom            Run interactively with more customization
    -f, --force         Remove all existing contents on path (Dangerous!)
    -h, --help          Show help message
    -v, --version       Show version info

To invoke %s interactively, run without any argument or with --custom
`, PROGRAM, PROGRAM)
}

func printLicenses() {
	fmt.Printf(`None (none)
Apache License 2.0 (apache2)
GNU General Public License v3.0 (gpl3)
MIT License (mit)
---
BSD 2-clause "Simplified" license (bsd2)
BSD 3-clause "New" or "Revised" license (bsd3)
Eclipse Public License 2.0 (epl2)
GNU Affero General Public License v3.0 (agpl3)
GNU General Public License v2.0 (gpl2)
GNU Lesser General Public License v2.1 (lgpl2)
GNU Lesser General Public License v3.0 (lgpl3)
Mozilla Public License 2.0 (mpl2)
The Unlicense (unlicense)
`)
}

func printStandards() {
	fmt.Printf(`c89
c90
c99
c11
c17
c18
gnu89
gnu90
gnu99
gnu11
gnu17
gnu18
c++98
c++03
c++11
c++14
c++17
gnu++98
gnu++03
gnu++11
gnu++14
gnu++17
`)
}
