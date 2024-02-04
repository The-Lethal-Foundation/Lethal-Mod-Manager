package handlers

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/The-Lethal-Foundation/lethal-core/config"
	"github.com/The-Lethal-Foundation/lethal-core/filesystem"
	"github.com/The-Lethal-Foundation/lethal-core/profile"
	"github.com/The-Lethal-Foundation/lethal-core/utils"
)

func handleInit() (string, error) {
	log.Println("Initializing file structure...")

	err := filesystem.InitializeStructure()
	if err != nil {
		log.Printf("Error initializing file structure: %v\n", err)
		return "", err
	}

	cfg, err := config.LoadConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName))
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		return "", err
	}

	// Check if other profiles were imported
	log.Println("Checking if other profiles were imported...")
	if cfg.OtherProfilesCloned == false {
		err := utils.CloneOtherProfiles(utils.KnownModManagersList)
		if err != nil {
			log.Printf("Error cloning other profiles: %v\n", err)
			return "", err
		}

		cfg.OtherProfilesCloned = true
		err = config.SaveConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName), cfg)
		if err != nil {
			log.Printf("Error saving config: %v\n", err)
			return "", err
		}
	}

	// Check if there are 0 profiles after initialization, if so create a default profile
	log.Println("Checking if there are 0 profiles after initialization...")
	profiles, err := profile.ListProfiles()
	if err != nil {
		log.Printf("Error listing profiles: %v\n", err)
		return "", err
	}
	if len(profiles) == 0 {
		log.Println("Creating default profile...")
		err := profile.CreateProfile("Default")
		if err != nil {
			log.Printf("Error creating default profile: %v\n", err)
			return "", err
		}

		// Update config
		cfg.LastUsedProfile = "Default"
		err = config.SaveConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName), cfg)
		if err != nil {
			log.Printf("Error saving config: %v\n", err)
			return "", err
		}

		return "Default", nil
	}

	// Check if the LastUsedProfile exists
	if cfg.LastUsedProfile != "" {
		log.Printf("Returning LastUsedProfile: %v\n", cfg.LastUsedProfile)
		return cfg.LastUsedProfile, nil
	}

	// If the LastUsedProfile does not exist, set it to the first profile
	cfg.LastUsedProfile = profiles[0]
	err = config.SaveConfig(filepath.Join(filesystem.GetDefaultPath(), config.ConfigFileName), cfg)
	if err != nil {
		log.Printf("Error saving config: %v\n", err)
		return "", err
	}

	return cfg.LastUsedProfile, nil
}

func handleRunGame(profile string) (string, error) {
	err := utils.LaunchGameProfile(profile)
	if err != nil {
		fmt.Printf("Error launching game: %v\n", err)
		return "", err
	}
	return "Game launched successfully", nil
}
