package main

import (
	"TridUI/trid"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Original TrID style arguments: trid <file> [-d:defs]
	// We'll also support more standard flags: -f <file> -d <defs>

	var defsPath string
	var filePath string
	var versionNumber = "1.0.0"

	// Check for standard flags first
	flag.StringVar(&defsPath, "d", "", "Path to TRD definitions file")
	flag.StringVar(&filePath, "f", "", "Path to file to analyze")
	help := flag.Bool("h", false, "Show help")
	version := flag.Bool("v", false, "Show version")
	flag.Parse()

	if *version {
		printVersion()
		return
	}

	if *help {
		printUsage()
		return
	}

	// Positional arguments handling (TrID style)
	args := flag.Args()
	if filePath == "" && len(args) > 0 {
		filePath = args[0]
	}

	// Check for -d:defs style in positional args if not set via flag
	if defsPath == "" {
		for _, arg := range args {
			if strings.HasPrefix(arg, "-d:") {
				defsPath = strings.TrimPrefix(arg, "-d:")
				break
			}
			if arg == "-v" || arg == "--version" {
				printVersion()
				return
			}
		}
	}

	if filePath == "" {
		printUsage()
		os.Exit(1)
	}

	analyzer := trid.NewAnalyzer()

	// Load definitions
	if defsPath != "" {
		err := analyzer.LoadDefinitions(defsPath)
		if err != nil {
			fmt.Printf("Error loading definitions from %s: %v\n", defsPath, err)
			os.Exit(1)
		}
	} else {
		// Try default locations
		err := loadDefaultDefinitions(analyzer)
		if err != nil {
			fmt.Printf("Error: No TRD definitions found. Use -d to specify path.\n")
			os.Exit(1)
		}
	}

	fmt.Printf("TridUI CLI v%s\n", versionNumber)
	fmt.Printf("Definitions loaded: %d\n", analyzer.GetDefinitionCount())
	fmt.Printf("Analyzing: %s\n\n", filePath)

	results, err := analyzer.AnalyzeFile(filePath)
	if err != nil {
		fmt.Printf("Error analyzing file: %v\n", err)
		os.Exit(1)
	}

	if len(results) == 0 {
		fmt.Println("Unknown file type.")
		return
	}

	// Output results in TrID style
	for _, res := range results {
		if res.Confidence < 0.1 {
			continue
		}
		fmt.Printf("%5.1f%% (.%s) %s\n", res.Confidence, res.Definition.FileType.Extension, res.Definition.FileType.Name)
	}
}

func loadDefaultDefinitions(a *trid.Analyzer) error {
	// Common locations for triddefs.trd
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)

	possiblePaths := []string{
		"triddefs.trd",
		filepath.Join(exeDir, "triddefs.trd"),
		// Standard TrID locations
		"/usr/share/trid/triddefs.trd",
		"/usr/local/share/trid/triddefs.trd",
	}

	// OS Specific TridUI paths
	home, _ := os.UserHomeDir()
	if home != "" {
		// Linux/macOS standard
		possiblePaths = append(possiblePaths, filepath.Join(home, ".trid", "triddefs.trd"))
		// Linux/macOS TridUI specific config
		possiblePaths = append(possiblePaths, filepath.Join(home, ".config", "TridUI", "triddefs.trd"))
		// macOS Application Support
		possiblePaths = append(possiblePaths, filepath.Join(home, "Library", "Application Support", "TridUI", "triddefs.trd"))
	}

	// Windows AppData (Roaming)
	appData := os.Getenv("APPDATA")
	if appData != "" {
		possiblePaths = append(possiblePaths, filepath.Join(appData, "TridUI", "triddefs.trd"))
	}

	// Windows LocalAppData
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData != "" {
		possiblePaths = append(possiblePaths, filepath.Join(localAppData, "TridUI", "triddefs.trd"))
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return a.LoadDefinitions(path)
		}
	}

	return fmt.Errorf("no definitions found in searched locations: %s", strings.Join(possiblePaths, ", "))
}

func printUsage() {
	printVersion()
	fmt.Println("\nUsage:")
	fmt.Println("  trid <file> [-d:defs]")
	fmt.Println("  trid -f <file> [-d <defs>]")
	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Println("TrID UI CLI v" + versionNumber)
	fmt.Println("GitHub: https://github.com/JMcrafter26/TridUI")
	fmt.Println("Author: Cufiy")
	fmt.Println("\nBased on TrID concept by Marco Pontello")
	fmt.Println("Original TrID website: https://mark0.net")
}
