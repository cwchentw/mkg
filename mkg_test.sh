#!/bin/sh

function assert {
    if [ $? -ne 0 ]; then echo "Failed program state"; exit 1; fi
}

PROGRAM=mkg

# Build executables
go build

# Create a flat application project for C.
./$PROGRAM -f --flat myapp

# Build the app.
cd myapp && make 2>&1 >/dev/null && assert && cd ..

# Create a flat application project for C++.
./$PROGRAM -f --flat -cxx myapp

# Build the app.
cd myapp && make 2>&1 >/dev/null && assert && cd ..

# Create a nested application project for C.
./$PROGRAM -f myapp

# Build the app.
cd myapp && make 2>&1 >/dev/null && assert && cd ..

# Create a nested application project for C++.
./$PROGRAM -f -cpp myapp

# Build the app.
cd myapp && make 2>&1 >/dev/null && assert && cd ..

# Remove the project.
rm -rf myapp

# Clean executables
go clean