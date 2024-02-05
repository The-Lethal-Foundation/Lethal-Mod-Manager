package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/The-Lethal-Foundation/lethal-core/config"
	"github.com/The-Lethal-Foundation/lethal-core/filesystem"
	"github.com/The-Lethal-Foundation/lethal-core/profile"
)

func handleGetProfiles() []string {
	profiles, err := profile.ListProfiles()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return profiles
}

func handleSaveLastUsedProfile(lastUsedProfile string) (string, error) {
	// load config
	cfg, err := config.LoadConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName))
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return "", err
	}

	// Update config
	cfg.LastUsedProfile = lastUsedProfile
	err = config.SaveConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName), cfg)
	if err != nil {
		fmt.Printf("Error saving config: %v\n", err)
		return "", err
	}

	return "Last used profile saved successfully", nil
}

func handleLoadLastUsedProfile() (string, error) {
	// load config
	cfg, err := config.LoadConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName))
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return "", err
	}

	return cfg.LastUsedProfile, nil
}

func handleRenameProfile(oldName, newName string) (string, error) {
	err := profile.RenameProfile(oldName, newName)
	if err != nil {
		fmt.Printf("Error renaming profile: %v\n", err)
		return "", err
	}
	return "Profile renamed successfully", nil
}
