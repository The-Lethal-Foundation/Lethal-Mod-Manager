package handlers

import (
	"fmt"

	profile "github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/profile"
)

func handleGetProfiles() []string {
	_, profiles, err := profile.EnumProfiles()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return profiles
}
