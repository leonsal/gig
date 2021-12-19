# This Makefile invokes the Makefile in imgui/libs to
# build and install the C/C++ libraries dependencies.
# It should work in Linux and Windows
# In Linux the libraries are installed in /usr/local/lib.
# In Windows they are installed in C:/giglibs.
all:
	${MAKE} -C imgui/libs install

