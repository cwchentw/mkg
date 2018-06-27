=========================
Using Generated Projects
=========================

Here we illustrate how to use the generated projects created by ``mkg``.

Currently, ``mkg`` supports either flat or nested project structure. The former
is suitable for simple projects while the latter fits more complex ones. All the
projects created by ``mkg`` are ready as Git repos.

Implicitly, ``mkg`` users need some knowledge to GNU Make to smoothly use these
generated projects.

--------------------
System Requirements
--------------------

* A recent C or C++ compiler
* GNU Make
* `Bats (Bash Automated Testing System) <https://github.com/sstephenson/bats>`_ on Unix-like systems

We assume the default compiler on each platform, namely

* Visual C++ on Windows
* Clang on Mac
* GCC on other Unix-like systems such as GNU/Linux

For Windows users, you may get a port of GNU Make at `MSYS2 <https://www.msys2.org/>`_
or `GnuWin32 <http://gnuwin32.sourceforge.net/>`_. Besides, MinGW, a Windows port of GCC,
is supported in these projects as well.

----------------------------------------
Flat Console Application Projects for C
----------------------------------------

Let's say that we want to create such a project *myapp*:

.. code-block:: console

   $ mkg --flat myapp
   $ cd myap

You may invoke these commands at the root of the project:

* ``make`` compiles the main application
* ``make test`` compiles the main application and run a test against it
* ``make clean`` cleans generated files such as executable and objects

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CC=gcc``
* ``make CC=gcc test``
* ``make CC=gcc clean`` 

------------------------------------------
Nested Console Application Projects for C
------------------------------------------

Let's say that we want to create such a project *myapp*:

.. code-block:: console

   $ mkg myapp
   $ cd myapp

You may invoke these commands at the root of the project:

* ``make`` compiles the main application
* ``make test`` compiles the main application and run a test against it
* ``make clean`` cleans generated files such as executable and objects

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CC=gcc``
* ``make CC=gcc test``
* ``make CC=gcc clean``

*myapp* owns a nested project structure like this:

.. code-block:: console

   $ tree
   .
   ├── dist
   ├── examples
   ├── include
   ├── Makefile
   ├── README.md
   ├── src
   │   ├── Makefile
   │   ├── Makefile.win
   │   └── myapp.c
   └── tests
       ├── myapp.bash
       └── myapp.vbs

* *dist* for generated executable
* *examples* for example code
* *include* for headers
* *src* for application source code
* *tests* for test programs

All these directory destinations are customizable.

------------------------------------------
Flat Console Application Projects for C++
------------------------------------------

Let's say that we want to create such a project *myapp*:

.. code-block:: console

   $ mkg -cxx --flat myapp
   $ cd myap

You may invoke these commands at the root of the project:

* ``make`` compiles the main application
* ``make test`` compiles the main application and run a test against it
* ``make clean`` cleans generated files such as executable and objects

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CXX=g++``
* ``make CXX=g++ test``
* ``make CXX=g++ clean``

--------------------------------------------
Nested Console Application Projects for C++
--------------------------------------------

Let's say that we want to create such a project *myapp*:

.. code-block:: console

   $ mkg -cxx myapp
   $ cd myapp

You may invoke these commands at the root of the project:

* ``make`` compiles the main application
* ``make test`` compiles the main application and run a test against it
* ``make clean`` cleans generated files such as executable and objects

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CXX=g++``
* ``make CXX=g++ test``
* ``make CXX=g++ clean``

*myapp* owns a nested project structure like this:

.. code-block:: console

   $ tree
   .
   ├── dist
   ├── examples
   ├── include
   ├── Makefile
   ├── README.md
   ├── src
   │   ├── Makefile
   │   ├── Makefile.win
   │   └── myapp.cpp
   └── tests
       ├── myapp.bash
       └── myapp.vbs

* *dist* for generated executable
* *examples* for example code
* *include* for headers
* *src* for application source code
* *tests* for test programs

All these directory destinations are customizable.

