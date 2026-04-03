# TridUI CLI

A fast, Go-native Command Line Interface for file type identification using TrID definitions. This tool provides a modern alternative to the original TrID utility with full **Unicode filename support**.

## Features

- **Unicode Filename Support**: Unlike the original binary, this version handles all characters in file paths correctly.
- **Fast Identification**: Powered by the same clean-room Go engine as the TridUI application.
- **Auto-Update**: Built-in support for downloading and updating TRD definitions directly from `mark0.net`.
- **Flexible Arguments**: Supports both traditional TrID syntax and standard CLI flags.
- **Cross-Platform**: Automatically searches for definitions in common system paths and TridUI application data folders.

## Installation

### From Source

Requires [Go](https://go.dev/dl/) 1.23 or higher.

```bash
# Clone the repository
git clone https://github.com/JMcrafter26/TridUI.git
cd TridUI

# Build the binary
go build -o trid.exe ./cli/main.go
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
