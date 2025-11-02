# Build Directory

This directory contains platform-specific build configurations and scripts for TridUI.

## Quick Start

**Windows:**
```cmd
.\build\build-windows.bat
```

**macOS:**
```bash
chmod +x build/build-darwin.sh
./build/build-darwin.sh
```

**Linux:**
```bash
chmod +x build/build-linux.sh
./build/build-linux.sh
```

## Build Scripts

The build scripts are fully automated and include:
- ✅ Prerequisite validation (Go, Wails, Node.js, pnpm)
- ✅ Version checking (Go 1.22+ required for Wails v2.10.2+)
- ✅ System dependency verification
- ✅ Optional tool detection (UPX, NSIS, create-dmg)
- ✅ Clean build environment preparation
- ✅ Architecture detection
- ✅ Comprehensive error reporting
- ✅ Build artifact verification

### Prerequisites

**All Platforms:**
- **Go 1.22+** - [Download](https://golang.org/dl/)
- **Wails CLI** - Install with: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- **Node.js 20+** - [Download](https://nodejs.org/)
- **pnpm 10+** - Install with: `npm install -g pnpm`

**Windows:**
- WebView2 Runtime (automatically installed by Wails if missing)
- Optional: [UPX](https://github.com/upx/upx/releases) for compression
- Optional: [NSIS](https://nsis.sourceforge.io/Download) for installers

**macOS:**
- Xcode Command Line Tools: `xcode-select --install`
- Optional: [UPX](https://github.com/upx/upx/releases) (not recommended on Apple Silicon)
- Optional: [create-dmg](https://github.com/create-dmg/create-dmg): `brew install create-dmg`

**Linux:**
```bash
# Debian/Ubuntu
sudo apt install libgtk-3-dev libwebkit2gtk-4.1-dev

# Fedora
sudo dnf install gtk3-devel webkit2gtk4.1-devel

# Arch
sudo pacman -S gtk3 webkit2gtk

# Optional: UPX for compression
```

### Build Outputs

**Windows:** `build/bin/windows/`
- `TridUI-win-amd64.exe` - Portable executable
- `TridUI-win-amd64-installer.exe` - NSIS installer (if available)
- `TridUI-win-arm64.exe` - ARM64 executable
- `TridUI-win-arm64-installer.exe` - ARM64 installer (if available)

**macOS:** `build/bin/darwin/`
- `TridUI.app` - Application bundle
- `TridUI-macOS-{arch}.zip` - Zipped app bundle
- `TridUI-macOS-{arch}.dmg` - DMG installer (if create-dmg available)
  - `{arch}` can be: `amd64`, `arm64`, or `universal`

**Linux:** `build/bin/linux/`
- `TridUI-linux-amd64` - x86_64 executable
- `TridUI-linux-arm64` - ARM64 executable

### Advanced Usage

**Manual Wails Commands:**
```bash
# Cross-platform builds
wails build -platform windows/amd64
wails build -platform darwin/universal
wails build -platform linux/amd64

# With compression
wails build -upx

# With Windows installer
wails build -platform windows/amd64 -nsis

# Development mode (hot reload)
wails dev
```

**Frontend Development:**
```bash
cd frontend
pnpm install        # Install dependencies
pnpm dev           # Start dev server
pnpm build         # Build for production
```

### Troubleshooting

**Build fails with "go: no such tool":**
- Update Go to 1.22 or later
- Run: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

**Frontend build errors:**
- Run: `cd frontend && pnpm install`
- Ensure Node.js 20+ is installed

**Linux: WebKit2GTK errors:**
- Install development headers (see Linux prerequisites above)
- Create symlink if using WebKit2GTK 4.1: `sudo ln -sf /usr/lib/x86_64-linux-gnu/pkgconfig/webkit2gtk-4.1.pc /usr/lib/x86_64-linux-gnu/pkgconfig/webkit2gtk-4.0.pc`

**macOS: Xcode errors:**
- Install Command Line Tools: `xcode-select --install`
- Accept license: `sudo xcodebuild -license accept`

**Windows: NSIS installer not created:**
- Install NSIS from https://nsis.sourceforge.io/Download
- Ensure `makensis.exe` is in PATH

## Directory Structure

```
build/
├── bin/                    # Build outputs
│   ├── windows/           # Windows binaries
│   ├── darwin/            # macOS app bundles
│   └── linux/             # Linux executables
├── darwin/                 # macOS-specific config
│   ├── Info.plist         # App bundle info (production)
│   ├── Info.dev.plist     # App bundle info (development)
│   └── gon-sign.json      # Code signing config
├── windows/                # Windows-specific config
│   ├── icon.ico           # Application icon
│   ├── info.json          # App metadata
│   ├── wails.exe.manifest # Windows manifest
│   └── installer/         # NSIS installer files
├── build-windows.bat       # Windows build script
├── build-darwin.sh         # macOS build script
└── build-linux.sh          # Linux build script
```

### Platform-Specific Files

**macOS (`darwin/`):**
- `Info.plist` - Used for production builds (`wails build`)
- `Info.dev.plist` - Used for development builds (`wails dev`)
- Delete these files to regenerate defaults

**Windows (`windows/`):**
- `icon.ico` - Application icon (regenerated from `appicon.png` if missing)
- `info.json` - Metadata for Windows Explorer properties
- `wails.exe.manifest` - Windows application manifest
- `installer/` - NSIS installer templates and scripts

## Distribution Notes

**Windows:**
- End users need WebView2 Runtime (usually pre-installed on Windows 11)
- Use the NSIS installer for automatic WebView2 installation

**macOS:**
- For public distribution, consider code signing and notarization
- Universal binaries work on both Intel and Apple Silicon Macs
- DMG files provide the best user experience

**Linux:**
- Users need runtime dependencies (GTK3, WebKit2GTK)
- Consider packaging as AppImage or .deb (see GitHub Actions workflows)
- Different distributions may need different dependency packages