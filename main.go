package main

import (
	"context"
	"embed"
	"os"
	"strings"

	"TridUI/internal/clirunner"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func isCLIMode(args []string) bool {
	if len(args) == 0 {
		return false
	}
	// If given more than 1 argument, or the first argument starts with '-', assume CLI mode.
	// Otherwise, it's likely a drag-and-drop of a single file, which goes to the GUI.
	if len(args) > 1 || strings.HasPrefix(args[0], "-") {
		return true
	}
	// Check if the single argument isn't an existing file, defaulting to CLI just in case
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if isCLIMode(args) {
		attachConsole()
		os.Exit(clirunner.Run(args))
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "TridUI",
		Width:       500,
		Height:      400,
		Frameless:   true,
		AlwaysOnTop: false,
		MaxHeight:   700,
		MaxWidth:    900,
		MinHeight:   300,
		MinWidth:    400,
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			// WindowIsTranslucent: true,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
			// WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		runtime.LogFatal(context.Background(), "Error: "+err.Error())
	}
}
