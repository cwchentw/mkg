package main

const program_app_test = `#!/usr/bin/env bats

PROGRAM=%s

@test "Test main program" {
    run ./$PROGRAM
    [ "$output" == "Hello World" ]
}
`

const program_app_test_nested = `#!/usr/bin/env bats

PROGRAM=%s
DIST_DIR=%s

@test "Test main program" {
    run ./$DIST_DIR/$PROGRAM
    [ "$output" == "Hello World" ]
}
`
