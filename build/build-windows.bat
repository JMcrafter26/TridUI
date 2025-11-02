@echo off
setlocal enabledelayedexpansion
title Build TridUI for Windows

echo.
echo ========================================
echo   TridUI Windows Build Script
echo ========================================
echo.

:: Step 1: Check for required tools
echo [1/6] Checking Prerequisites...
echo ----------------------------------------

:: Check for Go
where go >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Go is not installed or not in PATH.
    echo.
    echo Please install Go 1.22 or later from https://golang.org/dl/
    echo After installation, restart your terminal and try again.
    goto :error
)

:: Check Go version
for /f "tokens=3" %%i in ('go version') do set GO_VERSION=%%i
echo [OK] Go detected: %GO_VERSION%

:: Validate Go version (basic check for 1.22+)
echo %GO_VERSION% | findstr /R "go1\.2[2-9]\." >nul
if %errorlevel% neq 0 (
    echo %GO_VERSION% | findstr /R "go1\.[3-9][0-9]\." >nul
    if %errorlevel% neq 0 (
        echo [WARNING] Go version may be too old. Wails v2.10.2+ requires Go 1.22.0 or later.
        echo [WARNING] Current version: %GO_VERSION%
        echo.
        timeout /t 3 >nul
    )
)

:: Check for Wails CLI
where wails >nul 2>&1
if %errorlevel% neq 0 (
    echo [ERROR] Wails CLI is not installed or not in PATH.
    echo.
    echo Install it with:
    echo   go install github.com/wailsapp/wails/v2/cmd/wails@latest
    echo.
    echo After installation, ensure %%GOPATH%%\bin is in your PATH.
    goto :error
)
echo [OK] Wails CLI detected

:: Check for Node.js
where node >nul 2>&1
if %errorlevel% neq 0 (
    echo [WARNING] Node.js is not installed or not in PATH.
    echo [WARNING] Frontend dependencies may not build correctly.
    echo.
    timeout /t 2 >nul
) else (
    for /f "tokens=*" %%i in ('node -v') do set NODE_VERSION=%%i
    echo [OK] Node.js detected: !NODE_VERSION!
)

:: Check for pnpm
where pnpm >nul 2>&1
if %errorlevel% neq 0 (
    echo [WARNING] pnpm is not installed or not in PATH.
    echo [WARNING] Install it with: npm install -g pnpm
    echo.
    timeout /t 2 >nul
) else (
    for /f "tokens=*" %%i in ('pnpm -v') do set PNPM_VERSION=%%i
    echo [OK] pnpm detected: !PNPM_VERSION!
)

echo.

:: Navigate to project root
cd /d %~dp0\..
if %errorlevel% neq 0 (
    echo [ERROR] Failed to navigate to project root directory.
    goto :error
)

echo [2/6] Verifying Project Structure...
echo ----------------------------------------

:: Verify essential files exist
if not exist "wails.json" (
    echo [ERROR] wails.json not found. Are you in the TridUI project directory?
    goto :error
)
echo [OK] Project structure verified

if not exist "frontend\package.json" (
    echo [WARNING] frontend\package.json not found. Frontend build may fail.
    timeout /t 2 >nul
)

echo.

:: Step 2: Check for optional tools
echo [3/6] Detecting Optional Build Tools...
echo ----------------------------------------

set "useUpx="
set "useNsis="

:: Check if UPX is installed
where upx >nul 2>&1
if %errorlevel% neq 0 (
    echo [INFO] UPX not found - binaries will not be compressed
    echo       Install from: https://github.com/upx/upx/releases
) else (
    set "useUpx=-upx"
    for /f "tokens=*" %%i in ('upx --version 2^>^&1 ^| findstr /R "^upx"') do set UPX_VERSION=%%i
    echo [OK] UPX detected - compression enabled
    echo     !UPX_VERSION!
)

:: Check if NSIS is installed
where makensis >nul 2>&1
if %errorlevel% neq 0 (
    echo [INFO] NSIS not found - installer will not be created
    echo       Install from: https://nsis.sourceforge.io/Download
) else (
    set "useNsis=-nsis"
    for /f "tokens=2" %%i in ('makensis /VERSION') do set NSIS_VERSION=%%i
    echo [OK] NSIS detected - installer creation enabled
    echo     Version: !NSIS_VERSION!
)

