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
	$(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c $< 

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const makefile_lib_flat_c = `.PHONY: all dynamic static clean

all: dynamic

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CC),cl)
		for %%x in (*.c) do $(CC) $(CFLAGS) $(INCLUDE) $(LIBS) /c %%x
		link /DLL /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
	else
		for %%x in (*.c) do $(CC) $(CFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
		$(CC) $(CFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
	endif
else
	for x in ` + "`" + `ls *.c` + "`" + `; do $(CC) $(CFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
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

dynamic:
ifeq ($(detected_OS),Windows)
	ifeq ($(CXX),cl)
		for %%x in (*.cpp) do $(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c %%x
		link /DLL /out:$(DYNAMIC_LIB) $(INCLUDE) $(LIBS) $(OBJS)
	else
		for %%x in (*.cpp) do $(CXX) $(CXXFLAGS) -fPIC -c %%x $(INCLUDE) $(LIBS)
		$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
	endif
else
	for x in ` + "`" + `ls *.cpp` + "`" + `; do $(CXX) $(CXXFLAGS) -fPIC -c $$x $(INCLUDE) $(LIBS); done
	$(CXX) $(CXXFLAGS) -shared -o $(DYNAMIC_LIB) $(OBJS) $(INCLUDE) $(LIBS)
endif

static: $(OBJS)
ifeq ($(CC),cl)
	lib /out:$(STATIC_LIB) $(OBJS)
else
	$(AR) rcs -o $(STATIC_LIB) $(OBJS)
endif

%.obj: %.cpp
	$(CXX) $(CXXFLAGS) $(INCLUDE) $(LIBS) /c $<

%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< $(INCLUDE) $(LIBS)
`

const makefile_app_clean = `clean:
	$(RM) $(PROGRAM) $(OBJS)
`

const makefile_lib_clean = `clean:
	$(RM) $(DYNAMIC_LIB) $(STATIC_LIB) $(OBJS)
`
