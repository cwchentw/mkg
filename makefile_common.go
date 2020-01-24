package main

const MakefilePlatform = `# Detect underlying system.
ifeq ($(OS),Windows_NT)
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname -s 2>/dev/null || echo not')
endif

export detected_OS
`

const MakefileRM = `# Set proper RM on Windows.
ifeq ($(detected_OS),Windows)
	RM=del /q /f
endif

export RM
`

const MakefileSep = `# Set proper path separator.
ifeq ($(detected_OS),Windows)
	SEP=\\
else
	SEP=/
endif

export SEP
`
