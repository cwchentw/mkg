package main

const config_cc = `# Set default C compiler.
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

const config_cxx = `# Set default C++ compiler.
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

const config_cflags_debug = `# Set CFLAGS for Debug target.
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

const config_cxxflags_debug = `# Set CXXFLAGS for Debug target.
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

const config_cflags_release = `# Set CFLAGS for Release target.
ifndef CFLAGS_RELEASE
	ifeq ($(CC),cl)
		CFLAGS_RELEASE=/Wall /sdl /O2
	else
		CFLAGS_RELEASE=-Wall -Wextra -O2 -std=c99
	endif
endif  # CFLAGS_RELEASE

export CFLAGS_RELEASE
`

const config_cxxflags_release = `# Set CXXFLAGS for Release target.
ifndef CXXFLAGS_RELEASE
	ifeq ($(CC),cl)
		CXXFLAGS_RELEASE=/Wall /sdl /EHsc /std:c++11 /O2
	else
		CXXFLAGS_RELEASE=-Wall -Wextra -O2 -std=c++11
	endif
endif  # CXXFLAGS_RELEASE

export CXXFLAGS_DEBUG
`

const config_cflags = `# Set default CFLAGS
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

const config_cxxflags = `# Set default CXXFLAGS
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