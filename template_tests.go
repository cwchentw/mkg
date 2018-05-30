package main

const program_app_test = `#!/usr/bin/env bats

PROGRAM={{.Program}}

@test "Test main program" {
    run ./$PROGRAM
    [ "$output" == "Hello World" ]
}
`

const program_app_test_nested = `#!/usr/bin/env bats

PROGRAM={{.Program}}
DIST_DIR={{.DistDir}}

@test "Test main program" {
    run ./$DIST_DIR/$PROGRAM
    [ "$output" == "Hello World" ]
}
`
