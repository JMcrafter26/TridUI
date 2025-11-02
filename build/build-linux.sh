#!/bin/bash
# TridUI Linux Build Script
# Builds TridUI with comprehensive validation and error handling

set -e  # Exit on error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
print_header() {
    echo ""
    echo "========================================"
    echo "  $1"
    echo "========================================"
    echo ""
}

print_step() {
    echo ""
    echo -e "${BLUE}[$1]${NC} $2"
    echo "----------------------------------------"
}

print_success() {
    echo -e "${GREEN}[OK]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_info() {
    echo -e "[INFO] $1"
}

# Start build
print_header "TridUI Linux Build Script"

# Step 1: Check prerequisites
print_step "1/6" "Checking Prerequisites..."

# Check for Go
if ! command -v go >/dev/null 2>&1; then
    print_error "Go is not installed."
    echo ""
    echo "Please install Go 1.22 or later from https://golang.org/dl/"
    echo "After installation, restart your terminal and try again."
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}')
print_success "Go detected: $GO_VERSION"

# Validate Go version (basic check for 1.22+)
GO_MAJOR=$(echo $GO_VERSION | sed 's/go//' | cut -d. -f1)
GO_MINOR=$(echo $GO_VERSION | sed 's/go//' | cut -d. -f2)

if [ "$GO_MAJOR" -lt 1 ] || ([ "$GO_MAJOR" -eq 1 ] && [ "$GO_MINOR" -lt 22 ]); then
    print_warning "Go version may be too old. Wails v2.10.2+ requires Go 1.22.0 or later."
    print_warning "Current version: $GO_VERSION"
    sleep 2
fi

# Check for Wails CLI
if ! command -v wails >/dev/null 2>&1; then
    print_error "Wails CLI is not installed."
    echo ""
    echo "Install it with:"
    echo "  go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    echo ""
    echo "After installation, ensure \$GOPATH/bin is in your PATH."
    exit 1
fi
print_success "Wails CLI detected"

# Check for Node.js
if command -v node >/dev/null 2>&1; then
    NODE_VERSION=$(node -v)
    print_success "Node.js detected: $NODE_VERSION"
else
    print_warning "Node.js is not installed."
    print_warning "Frontend dependencies may not build correctly."
    sleep 2
fi

# Check for pnpm
if command -v pnpm >/dev/null 2>&1; then
    PNPM_VERSION=$(pnpm -v)
    print_success "pnpm detected: $PNPM_VERSION"
else
    print_warning "pnpm is not installed."
    print_warning "Install it with: npm install -g pnpm"
    sleep 2
fi

# Navigate to project root
cd "$(dirname "$0")/.."

# Step 2: Verify project structure
print_step "2/6" "Verifying Project Structure..."

if [ ! -f "wails.json" ]; then
    print_error "wails.json not found. Are you in the TridUI project directory?"
    exit 1
fi
print_success "Project structure verified"

if [ ! -f "frontend/package.json" ]; then
    print_warning "frontend/package.json not found. Frontend build may fail."
    sleep 2
fi

# Step 3: Check for system dependencies
print_step "3/6" "Checking System Dependencies..."

# Check for GTK3
if ! pkg-config --exists gtk+-3.0; then
    print_error "GTK3 development files not found."
    echo ""
    echo "Install with:"
    echo "  Ubuntu/Debian: sudo apt install libgtk-3-dev"
    echo "  Fedora:        sudo dnf install gtk3-devel"
    echo "  Arch:          sudo pacman -S gtk3"
    exit 1
fi
print_success "GTK3 detected"

# Check for WebKit2GTK
if pkg-config --exists webkit2gtk-4.1; then
    print_success "WebKit2GTK 4.1 detected"
elif pkg-config --exists webkit2gtk-4.0; then
    print_success "WebKit2GTK 4.0 detected"
else
    print_error "WebKit2GTK development files not found."
    echo ""
    echo "Install with:"
    echo "  Ubuntu/Debian: sudo apt install libwebkit2gtk-4.1-dev"
    echo "  Fedora:        sudo dnf install webkit2gtk4.1-devel"
    echo "  Arch:          sudo pacman -S webkit2gtk"
    exit 1
fi

# Step 4: Check for optional tools
print_step "4/6" "Detecting Optional Build Tools..."

UPX_FLAGS=""
if command -v upx >/dev/null 2>&1; then
    UPX_FLAGS="-upx"
    UPX_VERSION=$(upx --version 2>&1 | head -n1)
    print_success "UPX detected - compression enabled"
    echo "     $UPX_VERSION"
else
    print_info "UPX not found - binaries will not be compressed"
    echo "     Install from: https://github.com/upx/upx/releases"
fi

# Step 5: Prepare build environment
print_step "5/6" "Preparing Build Environment..."

# Clean previous builds
if [ -d "build/bin/linux" ]; then
    echo "Cleaning previous build artifacts..."
    rm -rf "build/bin/linux"
fi

# Create output directory
mkdir -p build/bin/linux
print_success "Build directory prepared"

# Detect architecture
ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    PLATFORM="linux/amd64"
    ARCH_NAME="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
    PLATFORM="linux/arm64"
    ARCH_NAME="arm64"
else
    print_error "Unsupported architecture: $ARCH"
    echo "       Only x86_64 and ARM64 are supported."
    exit 1
fi

print_success "Building for: $PLATFORM"

# Step 6: Build the application
print_step "6/6" "Building Application..."
echo "This may take several minutes on first build..."
echo ""

BUILD_CMD="wails build -platform $PLATFORM $UPX_FLAGS -clean"
echo "Command: $BUILD_CMD"
echo ""

if $BUILD_CMD; then
    echo ""
    print_success "Build completed successfully!"
else
    echo ""
    print_error "Build failed. Check the output above for details."
    echo ""
    echo "Common issues:"
    echo "- Frontend dependencies not installed: cd frontend && pnpm install"
    echo "- Outdated Wails CLI: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
    echo "- Missing system dependencies"
    exit 1
fi

# Move and rename binary
if [ -f "build/bin/TridUI" ]; then
    mv "build/bin/TridUI" "build/bin/linux/TridUI-linux-$ARCH_NAME"
    chmod +x "build/bin/linux/TridUI-linux-$ARCH_NAME"
    print_success "Binary moved to: build/bin/linux/TridUI-linux-$ARCH_NAME"
fi

# Build summary
echo ""
print_header "Build Summary"

OUTPUT_COUNT=0

if [ -f "build/bin/linux/TridUI-linux-$ARCH_NAME" ]; then
    print_success "Executable: build/bin/linux/TridUI-linux-$ARCH_NAME"
    OUTPUT_COUNT=$((OUTPUT_COUNT + 1))
    
    # Show file size
    FILE_SIZE=$(du -h "build/bin/linux/TridUI-linux-$ARCH_NAME" | cut -f1)
    echo "     Size: $FILE_SIZE"
fi

echo ""
echo "Output directory: build/bin/linux/"
echo "Total files: $OUTPUT_COUNT"

if [ $OUTPUT_COUNT -eq 0 ]; then
    print_warning "No output files detected. Build may have failed silently."
    exit 1
fi

# Additional information
echo ""
print_header "Next Steps"
echo ""
echo "- Test the application: ./build/bin/linux/TridUI-linux-$ARCH_NAME"
echo "- Runtime dependencies required:"
echo "    Ubuntu/Debian: libgtk-3-0 libwebkit2gtk-4.1-37"
echo "    Fedora:        gtk3 webkit2gtk4.1"
echo "    Arch:          gtk3 webkit2gtk"
echo ""
echo "Build completed at: $(date)"
echo ""

exit 0
