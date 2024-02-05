package tsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	types "github.com/KonstantinBelenko/lethal-mod-manager/pkg/types"
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
	downloadURL := fmt.Sprintf("https://gcdn.thunderstore.io/live/repository/packages/%s-%s-%s.zip", mod.Author, mod.Name, mod.Version)

	var resp *http.Response
	var err error
	maxRetries := 5

	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err = http.Get(downloadURL)
		if err != nil {
			return "", fmt.Errorf("error making download request: %w", err)
		}

		if resp.StatusCode == http.StatusOK {
			break
		} else if resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close() // Close the response body before retrying
			if attempt < maxRetries {
				time.Sleep(2 * time.Second)
				continue
			}
		} else {
			resp.Body.Close()
			return "", fmt.Errorf("received non-OK response status: %s", resp.Status)
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download after %d attempts", maxRetries)
	}

	tmpFile, err := os.CreateTemp("", fmt.Sprintf("%s-%s-%s-*.zip", mod.Author, mod.Name, mod.Version))
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %w", err)
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error writing to temp file: %w", err)
	}

	return tmpFile.Name(), nil
}
