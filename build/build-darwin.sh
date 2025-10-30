#!/bin/bash
# Build script for macOS version of TridscanUI
# 
# Windows users: This script is for macOS only. 
# For Windows builds, please use build-windows.bat instead.
# If you're trying to prepare this script for macOS use from Windows,
# ensure you convert line endings to LF (not CRLF) before transferring to a Mac.

# Title
echo "Building TridscanUI for macOS..."

# Navigate to the project root directory
cd "$(dirname "$0")/.."


# Check if UPX is installed
UPX_FLAGS=""
if command -v upx >/dev/null 2>&1; then
    echo "UPX is installed. Using UPX for compression."
    UPX_FLAGS="-upx"
else
    echo "UPX is not installed. Skipping UPX compression."
    echo "Note: There are known issues with UPX on Apple Silicon."
fi

# Clean the build directory
if [ -d "build/bin/darwin" ]; then
    rm -rf "build/bin/darwin"
fi

# Detect architecture
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    PLATFORM="darwin/amd64"
elif [ "$ARCH" = "arm64" ]; then
    PLATFORM="darwin/arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Detected platform: $PLATFORM"
echo "Starting build..."

# Create build output directory
mkdir -p build/bin/darwin

# Replace "darwin/" with "-" for the output filename
SAFE_PLATFORM="${PLATFORM/darwin\//"-"}"

# Build the application
wails build -platform $PLATFORM $UPX_FLAGS -clean

if [ $? -ne 0 ]; then
    echo "Build failed for $PLATFORM."
    exit 1
fi

# On macOS, the application is bundled as a .app
# If you want to create a universal binary, uncomment this:
# echo "Building universal binary..."
# wails build -platform darwin/universal $UPX_FLAGS -clean

# Copy the app bundle to the bin/darwin directory
if [ -d "build/bin/TridscanUI.app" ]; then
    cp -R build/bin/TridscanUI.app build/bin/darwin/
    echo "App bundle copied to build/bin/darwin/"
fi

echo "Build complete!"
exit 0
