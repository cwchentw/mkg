package main

const makefile_platform = `# Detect underlying system.
ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

const makefile_target = `# Set default target.
TARGET=

ifndef TARGET
	TARGET=Release
endif  # TARGET

export TARGET
`

const makefile_rm = `# Set proper RM on Windows.
ifeq ($(detected_OS),Windows)
	RM=del
endif

export RM
`

const makefile_sep = `# Set proper path separator.
ifeq ($(detected_OS),Windows)
	SEP=\\
else
	SEP=/
endif

export SEP
`
