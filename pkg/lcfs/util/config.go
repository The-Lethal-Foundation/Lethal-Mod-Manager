// util/config.go

package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CheckDirStructure() error {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return fmt.Errorf("APPDATA environment variable is not set")
	}

	dirs := []string{
		filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany"),
		filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany", "profiles"),
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return fmt.Errorf("required directory does not exist: %s", dir)
		}
	}

	return nil
}

func CreateDirStructure() error {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return fmt.Errorf("APPDATA environment variable is not set")
	}

	dirs := []string{
		filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany"),
		filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany", "profiles"),
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

type Config struct {
	LastUsedProfile string `json:"lastUsedProfile"`
}

func SaveConfig(config Config) error {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return fmt.Errorf("APPDATA environment variable is not set")
	}
	configPath := filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany", "LethalModManagerConfig.json")

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func LoadConfig() (Config, error) {
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return Config{}, fmt.Errorf("APPDATA environment variable is not set")
	}
	configPath := filepath.Join(appDataPath, "Thunderstore Mod Manager", "DataFolder", "LethalCompany", "LethalModManagerConfig.json")

	// check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// create empty config file
		config := Config{}
		err := SaveConfig(config)
		if err != nil {
			return Config{}, err
		}
		return config, nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
