@echo off
setlocal enabledelayedexpansion
title Build TridUI for Windows

echo ===== TridUI-Wails Windows Build Script =====
echo.

:: Check for required tools
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: Go is not installed or not in PATH.
    echo Please install Go 1.18 or later from https://golang.org/dl/
    goto :error
)

where wails >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: Wails CLI is not installed or not in PATH.
    echo Install it with: go install github.com/wailsapp/wails/v2/cmd/wails@latest
    goto :error
)

where deno >nul 2>&1
if %errorlevel% neq 0 (
    echo WARNING: Deno is not installed or not in PATH.
    echo The sandbox host build may fail.
    echo Install Deno from https://deno.land/
    echo.
    timeout /t 3 >nul
)

:: Navigate to script directory
cd /d %~dp0\..\

@REM echo Current Directory:
echo %cd%

:: Initialize build parameters
set useUpx=""
set useNsis=""

:: Check for optional tools
echo Step 2: Checking for Optional Tools...
echo ------------------------------

:: Check if UPX is installed
where upx >nul 2>&1
if %errorlevel% neq 0 (
    echo UPX is not installed. Binaries will not be compressed.
    echo To enable compression, install UPX from: https://github.com/upx/upx/releases
) else (
    set useUpx="-upx"
    echo UPX is installed. Binaries will be compressed.
)

:: Check if NSIS is installed
where makensis >nul 2>&1
if %errorlevel% neq 0 (
    echo NSIS is not installed. No installer will be created.
    echo To enable installer creation, install NSIS from: https://nsis.sourceforge.io/Download
) else (
    set useNsis="-nsis"
    echo NSIS is installed. An installer will be created.
)

echo.

:: Clean the build directory
echo Step 3: Preparing Build Environment...
echo ------------------------------

if exist "build\bin\windows" (
    echo Cleaning previous build files...
    rmdir /s /q "build\bin\windows"
)

:: Create build directory structure
mkdir "build\bin\windows" 2>nul

:: Detect system architecture
echo Step 4: Detecting System Architecture...
echo ------------------------------

set "ARCH=%PROCESSOR_ARCHITECTURE%"
if "%ARCH%"=="AMD64" (
    set "PLATFORM=windows/amd64"
) else if "%ARCH%"=="ARM64" (
    set "PLATFORM=windows/arm64"
) else (
    echo ERROR: Unsupported architecture: %ARCH%
    echo Only AMD64 and ARM64 architectures are supported.
    goto :error
)

echo Detected platform: %PLATFORM%
echo.


echo Step 5: Building Application...
echo ------------------------------
echo Building TridUI for %PLATFORM%...

:: Format the output filename by replacing "windows/" with "-"
set "SAFE_PLATFORM=%PLATFORM:windows/=-%"

:: Build options
set "BUILD_OPTS=-platform %PLATFORM% %useUpx% %useNsis% --o windows\TridUI%SAFE_PLATFORM%.exe -clean"

:: Execute build command with a more descriptive message
echo Running: wails build %BUILD_OPTS%
wails build %BUILD_OPTS%

:: Check for build errors
if %errorlevel% neq 0 (
    echo ERROR: Build failed for %PLATFORM%.
    goto :error
)

:: Create output directory if it doesn't exist (just in case)
if not exist "build\bin\windows" (
    mkdir "build\bin\windows"
)

:: Verify and relocate the binary
if exist "build\bin\TridUI%SAFE_PLATFORM%.exe" (
    echo Moving binary to build\bin\windows\...
    move "build\bin\TridUI%SAFE_PLATFORM%.exe" "build\bin\windows\TridUI%SAFE_PLATFORM%.exe" >nul 2>&1
    if %errorlevel% neq 0 (
        echo WARNING: Failed to move the executable. It may already be in the correct location.
    )
)

:: Move the installer if it was created
if exist "build\bin\TridUI%SAFE_PLATFORM%-installer.exe" (
    echo Moving installer to build\bin\windows\...
    move "build\bin\TridUI%SAFE_PLATFORM%-installer.exe" "build\bin\windows\TridUI%SAFE_PLATFORM%-installer.exe" >nul 2>&1
    if %errorlevel% neq 0 (
        echo WARNING: Failed to move the installer. It may already be in the correct location.
    )
)

:: Copy the sandbox host to the Windows directory if it exists
if exist "deno_host\compiled\sandbox_host.exe" (
    echo Copying Deno Sandbox Host to build\bin\windows\...
    copy "deno_host\compiled\sandbox_host.exe" "build\bin\windows\" >nul 2>&1
    if %errorlevel% neq 0 (
        echo WARNING: Failed to copy the sandbox host.
    )

    :: Zip compressed sandbox host
    echo Zipping Deno Sandbox Host...
    powershell -command "Compress-Archive -Path 'build\bin\windows\sandbox_host.exe' -DestinationPath 'build\bin\windows\sandbox_host_windows.zip' -Force" >nul 2>&1
)

:: Display successful completion message
echo.
echo ===== Build Process Complete =====
echo.
echo Build completed successfully!
echo.
echo Output files:
echo - Executable: build\bin\windows\TridUI%SAFE_PLATFORM%.exe

:: Check and report if installer was created
if exist "build\bin\windows\TridUI%SAFE_PLATFORM%-installer.exe" (
    echo - Installer: build\bin\windows\TridUI%SAFE_PLATFORM%-installer.exe
)

:: Check and report if sandbox host was copied
if exist "build\bin\windows\sandbox_host.exe" (
    echo - Sandbox Host: build\bin\windows\sandbox_host.exe
)

echo.
echo NOTE: If you distribute this application, users will need the WebView2 Runtime.
echo Wails will prompt users to install it if needed.
echo.

:: Exit successfully
goto :end

:error
echo.
echo ===== Build Process Failed =====
echo.
echo Please fix the errors above and try again.
echo.
exit /b 1

:end
:: Wait before exiting if script was double-clicked
if "%1"=="" (
    echo Press any key to exit...
    timeout /t 5 >nul
)
exit /b 0