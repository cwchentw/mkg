package main

const makefile_cc = `# Set default C compiler.
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

const makefile_cxx = `# Set default C++ compiler.
# Clean implict CXX variable.
CXX=

ifndef CXX
	ifeq ($(detected_OS),Windows)
		CXX=cl
	else ifeq ($(detected_OS),Darwin)
		CXX=clang++
	else
		CXX=g++
	endif
endif  # CXX

export CXX
`

const makefile_cflags_debug = `# Set CFLAGS for Debug target.
ifndef CFLAGS_DEBUG
	ifeq ($(CC),cl)
		CFLAGS_DEBUG=/Wall /sdl /Zi
	else ifeq ($(detected_OS),Darwin)
		ifeq ($(CC),clang)
			CFLAGS_DEBUG=-Wall -Wextra -O1 -g -std=c99 -fsanitize=address -fno-omit-frame-pointer
		else
			CFLAGS_DEBUG=-Wall -Wextra -g -std=c99
		endif
	else
		CFLAGS_DEBUG=-Wall -Wextra -g -std=c99
	endif
endif  # CFLAGS_DEBUG

export CFLAGS_DEBUG
`

const makefile_cxxflags_debug = `# Set CXXFLAGS for Debug target.
ifndef CXXFLAGS_DEBUG
	ifeq ($(CXX),cl)
		CXXFLAGS_DEBUG=/Wall /sdl /EHsc /std:c++11 /Zi
	else ifeq ($(detected_OS),Darwin)
		ifeq ($(CC),clang)
			CXXFLAGS_DEBUG=-Wall -Wextra -O1 -g -std=c++11 -fsanitize=address -fno-omit-frame-pointer
		else
			CXXFLAGS_DEBUG=-Wall -Wextra -g -std=c++11
		endif
	else
		CXXFLAGS_DEBUG=-Wall -Wextra -g -std=c++11
	endif
endif  # CXXFLAGS_DEBUG

export CXXFLAGS_DEBUG
`

const makefile_cflags_release = `# Set CFLAGS for Release target.
ifndef CFLAGS_RELEASE
	ifeq ($(CC),cl)
		CFLAGS_RELEASE=/Wall /sdl /O2
	else
		CFLAGS_RELEASE=-Wall -Wextra -O2 -std=c99
	endif
endif  # CFLAGS_RELEASE

export CFLAGS_RELEASE
`

const makefile_cxxflags_release = `# Set CXXFLAGS for Release target.
ifndef CXXFLAGS_RELEASE
	ifeq ($(CC),cl)
		CXXFLAGS_RELEASE=/Wall /sdl /EHsc /std:c++11 /O2
	else
		CXXFLAGS_RELEASE=-Wall -Wextra -O2 -std=c++11
	endif
endif  # CXXFLAGS_RELEASE

export CXXFLAGS_DEBUG
`

const makefile_cflags = `# Set default CFLAGS
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

const makefile_cxxflags = `# Set default CXXFLAGS
# Clean implict CXXFLAGS
CXXFLAGS=

ifndef CXXFLAGS
	ifeq ($(TARGET),Debug)
		CXXFLAGS=$(CXXFLAGS_DEBUG)
	else
		CXXFLAGS=$(CXXFLAGS_RELEASE)
	endif
endif  # CXXFLAGS

export CXXFLAGS
`

const makefile_program = `# Set proper program name.
ifeq ($(detected_OS),Windows)
	PROGRAM={{.Program}}.exe
else
	PROGRAM={{.Program}}
endif

export PROGRAM
`

const makefile_library = `# Set proper library name.
PROGRAM={{.Program}}

ifeq ($(detected_OS),Windows)
	ifeq ($(CC),cl)
		DYNAMIC_LIB=$(PROGRAM).dll
	else
		DYNAMIC_LIB=lib$(PROGRAM).dll
	endif
else
	ifeq ($(detected_OS),Darwin)
		DYNAMIC_LIB=lib$(PROGRAM).dylib
	else
		DYNAMIC_LIB=lib$(PROGRAM).so
	endif
endif

export DYNAMIC_LIB

ifeq ($(CC),cl)
	STATIC_LIB=$(PROGRAM).lib
else
	STATIC_LIB=lib$(PROGRAM).a
endif

export STATIC_LIB

# Add your own objects for the test programs
ifeq ($(CC),cl)
	TEST_OBJS=$(PROGRAM)_test.obj
else
	TEST_OBJS=$(PROGRAM)_test.o
endif

export TEST_OBJS
`

const makefile_objects = `# Set object files.
# Modify it if more than one source files.
ifeq ($(CC),cl)
	OBJS=$(PROGRAM).obj
else
	OBJS=$(PROGRAM).o
endif  # OBJS

export OBJS
`

const makefile_external_library = `# Set third-party include and library path
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
