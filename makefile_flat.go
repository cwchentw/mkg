package main

const makefileAppFlatC = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	for %%x in ($(TEST_PROGRAM)) do cscript %%x
else
	for t in $(TEST_PROGRAM); do bats $$t; done
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq ($(CC),cl)
	$(CC) /Fe:$(PROGRAM) $(OBJS) $(CFLAGS)
else
	$(CC) -o $(PROGRAM) $(OBJS) $(CFLAGS)
endif

%.obj: %.c
	$(CC) /c $< $(CFLAGS)

%.o: %.c
	$(CC) -c $< $(CFLAGS)
`

const makefileAppFlatCpp = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	for %%x in ($(TEST_PROGRAM)) do cscript %%x
else
	for t in $(TEST_PROGRAM); do bats $$t; done
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq ($(CXX),cl)
	$(CXX) /Fe:$(PROGRAM) $(OBJS) $(CXXFLAGS)
else
	$(CXX) -o $(PROGRAM) $(OBJS) $(CXXFLAGS)
endif

%.obj: %.cpp
	$(CXX) /c $< $(CXXFLAGS)

%.o: %.cpp
	$(CXX) -c $< $(CXXFLAGS)
`

const makefileLibFlatC = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.c)) do $(CC) $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS) -c %%x
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS) -o %%x.exe %%x.o
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. /L. $(LDFLAGS) $(LDLIBS) /c %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -o %%x.exe %%x.c -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(LDFLAGS) $(LDLIBS) $(OBJS)
else
	for %%x in ($(OBJS:.o=.c)) do $(CC) $(CFLAGS) -fPIC -c %%x -I. -L. $(LDFLAGS) $(LDLIBS)
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. -L. $(LDFLAGS) $(LDLIBS)
endif
else
	for x in $(OBJS:.o=.c); do $(CC) $(CFLAGS) -fPIC -c $$x -I. -L. $(LDFLAGS) $(LDLIBS); done
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. -L. $(LDFLAGS) $(LDLIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o $(STATIC_LIB) $(OBJS)
else
	$(AR) rcs $(STATIC_LIB) $(OBJS)
endif

%.obj: %.c
	$(CC) /c $< $(CFLAGS)

%.o: %.c
	$(CC) -c $< $(CFLAGS)
`

const makefileLibFlatCxx = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) /I. $(LDFLAGS) $(LDLIBS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) /I. $(LDFLAGS) $(LDLIBS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -c %%x -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) $(CXXFLAGS) -o %%x.exe %%x.o -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" -I. $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		$(CXX) -o "$${x%.*}" $$x -I. $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS) /c %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) -c %%x $(STATIC_LIB) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) -o %%x.exe %%x.o $(STATIC_LIB) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif  # $(CXX)
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		$(CXX) -o "$${x%.*}" $$x $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		.$(SEP)"$${x%.o}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif  # $(detected_OS)

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
else
	for %%x in ($(OBJS:.o=.cpp)) do $(CXX) -fPIC -c %%x $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
	$(CXX) -shared -o $(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
endif  # $(CXX)
else
	for x in $(OBJS:.o=.cpp); do $(CXX) -fPIC -c $$x $(CXXFLAGS) $(LDFLAGS) $(LDLIBS); done
	$(CXX) -shared -o $(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
endif  # $(detected_OS)

static: $(OBJS)
ifeq ($(CXX),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o $(STATIC_LIB) $(OBJS)
else
	$(AR) rcs $(STATIC_LIB) $(OBJS)
endif

%.obj: %.cpp
	$(CXX) /c $< $(CXXFLAGS)

%.o: %.cpp
	$(CXX) -c $< $(CXXFLAGS)
`

const makefileAppClean = `clean:
	$(RM) $(PROGRAM) $(OBJS)
`

const makefileLibClean = `clean:
	$(RM) $(DYNAMIC_LIB) $(STATIC_LIB) $(OBJS) $(TEST_OBJS)
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	$(RM) $(TEST_OBJS:.obj=.exe) $(TEST_OBJS:.obj=.lib) $(TEST_OBJS:.obj=.exp) $(OBJS:.obj=.exp) $(OBJS:.obj=.lib)
else
	$(RM) $(TEST_OBJS:.o=.exe)
endif
else
	$(RM) $(TEST_OBJS:.o=)
endif
`
