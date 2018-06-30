Set Shell = WScript.CreateObject("WScript.Shell")

' Build the main program.
Shell.Run "cmd /c go build", 1, True

' Test a flat application project for C (MSVC).
Err = Shell.Run("cmd /c .\mkg --flat -f myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat application project for C (MinGW).
Err = Shell.Run("cmd /c .\mkg --flat -f myapp " &_
    "&& cd myapp " &_
    "&& make CC=gcc test " &_
    "&& make CC=gcc clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested application project for C (MSVC).
Err = Shell.Run("cmd /c .\mkg -f myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested application project for C (MinGW).
Err = Shell.Run("cmd /c .\mkg -f myapp " &_
    "&& cd myapp " &_
    "&& make CC=gcc test " &_
    "&& make CC=gcc clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat application project for C++ (MSVC).
Err = Shell.Run("cmd /c .\mkg -cxx --flat --force myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat application project for C++ (MinGW).
Err = Shell.Run("cmd /c .\mkg -cxx --flat --force myapp " &_
    "&& cd myapp " &_
    "&& make CXX=g++ test " &_
    "&& make CXX=g++ clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested application project for C++ (MSVC).
Err = Shell.Run("cmd /c .\mkg -cpp --force myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested application project for C++ (MinGW).
Err = Shell.Run("cmd /c .\mkg -cxx -f myapp " &_
    "&& cd myapp " &_
    "&& make CXX=g++ test " &_
    "&& make CXX=g++ clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat library project for C (MSVC).
Err = Shell.Run("cmd /c .\mkg --library --flat -f mylib " &_
    "&& cd mylib " &_
    "&& make test " &_
    "&& make clean " &_
    "&& make testStatic " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat library project for C (MinGw).
Err = Shell.Run("cmd /c .\mkg --library --flat -f mylib " &_
    "&& cd mylib " &_
    "&& make CC=gcc test " &_
    "&& make CC=gcc clean " &_
    "&& make CC=gcc testStatic " &_
    "&& make CC=gcc clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested library project for C (MSVC).
Err = Shell.Run("cmd /c .\mkg --library --force mylib " &_
    "&& cd mylib " &_
    "&& make test " &_
    "&& make clean " &_
    "&& make testStatic " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested library project for C (MinGW).
Err = Shell.Run("cmd /c .\mkg --library -f mylib " &_
    "&& cd mylib " &_
    "&& make CC=gcc test " &_
    "&& make CC=gcc clean " &_
    "&& make CC=gcc testStatic " &_
    "&& make CC=gcc clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat library project for C++ (MSVC).
Err = Shell.Run("cmd /c .\mkg --library --flat -cxx --force mylib " &_
    "&& cd mylib " &_
    "&& make test " &_
    "&& make clean " &_
    "&& make testStatic " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a flat library project for C++ (MinGW).
Err = Shell.Run("cmd /c .\mkg --library --flat -cpp -f mylib " &_
    "&& cd mylib " &_
    "&& make CXX=g++ test " &_
    "&& make CXX=g++ clean " &_
    "&& make CXX=g++ testStatic " &_
    "&& make CXX=g++ clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested library project for C++ (MSVC).
Err = Shell.Run("cmd /c .\mkg --library -cxx -f mylib " &_
    "&& cd mylib " &_
    "&& make test " &_
    "&& make clean " &_
    "&& make testStatic " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a nested library project for C++ (MinGW).
Err = Shell.Run("cmd /c .\mkg --library -cpp -f mylib " &_
    "&& cd mylib " &_
    "&& make CXX=g++ test " &_
    "&& make CXX=g++ clean " &_
    "&& make CXX=g++ testStatic " &_
    "&& make CXX=g++ clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q mylib", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a project with -p parameter.
Err = Shell.Run("cmd /c .\mkg -p app myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Test a project with -a and -b parameter.
Err = Shell.Run("cmd /c .\mkg -a " & """ & Michael Chen & """ &_
    " -b " & """ & Hello World App & """ & " myapp " &_
    "&& cd myapp " &_
    "&& make test " &_
    "&& make clean " &_
    "&& cd .. " &_
    "&& rmdir /s /q myapp ", 1, True)

If Err <> 0 Then
    WScript.Quit
End If

' Clean the main program.
Shell.Run "cmd /c go clean && go fmt", 1, True