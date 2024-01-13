package mod

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func InstallModFromUrl(profile, modUrl string) error {
	// 1. Parse mod url into namespace and name
	mod, err := ParseModName(modUrl)
	if err != nil {
		return fmt.Errorf("error parsing mod name: %w", err)
	}

	err = InstallMod(profile, mod)
	if err != nil {
		return fmt.Errorf("error installing mod: %w", err)
	}

	return nil
}

// Installs / updates the specified mod in the specified profile
func InstallMod(profile string, mod ModName) error {

	// 1. Retrieve latest mod version
	modInfo, err := GetModInfo(mod)
	if err != nil {
		return fmt.Errorf("error getting mod info: %w", err)
	}
	mod.Version = modInfo.LatestVersion

	// 2. Check if such mod already exists or the version is outdated
	exists, err := LocalModExists(profile, mod)
	if err != nil {
		return fmt.Errorf("error checking if mod exists: %w", err)
	}

	outdated, err := LocalModOutdated(profile, mod)
	if err != nil {
		return fmt.Errorf("error checking if mod is outdated: %w", err)
	}

	if exists && !outdated {
		fmt.Printf("Mod %s already exists and is up to date\n", mod.Name)
		return nil
	}

	// 3. Download mod to temp folder
	zipName, err := DownloadMod(mod)
	if err != nil {
		return fmt.Errorf("error downloading mod: %w", err)
	}

	// 4. Unzip mod to profile folder
	err = UnzipMod(profile, zipName, mod)

	// 5. Go over the dependencies and install / update them

	return nil
}

// Downloads the mod zip file to a random temp folder
func DownloadMod(mod ModName) (string, error) {
	// Construct the download URL
	downloadURL := fmt.Sprintf("https://gcdn.thunderstore.io/live/repository/packages/%s-%s-%s.zip", mod.Author, mod.Name, mod.Version)

	// Make the HTTP request to download the mod
	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", fmt.Errorf("error making download request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	// Create a temporary file to save the downloaded mod
	tmpFile, err := os.CreateTemp("", fmt.Sprintf("%s-%s-%s-*.zip", mod.Author, mod.Name, mod.Version))
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %w", err)
	}
	defer tmpFile.Close()

	// Write the response content to the file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error writing to temp file: %w", err)
	}

	// Return the path to the downloaded file
	return tmpFile.Name(), nil
}

// ModInfo represents the information about a mod
type ModInfo struct {
	Downloads     int    `json:"downloads"`
	RatingScore   int    `json:"rating_score"`
	LatestVersion string `json:"latest_version"`
}

// GetModInfo fetches the information about a mod from the Thunderstore API
func GetModInfo(mod ModName) (*ModInfo, error) {
	url := fmt.Sprintf("https://thunderstore.io/api/v1/package-metrics/%s/%s", mod.Author, mod.Name)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var modInfo ModInfo
	if err := json.Unmarshal(body, &modInfo); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &modInfo, nil
}

func ParseModName(modUrl string) (ModName, error) {

	// Remove the trailing slash if exists
	modUrl = strings.TrimSuffix(modUrl, "/")

	// Parse the URL
	parsedUrl, err := url.Parse(modUrl)
	if err != nil {
		return ModName{}, fmt.Errorf("error parsing URL: %w", err)
	}

	// Split the path into segments
	segments := strings.Split(path.Clean(parsedUrl.Path), "/")[1:]

	// Assuming the URL format is like https://thunderstore.io/c/lethal-company/p/namespace/modname
	// and that there are at least 5 segments ("/c/lethal-company/p/namespace/modname")
	if len(segments) < 5 {
		return ModName{}, fmt.Errorf("invalid mod URL format")
	}

	return ModName{
		Author: segments[len(segments)-2],
		Name:   segments[len(segments)-1],
	}, nil
}
