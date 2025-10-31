package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"TridUI/trid"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	analyzer     *trid.Analyzer
	analyzerOnce sync.Once
	analyzerErr  error
)

// TridFileTypeResult represents a single file type match
type TridFileTypeResult struct {
	Name        string  `json:"name"`
	Extension   string  `json:"extension"`
	MimeType    string  `json:"mimeType"`
	Description string  `json:"description"`
	Confidence  float64 `json:"confidence"`
	URL         string  `json:"url"`
	Remarks     string  `json:"remarks"`
}

// TridScanResult represents the complete scan result
type TridScanResult struct {
	FileName     string               `json:"fileName"`
	FileSize     int64                `json:"fileSize"`
	Matches      []TridFileTypeResult `json:"matches"`
	BestMatch    *TridFileTypeResult  `json:"bestMatch"`
	TotalMatches int                  `json:"totalMatches"`
	Error        string               `json:"error,omitempty"`
}

// TridUpdateInfo represents information about definitions update
type TridUpdateInfo struct {
	CurrentMD5  string `json:"currentMD5"`
	LatestMD5   string `json:"latestMD5"`
	IsUpToDate  bool   `json:"isUpToDate"`
	LastUpdated string `json:"lastUpdated"`
	DefsCount   int    `json:"defsCount"`
	Error       string `json:"error,omitempty"`
}

// TridUpdateProgress represents download progress
type TridUpdateProgress struct {
	Downloaded int64  `json:"downloaded"`
	Total      int64  `json:"total"`
	Percentage int    `json:"percentage"`
	Message    string `json:"message"`
}

// initializeAnalyzer initializes the TrID analyzer with definitions
func initializeAnalyzer() (*trid.Analyzer, error) {
	analyzerOnce.Do(func() {
		analyzer = trid.NewAnalyzer()

		// Get the definitions file path from app data directory
		appDataDir, err := getAppDataDir()
		if err != nil {
			analyzerErr = fmt.Errorf("failed to get app data dir: %w", err)
			return
		}

		// Ensure the app data directory exists
		if err := os.MkdirAll(appDataDir, 0755); err != nil {
			analyzerErr = fmt.Errorf("failed to create app data dir: %w", err)
			return
		}

		defsPath := filepath.Join(appDataDir, "triddefs.trd")

		// Check if definitions file exists
		if _, err := os.Stat(defsPath); os.IsNotExist(err) {
			analyzerErr = fmt.Errorf("TrID definitions file not found at %s. Please download triddefs.trd and place it in the app directory", defsPath)
			return
		}

		// Load the definitions
		if err := analyzer.LoadDefinitions(defsPath); err != nil {
			analyzerErr = fmt.Errorf("failed to load TrID definitions: %w", err)
			return
		}
	})

	return analyzer, analyzerErr
}

// ProcessFile analyzes a file using TrID and returns the results
func (a *App) ProcessFile(filePath string) (*TridScanResult, error) {
	runtime.LogDebugf(a.ctx, "ProcessFile called with: %s", filePath)

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	result := &TridScanResult{
		FileName: filepath.Base(filePath),
		FileSize: fileInfo.Size(),
		Matches:  make([]TridFileTypeResult, 0),
	}

	// Initialize the analyzer
	analyzer, err := initializeAnalyzer()
	if err != nil {
		result.Error = err.Error()
		return result, nil // Return result with error instead of failing
	}

	// Analyze the file
	results, err := analyzer.AnalyzeFile(filePath)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to analyze file: %v", err)
		return result, nil
	}

	// Configuration for filtering
	const (
		minConfidence = 0.05 // Only include results with at least 5% confidence
		maxResults    = 50   // Limit to top 50 results
	)

	// Convert results to our format, filtering by confidence
	matchCount := 0
	for i, r := range results {
		// Skip results with very low confidence
		if r.Confidence < minConfidence {
			continue
		}

		// Limit total number of results
		if matchCount >= maxResults {
			break
		}

		match := TridFileTypeResult{
			Name:        r.Definition.FileType.Name,
			Extension:   r.Definition.FileType.Extension,
			MimeType:    r.Definition.FileType.MimeType,
			Description: r.Definition.FileType.Description,
			Confidence:  r.Confidence,
			URL:         r.Definition.FileType.URL,
			Remarks:     r.Definition.FileType.Remarks,
		}

		result.Matches = append(result.Matches, match)

		// Set best match (first result)
		if i == 0 {
			result.BestMatch = &match
		}

		matchCount++
	}

	result.TotalMatches = len(result.Matches)

	runtime.LogDebugf(a.ctx, "Filtered to %d matches (from %d total, min confidence: %.1f%%)",
		result.TotalMatches, len(results), minConfidence*100)

	runtime.LogDebugf(a.ctx, "Analysis complete: %d matches found", result.TotalMatches)

	return result, nil
}

// GetDefinitionsPath returns the path where TrID definitions should be placed
func (a *App) GetDefinitionsPath() (string, error) {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appDataDir, "triddefs.trd"), nil
}

// CheckDefinitionsExist checks if the TrID definitions file exists
func (a *App) CheckDefinitionsExist() (bool, error) {
	defsPath, err := a.GetDefinitionsPath()
	if err != nil {
		return false, err
	}
	_, err = os.Stat(defsPath)
	return err == nil, nil
}

// CheckForDefsUpdates checks if there are updates available for the definitions
func (a *App) CheckForDefsUpdates() (*TridUpdateInfo, error) {
	defsPath, err := a.GetDefinitionsPath()
	if err != nil {
		return nil, err
	}

	updater := trid.NewUpdater(defsPath)
	info, err := updater.CheckForUpdates()

	result := &TridUpdateInfo{
		CurrentMD5: info.CurrentMD5,
		LatestMD5:  info.LatestMD5,
		IsUpToDate: info.IsUpToDate,
		DefsCount:  info.DefsCount,
		Error:      info.Error,
	}

	if !info.LastUpdated.IsZero() {
		result.LastUpdated = info.LastUpdated.Format("2006-01-02 15:04:05")
	}

	return result, err
}

// UpdateDefinitions downloads and installs the latest definitions
func (a *App) UpdateDefinitions() error {
	defsPath, err := a.GetDefinitionsPath()
	if err != nil {
		return err
	}

	updater := trid.NewUpdater(defsPath)

	// Update with progress callback
	err = updater.Update(func(downloaded, total int64) {
		progress := &TridUpdateProgress{
			Downloaded: downloaded,
			Total:      total,
			Percentage: int((downloaded * 100) / total),
			Message:    fmt.Sprintf("Downloading: %d KB / %d KB", downloaded/1024, total/1024),
		}

		// Emit progress event to frontend
		runtime.EventsEmit(a.ctx, "trid:update:progress", progress)
	})

	if err != nil {
		return err
	}

	// Reset the analyzer to force reload with new definitions
	analyzerOnce = sync.Once{}
	analyzer = nil

	// Emit completion event
	runtime.EventsEmit(a.ctx, "trid:update:complete", true)

	return nil
}
