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
	ifeq ($(detected_OS),Windows)
		C_STD=
	else ifeq ($(detected_OS),Darwin)
		C_STD=c99
	else
		C_STD=c99
	endif
endif  # C_STD

export C_STD
`

const MakefileCFlagsDebug = `# Set CFLAGS for Debug target.
ifndef CFLAGS_DEBUG
	ifeq ($(CC),cl)
		CFLAGS_DEBUG=/Wall /sdl /Zi
	else ifeq ($(detected_OS),Darwin)
		ifeq ($(CC),clang)
			CFLAGS_DEBUG:=-Wall -Wextra -O1 -g -std=$(C_STD) -fsanitize=address -fno-omit-frame-pointer
		else
			CFLAGS_DEBUG:=-Wall -Wextra -g -std=$(C_STD)
		endif
	else
		CFLAGS_DEBUG:=-Wall -Wextra -g -std=$(C_STD)
	endif
endif  # CFLAGS_DEBUG

export CFLAGS_DEBUG
`

const MakefileCFlagsRelease = `# Set CFLAGS for Release target.
ifndef CFLAGS_RELEASE
	ifeq ($(CC),cl)
		CFLAGS_RELEASE=/Wall /sdl /O2
	else
		CFLAGS_RELEASE:=-Wall -Wextra -O2 -std=$(C_STD)
	endif
endif  # CFLAGS_RELEASE

export CFLAGS_RELEASE
`

const MakefileCFlags = `# Set default CFLAGS
# Clean implict CFLAGS
CFLAGS=

ifndef CFLAGS
	ifeq ($(TARGET),Debug)
		CFLAGS=$(CFLAGS_DEBUG)
	else
		CFLAGS=$(CFLAGS_RELEASE)
	endif
endif  # CFLAGS

export CFLAGS
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

# Set to VSVARS32.bat on Visual Studio 2015 or earlier version
SET_ENV=VsDevCmd.bat -arch=amd64

export SET_ENV
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

# Set to VSVARS32.bat on Visual Studio 2015 or earlier version
SET_ENV=VsDevCmd.bat -arch=amd64

export SET_ENV
`

const MakefileCExtLib = `# Set third-party include and library path
# Modify it as needed.
ifeq ($(CC),cl)
	INCLUDE=
	LIBS=
else
	INCLUDE=
	LIBS=
endif

export INCLUDE
export LIBS
`
