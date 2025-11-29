package main

import (
	"context"
	"encoding/json"
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
		// No filters specified - allows selecting any file including those without extensions
	})
	return filePath, err
}

func (a *App) GetOSName() string {
	return gort.GOOS
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
		wailsruntime.LogError(a.ctx, "Failed to get app data dir: "+err.Error())
		return "", fmt.Errorf("failed to get app data dir: %w", err)
	}
	configPath := filepath.Join(appDataDir, "config.json")
	configJSON, err := os.ReadFile(configPath)
	if err != nil {
		wailsruntime.LogErrorf(a.ctx, "Failed to read config.json from %s: %v", configPath, err)
		return "", fmt.Errorf("failed to read config.json: %w", err)
	}
	return string(configJSON), nil
}

// SaveConfig writes the provided JSON string to the config file (overwrites)
func SaveConfig(ctx context.Context, jsonData string) error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		wailsruntime.LogError(ctx, "Failed to get app data dir: "+err.Error())
		return fmt.Errorf("failed to get app data dir: %w", err)
	}
	configPath := filepath.Join(appDataDir, "config.json")
	// ensure directory exists
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		wailsruntime.LogErrorf(ctx, "Failed to create app data dir %s: %v", appDataDir, err)
		return err
	}
	if err := os.WriteFile(configPath, []byte(jsonData), 0644); err != nil {
		wailsruntime.LogErrorf(ctx, "Failed to write config.json to %s: %v", configPath, err)
		return fmt.Errorf("failed to write config.json: %w", err)
	}
	return nil
}

// UpdateConfigKey updates a single key inside the config JSON file and writes it back.
func UpdateConfigKey(ctx context.Context, key string, value interface{}) error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		wailsruntime.LogError(ctx, "Failed to get app data dir: "+err.Error())
		return fmt.Errorf("failed to get app data dir: %w", err)
	}
	configPath := filepath.Join(appDataDir, "config.json")

	// Read existing config; if missing, start with empty map
	var data map[string]interface{}
	raw, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			data = make(map[string]interface{})
		} else {
			wailsruntime.LogErrorf(ctx, "Failed to read config.json from %s: %v", configPath, err)
			return fmt.Errorf("failed to read config.json: %w", err)
		}
	} else {
		if err := json.Unmarshal(raw, &data); err != nil {
			// corrupted file: start fresh
			wailsruntime.LogErrorf(ctx, "Failed to parse config.json: %v", err)
			data = make(map[string]interface{})
		}
	}

	data[key] = value

	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := os.WriteFile(configPath, out, 0644); err != nil {
		wailsruntime.LogErrorf(ctx, "Failed to write config.json to %s: %v", configPath, err)
		return fmt.Errorf("failed to write config.json: %w", err)
	}
	return nil
}
