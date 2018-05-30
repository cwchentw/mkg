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

#ifndef __cplusplus
    #include <stdbool.h>
#endif

#ifdef __cplusplus
extern "C" {
#endif

bool is_even(int n);

#ifdef __cplusplus
}
#endif

#endif  // {{.Program}}_H
`

const program_lib_c = `#include <stdbool.h>
#include "{{.Program}}.h"

bool is_even(int n)
{
    return n % 2 == 0;
}
`

const program_lib_cpp = `#include "{{.Program}}.hpp"

bool is_even(int n)
{
    return n % 2 == 0;
}
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
    TEST(is_even(3) == false);
    TEST(is_even(4) == true);
    
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
    TEST(is_even(3) == false);
    TEST(is_even(4) == true);
    
    return 0;
}
`
