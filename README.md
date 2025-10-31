
<div style="display: flex; align-items: center; margin-bottom: 16px;">
  <img src="./icon.png" alt="TrID UI Icon" style="width: 64px; height: 64px; border-radius: 12px;" />
  <h1 style="margin-left: 16px;">TrID UI</h1>
</div>

<div style="text-align: center; margin-bottom: 16px;">
<img src="https://github.com/JMcrafter26/TridUI/blob/main/.github/screenshots/demonstration.gif?raw=true" alt="TrID UI Demonstration" style="width: 100%; border: 1px solid #ccc; border-radius: 8px; margin-bottom: 16px;" />
</div>

TrID UI is a lightweight desktop application that provides a user-friendly interface for TrID, a powerful tool for scanning and analyzing files. With TrID UI, users can easily select or drop files on the Home screen to initiate local scans, making it convenient to detect unrecognized filetypes.

The application uses a native Go implementation of the TrID file identification algorithm, providing fast and accurate file type detection without external dependencies.

> [!WARNING]  
> :warning: TrID UI is currently in early development. Features may be incomplete or unstable. Use at your own risk.

> [!TIP]
> Download TridUI from the [Releases Page](https://github.com/JMcrafter26/TridUI/releases)

## Features

- 🚀 Fast native Go-based file scanning
- 🎯 Accurate file type identification using TrID definitions
- 💻 Cross-platform desktop application (Windows, macOS, Linux)
- 🔒 100% local processing - no data leaves your computer
- 🎨 Modern, intuitive user interface
- 📊 Detailed match results with confidence scores
- 🔄 Drag-and-drop file support
- 🔁 Automatic definitions updates with one click
- 📅 Track last update date and definition count

## Table of Contents

<details>
<summary>Click to expand</summary>

- [Features](#features)
- [Table of Contents](#table-of-contents)
- [Demonstration and Screenshots](#demonstration-and-screenshots)
  - [Demonstration Videos](#demonstration-videos)
  - [Screenshots](#screenshots)
- [Setup](#setup)
  - [Prerequisites](#prerequisites)
    - [Option 1: Automatic Download (Recommended)](#option-1-automatic-download-recommended)
    - [Option 2: Manual Installation](#option-2-manual-installation)
  - [Building from Source](#building-from-source)
- [Usage](#usage)
- [Technical Details](#technical-details)
  - [Architecture](#architecture)
  - [TrID Scanner Implementation](#trid-scanner-implementation)
- [License and Attribution](#license-and-attribution)
- [Contributing](#contributing)
  - [Translations](#translations)

</details>

## Demonstration and Screenshots

### Demonstration Videos

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <video controls style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" loop>
    <source src="https://github.com/JMcrafter26/TridUI/raw/refs/heads/main/.github/videos/demo1.mp4" type="video/mp4">
    Your browser does not support the video tag.
  </video>
  <video controls style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" loop>
    <source src="https://github.com/JMcrafter26/TridUI/raw/refs/heads/main/.github/videos/demo2.mp4" type="video/mp4">
    Your browser does not support the video tag.
  </video>
  <video controls style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" loop>
    <source src="https://github.com/JMcrafter26/TridUI/raw/refs/heads/main/.github/videos/demo3.mp4" type="video/mp4">
    Your browser does not support the video tag.
  </video>
</div>

### Screenshots

<div style="overflow-x: scroll; display: flex; gap: 16px; padding-bottom: 16px; max-height: 400px; width: 100%;">
  <img src="./.github/screenshots/1.png" alt="TrID UI Screenshot 1" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="./.github/screenshots/2.png" alt="TrID UI Screenshot 2" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="./.github/screenshots/3.png" alt="TrID UI Screenshot 3" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="./.github/screenshots/4.png" alt="TrID UI Screenshot 4" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
  <img src="./.github/screenshots/5.png" alt="TrID UI Screenshot 5" style="width: 300px; border: 1px solid #ccc; border-radius: 8px;" />
</div>

## Setup

> [!TIP]
> You can find the pre-built binaries on the [Releases Page](https://github.com/JMcrafter26/TridUI/releases).

### Prerequisites

The application can automatically download and update the TrID definitions file for you!

#### Option 1: Automatic Download (Recommended)

1. Launch TrID UI
2. Go to Settings
3. Click "Download Definitions" or "Check for Updates"
4. The app will automatically download and install the latest definitions

#### Option 2: Manual Installation

1. Download the TrID definitions file (`triddefs.trd`) from [Mark0.net](https://mark0.net/soft-trid-deflist.html)
2. Place the `triddefs.trd` file in your application data directory:
   - **Windows**: `%APPDATA%\TridUI\triddefs.trd`
   - **macOS**: `~/Library/Application Support/TridUI/triddefs.trd`
   - **Linux**: `~/.local/share/TridUI/triddefs.trd`

You can use the "Open App Dir" button in Settings to navigate to the correct location.

### Building from Source

```bash
# Install dependencies
go mod download
cd frontend && pnpm install
cd ..

# Build the application
wails build
```

## Usage

1. Launch TrID UI
2. Click or drag-and-drop a file onto the interface
3. View the scan results with confidence scores
4. The best match is highlighted at the top
5. Additional possible matches are listed below

## Technical Details

### Architecture

- **Backend**: Go (Wails framework)
- **Frontend**: SvelteKit + TypeScript + DaisyUI (& Tailwind CSS)
- **TrID Engine**: Pure Go implementation (`/trid` package)

### TrID Scanner Implementation

The TrID scanner ([`/trid/trid.go`](https://github.com/JMcrafter26/TridUI/blob/main/trid/trid.go)) is a clean-room Go implementation that:

- Parses TRD (TrID Definition) files using the binary format specification
- Performs pattern matching at specified file offsets
- Supports string matching for enhanced accuracy
- Calculates confidence scores based on pattern weights
- Returns ranked results with detailed file type information

> You can find the TRD format specification on [Mark0.net](https://mark0.net/soft-trid-format.html).

## License and Attribution

TrID UI is open-source software licensed under the GNU AGPLv3 license. The UI is developed by Cufiy (aka JMcrafter26 - me) and is based on TrID by [Marco Pontello](https://mark0.net/).
Please refer to the LICENSE file for more details.

The trid.go scanner is a clean-room Go implementation by JMcrafter26 and is licensed under the GNU AGPLv3 license.

The app icon is based on eye icon by icons8.com

## Contributing

Contributions are welcome! If you'd like to contribute to TrID UI, please fork the repository and submit a pull request with your changes. For major changes, please open an issue first to discuss what you would like to change.

### Translations

TrID UI needs your help to reach a wider audience! The current translations are machine-generated and may contain inaccuracies.

If you'd like to contribute translations, please follow these steps:

1. Fork the repository
2. Create a new branch for your translation
3. Add your translation files to the `translations` directory
4. Submit a pull request with your changes

Thank you for helping us make TrID UI accessible to more users!
