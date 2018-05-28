#!/bin/sh

function assert {
    if [ $? -ne 0 ]; then echo "Failed program state"; exit 1; fi
}

function run {
    cd myapp && make 2>&1 >/dev/null && assert && make clean && \
        make test 2>&1 >/dev/null && assert && make clean && cd ..
}

PROGRAM=mkg

# Build executables
go build

# Create a flat application project for C.
./$PROGRAM -f --flat myapp

# Run the test
run

# Create a flat application project for C++.
./$PROGRAM -f --flat -cxx myapp

# Run the test
run

# Create a nested application project for C.
./$PROGRAM -f myapp

# Run the test
run

# Create a nested application project for C++.
./$PROGRAM -f -cpp myapp

# Run the test
run

# Remove the project.
rm -rf myapp

# Clean executables
go clean