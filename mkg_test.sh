#!/bin/sh

function assert {
    if [ $? -ne 0 ]; then echo "Failed program state"; exit 1; fi
}

function runApp {
    cd myapp && make 2>&1 >/dev/null && assert && make clean && \
        make test 2>&1 >/dev/null && assert && make clean && cd ..
}

function runLib {
    cd mylib && make 2>&1 >/dev/null && make test && assert && make clean && \
        make static 2>&1 >/dev/null && make testStatic && assert && make clean && cd ..
}

PROGRAM=mkg

# Build executables
go build

# Create a flat application project for C.
./$PROGRAM -f --flat myapp

# Run the test
runApp

# Create a flat application project for C++.
./$PROGRAM -f --flat -cxx myapp

# Run the test
runApp

# Create a nested application project for C.
./$PROGRAM -f myapp

# Run the test
runApp

# Create a nested application project for C++.
./$PROGRAM -f -cpp myapp

# Run the test
runApp

# Remove the project.
rm -rf myapp

# Create a flat library project for C.
./$PROGRAM -f --flat --library mylib

# Run the test.
runLib

# Create a nested library project for C.
./$PROGRAM -f --library mylib

# Run the test.
runLib

# Create a flat library project for C++.
./$PROGRAM -f -cxx --flat --library mylib

# Run the test.
runLib

# Create a nested library project for C++.
./$PROGRAM -f -cpp --library mylib

# Run the test.
runLib

# Remove the project.
rm -rf mylib

# Create a project with -p parameter.
./$PROGRAM -p app myapp

# Run the test.
runApp

# Create a project with -a and -b parameter.
./$PROGRAM -a "Michael Chen" -b "Hello World App" -f myapp

# Run the test
runApp

# Remove the project.
rm -rf myapp

# Clean executables
go clean

# Formatting go code.
go fmt
