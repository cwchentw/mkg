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
	$(CC) /Fe:$(PROGRAM) $(OBJS) $(CFLAGS) $(LDFLAGS) $(LDLIBS)
else
	$(CC) -o $(PROGRAM) $(OBJS) $(CFLAGS) $(LDFLAGS) $(LDLIBS)
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
	$(CXX) /Fe:$(PROGRAM) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
else
	$(CXX) -o $(PROGRAM) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
endif

%.obj: %.cpp
	$(CXX) /c $< $(CXXFLAGS)

%.o: %.cpp
	$(CXX) -c $< $(CXXFLAGS)
`

const makefileLibFlatC = `DYNAMIC := all test dynamic


.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /D{{.PROGRAM}}_IMPORT_SYMBOLS /MD /I. /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CC) %%x $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.c)) do $(CC) -c %%x $(CFLAGS) -I.
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) -o %%x.exe %%x.o $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS) 
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) -c "$${x%.*}.c" -I. $(CFLAGS); \
		$(CC) -o "$${x%.*}" $(CFLAGS) $$x -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) /c %%x $(CFLAGS) /I. /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CC) %%x $(CFLAGS) /I. $(LDFLAGS) $(LDLIBS) /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) -o %%x.exe %%x.c $(CC) $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) -c "$${x%.*}.c" $(CFLAGS) -I.; \
		$(CC) -o "$${x%.*}" $$x $(CFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

dynamic: $(OBJS)
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	link /DLL /out:$(DYNAMIC_LIB) $(LDFLAGS) $(LDLIBS) $(OBJS)
else
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. -L. $(LDFLAGS) $(LDLIBS)
endif
else
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
ifneq (,$(findstring $(MAKECMDGOALS),$(DYNAMIC)))
	$(CC) /c $< $(CFLAGS) /D{{.PROGRAM}}_EXPORT_SYMBOLS /MD
else
	$(CC) /c $< $(CFLAGS) /MT
endif

%.o: %.c
ifneq (,$(findstring $(MAKECMDGOALS),$(DYNAMIC)))
	$(CC) -c $< $(CFLAGS) -fPIC
else
	$(CC) -c $< $(CFLAGS)
endif
`

const makefileLibFlatCxx = `DYNAMIC := all test dynamic


.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) /c %%x $(CXXFLAGS) /D{{.PROGRAM}}_IMPORT_SYMBOLS /MD /I. /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS)) do $(CXX) %%x $(CXXFLAGS) /I. $(LDFLAGS) $(LDLIBS) /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) -c %%x $(CXXFLAGS) -I.
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) -o %%x.exe %%x.o $(CXXFLAGS) -I. -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" -I. $(CXXFLAGS); \
		$(CXX) -o "$${x%.*}" $$x -I. $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) /c %%x $(CXXFLAGS) /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS)) do $(CXX) %%x $(CXXFLAGS) $(LDFLAGS) $(LDLIBS) /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) -c %%x $(STATIC_LIB) $(CXXFLAGS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) -o %%x.exe %%x.o $(STATIC_LIB) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif  # $(CXX)
else
	for x in $(TEST_OBJS); do \
		$(CXX) -c "$${x%.*}.cpp" $(CXXFLAGS); \
		$(CXX) -o "$${x%.*}" $$x $(CXXFLAGS) -L. -l{{.Program}} $(LDFLAGS) $(LDLIBS); \
		.$(SEP)"$${x%.o}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif  # $(detected_OS)

dynamic: $(OBJS)
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	link /DLL /out:$(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
else
	$(CXX) -shared -o $(DYNAMIC_LIB) $(OBJS) $(CXXFLAGS) $(LDFLAGS) $(LDLIBS)
endif  # $(CXX)
else
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
ifneq (,$(findstring $(MAKECMDGOALS),$(DYNAMIC)))
	$(CXX) /c $< $(CXXFLAGS) /D{{.PROGRAM}}_EXPORT_SYMBOLS /MD
else
	$(CXX) /c $< $(CXXFLAGS) /MT
endif

%.o: %.cpp
ifneq (,$(findstring $(MAKECMDGOALS),$(DYNAMIC)))
	$(CXX) -c $< $(CXXFLAGS) -fPIC
else
	$(CXX) -c $< $(CXXFLAGS)
endif
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
