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
ifndef CXXFLAGS_DEBUG
	ifeq ($(CXX),cl)
		CXXFLAGS_DEBUG:=/Wall /sdl /EHsc /std:$(CXX_STD) /Zi
	else ifeq ($(detected_OS),Darwin)
		ifeq ($(CXX),clang)
			CXXFLAGS_DEBUG:=-Wall -Wextra -O1 -g -std=$(CXX_STD) -fsanitize=address -fno-omit-frame-pointer
		else
			CXXFLAGS_DEBUG:=-Wall -Wextra -g -std=$(CXX_STD)
		endif
	else
		CXXFLAGS_DEBUG:=-Wall -Wextra -g -std=$(CXX_STD)
	endif
endif  # CXXFLAGS_DEBUG

export CXXFLAGS_DEBUG
`

const makefile_cxxflags_release = `# Set CXXFLAGS for Release target.
ifndef CXXFLAGS_RELEASE
	ifeq ($(CXX),cl)
		CXXFLAGS_RELEASE:=/Wall /sdl /EHsc /std:$(CXX_STD) /O2
	else
		CXXFLAGS_RELEASE:=-Wall -Wextra -O2 -std=$(CXX_STD)
	endif
endif  # CXXFLAGS_RELEASE

export CXXFLAGS_DEBUG
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

# Set to VSVARS32.bat on Visual Studio 2015 or earlier version
SET_ENV=VsDevCmd.bat -arch=amd64

export SET_ENV
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

# Set to VSVARS32.bat on Visual Studio 2015 or earlier version
SET_ENV=VsDevCmd.bat -arch=amd64

export SET_ENV
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
