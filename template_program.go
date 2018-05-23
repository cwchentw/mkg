package main

const program_app_c = `#include <stdio.h>

int main(int argc, char *argv[])
{
    printf("Hello World\n");
    
    return 0;
}
`

const program_app_cpp = `#include <iostream>

using std::cout;
using std::endl;

int main(int argc, char *argv[])
{
    cout << "Hello World" << endl;
    
    return 0;
}
`
