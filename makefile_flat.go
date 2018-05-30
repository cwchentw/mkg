package main

const makefile_app_flat_c = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	@echo "Unsupported"
else
	./$(PROGRAM).bash
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq (($CC),cl)
	$(CC) $(CFLAGS) /Fe $(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS)
else
	$(CC) $(CFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%.obj: %.c
	$(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c $< 

%.o: %.c
	$(CC) $(CFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const makefile_app_flat_cpp = `.PHONY: all clean

all: run

test: $(PROGRAM)
ifeq ($(detected_OS),Windows)
	@echo "Unsupported"
else
	./$(PROGRAM).bash
endif

run: $(PROGRAM)
	.$(SEP)$(PROGRAM)
	echo $$?

$(PROGRAM): $(OBJS)
ifeq (($CXX),cl)
	$(CXX) $(CXXFLAGS) /Fe $(PROGRAM) $(INCLUDE) $(LIBS) $(OBJS) 
else
	$(CXX) $(CXXFLAGS) -o $(PROGRAM) $(OBJS) $(INCLUDE) $(LIBS)
endif

%.obj: %.cpp
	$(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c $< 

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I. $(INCLUDE) $(LIBS)
`

const makefile_lib_flat_c = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic $(TEST_OBJS)
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

testStatic: static $(TEST_OBJS)
	for x in $(TEST_OBJS); do \
		$(CC) $(CFLAGS) -c "$${x%.*}.c" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CC) $(CFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CC),cl)
		for %%x in ($(OBJS:.o=.c)) do $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c %%x
		link /DLL /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
	else
		for %%x in ($(OBJS:.o=.c)) do $(CC) $(CFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
		$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
	endif
else
	for x in $(OBJS:.o=.c); do $(CC) $(CFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
	$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%.obj: %.c
	$(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c $<

%.o: %.c
	$(CC) $(CFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const makefile_lib_flat_cxx = `.PHONY: all dynamic static clean

all: dynamic

test: dynamic
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		LD_LIBRARY_PATH=. .$(SEP)"$${x%.*}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

testStatic: static
	for x in $(TEST_OBJS); do \
		$(CXX) $(CXXFLAGS) -c "$${x%.*}.cpp" -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		$(CXX) $(CXXFLAGS) -o "$${x%.*}" $$x -I. $(INCLUDE) -L. -l{{.Program}} $(LIBS); \
		.$(SEP)"$${x%.o}"; \
		if [ $$? -ne 0 ]; then echo "Failed program state"; exit 1; fi \
	done

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CXX),cl)
		for %%x in ($(OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c %%x
		link /DLL /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
	else
		for %%x in ($(OBJS:.o=.cpp)) do $(CXX) $(CXXFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
		$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
	endif
else
	for x in $(OBJS:.o=.cpp); do $(CXX) $(CXXFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
	$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%.obj: %.cpp
	$(CXX) $(CXXFLAGS) /I. $(INCLUDE) $(LIBS) /c $<

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -I. $(INCLUDE) $(LIBS)
`

const makefile_app_clean = `clean:
	$(RM) $(PROGRAM) $(OBJS)
`

const makefile_lib_clean = `clean:
	$(RM) $(DYNAMIC_LIB) $(STATIC_LIB) $(OBJS)
`
