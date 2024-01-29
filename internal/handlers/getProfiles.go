package handlers

import (
	"fmt"

	profile "github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/profile"
	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
)

func handleGetProfiles() []string {
	_, profiles, err := profile.EnumProfiles()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return profiles
}

func handleSaveLastUsedProfile(lastUsedProfile string) (string, error) {
	config := util.Config{LastUsedProfile: lastUsedProfile}
	if err := util.SaveConfig(config); err != nil {
		fmt.Printf("Error saving last used profile: %v\n", err)
		return "Error saving last used profile", err
	}
	return "Last used profile saved successfully", nil
}

func handleLoadLastUsedProfile() (string, error) {
	config, err := util.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading last used profile: %v\n", err)
		return "", err
	}
	return config.LastUsedProfile, nil
}