echo.

:: Step 3: Prepare build environment
echo [4/6] Preparing Build Environment...
echo ----------------------------------------

:: Clean previous builds
if exist "build\bin\windows" (
    echo Cleaning previous build artifacts...
    rmdir /s /q "build\bin\windows" 2>nul
    if %errorlevel% neq 0 (
        echo [WARNING] Could not fully clean build directory. Some files may be in use.
    )
)

:: Create output directory
mkdir "build\bin\windows" 2>nul
if not exist "build\bin\windows" (
    echo [ERROR] Failed to create output directory: build\bin\windows
    goto :error
)
echo [OK] Build directory prepared

echo.

:: Step 4: Detect system architecture
echo [5/6] Detecting System Architecture...
echo ----------------------------------------

set "ARCH=%PROCESSOR_ARCHITECTURE%"
if "%ARCH%"=="AMD64" (
    set "PLATFORM=windows/amd64"
    set "ARCH_NAME=amd64"
) else if "%ARCH%"=="ARM64" (
    set "PLATFORM=windows/arm64"
    set "ARCH_NAME=arm64"
) else (
    echo [ERROR] Unsupported architecture: %ARCH%
    echo        Only AMD64 and ARM64 are supported.
    goto :error
)

echo [OK] Building for: %PLATFORM%
echo.

:: Step 5: Build the application
echo [6/6] Building Application...
echo ----------------------------------------
echo This may take several minutes on first build...
echo.

:: Build with Wails CLI
set "BUILD_CMD=wails build -platform %PLATFORM% %useUpx% %useNsis% -clean"
echo Command: %BUILD_CMD%
echo.

%BUILD_CMD%

if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Build failed. Check the output above for details.
    echo.
    echo Common issues:
    echo - Frontend dependencies not installed: cd frontend ^&^& pnpm install
    echo - Outdated Wails CLI: go install github.com/wailsapp/wails/v2/cmd/wails@latest
    echo - Missing system dependencies
    goto :error
)

echo.
echo [OK] Build completed successfully!
echo.

:: Verify output files
echo ========================================
echo   Build Summary
echo ========================================
echo.

set "OUTPUT_COUNT=0"

if exist "build\bin\TridUI.exe" (
    echo [OK] Executable: build\bin\TridUI.exe
    set /a OUTPUT_COUNT+=1
    
    :: Move to windows directory with proper naming
    move /Y "build\bin\TridUI.exe" "build\bin\windows\TridUI-win-%ARCH_NAME%.exe" >nul 2>&1
    if %errorlevel% equ 0 (
        echo     Moved to: build\bin\windows\TridUI-win-%ARCH_NAME%.exe
    )
)

if exist "build\bin\TridUI-%ARCH_NAME%-installer.exe" (
    echo [OK] Installer: build\bin\TridUI-%ARCH_NAME%-installer.exe
    set /a OUTPUT_COUNT+=1
    
    :: Move installer
    move /Y "build\bin\TridUI-%ARCH_NAME%-installer.exe" "build\bin\windows\TridUI-win-%ARCH_NAME%-installer.exe" >nul 2>&1
    if %errorlevel% equ 0 (
        echo     Moved to: build\bin\windows\TridUI-win-%ARCH_NAME%-installer.exe
    )
)

echo.
echo Output directory: build\bin\windows\
echo Total files: %OUTPUT_COUNT%
echo.

if %OUTPUT_COUNT% equ 0 (
    echo [WARNING] No output files detected. Build may have failed silently.
    goto :error
)

:: Additional information
echo ========================================
echo   Next Steps
echo ========================================
echo.
echo - Test the application: build\bin\windows\TridUI-win-%ARCH_NAME%.exe
echo - WebView2 Runtime required for end users
echo - Distribute the installer for easier deployment
echo.
echo Build completed at: %date% %time%
echo.

goto :end

:error
echo.
echo ========================================
echo   BUILD FAILED
echo ========================================
echo.
echo Please review the errors above and:
echo 1. Ensure all prerequisites are installed
echo 2. Check the Wails documentation: https://wails.io
echo 3. Review the build logs for specific errors
echo.
exit /b 1

:end
echo Press any key to exit...
pause >nul
exit /b 0
