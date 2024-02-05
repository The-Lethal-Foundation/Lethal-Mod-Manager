package util

import (
	"fmt"
	"os"
)

/*
util.go

Utilities for lethal ocmpany file system access.
*/

func GetModsPath(profileName string) (string, error) {
	profilePath, err := GetProfilePath(profileName)
	if err != nil {
		return "", err
	}

	return profilePath + "\\BepInEx\\plugins", nil
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

// Returns the path to the profiles folder
func GetProfilesPath() (string, error) {
	// path = %AppData% / Roaming / Thunderstore Mod Manager / DataFolder / LethalCompany / profiles / ...
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return "", fmt.Errorf("APPDATA environment variable is not set")
	}

	// 1. Check if the path exists
	path := appDataPath + "\\Thunderstore Mod Manager\\DataFolder\\LethalCompany\\profiles"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist: %s", path)
	}

	return path, nil
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}

	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == nil {
		return false, nil
	}

	return true, nil
}
