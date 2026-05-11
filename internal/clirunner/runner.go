package clirunner

import (
	"TridUI/trid"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const versionNumber = "1.1.0"

type ScanResult struct {
	File  string
	Score int
	Perc  string
	Ext   string
	Type  string
	Mime  string
}

// Run executes the CLI logic. It takes the arguments slice (usually os.Args[1:])
// and returns an exit code.
func Run(args []string) int {
	var defsPath string
	var filePath string
	var fileListPath string
	var outFilePath string
	var jsonFilePath string
	var resNum int

	// Create a new FlagSet to allow passing specific args
	fs := flag.NewFlagSet("TridUI", flag.ContinueOnError)

	// Define all flags inspired by original TrID
	fs.StringVar(&defsPath, "d", "", "Use specified definitions package")
	fs.StringVar(&defsPath, "defs", "", "Use specified definitions package (long)")
	fs.StringVar(&filePath, "f", "", "File to scan (standard flag)")
	fs.StringVar(&filePath, "file", "", "File to scan (long)")
	fs.StringVar(&fileListPath, "file-list", "", "Specify a text file containing a list of files to scan (use - to read from stdin)")
	fs.StringVar(&fileListPath, "fl", "", "Specify a text file containing a list of files to scan (use - to read from stdin) (short)")
	fs.StringVar(&outFilePath, "o", "", "Create a CSV file with the results")
	fs.StringVar(&outFilePath, "out", "", "Create a CSV file with the results (long)")
	fs.StringVar(&jsonFilePath, "j", "", "Create a JSON file with the results")
	fs.StringVar(&jsonFilePath, "json", "", "Create a JSON file with the results (long)")
	fs.IntVar(&resNum, "n", 5, "Show the first RESNUM matches")
	fs.IntVar(&resNum, "num", 5, "Show the first RESNUM matches (long)")

	addExt := fs.Bool("ae", false, "Add guessed extension to filenames")
	addExtLong := fs.Bool("add-ext", false, "Add guessed extension to filenames (long)")
	changeExt := fs.Bool("ce", false, "Change filenames extensions")
	changeExtLong := fs.Bool("change-ext", false, "Change filenames extensions (long)")
	recursive := fs.Bool("r", false, "Recursively include files in subdirectories")
	recursiveLong := fs.Bool("recursive", false, "Recursively include files in subdirectories (long)")
	noStrings := fs.Bool("ns", false, "Disable strings check")
	noStringsLong := fs.Bool("no-strings", false, "Disable strings check (long)")
	verbose := fs.Bool("v", false, "Verbose output")
	verboseLong := fs.Bool("verbose", false, "Verbose output (long)")
	wait := fs.Bool("w", false, "Wait for a key press at the end")
	waitLong := fs.Bool("wait", false, "Wait for a key press at the end (long)")
	update := fs.Bool("u", false, "Update definitions package")
	updateLong := fs.Bool("update", false, "Update definitions package (long)")
	help := fs.Bool("h", false, "Show help")
	helpLong := fs.Bool("help", false, "Show help (long)")

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			printUsage(fs)
			return 0
		}
		fmt.Printf("Error string: %v\n", err)
		return 1
	}

	*addExt = *addExt || *addExtLong
	*changeExt = *changeExt || *changeExtLong
	*recursive = *recursive || *recursiveLong
	*noStrings = *noStrings || *noStringsLong
	*verbose = *verbose || *verboseLong
	*wait = *wait || *waitLong
	*update = *update || *updateLong
	*help = *help || *helpLong

	if *help {
		printUsage(fs)
		return 0
	}

	if *update {
		runUpdate(defsPath)
		return 0
	}

	_ = noStrings

	cmdArgs := fs.Args()
	var targets []string

	if filePath != "" {
		targets = append(targets, filePath)
	}

	for _, arg := range cmdArgs {
		if strings.HasPrefix(arg, "-d:") {
			if defsPath == "" {
				defsPath = strings.TrimPrefix(arg, "-d:")
			}
		} else {
			targets = append(targets, arg)
		}
	}

	if fileListPath != "" {
		list, err := readFilesFromList(fileListPath)
		if err != nil {
			fmt.Printf("Error reading file list: %v\n", err)
			return 1
		}
		targets = append(targets, list...)
	}

	if len(targets) == 0 {
		printUsage(fs)
		return 1
	}

	files := expandFiles(targets, *recursive)
	if len(files) == 0 {
		fmt.Println("Error: no file(s) to analyze!")
		return 1
	}

	analyzer := trid.NewAnalyzer()

	if defsPath != "" {
		err := analyzer.LoadDefinitions(defsPath)
		if err != nil {
			fmt.Printf("Error loading definitions from %s: %v\n", defsPath, err)
			return 1
		}
	} else {
		err := loadDefaultDefinitions(analyzer)
		if err != nil {
			fmt.Printf("Error: No TRD definitions found.\nRun 'trid -u' to download them automatically.\n")
			return 1
		}
	}

	if *verbose {
		fmt.Printf("TrID - File Identifier v%s - (C) 2003-2025 By M.Pontello\n", "2.43")
		fmt.Printf("TridUI CLI v%s - (C) 2025 By Cufiy\n\n", versionNumber)
		fmt.Printf("Definitions loaded: %d\n", analyzer.GetDefinitionCount())
	}

	var allResults []ScanResult

	fmt.Println("Analyzing...")
	for _, file := range files {
		fmt.Printf("\nFile: %s\n", file)

		results, err := analyzer.AnalyzeFile(file)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		if len(results) == 0 {
			fmt.Println("      Unknown!")
			allResults = append(allResults, ScanResult{File: file})
			continue
		}

		limit := resNum
		if len(results) < limit {
			limit = len(results)
		}

		for i := 0; i < limit; i++ {
			res := results[i]
			if res.Confidence < 0.1 && i > 0 {
				break
			}

			fmt.Printf("%5.1f%% (.%s) %s", res.Confidence, res.Definition.FileType.Extension, res.Definition.FileType.Name)
			if *verbose {
				fmt.Printf(" (%d/%d/%d)", res.Score, len(res.Definition.Patterns), len(res.Definition.Strings))
			}
			fmt.Println()

			if *verbose {
				if res.Definition.FileType.MimeType != "" {
					fmt.Printf("         Mime type  : %s\n", res.Definition.FileType.MimeType)
				}
				if res.Definition.FileType.URL != "" {
					fmt.Printf("         Related URL: %s\n", res.Definition.FileType.URL)
				}
				if res.Definition.FileType.Remarks != "" {
					fmt.Printf("         Remarks    : %s\n", res.Definition.FileType.Remarks)
				}
				if res.Definition.FileType.Author != "" {
					fmt.Printf("         Author     : %s\n", res.Definition.FileType.Author)
				}
			}

			if i == 0 {
				allResults = append(allResults, ScanResult{
					File:  file,
					Score: res.Score,
					Perc:  fmt.Sprintf("%.1f%%", res.Confidence),
					Ext:   res.Definition.FileType.Extension,
					Type:  res.Definition.FileType.Name,
					Mime:  res.Definition.FileType.MimeType,
				})
			}
		}

		if (*addExt || *changeExt) && len(results) > 0 {
			handleExtension(file, results[0], *addExt, *changeExt)
		}
	}

	if outFilePath != "" {
		saveCSV(outFilePath, allResults)
	}

	if jsonFilePath != "" {
		saveJSON(jsonFilePath, allResults)
	}

	if *wait {
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}

	return 0
}

