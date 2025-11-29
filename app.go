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
