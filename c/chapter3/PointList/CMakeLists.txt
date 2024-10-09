cmake_minimum_required(VERSION 3.10)

# 프로젝트 이름 설정
project(PointList C)

# C 표준 설정 (필요에 따라 수정 가능)
set(CMAKE_C_STANDARD 99)

# include 디렉토리를 include 경로에 추가
include_directories(app/include)

set(SOURCES
    app/main.c
    app/src/ArrayList.c
    app/src/Point.c)

add_executable(PointList ${SOURCES})