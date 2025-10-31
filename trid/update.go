package trid

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	DefsURL      = "http://mark0.net/download/triddefs.zip"
	DefsMD5URL   = "http://mark0.net/download/triddefs.trd.md5"
	DefsFilename = "triddefs.trd"
)

// UpdateInfo contains information about the definitions update
type UpdateInfo struct {
	CurrentMD5  string
	LatestMD5   string
	IsUpToDate  bool
	LastUpdated time.Time
	DefsCount   int
	Error       string
}

// Updater handles TrID definitions updates
type Updater struct {
	defsPath     string
	metadataPath string
}

// DefsMetadata stores metadata about the definitions file
type DefsMetadata struct {
	LastUpdated time.Time
	MD5Hash     string
	DefsCount   int
}

// NewUpdater creates a new definitions updater
func NewUpdater(defsPath string) *Updater {
	return &Updater{
		defsPath:     defsPath,
		metadataPath: defsPath + ".meta",
	}
}

// CheckForUpdates checks if there are updates available for the definitions
func (u *Updater) CheckForUpdates() (*UpdateInfo, error) {
	info := &UpdateInfo{
		LastUpdated: u.getLastUpdateTime(),
	}

	// Get current MD5 if file exists
	if _, err := os.Stat(u.defsPath); err == nil {
		currentMD5, err := u.calculateMD5File(u.defsPath)
		if err != nil {
			info.Error = fmt.Sprintf("Failed to calculate current MD5: %v", err)
			return info, err
		}
		info.CurrentMD5 = currentMD5

		// Count current definitions
		analyzer := NewAnalyzer()
		if err := analyzer.LoadDefinitions(u.defsPath); err == nil {
			info.DefsCount = analyzer.GetDefinitionCount()
		}
	} else {
		info.CurrentMD5 = "none"
	}

	// Get latest MD5 from server
	latestMD5, err := u.fetchLatestMD5()
	if err != nil {
		info.Error = fmt.Sprintf("Failed to fetch latest MD5: %v", err)
		return info, err
	}
	info.LatestMD5 = latestMD5

	// Compare MD5s
	info.IsUpToDate = (info.CurrentMD5 == info.LatestMD5)

	return info, nil
}

// Update downloads and installs the latest definitions
func (u *Updater) Update(progressCallback func(downloaded, total int64)) error {
	// Get latest MD5 first
	latestMD5, err := u.fetchLatestMD5()
	if err != nil {
		return fmt.Errorf("failed to fetch latest MD5: %w", err)
	}

	// Download the zip file
	trdData, err := u.downloadDefs(progressCallback)
	if err != nil {
		return fmt.Errorf("failed to download definitions: %w", err)
	}

	// Verify integrity
	downloadedMD5 := u.calculateMD5Data(trdData)
	if downloadedMD5 != latestMD5 {
		return errors.New("MD5 mismatch - downloaded file is corrupted")
	}

	// Save the file
	if err := os.WriteFile(u.defsPath, trdData, 0644); err != nil {
		return fmt.Errorf("failed to save definitions: %w", err)
	}

	// Update metadata
	metadata := DefsMetadata{
		LastUpdated: time.Now(),
		MD5Hash:     downloadedMD5,
	}

	// Count definitions
	analyzer := NewAnalyzer()
	if err := analyzer.LoadDefinitionsFromBytes(trdData); err == nil {
		metadata.DefsCount = analyzer.GetDefinitionCount()
	}

	u.saveMetadata(metadata)

	return nil
}

// fetchLatestMD5 retrieves the latest MD5 hash from the server
func (u *Updater) fetchLatestMD5() (string, error) {
	resp, err := http.Get(DefsMD5URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// The MD5 file contains just the hash
	md5Hash := string(bytes.TrimSpace(data))
	return md5Hash, nil
}

// downloadDefs downloads the definitions zip file and extracts the TRD file
func (u *Updater) downloadDefs(progressCallback func(downloaded, total int64)) ([]byte, error) {
	resp, err := http.Get(DefsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	// Read the response with progress tracking
	var buf bytes.Buffer
	totalSize := resp.ContentLength
	downloaded := int64(0)

	buffer := make([]byte, 8192)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			buf.Write(buffer[:n])
			downloaded += int64(n)
			if progressCallback != nil {
				progressCallback(downloaded, totalSize)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	// Extract triddefs.trd from the zip
	zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return nil, fmt.Errorf("failed to read zip: %w", err)
	}

	// Find and extract triddefs.trd
	for _, file := range zipReader.File {
		if file.Name == DefsFilename {
			rc, err := file.Open()
			if err != nil {
				return nil, fmt.Errorf("failed to open file in zip: %w", err)
			}
			defer rc.Close()

			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("failed to read file from zip: %w", err)
			}

			return data, nil
		}
	}

	return nil, errors.New("triddefs.trd not found in zip file")
}

// calculateMD5File calculates MD5 hash of a file
func (u *Updater) calculateMD5File(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return u.calculateMD5Data(data), nil
}

// calculateMD5Data calculates MD5 hash of data
func (u *Updater) calculateMD5Data(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

// getLastUpdateTime retrieves the last update time from metadata
func (u *Updater) getLastUpdateTime() time.Time {
	metadata := u.loadMetadata()
	return metadata.LastUpdated
}

// saveMetadata saves the metadata to disk
func (u *Updater) saveMetadata(metadata DefsMetadata) error {
	data := fmt.Sprintf("%s\n%d\n%d",
		metadata.LastUpdated.Format(time.RFC3339),
		metadata.DefsCount,
		metadata.LastUpdated.Unix())

	return os.WriteFile(u.metadataPath, []byte(data), 0644)
}

// loadMetadata loads the metadata from disk
func (u *Updater) loadMetadata() DefsMetadata {
	data, err := os.ReadFile(u.metadataPath)
	if err != nil {
		return DefsMetadata{}
	}

	lines := bytes.Split(data, []byte("\n"))
	if len(lines) < 1 {
		return DefsMetadata{}
	}

	metadata := DefsMetadata{}
	if t, err := time.Parse(time.RFC3339, string(lines[0])); err == nil {
		metadata.LastUpdated = t
	}

	if len(lines) >= 2 {
		fmt.Sscanf(string(lines[1]), "%d", &metadata.DefsCount)
	}

	return metadata
}
