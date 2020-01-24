package main

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

const MakefileCXXStandard = `# Clean CXX_STD variable.
CXX_STD=

ifndef CXX_STD
	ifeq ($(detected_OS),Windows)
		CXX_STD={{ .StandardWin }}
	else ifeq ($(detected_OS),Darwin)
		CXX_STD={{ .Standard }}
	else
		CXX_STD={{ .Standard }}
	endif
endif  # CXX_STD

export CXX_STD
`

const makefile_cxxflags_debug = `# Set CXXFLAGS for Debug target.
ifneq (,$(DEBUG))
	ifeq ($(CXX),cl)
		CXXFLAGS+=/DDEBUG /Zi /Od
	else
		CXXFLAGS+=-DDEBUG -g -O0
	endif
else
	ifeq ($(CXX),cl)
		CXXFLAGS+=/O2
	else
		CXXFLAGS+=-O2
	endif
endif

export CXXFLAGS
`

const makefile_cxxflags_release = `# Set CXXFLAGS
CXXFLAGS=
ifndef CXXFLAGS
	ifeq ($(CXX),cl)
		CXXFLAGS:=/W4 /sdl /EHsc /std:$(CXX_STD)
	else
		CXXFLAGS:=-Wall -Wextra -std=$(CXX_STD)
	endif
endif
`

const makefileLibCpp = `# Set proper library name.
PROGRAM={{.Program}}

ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	DYNAMIC_LIB=$(PROGRAM).dll
else
	DYNAMIC_LIB=lib$(PROGRAM).dll
endif  # $(CXX)
else
ifeq ($(detected_OS),Darwin)
	DYNAMIC_LIB=lib$(PROGRAM).dylib
else
	DYNAMIC_LIB=lib$(PROGRAM).so
endif  # $(detected_OS),Darwin
endif  # $(detected_OS),Windows

export DYNAMIC_LIB

ifeq ($(CXX),cl)
	STATIC_LIB=$(PROGRAM).lib
else
	STATIC_LIB=lib$(PROGRAM).a
endif

export STATIC_LIB

# Add your own test programs as needed.
TEST_SOURCE=$(PROGRAM)_test.cpp

ifeq ($(CXX),cl)
	TEST_OBJS=$(TEST_SOURCE:.cpp=.obj)
else
	TEST_OBJS=$(TEST_SOURCE:.cpp=.o)
endif

export TEST_OBJS
`

const makefileObjectCpp = `# Modify it if more than one source files.
SOURCE=$(PROGRAM:.exe=).cpp

# Set object files.
ifeq ($(CXX),cl)
	OBJS=$(SOURCE:.cpp=.obj)
else
	OBJS=$(SOURCE:.cpp=.o)
endif  # OBJS

export OBJS
`

const makefileObjCppLib = `# Modify it if more than one source files.
SOURCE=$(PROGRAM:.exe=).cpp

# Set object files.
ifeq ($(CXX),cl)
	OBJS=$(SOURCE:.cpp=.obj)
else
	OBJS=$(SOURCE:.cpp=.o)
endif  # OBJS

export OBJS
`

const makefileExternalLibraryCpp = `# Set third-party include and library path
# Modify it as needed.
ifeq ($(CXX),cl)
	INCLUDE=
	LIBS=
else
	INCLUDE=
	LIBS=
endif

export INCLUDE
export LIBS
`
