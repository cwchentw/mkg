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
	$(CC) /Fe:$(PROGRAM) $(OBJS) $(CFLAGS) $(LDFLAGS)
else
	$(CC) -o $(PROGRAM) $(OBJS) $(CFLAGS) $(LDFLAGS)
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
	$(CXX) /Fe:$(PROGRAM) $(OBJS) $(CXXFLAGS) $(LDFLAGS)
else
	$(CXX) -o $(PROGRAM) $(OBJS) $(CXXFLAGS) $(LDFLAGS)
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
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.c)) do $(CC) $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) -c %%x
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) -o %%x.exe %%x.o
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. -L. -l{{.Program}} $(LDFLAGS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. -L. -l{{.Program}} $(LDFLAGS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. /L. $(LDFLAGS) /c %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -o %%x.exe %%x.c -I. -L. -l{{.Program}} $(LDFLAGS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. -L. -l{{.Program}} $(LDFLAGS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. -L. -l{{.Program}} $(LDFLAGS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(LDFLAGS) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(LDFLAGS) $(OBJS)
else
	for %%x in ($(OBJS:.o=.c)) do $(CC) $(CFLAGS) -fPIC -c %%x -I. -L. $(LDFLAGS)
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. -L. $(LDFLAGS)
endif
else
	for x in $(OBJS:.o=.c); do $(CC) $(CFLAGS) -fPIC -c $$x -I. -L. $(LDFLAGS); done
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. -L. $(LDFLAGS)
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
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) /I. $(LDFLAGS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) /I. $(LDFLAGS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -c %%x -I. -L. -l{{.Program}} $(LDFLAGS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) $(CXXFLAGS) -o %%x.exe %%x.o -I. -L. -l{{.Program}} $(LDFLAGS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" -I. $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS); \
		$(CXX) -o "$${x%.*}" $$x -I. $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) /c %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) -c %%x $(STATIC_LIB) $(CXXFLAGS) $(LDFLAGS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) -o %%x.exe %%x.o $(STATIC_LIB) $(CXXFLAGS) $(LDFLAGS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif  # $(CXX)
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" $(CXXFLAGS) -L. -l{{.Program}} $(LDFALGS); \
		$(CXX) -o "$${x%.*}" $$x $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS); \
		.$(SEP)"$${x%.o}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif  # $(detected_OS)

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) $(LDFLAGS) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS)
else
	for %%x in ($(OBJS:.o=.cpp)) do $(CXX) -fPIC -c %%x $(CXXFLAGS) $(LDFLAGS)
	$(CXX) -shared -o $(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS)
endif  # $(CXX)
else
	for x in $(OBJS:.o=.cpp); do $(CXX) -fPIC -c $$x $(CXXFLAGS) $(LDFLAGS); done
	$(CXX) -shared -o $(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS)
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
