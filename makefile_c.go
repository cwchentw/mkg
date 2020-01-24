package main

const MakefileCC = `# Set default C compiler.
# Clean implict CC variable.
CC=

ifndef CC
	ifeq ($(detected_OS),Windows)
		CC=cl
	else ifeq ($(detected_OS),Darwin)
		CC=clang
	else
		CC=gcc
	endif
endif  # CC

export CC
`

const MakefileCStandard = `# Clean C_STD variable.
C_STD=

ifndef C_STD
	ifeq ($(CC),cl)
		C_STD=
	else
		C_STD={{ .Standard }}
	endif
endif  # C_STD

export C_STD
`

const MakefileCFlagsDebug = `# Set CFLAGS for Debug target.
ifneq (,$(DEBUG))
	ifeq ($(CC),cl)
		CFLAGS+=/DDEBUG /Zi /Od
	else
		CFLAGS+=-DDEBUG -g -O0
	endif
else
	ifeq ($(CC),cl)
		CFLAGS+=/O2
	else
		CFLAGS+=-O2
	endif
endif

export CFLAGS
`

const MakefileCFlagsRelease = `# Set CFLAGS for Release target.
CFLAGS=
ifndef CFLAGS
	ifeq ($(CC),cl)
		CFLAGS=/W4 /sdl
	else
		CFLAGS:=-Wall -Wextra -std=$(C_STD)
	endif
endif
`

const MakefileProgram = `# Set proper program name.
ifeq ($(detected_OS),Windows)
	PROGRAM={{.Program}}.exe
else
	PROGRAM={{.Program}}
endif

export PROGRAM

# Add your own test programs as needed.
ifeq ($(detected_OS),Windows)
	TEST_PROGRAM={{.Program}}.vbs
else
	TEST_PROGRAM={{.Program}}.bash
endif
`

const MakefileCLib = `# Set proper library name.
PROGRAM={{.Program}}

ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	DYNAMIC_LIB=$(PROGRAM).dll
else
	DYNAMIC_LIB=lib$(PROGRAM).dll
endif  # $(CC)
else
ifeq ($(detected_OS),Darwin)
	DYNAMIC_LIB=lib$(PROGRAM).dylib
else
	DYNAMIC_LIB=lib$(PROGRAM).so
endif  # $(detected_OS),Darwin
endif  # $(detected_OS),Windows

export DYNAMIC_LIB

ifeq ($(CC),cl)
	STATIC_LIB=$(PROGRAM).lib
else
	STATIC_LIB=lib$(PROGRAM).a
endif

export STATIC_LIB

# Add your own test programs as needed.
TEST_SOURCE=$(PROGRAM)_test.c

ifeq ($(CC),cl)
	TEST_OBJS=$(TEST_SOURCE:.c=.obj)
else
	TEST_OBJS=$(TEST_SOURCE:.c=.o)
endif

export TEST_OBJS
`

const MakefileCObj = `# Modify it if more than one source files.
SOURCE=$(PROGRAM:.exe=).c

# Set object files.
ifeq ($(CC),cl)
	OBJS=$(SOURCE:.c=.obj)
else
	OBJS=$(SOURCE:.c=.o)
endif  # OBJS

export OBJS
`

const MakefileCObjLib = `# Modify it if more than one source files.
SOURCE=$(PROGRAM:.exe=).c

# Set object files.
ifeq ($(CC),cl)
	OBJS=$(SOURCE:.c=.obj)
else
	OBJS=$(SOURCE:.c=.o)
endif  # OBJS

export OBJS
`

const MakefileCExtLib = `# Set third-party include and library path
# Modify it as needed.
ifeq ($(CC),cl)
	LDFLAGS=
	LDLIBS=
else
	LDFLAGS=
	LDLIBS=
endif

export LDFLAGS
export LDLIBS=
`
