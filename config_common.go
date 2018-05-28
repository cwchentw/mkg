package main

const config_platform = `# Detect underlying system.
ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

const config_target = `# Set default target.
TARGET=

ifndef TARGET
	TARGET=Release
endif  # TARGET

export TARGET
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

export SEP
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

export DYNAMIC_LIB

ifeq ($(CC),cl)
	STATIC_LIB=$(PROGRAM).lib
else
	STATIC_LIB=lib$(PROGRAM).a
endif

export STATIC_LIB
`

const config_objects = `# Set object files.
# Modify it if more than one source files.
ifeq ($(CC),cl)
	OBJS=$(PROGRAM).obj
else
	OBJS=$(PROGRAM).o
endif  # OBJS

export OBJS
`

const config_external_library = `# Set third-party include and library path
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
