# mkg - Opinionated GNU Make-based Project Generator

`mkg` generates GNU Make-based projects for either C or C++.

## System Requirements

To use `mkg`-generated projects, you need

* A recent C (or C++) compiler
* GNU Make

To compile `mkg` from source, you need

* A recent Go compiler

We provide pre-compiled executables. If in doubt, check our source and compile it by yourself.

## Install

Just move pre-compiled `mkg` executable to any valid system **PATH**.

Alternatively, install it with Go:

```
$ go get https://github.com/cwchentw/mkg.git
```

## Synposis

Run it in batch mode:

```
$ mkg [option] /path/to/project
```

Rut it interactively with a sensible project structure:

```
$ mkg
```

Rut it interactively with more customization:

```
$ mkg --custom
```

## Usage

`mkg` generated projects utilizes system default C or C++ compiler, i.e. Visual C++, Clang, GCC. Nevertheless, `mkg` users may choose their favored compiler by setting environment variables.

Make is a part of POSIX standard and many Unix-like systems adopt GNU Make as their Make implementations. In Windows, you may get a GNU Make port from either [GnuWin32](http://gnuwin32.sourceforge.net/) or [MSYS2](https://www.msys2.org/).

By default, `mkg` will generate a nested C application project to the target path:

```
$ mkg myapp
```

You may adjust `mkg` with some parameters:

```
$ mkg -cpp --library --flat mylib
```

To invoke `mkg` interactively, run it without any argument:

```
$ mkg
Program name [myapp]:
Project path [myapp]:
Project author [somebody]: Michael Chen
Project brief description [something]: Yet Another Application
Project language (c/cpp) [c]:
Project type (app/lib) [app]:

None (none)
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

Project licensing [none]:
```

In this case, `mkg` will generate a project with a sensible project structure.

Alternatively, run it interactively with more customization:

```
$ mkg --custom
Program name [myapp]:
Project path [myapp]:
Project author [somebody]: Michael Chen
Project brief description [something]: Yet Another Application
Project language (c/cpp) [c]:
Project type (app/lib) [app]:

(Choose licensing as above...)

Project structure (nested/flat) [nested]:
Project source directory [src]:
Project include directory [include]:
Project test directory [test]:
Project example directory [examples]:
Project config file [Makefile]:
```

## Options

### Program metadata

* `-v` or `--version`: Show version info and exit the program
* `-h` or `--help`: Show help message and exit the program
* `--licenses`: Show the available open-source licenses and exit the program

### Project metadata

* `-p _prog_` or `--program _prog_`: Set program name to _prog_, default to directory name
* `-a _author_` or `--author _author_`: Set project author to _author_
* `-b _brief_` or `--brief _brief_`: Set a brief description to _brief_ for the project
* `-o _makefile_` or `--output _makefile_`: Set Make configuration file to _makefile_, default to *Makefile*
* `-l _license_` or `--license _license_`: Choose a open-source _license_ for the project

Here are the available licenses in our program:

* Recommended
  * Apache License 2.0 (apache2)
  * GNU General Public License v3.0 (gpl3)
  * MIT License (mit)
* Alternative
  * BSD 2-clause "Simplified" license (bsd2)
  * BSD 3-clause "New" or "Revised" license (bsd3)
  * Eclipse Public License 2.0 (epl2)
  * GNU Affero General Public License v3.0 (agpl3)
  * GNU General Public License v2.0 (gpl2)
  * GNU Lesser General Public License v2.1 (lgpl2)
  * GNU Lesser General Public License v3.0 (lgpl3)
  * Mozilla Public License 2.0 (mpl2)
  * The Unlicense (unlicense)

### Behavior modifiers

* `-c` or `-C`: generate a C project (default)
* `-cpp` or `-cxx`: generate a C++ project
* `--console`: generate an console application project (default)
* `--library`: generate a library project
* `--nested`: generate a nested project (default)
* `--flat`: generate a flat project
* `-f` or `--force`: Remove all existing contents on path (Dangerous!)
* `--custom`: run it interactively with more customization

### Project structure

These parameters only make effects in nested projects.

* `-s _dir_` or `--source _dir_`: set source directory, default to *src*
* `-i _dir_` or `--include _dir_`: set include directory, default to *include*
* `-d _dir_` or `--dist _dir_`: set dist directory, default to *dist*
* `-t _dir_` or `--test _dir_`: set test programs directory, default to *tests*
* `-e _dir_` or `--example _dir_`: set example programs directory, default to *examples*

## Philosophy

`mkg` is a Makefile-based project generator that is

* Green: `mkg` is a statically-compiled executable without any external runtime environment
* Portable: `mkg` itself and the generated projects are portable on the big three desktop systems
* Simple: no yet another Makefile-generating language but only the dead-simple Makefile mini-language

[Autotools](https://www.gnu.org/savannah-checkouts/gnu/autoconf/manual/autoconf-2.69/html_node/The-GNU-Build-System.html#The-GNU-Build-System) is a well-known Makefile generating tool, but only feasible on Unix-like systems. [CMake](https://cmake.org/) is famous and cross-platform, but CMake users need a full language to utilize CMake. [Bakefile](https://bakefile.org/) is a less famous CMake alternative, but you still need yet another high-level language to utilize it. There have been some community projects like [PyMake](https://github.com/Melinysh/PyMake) or [vfnmake](https://github.com/Vifon/vfnmake), but they rely on some runtime environments and provide no support to Windows-family systems. Therefore, we made our own wheel.

## TODO

Add supports to the following compilers or toolchains:

* Fortran (gFortran-based)
* Objective-C (either GCC or Clang-based)
* LaTeX
* (May or may not) Vala

(May or may not) add support to CMake.

## Author

2018 Michael Chen

## License

MIT
