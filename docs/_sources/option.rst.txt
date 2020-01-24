========
Options
========

Here we list the optional parameters in ``mkg``.

-----------------
Program Metadata
-----------------

* ``-v`` or ``--version`` shows version number and exits the program
* ``-h`` or ``--help`` shows help message and exits the program
* ``--licenses`` show the available open-source licenses and exits the program
* ``--standards``: Show available language standards and exit the program

-----------------
Project Metadata
-----------------

* ``-p _prog_`` or ``--prog _prog_`` sets project program name to *prog*, default to directory name
* ``-a _author_`` or ``--author _author_`` sets project author to *author*, default to *somebody*
* ``-b _brief_`` or ``--brief _brief_`` sets project description to *brief*, default to *something*
* ``-o _config_`` or ``--config _config_`` sets project configuration to *config*, default to *Makefile*
* ``-l _license_`` or ``--license _license_`` chooses a open-source *license*, default to *none*

Here are the available licenses in our program:

* Recommended

  * Apache License 2.0 (``apache2``)
  * GNU General Public License v3.0 (``gpl3``)
  * MIT License (``mit``)
  
* Alternative

  * BSD 2-clause "Simplified" license (``bsd2``)
  * BSD 3-clause "New" or "Revised" license (``bsd3``)
  * Eclipse Public License 2.0 (``epl2``)
  * GNU Affero General Public License v3.0 (``agpl3``)
  * GNU General Public License v2.0 (``gpl2``)
  * GNU Lesser General Public License v2.1 (``lgpl2``)
  * GNU Lesser General Public License v3.0 (``lgpl3``)
  * Mozilla Public License 2.0 (``mpl2``)
  * The Unlicense (``unlicense``)

-------------------
Behavior Modifiers
-------------------

* ``-c`` or ``-C`` generates a C project (default)
* ``-cpp`` or ``-cxx`` generates a C++ project
* ``-std _std_`` or ``--standard _std_`` set the language standard to *std*
* ``--console`` generates a console application project (default)
* ``--library`` generates a library project
* ``--nested`` generates a nested project (default)
* ``--flat`` generates a flat project
* ``-f`` or ``--force`` removes all existing contents on path (**Dangerous!**)
* ``--custom`` runs it interactively with more customization

Here are the available language standard for C:

* ``c89`` or ``c90``
* ``c99``
* ``c11`` (default)
* ``c17`` or ``c18``
* ``gnu89`` or ``gnu90``
* ``gnu99``
* ``gnu11``
* ``gnu17`` or ``gnu18``

Due to the limitation from Visual C++, this setting won't take effect when using Visual C++.

Here are the available language standard for C++:

* ``c++98`` or ``c++03``
* ``c++11``
* ``c++14``
* ``c++17`` (default)
* ``gnu++98`` or ``gnu++03``
* ``gnu++11``
* ``gnu++14``
* ``gnu++17``

Due to the limitation from Visual C++, ``mkg`` will automatically set to the most appropriate language standard for C++ when using Visual C++.

------------------
Project Structure
------------------

These parameters only make effects in nested projects.

* ``-s _dir_`` or ``--source _dir_`` sets source directory, default to *src*
* ``-i _dir_`` or ``--include _dir_`` sets include directory, default to *include*
* ``-d _dir_`` or ``--dist _dir_`` sets dist directory, default to *dist*
* ``-t _dir_`` or ``--test _dir_`` sets test programs directory, default to *tests*
* ``-e _dir_`` or ``--example _dir_`` sets example programs directory, default to *examples*
