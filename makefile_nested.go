package main

const makefileProjectStructure = `# Set project structure.
SOURCE_DIR={{.SrcDir}}
INCLUDE_DIR={{.IncludeDir}}
DIST_DIR={{.DistDir}}
TEST_DIR={{.TestDir}}
EXAMPLE_DIR={{.ExampleDir}}

export SOURCE_DIR
export INCLUDE_DIR
export DIST_DIR
export TEST_DIR
export EXAMPLE_DIR
`

const makefileAppNested = `.PHONY: all test run clean

all: run

test: .$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
ifeq ($(detected_OS),Windows)
	for %%x in ($(TEST_PROGRAM) do cscript $(TEST_DIR)/%%x
else
	for t in $(TEST_PROGRAM); do bats $(TEST_DIR)/$$t; done
endif

run: .$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	.$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	echo $$?

.$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR) -f Makefile.win
else
	$(MAKE) -C $(SOURCE_DIR)
endif
`

const makefileLibNested = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(TEST_DIR) -f Makefile.win test
else
	$(MAKE) -C $(TEST_DIR) test
endif

testStatic: static
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(TEST_DIR) -f Makefile.win testStatic
else
	$(MAKE) -C $(TEST_DIR) testStatic
endif

dynamic: .$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)

.$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR) -f Makefile.win
else
	$(MAKE) -C $(SOURCE_DIR)
endif

static: .$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)

.$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB):
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR) -f Makefile.win static
else
	$(MAKE) -C $(SOURCE_DIR) static
endif
`

const makefileAppNestedClean = `clean:
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR) -f Makefile.win clean
else
	$(MAKE) -C $(SOURCE_DIR) clean
endif
	$(RM) $(DIST_DIR)$(SEP)$(PROGRAM)
`

const makefileLibNestedClean = `clean:
ifeq ($(detected_OS),Windows)
	$(MAKE) -C $(SOURCE_DIR) -f Makefile.win clean
	$(MAKE) -C $(TEST_DIR) -f Makefile.win clean
else
	$(MAKE) -C $(SOURCE_DIR) clean
	$(MAKE) -C $(TEST_DIR) clean
endif
	$(RM) $(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
		$(DIST_DIR)$(SEP)$(STATIC_LIB) \
		$(DIST_DIR)$(SEP)$(DYNAMIC_LIB:.dll=.exp)
`

const makefileInternalAppC = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	$(CC) $(CFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalAppCWin = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
ifeq ($(CC),cl)
	$(SET_ENV) && $(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) \
		/Fe:..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS)
else
	$(CC) $(CFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
endif

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%.obj: %.c
	$(SET_ENV) && $(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%.o: %.c
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalAppCxx = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
	$(CXX) $(CXXFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalAppCxxWin = `.SUFFIXES:

.PHONY: all clean

all: ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM)
ifeq ($(CXX),cl)
	$(SET_ENV) && $(CXX) $(CXXFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) \
		/Fe:..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS)
else
	$(CXX) $(CXXFLAGS) -o ..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
endif

..$(SEP)$(DIST_DIR)$(SEP)$(PROGRAM): $(OBJS)

%.obj: %.cpp
	$(SET_ENV) && $(CXX) $(CXXFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalLibC = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
	for x in ` + "`" + `ls *.c` + "`" + `; do $(CC) $(CFLAGS) -fPIC -c $$x \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS); done
	$(CC) $(CFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)

static: $(OBJS)
ifeq ($(detected_OS),Darwin)
	libtool -static -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
endif

%.o: %.c
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalLibCWin = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in (*.c) do $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) \
		/I ..$(SEP)$(INCLUDE_DIR) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
		$(INCLUDE) $(LIBS) $(OBJS)
else
	for %%x in (*.c) do $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) \
		-I ..$(SEP)$(INCLUDE_DIR) -c %%x
	$(CC) $(CFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
		$(OBJS) $(INCLUDE) $(LIBS) -I ..$(SEP)$(INCLUDE_DIR)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
endif

%.obj: %.c
	$(SET_ENV) && $(CC) $(CFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%.o: %.c
	$(CC) $(CFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalLibCxx = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
	for x in ` + "`" + `ls *.cpp` + "`" + `; do $(CXX) $(CXXFLAGS) -fPIC -c $$x \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS); done
	$(CXX) $(CXXFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) $(OBJS) \
		-I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)

static: $(OBJS)
ifeq ($(detected_OS),Darwin)
	libtool -static -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
endif

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalLibCxxWin = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
ifeq ($(CXX),cl)
	$(SET_ENV) && for %%x in (*.cpp) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) \
		/I ..$(SEP)$(INCLUDE_DIR) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
		$(INCLUDE) $(LIBS) $(OBJS)
