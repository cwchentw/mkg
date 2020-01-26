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

const program_header = `#ifndef {{.Program}}_H
#define {{.Program}}_H

#define  {{.Program}}_VERSION  "0.1.0"

#define  {{.Program}}_VERSION_MAJOR  0
#define  {{.Program}}_VERSION_MINOR  1
#define  {{.Program}}_VERSION_PATCH  0

#if _MSC_VER
    #if defined({{.Program}}_IMPORT_SYMBOLS)
        #define {{.Program}}_PUBLIC __declspec(dllimport)
    #elif defined({{.Program}}_EXPORT_SYMBOLS)
        #define {{.Program}}_PUBLIC __declspec(dllexport)
    #else
        #define {{.Program}}_PUBLIC
    #endif
#elif __GNUC__ >= 4 || __clang__
    #define {{.Program}}_PUBLIC __attribute__((__visibility__("default")))
#else
    #define {{.Program}}_PUBLIC
#endif

#if __GNUC__ >= 4 || __clang__
    #define {{.Program}}_PRIVATE __attribute__((__visibility__("hidden")))
#else
    #define {{.Program}}_PRIVATE
#endif

#if _MSC_VER
    #include <windows.h>
#else  /* !_MSC_VER */
#ifdef __cplusplus
    #ifndef _BOOL_IS_DEFINED
        typedef bool BOOL;
        #define FALSE  false
        #define TRUE   true
        #define _BOOL_IS_DEFINED
    #endif  /* BOOL */
#else
    #if __STDC_VERSION__ < 199901L
        #ifndef _BOOL_IS_DEFINED
            typedef char BOOL;
            #define FALSE  0
            #define TRUE   1
            #define _BOOL_IS_DEFINED
        #endif  /* BOOL */
    #else
        #ifndef _BOOL_IS_DEFINED
            #include <stdbool.h>
            typedef bool BOOL;
            #define FALSE  false
            #define TRUE   true
            #define _BOOL_IS_DEFINED
        #endif  /* BOOL */
    #endif  /* C89 */
#endif  /* __cplusplus */
#endif  /* _MSC_VER */

#ifdef __cplusplus
extern "C" {
#endif

{{.Program}}_PUBLIC BOOL is_even(int n);

#ifdef __cplusplus
}
#endif

#endif  /* {{.Program}}_H */
`

const program_lib_c = `#include "{{.Program}}.h"

BOOL is_even(int n)
{
    return n % 2 == 0;
}
`

const program_lib_cpp = `#include "{{.Program}}.hpp"

BOOL is_even(int n)
{
    return n % 2 == 0;
}
`
