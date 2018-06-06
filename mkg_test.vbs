Set Shell = WScript.CreateObject("WScript.Shell")

' Compile the project.
Shell.Run "cmd /c go build " &_
    "&& .\mkg --flat -f myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp " &_
    "&& go clean" &_
    "&& go fmt"
