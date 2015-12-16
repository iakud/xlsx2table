cmake_minimum_required(VERSION 2.6)

set(CXX_FLAGS
	-g
	-DCHECK_PTHREAD_RETURN_VALUE
	-D_FILE_OFFSET_BITS=64
	-Wall
	-Wextra
	-Werror
	-Wconversion
	-Wno-unused-parameter
	-Wold-style-cast
	-Woverloaded-virtual
	-Wpointer-arith
	-Wshadow
	-Wwrite-strings
	-march=native
	-rdynamic
	)

string(REPLACE ";" " " CMAKE_CXX_FLAGS "${CXX_FLAGS}")

set(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/bin)
set(LIBRARY_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/lib)

include_directories(${PROJECT_SOURCE_DIR}/include)
aux_source_directory(${PROJECT_SOURCE_DIR}/src SOURCE_FILES)

add_library(table ${SOURCE_FILES})

if (NOT CMAKE_BUILD_NO_EXAMPLES)
	add_subdirectory(example)
endif()