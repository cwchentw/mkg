package main

const config_project_structure = `# Set project structure.
SOURCE_DIR=%s
INCLUDE_DIR=%s
DIST_DIR=%s
TEST_DIR=%s
EXAMPLE_DIR=%s

export SOURCE_DIR
export INCLUDE_DIR
export DIST_DIR
export TEST_DIR
export EXAMPLE_DIR
`

const config_app_nested = `.PHONY: all test run clean

all: run

test: .$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	bats $(TEST_DIR)/$(PROGRAM).bash
	echo $$?

run: .$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	.$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	echo $$?

.$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win
else
	$(MAKE) -C $(SOURCE_DIR)
endif
`

const config_lib_nested = `.PHONY: all dynamic static clean

all: dynamic

dynamic: .$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)

.$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win
else
	$(MAKE) -C $(SOURCE_DIR)
endif

static: .$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)

.$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win static
else
	$(MAKE) -C $(SOURCE_DIR) static
endif
`

const config_app_nested_clean = `clean:
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win clean
else
	$(MAKE) -C $(SOURCE_DIR) clean
endif
	$(RM) $(DIST_DIR)$(SEP)$(PROGRAM)
`

const config_lib_nested_clean = `clean:
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win clean
else
	$(MAKE) -C $(SOURCE_DIR) clean
endif
	$(RM) $(DIST_DIR)$(SEP)$(DYNAMIC_LIB)
	$(RM) $(DIST_DIR)$(SEP)$(STATIC_LIB)
`

const config_internal_app_c = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
ifeq ($(CC),cl)
	$(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) \
		/Fe ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS)
else
	$(CC) $(CFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
endif

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%s: %s
	$(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%s: %s
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const config_internal_app_cxx = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
ifeq ($(CXX),cl)
	$(CXX) $(CXXFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) \
		/Fe ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS)
else
	$(CXX) $(CXXFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
endif

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%s: %s
	$(CXX) $(CXXFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%s: %s
	$(CXX) $(CXXFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const config_internal_lib_c = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CC),cl)
		for %%x in (*.c) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) \
			/I ..$(SEP)$(INCLUDE_DIR) /c %%x
		link /DLL /out:..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
			$(INCLUDE) $(LIBS) /I ..$(SEP)$(INCLUDE_DIR) $(OBJS)
	else
		for %%x in (*.c) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) \
			-I ..$(SEP)$(INCLUDE_DIR) /c %%x
		$(CC) $(CFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
			$(OBJS) $(INCLUDE) $(LIBS) -I ..$(SEP)$(INCLUDE_DIR)
	endif
else
	for x in ` + "`" + `ls *.c` + "`" + `; do $(CC) $(CFLAGS) -fPIC -c $$x \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS); done
	$(CC) $(CFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /I ..$(SEP)$(INCLUDE_DIR) /out:..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) \
		$(OBJS)
else
	$(AR) rcs -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
endif

%s: %s
	$(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%s: %s
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const config_internal_clean = `clean:
	$(RM) $(OBJS)
`
