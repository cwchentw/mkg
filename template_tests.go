package main

const program_app_test = `#!/usr/bin/env bats

PROGRAM=%s

@test "Test main program" {
    run ./$PROGRAM
    [ "$output" == "Hello World" ]
}
`
