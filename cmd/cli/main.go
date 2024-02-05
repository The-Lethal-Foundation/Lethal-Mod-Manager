package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/profile"
)

func main() {
	// Attempt to load saved profile
	savedProfile, err := profile.LoadProfile()
	if err != nil || savedProfile == "" {
		// If no profile is saved, enumerate profiles and ask the user to choose one
		savedProfile, err = promptForProfile()
		if err != nil {
			fmt.Println("Failed to select a profile:", err)
			return
		}

		// Save the selected profile
		if err := profile.SaveProfile(savedProfile); err != nil {
			fmt.Println("Failed to save profile:", err)
			return
		}
	}

	fmt.Printf("Using profile: %s\n", savedProfile)

	for {
		fmt.Println("Enter command ('install mod', 'search mods', 'list mods', 'zip mods', or 'quit'):")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "install mod":
			CLI_InstallMod(savedProfile)
		case "list mods":
			CLI_ListMods(savedProfile)
		case "search mods":
			CLI_SearchMods(savedProfile)
		case "zip mods":
			CLI_ZipMods(savedProfile)
		case "quit":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func promptForMod() (string, error) {
	fmt.Print("Enter mod URL: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input), nil
}

func promptForProfile() (string, error) {
	_, profileNames, err := profile.EnumProfiles()
	if err != nil {
		return "", err
	}

	fmt.Println("Available profiles:")
	for i, name := range profileNames {
		fmt.Printf("%d: %s\n", i+1, name)
	}
	fmt.Print("Choose a profile (number): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || choice < 1 || choice > len(profileNames) {
		return "", errors.New("invalid choice")
	}

	return profileNames[choice-1], nil
}
