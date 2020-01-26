package main

const programAppTest = `#!/usr/bin/env bats

PROGRAM={{.Program}}

@test "Test main program" {
    run ./$PROGRAM
    [ "$output" == "Hello World" ]
}
`

const programAppTestNested = `#!/usr/bin/env bats

PROGRAM={{.Program}}
DIST_DIR={{.DistDir}}

@test "Test main program" {
    run ./$DIST_DIR/$PROGRAM
    [ "$output" == "Hello World" ]
}
`

const programAppTestWin = `' Set Program States
Dim Program : Program = "{{.Program}}"

Assert Capture(".\" & Program)(0) = "Hello World" & vbNewLine, "Wrong value"

' Home-made assert for VBScript.
Sub Assert( boolExpr, strOnFail )
	If not boolExpr then
		Err.Raise vbObjectError + 99999, , strOnFail
	End If
End Sub

' Capture stdout and stderr from cmd
Function Capture(cmd)
	Set WshShell = WScript.CreateObject("WScript.Shell")

	Set output = WshShell.Exec("cmd.exe /c " & cmd)

	Dim arr(2)
	
	arr(0) = ""

	Do
		o = output.StdOut.ReadLine()
		arr(0) = arr(0) & o & vbNewLine
	Loop While Not output.Stdout.atEndOfStream
	
	arr(1) = ""
	
	Do
		e = output.StdErr.ReadLine()
		arr(1) = arr(1) & e & vbNewLine
	Loop While Not output.StdErr.atEndOfStream
	
	Capture = arr
End Function
`

const programAppTestNestedWin = `' Set Program States
Dim Program : Program = "{{.Program}}"
Dim DistDir : DistDir = "{{.DistDir}}"

Assert _
	Capture(DistDir & "\" & Program)(0) = "Hello World" & vbNewLine, _
	"Wrong value"

' Home-made assert for VBScript.
Sub Assert( boolExpr, strOnFail )
	If not boolExpr then
		Err.Raise vbObjectError + 99999, , strOnFail
	End If
End Sub

' Capture stdout and stderr from cmd
Function Capture(cmd)
	Set WshShell = WScript.CreateObject("WScript.Shell")

	Set output = WshShell.Exec("cmd.exe /c " & cmd)

	Dim arr(2)
	
	arr(0) = ""

	Do
		o = output.StdOut.ReadLine()
		arr(0) = arr(0) & o & vbNewLine
	Loop While Not output.Stdout.atEndOfStream
	
	arr(1) = ""
	
	Do
		e = output.StdErr.ReadLine()
		arr(1) = arr(1) & e & vbNewLine
	Loop While Not output.StdErr.atEndOfStream
	
	Capture = arr
End Function
`

const program_lib_test_c = `#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
#include "{{.Program}}.h"

#define TEST(cond) { \
        if (!cond) { \
            fprintf(stderr, "%s %d: Failed on %s\n", __FILE__, __LINE__, #cond); \
            exit(1); \
        } \
    }

int main(void)
{
    TEST(is_even(3) == FALSE);
    TEST(is_even(4) == TRUE);
    
    return 0;
}
`

const program_lib_test_cxx = `#include <cstdlib>
#include <cstdio>
#include "{{.Program}}.hpp"

#define TEST(cond) { \
        if (!cond) { \
            fprintf(stderr, "%s %d: Failed on %s\n", __FILE__, __LINE__, #cond); \
            exit(1); \
        } \
    }

int main(void)
{
    TEST(is_even(3) == FALSE);
    TEST(is_even(4) == TRUE);
    
    return 0;
}
`