else
	for %%x in (*.cpp) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) \
		-I ..$(SEP)$(INCLUDE_DIR) -c %%x
	$(CXX) $(CXXFLAGS) -shared -o ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) \
		$(OBJS) $(INCLUDE) $(LIBS) -I ..$(SEP)$(INCLUDE_DIR)
endif

static: $(OBJS)
ifeq ($(CXX),cl)
	lib /out:..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(OBJS)
endif

%.obj: %.cpp
	$(SET_ENV) && $(CXX) $(CXXFLAGS) /I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS) /c $<

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I ..$(SEP)$(INCLUDE_DIR) $(INCLUDE) $(LIBS)
`

const makefileInternalClean = `clean:
	$(RM) $(OBJS)
`

const makefile_internal_lib_test_c = `.PHONY: all test testStatic dynamic static clean
all: test

test: dynamic
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=..$(SEP)$(DIST_DIR) .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

testStatic: static
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

dynamic: ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) dynamic

static: ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) static
`

const makefileInternalLibTestCWin = `.PHONY: all clean
all: test
	
test: dynamic
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in (*.c) do $(CC) $(CFLAGS) \
		$(INCLUDE) $(LIBS) /I..$(SEP)$(INCLUDE_DIR) %%x \
		..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB:.dll=.lib)
	copy ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) . \
		&& for %%x in ($(TEST_OBJS:.obj=.exe)) do .$(SEP)%%x \
		&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -o %%x.exe %%x.c \
		-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
		-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS)
	copy ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) . \
		&& for %%x in ($(TEST_OBJS:.o=.exe)) do .$(SEP)%%x \
		&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif

dynamic: ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)

..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB):
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) -f Makefile.win dynamic

testStatic: $(TEST_OBJS:.obj=.exe)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .$(SEP)%%x \
	&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.exe)) do .$(SEP)%%x \
	&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif

$(TEST_OBJS:.obj=.exe): static
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.c)) do \
		$(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /I..$(SEP)$(INCLUDE_DIR) %%x \
		..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)
else
	for %%x in ($(TEST_OBJS:.o=)) do \
		$(CC) $(CFLAGS) -o %%x.exe %%x.c ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) \
		-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE)
endif

static: ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)

..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB):
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) -f Makefile.win static
`

const makefile_internal_lib_test_cxx = `.PHONY: all test testStatic dynamic static clean
all: test

test: dynamic
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=..$(SEP)$(DIST_DIR) .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

testStatic: static
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x \
			-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
			-L..$(SEP)$(DIST_DIR) -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

dynamic: ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) dynamic

static: ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) static
`

const makefileInternalLibTestCxxWin = `.PHONY: all test dynamic clean

all: test

test: dynamic
ifeq ($(CXX),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.cpp)) do \
		$(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) \
		/I..$(SEP)$(INCLUDE_DIR) %%x \
		..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB:.dll=.lib)
	copy ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) . \
		&& for %%x in ($(TEST_OBJS:.obj=.exe)) do .$(SEP)%%x \
		&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do \
		$(CXX) $(CXXFLAGS) -o %%x.exe %%x.cpp \
		-I..$(SEP)$(INCLUDE_DIR) \
		-L..$(SEP)$(DIST_DIR) -l{{.Program}}
	copy ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB) . \
		&& for %%x in ($(TEST_OBJS:.o=.exe)) do .$(SEP)%%x \
		&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif

dynamic: ..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB)

..$(SEP)$(DIST_DIR)$(SEP)$(DYNAMIC_LIB):
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) -f Makefile.win dynamic

testStatic: $(TEST_OBJS:.obj=.exe)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .$(SEP)%%x \
	&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.exe)) do .$(SEP)%%x \
	&& if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif

$(TEST_OBJS:.obj=.exe): static
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.cpp)) do \
		$(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /I..$(SEP)$(INCLUDE_DIR) %%x \
		..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)
else
	for %%x in ($(TEST_OBJS:.o=)) do \
		$(CXX) $(CXXFLAGS) -o %%x.exe %%x.cpp \
		-I..$(SEP)$(INCLUDE_DIR) $(INCLUDE) \
		..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB) $(LIBS)
endif

static: ..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB)

..$(SEP)$(DIST_DIR)$(SEP)$(STATIC_LIB):
	$(MAKE) -C ..$(SEP)$(SOURCE_DIR) -f Makefile.win static
`

const makefile_internal_lib_test_clean = `clean:
	$(RM) $(TEST_OBJS) $(TEST_OBJS:.o=)
`
const makefileInternalLibTestCleanWin = `clean:
	$(RM) $(TEST_OBJS) $(TEST_OBJS:.obj=.exe) $(TEST_OBJS:.o=.exe) $(DYNAMIC_LIB)
`
