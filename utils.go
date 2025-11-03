package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	gort "runtime"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// OpenFileDialog opens a native file dialog and returns the selected file path
func (a *App) OpenFileDialog() (string, error) {
	filePath, err := wailsruntime.OpenFileDialog(a.ctx, wailsruntime.OpenDialogOptions{
		Title: "Select File to Analyze",
		Filters: []wailsruntime.FileFilter{
			{
				DisplayName: "All Files (*)",
				Pattern:     "*",
			},
		},
	})
	return filePath, err
}

// getAppDataDir returns the appropriate application data directory for the current OS
func getAppDataDir() (string, error) {
	var appDataDir string

	switch gort.GOOS {
	case "windows":
		appDataDir = os.Getenv("APPDATA")
		if appDataDir == "" {
			return "", fmt.Errorf("APPDATA environment variable not set")
		}
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		appDataDir = filepath.Join(homeDir, "Library", "Application Support")
	case "linux":
		appDataDir = os.Getenv("XDG_DATA_HOME")
		if appDataDir == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			appDataDir = filepath.Join(homeDir, ".local", "share")
		}
	default:
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		appDataDir = filepath.Join(homeDir, ".config")
	}

	appDataDir = filepath.Clean(appDataDir)
	appDataDir = filepath.Join(appDataDir, "TridUI")

	// if the directory doesn't exist, create it
	if _, err := os.Stat(appDataDir); os.IsNotExist(err) {
		err := os.MkdirAll(appDataDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return appDataDir, nil
}

// OpenAppDir opens the application data directory in the system's file explorer
func (a *App) OpenAppDir() error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	// based on OS, open the app data directory in the file explorer
	var cmd *exec.Cmd
	if gort.GOOS == "windows" {
		cmd = exec.Command("explorer", appDataDir)
	} else {
		cmd = exec.Command("open", appDataDir)
	}
	return cmd.Start()
}

// GetConfig returns the current application configuration (json string)
func (a *App) GetConfig() (string, error) {
	// config is in a file in getAppDataDir()/config.json
	appDataDir, err := getAppDataDir()
	if err != nil {
		return "", fmt.Errorf("failed to get app data dir: %w", err)
	}
	configPath := filepath.Join(appDataDir, "config.json")
	configJSON, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config.json: %w", err)
	}
	return string(configJSON), nil
}
