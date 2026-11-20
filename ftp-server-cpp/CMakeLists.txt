cmake_minimum_required(VERSION 3.28)
project(ftp_server_cpp CXX)

find_package(Boost REQUIRED CONFIG COMPONENTS headers system)

add_executable(ftp_server_cpp src/main.cpp)
target_compile_features(ftp_server_cpp PRIVATE cxx_std_23)
set(CMAKE_CXX_EXTENSIONS OFF)
target_link_libraries(ftp_server_cpp PRIVATE Boost::headers Boost::system)
