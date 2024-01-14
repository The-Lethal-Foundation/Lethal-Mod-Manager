package tsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	types "github.com/KonstantinBelenko/lethal_mod_manager/pkg/types"
)

/*
	tsapi.go

	Thunderstore API access.
*/

// GetModInfo fetches the information about a mod from the Thunderstore API
func GetModInfo(mod types.ModName) (*types.ModInfoResponse, error) {
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

	var modInfo types.ModInfoResponse
	if err := json.Unmarshal(body, &modInfo); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &modInfo, nil
}

// Downloads the mod zip file to a random temp folder
func DownloadMod(mod types.ModName) (string, error) {
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
