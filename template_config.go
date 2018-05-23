package main

const config_platform = `ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

const config_cc = `CC=

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

const config_cxx = `CXX=

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

const config_cflags_debug = `ifndef CFLAGS_DEBUG
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

const config_cxxflags_debug = `ifndef CXXFLAGS_DEBUG
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

const config_cflags_release = `ifndef CFLAGS_RELEASE
	ifeq ($(CC),cl)
		CFLAGS_RELEASE=/Wall /sdl /O2
	else
		CFLAGS_RELEASE=-Wall -Wextra -O2 -std=c99
	endif
endif  # CFLAGS_RELEASE

export CFLAGS_RELEASE
`

const config_cxxflags_release = `ifndef CXXFLAGS_RELEASE
	ifeq ($(CC),cl)
		CXXFLAGS_RELEASE=/Wall /sdl /EHsc /std:c++11 /O2
	else
		CXXFLAGS_RELEASE=-Wall -Wextra -O2 -std=c++11
	endif
endif  # CXXFLAGS_RELEASE

export CXXFLAGS_DEBUG
`

const config_target = `TARGET=

ifndef TARGET
	TARGET=Release
endif  # TARGET

export TARGET
`

const config_cflags = `CFLAGS=

ifndef CFLAGS
	ifeq ($(TARGET),Debug)
		CFLAGS=$(CFLAGS_DEBUG)
	else
		CFLAGS=$(CFLAGS_RELEASE)
	endif
endif  # CFLAGS

export CFLAGS
`

const config_cxxflags = `CXXFLAGS=

ifndef CXXFLAGS
	ifeq ($(TARGET),Debug)
		CXXFLAGS=$(CXXFLAGS_DEBUG)
	else
		CXXFLAGS=$(CXXFLAGS_RELEASE)
	endif
endif  # CXXFLAGS

export CXXFLAGS
`

const config_rm = `ifeq ($(detected_OS),Windows)
	RM=del
endif

export RM
`

const config_program = `ifeq ($(detected_OS),Windows)
	PROGRAM=%s.exe
else
	PROGRAM=%s
endif

export PROGRAM

ifeq ($(CC),cl)
	OBJS=$(PROGRAM).obj
else
	OBJS=$(PROGRAM).o
endif  # OBJS

export OBJS
`

const config_app_flat_c = `.PHONY: all clean

all: $(PROGRAM)

$(PROGRAM): $(OBJS)
ifeq (($CC),cl)
	$(CC) $(CFLAGS) /Fe $(PROGRAM) $(OBJS)
else
	$(CC) $(CFLAGS) -o $(PROGRAM) $(OBJS)
endif

%s: %s
	$(CC) $(CFLAGS) /c $<

%s: %s
	$(CC) $(CFLAGS) -c $<
`

const config_app_flat_cpp = `.PHONY: all clean

all: $(PROGRAM)

$(PROGRAM): $(OBJS)
ifeq (($CXX),cl)
	$(CXX) $(CXXFLAGS) /Fe $(PROGRAM) $(OBJS)
else
	$(CXX) $(CXXFLAGS) -o $(PROGRAM) $(OBJS)
endif

%s: %s
	$(CXX) $(CXXFLAGS) /c $<

%s: %s
	$(CXX) $(CXXFLAGS) -c $<
`

const config_clean = `clean:
	$(RM) $(PROGRAM) $(OBJS)
`