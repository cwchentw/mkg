package main

import (
	"fmt"
	"io"
)

func printVersion() {
	fmt.Println(VERSION)
}

func printHelp(stream io.Writer) {
	fmt.Fprintf(stream, `Usage: %s [option] /path/to/project

Options:
    -p _prog_           Set program name as _prog_
    --program _prog_
    -o _config_         Set config file to _config_, default to Makefile
    --output _config_
    -l _cert_           Choose a open-source license from _cert_
    --license _cert_
    --licenses          Show available licenses

    -c, -C              Generate a C project (default)
    -cpp, -cxx          Generate a C++ project
    -app                Generate an application project (default)
    --application
    -lib                Generate a library project
    --library
    --nested            Generate a nested project (default)
    --flat              Generate a flat project

    -s _dir_            Make a custom source directory at _dir_
    --source _dir_
    -i _dir_            Make a custom include directory at _dir_
    --include _dir_
    -t _dir_            Make a custom test directory at _dir_
    --test _dir_
    -e _dir_            Make a custom example directory at _dir_
    --example _dir_
    
    Custom directories only make effects in nested projects.

    --custom            Run interactively with more customization
    -h, --help          Show help message
    -v, --version       Show version info

To invoke %s interactively, run without any argument or with --custom
`, PROGRAM, PROGRAM)
}
