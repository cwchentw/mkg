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

const config_app_nested_c = `.PHONY: all run clean

all: run

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

const config_app_nested_clean = `clean:
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR)$(SEP)Makefile.win clean
else
	$(MAKE) -C $(SOURCE_DIR) clean
endif
	$(RM) $(DIST_DIR)$(SEP)$(PROGRAM)
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

const config_internal_clean = `clean:
	$(RM) $(OBJS)
`
