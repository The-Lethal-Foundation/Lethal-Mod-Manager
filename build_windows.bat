@echo off
SETLOCAL

SET BIN_DIR=%CD%\bin
IF NOT EXIST "%BIN_DIR%" MKDIR "%BIN_DIR%"

SET APPS=app cli

FOR %%a IN (%APPS%) DO (
    ECHO Building %%a...
    go build -ldflags "-H windowsgui" -o "%BIN_DIR%\%%a.exe" %CD%\cmd\%%a
)

ECHO Build complete.
ENDLOCAL