func readFilesFromList(path string) ([]string, error) {
	var scanner *bufio.Scanner
	if path == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		scanner = bufio.NewScanner(f)
	}

	var result []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			result = append(result, line)
		}
	}
	return result, scanner.Err()
}

func expandFiles(patterns []string, recursive bool) []string {
	var result []string
	for _, pattern := range patterns {
		matches, _ := filepath.Glob(pattern)
		if len(matches) == 0 {
			info, err := os.Stat(pattern)
			if err == nil {
				if info.IsDir() {
					result = append(result, walkDir(pattern, recursive)...)
				} else {
					result = append(result, pattern)
				}
			}
			continue
		}
		for _, m := range matches {
			info, _ := os.Stat(m)
			if info != nil && info.IsDir() {
				result = append(result, walkDir(m, recursive)...)
			} else {
				result = append(result, m)
			}
		}
	}
	unique := make(map[string]bool)
	var final []string
	for _, f := range result {
		if !unique[f] {
			unique[f] = true
			final = append(final, f)
		}
	}
	sort.Strings(final)
	return final
}

func walkDir(root string, recursive bool) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || path == root {
			return nil
		}
		if info.IsDir() {
			if !recursive {
				return filepath.SkipDir
			}
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

func handleExtension(file string, res trid.Result, add, change bool) {
	ext := res.Definition.FileType.Extension
	if ext == "" {
		return
	}
	if idx := strings.Index(ext, "/"); idx != -1 {
		ext = ext[:idx]
	}
	ext = "." + strings.ToLower(strings.TrimPrefix(ext, "."))
	currentExt := filepath.Ext(file)

	newPath := ""
	if add {
		newPath = file + ext
	} else if change && !strings.EqualFold(currentExt, ext) {
		newPath = file[:len(file)-len(currentExt)] + ext
	}

	if newPath != "" && newPath != file {
		finalPath := newPath
		counter := 1
		for {
			if _, err := os.Stat(finalPath); os.IsNotExist(err) {
				break
			}
			extOnly := filepath.Ext(newPath)
			base := newPath[:len(newPath)-len(extOnly)]
			finalPath = fmt.Sprintf("%s (%d)%s", base, counter, extOnly)
			counter++
		}
		if err := os.Rename(file, finalPath); err == nil {
			fmt.Printf("Renamed: %s -> %s\n", file, finalPath)
		}
	}
}

func saveCSV(path string, results []ScanResult) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating CSV: %v\n", err)
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	w.Write([]string{"File", "TrID-Score", "%", "Extension(s)", "Filetype", "MIME type"})
	for _, r := range results {
		w.Write([]string{r.File, fmt.Sprint(r.Score), r.Perc, r.Ext, r.Type, r.Mime})
	}
	fmt.Printf("\nCSV file '%s' written (%d rows).\n", path, len(results))
}

