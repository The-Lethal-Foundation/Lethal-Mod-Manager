package handlers

import (
	"fmt"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
)

func handleInit() (string, error) {
	fmt.Println("init")

	err := util.CheckDirStructure()
	if err != nil {
		util.CreateDirStructure()
	}

	// load default profile
	config, err := util.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading last used profile: %v\n", err)
		return "", err
	}

	// check if default profile is ""
	if config.LastUsedProfile == "" {
		// TODO: make sure the Default profile exists
		config.LastUsedProfile = "Default"
		err := util.SaveConfig(config)
		if err != nil {
			fmt.Printf("Error saving last used profile: %v\n", err)
			return "", err
		}

		return config.LastUsedProfile, nil
	}

	return config.LastUsedProfile, nil
}
