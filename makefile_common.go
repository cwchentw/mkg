package main

const makefilePlatform = `# Detect underlying system.
ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

const makefileTarget = `# Set default target.
TARGET=

ifndef TARGET
	TARGET=Release
endif  # TARGET

export TARGET
`

const makefileRm = `# Set proper RM on Windows.
ifeq ($(detected_OS),Windows)
	RM=del
endif

export RM
`

const makefileSep = `# Set proper path separator.
ifeq ($(detected_OS),Windows)
	SEP=\\
else
	SEP=/
endif

export SEP
`