func saveJSON(path string, results []ScanResult) {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		return
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return
	}
	fmt.Printf("\nJSON file '%s' written (%d rows).\n", path, len(results))
}

func loadDefaultDefinitions(a *trid.Analyzer) error {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	possiblePaths := []string{
		"triddefs.trd",
		filepath.Join(exeDir, "triddefs.trd"),
		"/usr/share/trid/triddefs.trd",
		"/usr/local/share/trid/triddefs.trd",
	}
	home, _ := os.UserHomeDir()
	if home != "" {
		possiblePaths = append(possiblePaths, filepath.Join(home, ".trid", "triddefs.trd"))
		possiblePaths = append(possiblePaths, filepath.Join(home, ".config", "TridUI", "triddefs.trd"))
		possiblePaths = append(possiblePaths, filepath.Join(home, "Library", "Application Support", "TridUI", "triddefs.trd"))
	}
	if appData := os.Getenv("APPDATA"); appData != "" {
		possiblePaths = append(possiblePaths, filepath.Join(appData, "TridUI", "triddefs.trd"))
	}
	if localAppData := os.Getenv("LOCALAPPDATA"); localAppData != "" {
		possiblePaths = append(possiblePaths, filepath.Join(localAppData, "TridUI", "triddefs.trd"))
	}
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return a.LoadDefinitions(path)
		}
	}
	return fmt.Errorf("no definitions found")
}

func printUsage(fs *flag.FlagSet) {
	fmt.Println("TrID - File Identifier v2.43 - (C) 2003-2025 By M.Pontello")
	fmt.Println("TridUI CLI v" + versionNumber + " - (C) 2025 By Cufiy")
	fmt.Printf("https://github.com/JMcrafter26/TridUI/\n\n")
	fmt.Println("\nUsage:")
	fmt.Println("  trid [files ...] [options]")
	fmt.Println("\nOptions:")
	fs.PrintDefaults()
}

func runUpdate(defsPath string) {
	if defsPath == "" {
		appData := os.Getenv("APPDATA")
		if appData != "" {
			defsPath = filepath.Join(appData, "TridUI", "triddefs.trd")
		} else {
			home, _ := os.UserHomeDir()
			if home != "" {
				defsPath = filepath.Join(home, ".trid", "triddefs.trd")
			} else {
				defsPath = "triddefs.trd"
			}
		}
	}
	fmt.Printf("Updating definitions at: %s\n", defsPath)
	dir := filepath.Dir(defsPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	updater := trid.NewUpdater(defsPath)
	info, err := updater.CheckForUpdates()
	if err != nil {
		fmt.Printf("Error checking for updates: %v\n", err)
		return
	}
	if info.IsUpToDate && info.CurrentMD5 != "none" {
		fmt.Println("Definitions are already up to date.")
		return
	}
	fmt.Println("New definitions available. Downloading...")
	err = updater.Update(func(downloaded, total int64) {
		if total > 0 {
			fmt.Printf("\rDownloading: %d%% (%d/%d bytes)", downloaded*100/total, downloaded, total)
		} else {
			fmt.Printf("\rDownloading: %d bytes", downloaded)
		}
	})
	fmt.Println()
	if err != nil {
		fmt.Printf("Error updating: %v\n", err)
		return
	}
	fmt.Println("Successfully updated definitions!")
}
