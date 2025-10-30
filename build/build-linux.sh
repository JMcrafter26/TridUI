#!/bin/bash
# Build script for Linux version of TridUI
# 
# This script builds the Linux version of TridUI.
# Ensure this script has execute permissions with: chmod +x build-linux.sh

# Title
echo "Building TridUI for Linux..."

# Navigate to the project root directory
cd "$(dirname "$0")/.."

# Check if UPX is installed
UPX_FLAGS=""
if command -v upx >/dev/null 2>&1; then
    echo "UPX is installed. Using UPX for compression."
    UPX_FLAGS="-upx"
else
    echo "UPX is not installed. Skipping UPX compression."
fi

# Clean the build directory
if [ -d "build/bin/linux" ]; then
    rm -rf "build/bin/linux"
fi

# Create build output directory
mkdir -p build/bin/linux

# Detect architecture
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    PLATFORM="linux/amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
    PLATFORM="linux/arm64"
else
    echo "Unsupported architecture: $ARCH"
    exit 1
fi

echo "Detected platform: $PLATFORM"
echo "Starting build..."

# Replace "linux/" with "-" for the output filename
SAFE_PLATFORM="${PLATFORM/linux\//"-"}"

# Build the application
wails build -platform $PLATFORM $UPX_FLAGS -clean

if [ $? -ne 0 ]; then
    echo "Build failed for $PLATFORM."
    exit 1
fi

# Move the binary to the linux directory
if [ -f "build/bin/TridUI" ]; then
    echo "Moving binary to build/bin/linux/"
    cp "build/bin/TridUI" "build/bin/linux/TridUI$SAFE_PLATFORM"
    # Make sure it's executable
    chmod +x "build/bin/linux/TridUI$SAFE_PLATFORM"
    echo "Binary copied to build/bin/linux/TridUI$SAFE_PLATFORM"
fi

echo "Build complete!"
echo ""
echo "Note: On Linux, you may need the following dependencies to run the application:"
echo "- libgtk-3"
echo "- libwebkit2gtk-4.0"
echo "- gstreamer1.0-plugins-good (for audio/video playback)"
echo ""
echo "These can be installed with your distribution's package manager."
echo "For example, on Ubuntu/Debian: sudo apt install libgtk-3-0 libwebkit2gtk-4.0-37 gstreamer1.0-plugins-good"
echo "On Fedora: sudo dnf install gtk3 webkit2gtk4.0 gstreamer1-plugins-good"
echo "On Arch: sudo pacman -S gtk3 webkit2gtk gst-plugins-good"

exit 0
