@echo off
setlocal enabledelayedexpansion
set DAY=%~1
mkdir day%DAY%
cd day%DAY%
if not exist "day%DAY%.go" (
    echo package day%DAY% > day%DAY%.go
)

if not exist "day%DAY%_test.go" (
    echo package day%DAY% > day%DAY%_test.go
)

if not exist "example.txt" (
    type nul > example.txt
)

if not exist "input.txt" (
    type nul > input.txt
)