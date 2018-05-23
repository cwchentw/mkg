package main

const config_platform = `# Detect underlying system.
ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

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

const config_target = `# Set default target.
TARGET=

ifndef TARGET
	TARGET=Release
endif  # TARGET

export TARGET
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

const config_rm = `# Set proper RM on Windows.
ifeq ($(detected_OS),Windows)
	RM=del
endif

export RM
`

const config_sep = `# Set proper path separator.
ifeq ($(detected_OS),Windows)
	SEP=\\
else
	SEP=/
endif
`

const config_program = `# Set proper program name.
ifeq ($(detected_OS),Windows)
	PROGRAM=%s.exe
else
	PROGRAM=%s
endif

export PROGRAM
`

const config_library = `# Set proper library name.
PROGRAM=%s

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

ifeq ($(CC),cl)
	STATIC_LIB=$(PROGRAM).lib
else
	STATIC_LIB=lib$(PROGRAM).a
endif
`

const config_objects = `ifeq ($(CC),cl)
	OBJS=$(PROGRAM).obj
else
	OBJS=$(PROGRAM).o
endif  # OBJS

export OBJS
`

const config_external_library = `# Set third-party include and library path
ifeq ($(CC),cl)
	INCLUDE=
	LIBS=
else
	INCLUDE=
	LIBS=
endif
`

const config_app_flat_c = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	@echo "Unsupported"
else
	./$(PROGRAM).bash
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq (($CC),cl)
	$(CC) $(CFLAGS) /Fe $(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS)
else
	$(CC) $(CFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%s: %s
	$(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c $< 

%s: %s
	$(CC) $(CFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const config_app_flat_cpp = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	@echo "Unsupported"
else
	./$(PROGRAM).bash
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq (($CXX),cl)
	$(CXX) $(CXXFLAGS) /Fe $(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS) 
else
	$(CXX) $(CXXFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%s: %s
	$(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c $< 

%s: %s
	$(CXX) $(CXXFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const config_lib_flat_c = `.PHONY: all clean

all: dynamic

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CC),cl)
		for %%x in (*.c) do $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c %%x
		link /DLL /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
	else
		for %%x in (*.c) do $(CC) $(CFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
		$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
	endif
else
	for x in ` + "`" + `ls *.c` + "`" + `; do $(CC) $(CFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%s: %s
	$(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c $<

%s: %s
	$(CC) $(CFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const config_app_clean = `clean:
	$(RM) $(PROGRAM) $(OBJS)
`

const config_lib_clean = `clean:
	$(RM) $(DYNAMIC_LIB) $(STATIC_LIB) $(OBJS)
`
