package handlers

import (
	"fmt"
	"log"

	"github.com/The-Lethal-Foundation/lethal-core/api"
	"github.com/The-Lethal-Foundation/lethal-core/modmanager"
)

type GetModsResponse struct {
	ModName        string `json:"mod_name"`
	ModAuthor      string `json:"mod_author"`
	ModVersion     string `json:"mod_version"`
	ModDescription string `json:"mod_description"`
	ModPathName    string `json:"mod_path_name"`
}

func handleGetMods(profileName string) ([]GetModsResponse, error) {
	modsDetails, err := modmanager.ListMods(profileName)
	if err != nil {
		log.Printf("Error enumerating mods: %v", err)
		return nil, err
	}

	// Go over each mod and grab manifest if it exists
	mods := []GetModsResponse{}
	for _, modDetails := range modsDetails {

		mods = append(mods, GetModsResponse{
			ModName:        modDetails.Manifest.Name,
			ModAuthor:      modDetails.Author,
			ModVersion:     modDetails.Manifest.Version,
			ModDescription: modDetails.Manifest.Description,
			ModPathName:    modDetails.ModDirName,
		})
	}

	return mods, nil
}

func handleDeleteMod(profileName, modDirName string) (string, error) {
	err := modmanager.DeleteMod(profileName, modDirName)
	if err != nil {
		log.Printf("Error deleting mod: %v", err)
		return "", err
	}

	return "Mod deleted", nil
}

func handleGetGlobalMods(ordering string, section string, query string, page int) ([]api.GlobalModView, error) {
	mods, err := api.GlobalListMods(api.OrderingType(ordering), api.SectionType(section), query, page)
	if err != nil {
		log.Printf("Error getting global mods: %v", err)
		return nil, err
	}

	for _, mod := range mods {
		fmt.Printf("")
		fmt.Printf("Mod: %v", mod)
		fmt.Printf("")
	}

	return mods, nil
}

func handleInstallMod(profileName, authorName, modName string) (string, error) {
	err := modmanager.InstallMod(profileName, authorName, modName)
	if err != nil {
		log.Printf("Error installing mod: %v", err)
		return "", err
	}

	return "Mod installed", nil
}

func handleInstallModFromUrl(profileName, url string) (string, error) {
	err := modmanager.InstallModFromUrl(profileName, url)

	if err != nil {
		log.Printf("Error installing mod: %v", err)
		return "", err
	}

	return "Mod installed", nil
}
