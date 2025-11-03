package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Version constants
const (
	AppVersion = "1.1.0"
	RepoOwner  = "JMcrafter26"
	RepoName   = "TridUI"
)

// UpdateInfo contains information about available updates
type UpdateInfo struct {
	CurrentVersion  string `json:"currentVersion"`
	LatestVersion   string `json:"latestVersion"`
	UpdateAvailable bool   `json:"updateAvailable"`
	ReleaseURL      string `json:"releaseUrl"`
	ReleaseNotes    string `json:"releaseNotes"`
	PublishedAt     string `json:"publishedAt"`
}

// GitHubRelease represents a GitHub release
type GitHubRelease struct {
	TagName     string `json:"tag_name"`
	Name        string `json:"name"`
	Body        string `json:"body"`
	HTMLURL     string `json:"html_url"`
	PublishedAt string `json:"published_at"`
	Prerelease  bool   `json:"prerelease"`
	Draft       bool   `json:"draft"`
}

// GitHubErrorResponse represents a GitHub API error response
type GitHubErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

// GetVersion returns the current application version
func (a *App) GetVersion() string {
	return AppVersion
}

// CheckForUpdates checks GitHub for the latest release and compares versions
func (a *App) CheckForUpdates() (*UpdateInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", RepoOwner, RepoName)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set User-Agent header as required by GitHub API
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", RepoName, AppVersion))
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		// Try to parse GitHub error response
		var ghError GitHubErrorResponse
		if err := json.Unmarshal(body, &ghError); err == nil && ghError.Message != "" {
			// Return user-friendly error message for common cases
			if resp.StatusCode == http.StatusForbidden && strings.Contains(ghError.Message, "rate limit") {
				return nil, fmt.Errorf("API rate limit exceeded. Please try again later or use authentication for higher limits")
			}
			if resp.StatusCode == http.StatusNotFound {
				return nil, fmt.Errorf("no releases found for this repository")
			}
			// Return the GitHub error message
			return nil, fmt.Errorf("%s", ghError.Message)
		}

		// Fallback to status code if we can't parse the error
		return nil, fmt.Errorf("GitHub API error (status %d)", resp.StatusCode)
	}

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Skip draft and prerelease versions
	if release.Draft || release.Prerelease {
		return &UpdateInfo{
			CurrentVersion:  AppVersion,
			LatestVersion:   AppVersion,
			UpdateAvailable: false,
			ReleaseURL:      "",
			ReleaseNotes:    "",
			PublishedAt:     "",
		}, nil
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	updateAvailable := compareVersions(AppVersion, latestVersion)

	return &UpdateInfo{
		CurrentVersion:  AppVersion,
		LatestVersion:   latestVersion,
		UpdateAvailable: updateAvailable,
		ReleaseURL:      release.HTMLURL,
		ReleaseNotes:    release.Body,
		PublishedAt:     release.PublishedAt,
	}, nil
}

// compareVersions returns true if remoteVersion is newer than currentVersion
func compareVersions(currentVersion, remoteVersion string) bool {
	// Simple semantic version comparison
	current := parseVersion(currentVersion)
	remote := parseVersion(remoteVersion)

	for i := 0; i < 3; i++ {
		if remote[i] > current[i] {
			return true
		}
		if remote[i] < current[i] {
			return false
		}
	}

	return false
}

// parseVersion parses a semantic version string into [major, minor, patch]
func parseVersion(version string) [3]int {
	var parts [3]int
	cleaned := strings.TrimPrefix(version, "v")

	// Split by dots
	segments := strings.Split(cleaned, ".")
	for i := 0; i < len(segments) && i < 3; i++ {
		// Parse each segment, ignoring errors (defaults to 0)
		fmt.Sscanf(segments[i], "%d", &parts[i])
	}

	return parts
}
