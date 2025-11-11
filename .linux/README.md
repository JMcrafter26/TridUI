# Linux Distribution Files

This folder contains packaging files for distributing TridUI on Linux app stores.

## Files

### Flathub (Flatpak)
- `com.github.jmcrafter26.TridUI.yml` - Flatpak manifest
- `com.github.jmcrafter26.TridUI.desktop` - Desktop entry
- `com.github.jmcrafter26.TridUI.metainfo.xml` - AppStream metadata

### Snap Store
- `snapcraft.yaml` - Snap package manifest
- `snap/local/tridui.desktop` - Desktop file for snap

## Documentation

- **`PUBLISHING_GUIDE.md`** - Beginner-friendly guide (START HERE!)
- **`PUBLISHING_TECHNICAL.md`** - Technical reference

## Quick Start

If this is your first time publishing to Linux stores, read `PUBLISHING_GUIDE.md` first. It walks you through everything step-by-step.

## Testing Locally

### Flatpak
```bash
flatpak-builder --force-clean build-dir com.github.jmcrafter26.TridUI.yml
flatpak-builder --run build-dir com.github.jmcrafter26.TridUI.yml TridUI
```

### Snap
```bash
snapcraft
sudo snap install --dangerous tridui_*.snap
```

## Notes

- All paths in the manifest files assume they're in the `.linux/` folder
- Keep version numbers in sync with `wails.json` and `version.go`
- Update release notes in `metainfo.xml` for each new version
