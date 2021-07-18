==================
Using the Program
==================

You may invoke ``mkg`` in either interactive or batch mode. In the former mode,
you will enter interactive prompts to enter project-related information; in the
latter mode, ``mkg`` will create a project on-th-fly, which is customizable 
through command-line arguments.

By default, ``mkg`` will create a console application project for C with a
sensible project structure on the destination path:

.. code-block:: console

   $ mkg /path/to/myapp

You may run ``mkg`` with different parameters to create different types of projects:

.. code-block:: console

   $ mkg --library -cxx /path/to/mylib

When you run ``mkg`` without any argument, it will enter interactive mode:

.. code-block:: console

   $ mkg
   Program name [myapp]:
   Project path [myapp]:
   Project author [somebody]: Michelle Chen
   Project brief description [something]: Yet Another Application
   Project language (c/cpp) [c]:
   Project type (app/lib) [app]:

   None (none)
   Apache License 2.0 (apache2)
   GNU General Public License v3.0 (gpl3)
   MIT License (mit)
   ---
   BSD 2-clause "Simplified" license (bsd2)
   BSD 3-clause "New" or "Revised" license (bsd3)
   Eclipse Public License 2.0 (epl2)
   GNU Affero General Public License v3.0 (agpl3)
   GNU General Public License v2.0 (gpl2)
   GNU Lesser General Public License v2.1 (lgpl2)
   GNU Lesser General Public License v3.0 (lgpl3)
   Mozilla Public License 2.0 (mpl2)
   The Unlicense (unlicense)

   Project licensing [none]:

You may execute ``mkg`` interactively with more customization:

.. code-block:: console

   $ mkg --custom
   Program name [myapp]:
   Project path [myapp]:
   Project author [somebody]: Michelle Chen
   Project brief description [something]: Yet Another Application
   Project language (c/cpp) [c]:
   Project type (app/lib) [app]:

   (Choose licensing as above...)

   Project structure (nested/flat) [nested]:
   Project source directory [src]:
   Project include directory [include]:
   Project test directory [tests]:
   Project example directory [examples]:
   Project config file [Makefile]:

Currently, ``mkg`` generates the following types of projects:

* Console application projects for C or C++
* Library projects for C or C++