----------------------------
Flat Library Projects for C
----------------------------

Let's say that we want to create such a project *mylib*:

.. code-block:: console

   $ mkg --library --flat mylib
   $ cd mylib

You may invoke these commands at the root of the project:

* ``make`` or ``make dynamic`` compiles the dynamic library
* ``make static`` compiles the static library
* ``make test`` compiles and tests against the dynamic library
* ``make testStatic`` compiles and tests against the static library
* ``make clean`` cleans generated files

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CC=gcc`` or ``make CC=gcc dynamic``
* ``make CC=gcc static``
* ``make CC=gcc test``
* ``make CC=gcc testStatic``
* ``make CC=gcc clean``

------------------------------
Nested Library Projects for C
------------------------------

Let's say that we want to create such a project *mylib*:

.. code-block:: console

   $ mkg --library mylib
   $ cd mylib

You may invoke these commands at the root of the project:

* ``make`` or ``make dynamic`` compiles the dynamic library
* ``make static`` compiles the static library
* ``make test`` compiles and tests against the dynamic library
* ``make testStatic`` compiles and tests against the static library
* ``make clean`` cleans generated files

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CC=gcc`` or ``make CC=gcc dynamic``
* ``make CC=gcc static``
* ``make CC=gcc test``
* ``make CC=gcc testStatic``
* ``make CC=gcc clean``

*mylib* owns a nested project structure like this:

.. code-block:: console

   $ tree
   .
   ├── dist
   ├── examples
   ├── include
   │   └── mylib.h
   ├── Makefile
   ├── README.md
   ├── src
   │   ├── Makefile
   │   ├── Makefile.win
   │   ├── mylib.c
   │   └── mylib.def
   └── tests
       ├── Makefile
       ├── Makefile.win
       └── mylib_test.c

* *dist* for generated executable
* *examples* for example code
* *include* for headers
* *src* for application source code
* *tests* for test programs

All these directory destinations are customizable.

------------------------------
Flat Library Projects for C++
------------------------------

Let's say that we want to create such a project *mylib*:

.. code-block:: console

   $ mkg --library -cxx --flat mylib
   $ cd mylib

You may invoke these commands at the root of the project:

* ``make`` or ``make dynamic`` compiles the dynamic library
* ``make static`` compiles the static library
* ``make test`` compiles and tests against the dynamic library
* ``make testStatic`` compiles and tests against the static library
* ``make clean`` cleans generated files

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CXX=g++`` or ``make CC=g++ dynamic``
* ``make CXX=g++ static``
* ``make CXX=g++ test``
* ``make CXX=g++ testStatic``
* ``make CXX=g++ clean``

--------------------------------
Nested Library Projects for C++
--------------------------------

Let's say that we want to create such a project *mylib*:

.. code-block:: console

   $ mkg --library -cxx mylib
   $ cd mylib

You may invoke these commands at the root of the project:

* ``make`` or ``make dynamic`` compiles the dynamic library
* ``make static`` compiles the static library
* ``make test`` compiles and tests against the dynamic library
* ``make testStatic`` compiles and tests against the static library
* ``make clean`` cleans generated files

On Windows, the projects assume MSVC; however, MinGW is supported as well:

* ``make CXX=g++`` or ``make CC=g++ dynamic``
* ``make CXX=g++ static``
* ``make CXX=g++ test``
* ``make CXX=g++ testStatic``
* ``make CXX=g++ clean``

*mylib* owns a nested project structure like this:

.. code-block:: console

   $ tree
   .
   ├── dist
   ├── examples
   ├── include
   │   └── mylib.hpp
   ├── Makefile
   ├── README.md
   ├── src
   │   ├── Makefile
   │   ├── Makefile.win
   │   ├── mylib.cpp
   │   └── mylib.def
   └── tests
       ├── Makefile
       ├── Makefile.win
       └── mylib_test.cpp

* *dist* for generated executable
* *examples* for example code
* *include* for headers
* *src* for application source code
* *tests* for test programs

All these directory destinations are customizable.