package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func (a *App) ProcessFile(filePath string) error {
	// TODO: Implement actual Tridscan processing
	// For now, just log the file path to confirm the integration works
	fmt.Printf("ProcessFile called with: %s\n", filePath)

	// Example of how you would call your Python script when ready:
	// cmd := exec.Command("python", "path/to/tridscan.py", filePath)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	//     return fmt.Errorf("failed to process file: %w", err) n
	// }
	// fmt.Printf("Tridscan output: %s\n", output)

	return nil
}

func (a *App) CheckTrid() error {
	// check if tridscan.exe is in getAppDataDir()
	appDataDir, err := getAppDataDir()
	if err != nil {
		return fmt.Errorf("failed to get app data dir: %w", err)
	}

	tridPath := ""

	// check if tridscan is in appDataDir (based on OS)
	// first check if tridscan.py exists
	if _, err := os.Stat(filepath.Join(appDataDir, "tridscan.py")); os.IsExist(err) {
		tridPath = filepath.Join(appDataDir, "tridscan.py")
	}
	switch runtime.GOOS {
	case "windows":
		tridPath = filepath.Join(appDataDir, "tridscan.exe")

	case "darwin":
		tridPath = filepath.Join(appDataDir, "tridscan")

	case "linux":
		tridPath = filepath.Join(appDataDir, "tridscan")
	}

	if _, err := os.Stat(tridPath); os.IsNotExist(err) {
		fmt.Println("tridscan not found in app data dir. Prompting to select")
	}
	return nil
}
