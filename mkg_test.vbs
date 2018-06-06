Set Shell = WScript.CreateObject("WScript.Shell")

' Build the main program.
Shell.Run "cmd /c go build", 1, True

' Test a flat application project for C.
Shell.Run "cmd /c .\mkg --flat -f myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True

' Test a flat application project for C++.
Shell.Run "cmd /c .\mkg --flat -cxx --force myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True

' Test a flat library project for C.
Shell.Run "cmd /c .\mkg --flat --library -f mylib " &_
    "&& cd mylib " &_
    "&& make test " &_
    "&& make clean " &_
    "&& make testStatic " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib ", 1, True

' Clean the main program.
Shell.Run "cmd /c go clean && go fmt", 1, True