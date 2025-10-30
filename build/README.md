# Build Directory

The build directory is used to house all the build files and assets for your application. 

The structure is:

* bin - Output directory
* darwin - macOS specific files
* windows - Windows specific files

## Build Scripts

This directory contains scripts for building TridscanUI on different platforms.

### Windows Build

#### Prerequisites

Before building on Windows, ensure you have the following installed:

1. Go (1.18 or newer)
2. Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
3. Deno (for sandbox host)
4. Node.js and npm (for frontend)
5. UPX (optional, for binary compression)
6. NSIS (optional, for installer creation)

#### Building on Windows

To build TridscanUI for Windows:

1. Open a PowerShell or Command Prompt
2. Navigate to the project directory
3. Run the build script:
   ```powershell
   .\build\build-windows.bat
   ```

4. The built application will be available in `build\bin\windows\`

#### Build Output on Windows

- `TridscanUI-amd64.exe` - The standalone application executable
- `TridscanUI-amd64-installer.exe` - The installer (if NSIS is installed)

### macOS Build

#### Prerequisites

Before building on macOS, ensure you have the following installed:

1. Go (1.18 or newer)
2. Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
3. Deno (for sandbox host)
4. Node.js and npm (for frontend)
5. UPX (optional, for binary compression) - Note: There are known issues with UPX on Apple Silicon

#### Building on macOS

To build TridscanUI for macOS:

1. Make the build script executable (only needed once):
   ```bash
   chmod +x build/build-darwin.sh
   ```

2. Run the build script:
   ```bash
   ./build/build-darwin.sh
   ```

3. The built application will be available in `build/bin/darwin/TridscanUI.app`

#### Advanced macOS Build Options

- For a universal binary (works on both Intel and Apple Silicon), uncomment the relevant section in the build-darwin.sh script
- If you want to customize the minimum macOS version, you can use:
  ```bash
  CGO_CFLAGS=-mmacosx-version-min=10.15.0 CGO_LDFLAGS=-mmacosx-version-min=10.15.0 ./build/build-darwin.sh
  ```

### Linux Build

#### Prerequisites

Before building on Linux, ensure you have the following installed:

1. Go (1.18 or newer)
2. Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)
3. Deno (for sandbox host)
4. Node.js and npm (for frontend)
5. UPX (optional, for binary compression)
6. Required system dependencies:
   - For Ubuntu/Debian:
     ```bash
     sudo apt install build-essential libgtk-3-dev libwebkit2gtk-4.0-dev
     ```
   - For Fedora:
     ```bash
     sudo dnf install gcc-c++ gtk3-devel webkit2gtk3-devel
     ```
   - For Arch:
     ```bash
     sudo pacman -S gcc gtk3 webkit2gtk
     ```

#### Building on Linux

To build TridscanUI for Linux:

1. Make the build script executable (only needed once):
   ```bash
   chmod +x build/build-linux.sh
   ```

2. Run the build script:
   ```bash
   ./build/build-linux.sh
   ```

3. The built application will be available in `build/bin/linux/`

#### Build Output on Linux

- `TridscanUI-amd64` - The executable for x86_64 systems
- `TridscanUI-arm64` - The executable for ARM64 systems (if built on such a system)

#### Running on Linux

For the application to run properly, users may need to install runtime dependencies:

```bash
# Ubuntu/Debian
sudo apt install libgtk-3-0 libwebkit2gtk-4.0-37 gstreamer1.0-plugins-good

# Fedora
sudo dnf install gtk3 webkit2gtk4.0 gstreamer1-plugins-good

# Arch
sudo pacman -S gtk3 webkit2gtk gst-plugins-good
```

## Cross-Platform Development Notes

### Preparing Scripts for Different Platforms

If you're developing on one platform and preparing scripts for another:

#### Windows to Unix (macOS/Linux) Transfer

1. Make sure to convert the line endings of shell scripts from CRLF to LF before transferring to macOS or Linux
2. You can use tools like VS Code (set to LF line endings) or `dos2unix` for this conversion
3. Shell scripts must have execute permissions set after transfer: `chmod +x script.sh`

#### Unix to Windows Transfer

1. Windows batch files should use CRLF line endings
2. If transferring files from macOS or Linux, ensure your Git configuration is set properly to handle line endings

### System-Specific Considerations

#### Windows
- Make sure WebView2 runtime is installed (Wails will help detect this)
- NSIS is required only for creating installers

#### macOS
- For distribution, consider code signing and notarization
- Universal binaries (targeting both Intel and Apple Silicon) require additional build settings

#### Linux
- System dependencies (GTK3, WebKit) must be installed before building
- Additional media playback libraries may be needed for audio/video elements
- Different distributions may have different package names for dependencies

## Directory Structure Details

### Mac

The `darwin` directory holds files specific to Mac builds.
These may be customised and used as part of the build. To return these files to the default state, simply delete them
and build with `wails build`.

The directory contains the following files:

- `Info.plist` - the main plist file used for Mac builds. It is used when building using `wails build`.
- `Info.dev.plist` - same as the main plist file but used when building using `wails dev`.

### Windows

The `windows` directory contains the manifest and rc files used when building with `wails build`.
These may be customised for your application. To return these files to the default state, simply delete them and
build with `wails build`.

- `icon.ico` - The icon used for the application. This is used when building using `wails build`. If you wish to
  use a different icon, simply replace this file with your own. If it is missing, a new `icon.ico` file
  will be created using the `appicon.png` file in the build directory.
- `installer/*` - The files used to create the Windows installer. These are used when building using `wails build`.
- `info.json` - Application details used for Windows builds. The data here will be used by the Windows installer,
  as well as the application itself (right click the exe -> properties -> details)
- `wails.exe.manifest` - The main application manifest file.