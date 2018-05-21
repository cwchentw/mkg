# mkg - Opinioned GNU Make-based Project Generator for C or C++

`mkg` generates cross-platform, GNU Make-based C or C++ projects.

## System Requirements

* A recent Go compiler

We provide pre-compiled executables. If in doubt, check our source and compile it by yourself. 

## Install

Just move pre-compiled `mkg` executable to any valid system **PATH**.

Alternatively, get it with Go:

```
$ go get https://github.com/cwchentw/mkg.git
```

## Synposis

```
$ mkg [option] /path/to/project
```

## Usage

`mkg` generated Makefile utilizes system default C or C++ compiler, i.e. Visual C++, Clang, GCC. Nevertheless, `mkg` users may choose their favored compiler by setting environment variables.

Make is a part of POSIX standard and many Unix-like systems adopt GNU Make as their Make implementations. In Windows, you may get a GNU Make port from either [GnuWin32](http://gnuwin32.sourceforge.net/) or [MSYS2](https://www.msys2.org/).

By default, `mkg` will generate a nested C application project to the target path:

```
$ mkg hello
```

You may adjust `mkg` with some parameters:

```
$ mkg -cpp -lib --flat greet
```

To invoke `mkg` interactively, run it without any argument:

```
$ mkg
Program name [hello]:
Project path [hello]:
Project language (c/cpp) [c]:
Project type (app/lib) [app]:
```

In this case, `mkg` will generate a project with a sensible project structure.

Alternatively, run it interactively with more customization:

```
$ mkg --custom
Program name [hello]:
Project path [hello]:
Project language (c/cpp) [c]:
Project type (app/lib) [app]:
Project structure (nested/flat) [nested]:
Project source directory [src]:
Project include directory [include]:
Project test directory [test]:
Project example directory [examples]:
Project config file [Makefile]:
```

## Options

### Program metadata

* `-v` or `--version`: Show version info
* `--license`: Show license info
* `-h` or `--help`: Show help message

### Project metadata

* `-p _prog_` or `--program _prog_`: Set program name to _prog_, default to directory name
* `-o _config_` or `--output _config_`: Set Make configuration to _config_, default to *Makefile*

### Behavior modifiers

* `-c` or `-C`: generate a C project (default)
* `-cpp` or `-cxx`: generate a C++ project
* `-app` or `--application`: generate an application project (default)
* `-lib` or `--library`: generate a library project
* `--nested`: generate a nested project (default)
* `--flat`: generate a flat project
* `--custom`: run it interactively with more customization

### Project structure

These parameters only make effects in nested projects.

* `-s _dir_` or `--source _dir_`: set source directory, default to *src*
* `-i _dir_` or `--include _dir_`: set include directory, default to *include*
* `-t _dir_` or `--test _dir_`: set test programs directory, default to *test*
* `-e _dir_` or `--example _dir_`: set example programs directory, default to *examples*

## Philosophy

`mkg` is a C or C++ project generator that is

* Green: `mkg` is a statically-compiled executable without any external runtime environment
* Portable: `mkg` itself and the generated projects are portable on the big three desktop systems
* Simple: no yet another Makefile-generating language but only the dead-simple Makefile mini-language

[Autotools](https://www.gnu.org/savannah-checkouts/gnu/autoconf/manual/autoconf-2.69/html_node/The-GNU-Build-System.html#The-GNU-Build-System) is a well-known Makefile generating tool, but only feasible on Unix-like systems. [CMake](https://cmake.org/) is famous and cross-platform, but CMake users need a full language to utilize CMake. [Bakefile](https://bakefile.org/) is a less famous CMake alternative, but you still need yet another high-level language to utilize it. There have been some community projects like [PyMake](https://github.com/Melinysh/PyMake) or [vfnmake](https://github.com/Vifon/vfnmake), but they rely on some runtime environments. Therefore, we made our own wheel.

## Author

Michael Chen, 2018.

## License

MIT
