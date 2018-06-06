Set Shell = WScript.CreateObject("WScript.Shell")

' Test a flat application project for C.
Shell.Run "cmd /c go build " &_
    "&& .\mkg --flat -f myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp " &_
    "&& go clean " &_
    "&& go fmt", 1, True

' Test a flat application project for C++.
Shell.Run "cmd /c go build " &_
    "&& .\mkg --flat -cxx --force myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp " &_
    "&& go clean " &_
    "&& go fmt", 1, True