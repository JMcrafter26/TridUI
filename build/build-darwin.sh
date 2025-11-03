#!/bin/bash
# TridUI macOS Build Script
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
print_header "TridUI macOS Build Script"

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

# Check for Xcode Command Line Tools
if ! xcode-select -p >/dev/null 2>&1; then
    print_error "Xcode Command Line Tools not found."
    echo ""
    echo "Install with: xcode-select --install"
    exit 1
fi
print_success "Xcode Command Line Tools detected"

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

# Step 3: Detect architecture and build type
print_step "3/6" "Detecting System Architecture..."

ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    PLATFORM="darwin/amd64"
    ARCH_NAME="amd64"
elif [ "$ARCH" = "arm64" ]; then
    PLATFORM="darwin/arm64"
    ARCH_NAME="arm64"
else
    print_error "Unsupported architecture: $ARCH"
    echo "       Only x86_64 and ARM64 are supported."
    exit 1
fi

print_success "Building for: $PLATFORM"

# Ask if user wants universal binary
echo ""
echo "Build options:"
echo "  1) Native binary ($ARCH_NAME only)"
echo "  2) Universal binary (Intel + Apple Silicon)"
echo ""
read -p "Select build type [1]: " BUILD_TYPE
BUILD_TYPE=${BUILD_TYPE:-1}

if [ "$BUILD_TYPE" = "2" ]; then
    PLATFORM="darwin/universal"
    ARCH_NAME="universal"
    print_info "Building universal binary for both Intel and Apple Silicon"
fi

# Step 4: Check for optional tools
print_step "4/6" "Detecting Optional Build Tools..."

UPX_FLAGS=""
if command -v upx >/dev/null 2>&1; then
    if [ "$ARCH" = "arm64" ]; then
        print_warning "UPX detected but known to have issues on Apple Silicon"
        print_info "Skipping UPX compression for safety"
    else
        UPX_FLAGS="-upx"
        UPX_VERSION=$(upx --version 2>&1 | head -n1)
        print_success "UPX detected - compression enabled"
        echo "     $UPX_VERSION"
    fi
else
    print_info "UPX not found - binaries will not be compressed"
    echo "     Install with: brew install upx"
fi

# Check for create-dmg
if command -v create-dmg >/dev/null 2>&1; then
    print_success "create-dmg detected - DMG creation available"
    HAS_DMG_TOOL=true
else
    print_info "create-dmg not found - only .app bundle will be created"
    echo "     Install with: brew install create-dmg"
    HAS_DMG_TOOL=false
fi

# Step 5: Prepare build environment
print_step "5/6" "Preparing Build Environment..."

# Clean previous builds
if [ -d "build/bin/darwin" ]; then
    echo "Cleaning previous build artifacts..."
    rm -rf "build/bin/darwin"
fi

# Create output directory
mkdir -p build/bin/darwin
print_success "Build directory prepared"

# Step 6: Build the application
print_step "6/6" "Building Application..."
echo "This may take several minutes on first build..."
echo ""

BUILD_CMD="wails build -platform $PLATFORM $UPX_FLAGS -clean -darwin-entitlements build/darwin/TridUI.entitlements"
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
    echo "- Missing Xcode Command Line Tools"
    exit 1
fi

# Move app bundle
if [ -d "build/bin/TridUI.app" ]; then
    cp -R "build/bin/TridUI.app" "build/bin/darwin/"
    print_success "App bundle moved to: build/bin/darwin/TridUI.app"
    
    # Create ZIP using ditto (preserves macOS permissions and attributes)
    echo "Creating ZIP archive..."
    cd build/bin/darwin
    ditto -c -k --keepParent TridUI.app "TridUI-macOS-$ARCH_NAME.zip"
    cd ../../..
    print_success "ZIP created: build/bin/darwin/TridUI-macOS-$ARCH_NAME.zip"
    
    # Create DMG if tool is available
    if [ "$HAS_DMG_TOOL" = true ]; then
        echo "Creating DMG..."
        
        ICON_PATH="build/appicon.png"
        if [ ! -f "$ICON_PATH" ]; then
            ICON_PATH="icon.png"
        fi
        
        create-dmg \
            --volname "TrID UI" \
            --window-pos 200 120 \
            --window-size 600 400 \
            --icon-size 100 \
            --icon "TridUI.app" 175 190 \
            --hide-extension "TridUI.app" \
            --app-drop-link 425 190 \
            "build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg" \
            "build/bin/darwin/TridUI.app" 2>/dev/null || print_warning "DMG creation failed, but ZIP is available"
        
        if [ -f "build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg" ]; then
            print_success "DMG created: build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg"
        fi
    fi
fi

# Build summary
echo ""
print_header "Build Summary"

OUTPUT_COUNT=0

if [ -d "build/bin/darwin/TridUI.app" ]; then
    print_success "App bundle: build/bin/darwin/TridUI.app"
    OUTPUT_COUNT=$((OUTPUT_COUNT + 1))
fi

if [ -f "build/bin/darwin/TridUI-macOS-$ARCH_NAME.zip" ]; then
    print_success "ZIP archive: build/bin/darwin/TridUI-macOS-$ARCH_NAME.zip"
    FILE_SIZE=$(du -h "build/bin/darwin/TridUI-macOS-$ARCH_NAME.zip" | cut -f1)
    echo "     Size: $FILE_SIZE"
    OUTPUT_COUNT=$((OUTPUT_COUNT + 1))
fi

if [ -f "build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg" ]; then
    print_success "DMG image: build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg"
    FILE_SIZE=$(du -h "build/bin/darwin/TridUI-macOS-$ARCH_NAME.dmg" | cut -f1)
    echo "     Size: $FILE_SIZE"
    OUTPUT_COUNT=$((OUTPUT_COUNT + 1))
fi

echo ""
echo "Output directory: build/bin/darwin/"
echo "Total artifacts: $OUTPUT_COUNT"

if [ $OUTPUT_COUNT -eq 0 ]; then
    print_warning "No output files detected. Build may have failed silently."
    exit 1
fi

# Additional information
echo ""
print_header "Next Steps"
echo ""
echo "- Test the application: open build/bin/darwin/TridUI.app"
echo "- Distribute the DMG or ZIP file"
echo ""
if [ "$ARCH_NAME" != "universal" ]; then
    print_info "Note: This build only runs on $ARCH_NAME Macs"
    print_info "Run the script again and select option 2 for a universal binary"
fi
echo ""
echo "For distribution, consider:"
echo "- Code signing: https://developer.apple.com/support/code-signing/"
echo "- Notarization: https://developer.apple.com/documentation/security/notarizing_macos_software_before_distribution"
echo ""
echo "Build completed at: $(date)"
echo ""

exit 0
