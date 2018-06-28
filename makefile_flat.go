package main

const makefileAppFlatC = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	for %%x in ($(TEST_PROGRAM) do cscript %%x
else
	for t in $(TEST_PROGRAM); do bats $$t; done
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq ($(CC),cl)
	$(SET_ENV) && $(CC) $(CFLAGS) /Fe:$(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS)
else
	$(CC) $(CFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%.obj: %.c
	$(SET_ENV) && $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c $< 

%.o: %.c
	$(CC) $(CFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const makefileAppFlatCpp = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	for %%x in ($(TEST_PROGRAM) do cscript %%x
else
	for t in $(TEST_PROGRAM); do bats $$t; done
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq ($(CXX),cl)
	$(SET_ENV) && $(CXX) $(CXXFLAGS) /Fe:$(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS) 
else
	$(CXX) $(CXXFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%.obj: %.cpp
	$(SET_ENV) && $(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c $< 

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I. $(INCLUDE) $(LIBS)
`

const makefileLibFlatC = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(INCLUDE) $(LIBS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	$(SET_ENV) && for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(INCLUDE) $(LIBS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.c)) do $(CC) $(CFLAGS) -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS) -c %%x
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS) -o %%x.exe %%x.o
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(INCLUDE) /L. $(LIBS) /c %%x /link $(STATIC_LIB)
	$(SET_ENV) && for %%x in ($(TEST_OBJS)) do $(CC) $(CFLAGS) /I. $(INCLUDE) $(LIBS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=)) do $(CC) $(CFLAGS) -o %%x.exe %%x.c -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CC),cl)
	$(SET_ENV) && for %%x in ($(OBJS:.obj=.c)) do $(CC) $(CFLAGS) /I. $(INCLUDE) $(LIBS) /c %%x
	$(SET_ENV) && link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
else
	for %%x in ($(OBJS:.o=.c)) do $(CC) $(CFLAGS) -fPIC -c %%x -I. $(INCLUDE) -L. $(LIBS)
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. $(INCLUDE) -L. $(LIBS)
endif
else
	for x in $(OBJS:.o=.c); do $(CC) $(CFLAGS) -fPIC -c $$x -I. $(INCLUDE) -L. $(LIBS); done
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) -I. $(INCLUDE) -L. $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o $(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%.obj: %.c
	$(SET_ENV) && $(CC) $(CFLAGS) /I. $(INCLUDE) $(LIBS) /c $<

%.o: %.c
	$(CC) $(CFLAGS) -c $< -I. $(INCLUDE) -L. $(LIBS)
`

const makefileLibFlatCxx = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c %%x /link $(DYNAMIC_LIB:.dll=.lib)
	$(SET_ENV) && for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) %%x /link $(DYNAMIC_LIB:.dll=.lib)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -c %%x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) $(CXXFLAGS) -o %%x.exe %%x.o -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif
else
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif

testStatic: static
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	$(SET_ENV) && for %%x in ($(TEST_OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c %%x /link $(STATIC_LIB)
	$(SET_ENV) && for %%x in ($(TEST_OBJS)) do $(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) %%x /link $(STATIC_LIB)
	for %%x in ($(TEST_OBJS:.obj=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
else
	for %%x in ($(TEST_OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -c %%x $(STATIC_LIB) -I. $(INCLUDE) $(LIBS)
	for %%x in ($(TEST_OBJS:.o=)) do $(CXX) $(CXXFLAGS) -o %%x.exe %%x.o $(STATIC_LIB) -I. $(INCLUDE) $(LIBS)
	for %%x in ($(TEST_OBJS:.o=.exe)) do .\%%x && if %%errorlevel%% neq 0 exit /b %%errorlevel%%
endif  # $(CXX)
else
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.o}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done
endif  # $(detected_OS)

dynamic:
ifeq ($(detected_OS),Windows)
ifeq ($(CXX),cl)
	for %%x in ($(OBJS:.obj=.cpp)) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c %%x
	link /DLL /DEF:$(DYNAMIC_LIB:.dll=.def) /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
else
	for %%x in ($(OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
	$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif  # $(CXX)
else
	for x in $(OBJS:.o=.cpp); do $(CXX) $(CXXFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
	$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif  # $(detected_OS)

static: $(OBJS)
ifeq ($(CXX),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else ifeq ($(detected_OS),Darwin)
	libtool -static -o $(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%.obj: %.cpp
	$(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c $<

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I. $(INCLUDE) -L. $(LIBS)
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
