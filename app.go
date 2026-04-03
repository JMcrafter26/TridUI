package main

import (
	"context"
	"os"
	"path/filepath"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// If the application was launched with a file path (e.g. user dropped a file
	// onto the executable/shortcut), emit an event so the frontend can handle it.
	// os.Args[0] is the executable; any additional args are file paths.
	args := os.Args[1:]
	if len(args) > 0 {
		// Use the first argument as the file to open
		filePath := args[0]
		// Verify the file exists before emitting
		if _, err := os.Stat(filePath); err == nil {
			go func(p string) {
				wailsruntime.LogInfof(a.ctx, "Startup file provided: %s", p)
				wailsruntime.EventsEmit(a.ctx, "app:openFile", p)
			}(filePath)
		} else {
			wailsruntime.LogInfof(a.ctx, "Startup arg not a valid file: %s", filePath)
		}
	}
}


// ConfigExists returns true if config.json exists in the app data directory
func (a *App) ConfigExists() (bool, error) {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return false, err
	}
	configPath := filepath.Join(appDataDir, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// SaveConfig saves the full settings JSON to config.json (overwrites). This is callable from frontend.
func (a *App) SaveConfig(jsonData string) error {
	// run write synchronously here but can be called from frontend async side
	return SaveConfig(a.ctx, jsonData)
}

// SaveSetting updates a single key in the config file asynchronously to avoid blocking the UI.
func (a *App) SaveSetting(key string, value interface{}) {
	go func() {
		if err := UpdateConfigKey(a.ctx, key, value); err != nil {
			wailsruntime.LogErrorf(a.ctx, "SaveSetting failed for %s: %v", key, err)
		}
	}()
}
