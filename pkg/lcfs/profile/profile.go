package profile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const saveFilePath = "selected_profile.txt"

// Enumerates the available thuenderstore profiles for Lethal Company.
// Returns a slice of profile paths and a slice of profile names, or an error.
func EnumProfiles() ([]string, []string, error) {

	// path = %AppData% / Roaming / Thunderstore Mod Manager / DataFolder / LethalCompany / profiles / ...
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return nil, nil, fmt.Errorf("APPDATA environment variable is not set")
	}

	// 1. Check if the path exists
	path := appDataPath + "\\Thunderstore Mod Manager\\DataFolder\\LethalCompany\\profiles"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("path does not exist: %s", path)
	}

	// 2. Enumerate the profiles in folder
	profile_names := []string{}
	profile_paths := []string{}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to enumerate profiles: %s", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			profile_paths = append(profile_paths, filepath.Join(path, entry.Name()))
			profile_names = append(profile_names, entry.Name())
		}
	}

	return profile_paths, profile_names, nil
}

func SaveProfile(profile string) error {
	return os.WriteFile(saveFilePath, []byte(profile), 0644)
}

func LoadProfile() (string, error) {
	data, err := os.ReadFile(saveFilePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func GetProfilePath(profile string) (string, error) {
	// path = %AppData% / Roaming / Thunderstore Mod Manager / DataFolder / LethalCompany / profiles / ...
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return "", fmt.Errorf("APPDATA environment variable is not set")
	}

	// 1. Check if the path exists
	path := appDataPath + "\\Thunderstore Mod Manager\\DataFolder\\LethalCompany\\profiles\\" + profile
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %s", path)
	}

	return path, nil
}
