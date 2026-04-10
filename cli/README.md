# TridUI CLI

A fast, Go-native Command Line Interface for file type identification using TrID definitions. This tool provides a modern alternative to the original TrID utility with full **Unicode filename support**.

## Features

- **Unicode Filename Support**: Unlike the original binary, this version handles all characters in file paths correctly.
- **Fast Identification**: Powered by the same clean-room Go engine as the TridUI application.
- **Auto-Update**: Built-in support for downloading and updating TRD definitions directly from `mark0.net`.
- **Flexible Arguments**: Supports both traditional TrID syntax and standard CLI flags.
- **Cross-Platform**: Automatically searches for definitions in common system paths and TridUI application data folders.

## Installation

### Download

> [!IMPORTANT]
> The CLI will be included in the next official release. In the meantime, you can download the latest automated builds from the **[GitHub Actions page](https://github.com/JMcrafter26/TridUI/actions/workflows/build-cli.yml)**.

#### Latest Build Reference
- **Production Build (Optimized)**: [GitHub Run 24255339319](https://github.com/JMcrafter26/TridUI/actions/runs/24255339319) ([VirusTotal](https://www.virustotal.com/gui/file-analysis/ZTZkNGU3ZGQ1M2I2ZmY5MmFjNjVjZWM5MjkyZWViMzg6MTc3NTg0MjA4NA==))
- **Small Build (UPX Compressed)**: [GitHub Run 24255516178](https://github.com/JMcrafter26/TridUI/actions/runs/24255516178) ([VirusTotal](https://www.virustotal.com/gui/file-analysis/NDJmOTczYzQxNzhhNDdlYjEwZWRmY2MzOWQ2ZDEyZWE6MTc3NTg0MjI3NA==) | [Threat.rip](https://www.threat.rip/file/d0dfa0c020295d0ca589f534aac029783fbcd317279ea5be097f962e3676e79f))
  - *Note: UPX compression significantly reduces file size (e.g., 2MB vs 6MB), but often causes Antivirus (AV) software to flag the binary as a **false positive** due to its packing nature.*

### From Source

Requires [Go](https://go.dev/dl/) 1.23 or higher.

```powershell
# Clone the repository
git clone https://github.com/JMcrafter26/TridUI.git
cd TridUI

# Build the binary (optimized for size)
go build -ldflags="-s -w" -o trid.exe ./cli/main.go

# Optional: Further compress using UPX
upx -9 trid.exe
```

## Usage

### Analyze a File

Basic usage (standard flag):
```bash
trid -f my_document.pdf
```

TrID-compatible syntax:
```bash
trid image.png
```

### Definitions Management

By default, the tool searches in current directories and common TridUI data folders. To specify a custom definitions file:
```bash
trid document.docx -d:C:\Custom\triddefs.trd
# OR
trid -f document.docx -d C:\Custom\triddefs.trd
```

### Advanced Usage

**Read file list from stdin**:
```bash
# Windows
Get-ChildItem *.bin -Name | trid -file-list -

# Linux / WSL
find . -name "*.bin" | trid -file-list -
```

**Export Results**:
```bash
# JSON output
trid -f my_file.iso -j results.json

# CSV output
trid -f my_file.iso -o results.csv
# OR (long version)
trid -f my_file.iso -out results.csv
```

**Recursion & Extensions**:
```bash
# Recursively scan a folder
trid -r ./my_folder
# OR
trid --recursive ./my_folder

# Automatically add guessed extensions to filenames
trid -ae image_without_ext
# OR
trid --add-ext image_without_ext

# Change extensions to the best match
trid -ce file.wrong_ext
# OR
trid --change-ext file.wrong_ext
```

### Options Reference

| Short | Long | Description |
| :--- | :--- | :--- |
| `-ae` | `-add-ext` | Add guessed extension to filenames |
| `-ce` | `-change-ext` | Change filenames extensions |
| `-r` | `-recursive` | Recursively include files in subdirectories |
| `-d` | `-defs` | Use specified definitions package |
| `-f` | `-file` | File to scan (standard flag) |
| `-fl` | `-file-list` | File list path (use `-` for stdin) |
| `-n` | `-num` | Show the first RESNUM matches (default: 5) |
| `-ns` | `-no-strings` | Disable strings check |
| `-o` | `-out` | Create a CSV file with results |
| `-j` | `-json` | Create a JSON file with results |
| `-v` | `-verbose` | Verbose output |
| `-w` | `-wait` | Wait for key press at end |
| `-u` | `-update` | Update definitions package |
| `-h` | `-help` | Show help |

### Updating Definitions

Download or update the `triddefs.trd` file automatically:
```bash
trid -u
```

## Credits

- **Current Author**: [Cufiy](https://github.com/JMcrafter26)
- **Original Concept**: [Marco Pontello](https://mark0.net) (TrID creator)
- **Engine**: Native Go implementation from the TridUI project.

## License

This project is licensed under the same terms as the parent TridUI project.
