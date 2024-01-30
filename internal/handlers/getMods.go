package handlers

import (
	"fmt"
	"log"

	mod "github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/mod"
	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/tsapi"
	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/types"
)

type GetModsResponse struct {
	ModName        string `json:"mod_name"`
	ModAuthor      string `json:"mod_author"`
	ModVersion     string `json:"mod_version"`
	ModDescription string `json:"mod_description"`
	ModPathName    string `json:"mod_path_name"`
}

func handleGetMods(profileName string) ([]GetModsResponse, error) {
	_, modNames, err := mod.EnumMods(profileName)
	if err != nil {
		log.Printf("Error enumerating mods: %v", err)
		return nil, err
	}

	// Go over each mod and grab manifest if it exists
	mods := []GetModsResponse{}
	for _, modName := range modNames {
		parsedModName, err := mod.LocalParseModName(modName)
		if err != nil {
			log.Printf("Error parsing mod name: %v", err)
			return nil, err
		}

		hasManifest, err := mod.LocalModHasManifest(profileName, modName)
		if err != nil {
			log.Printf("Error checking manifest: %v", err)
			continue
		}

		if !hasManifest {
			continue
		}

		manifest, err := mod.LocalGetModManifest(profileName, modName)
		if err != nil {
			log.Printf("Error getting manifest: %v", err)
			return nil, err
		}

		mods = append(mods, GetModsResponse{
			ModName:        manifest.Name,
			ModAuthor:      parsedModName.Author,
			ModVersion:     manifest.VersionNumber,
			ModDescription: manifest.Description,
			ModPathName:    modName,
		})
	}

	return mods, nil
}

func handleDeleteMod(profileName, modPathName string) (string, error) {
	err := mod.DeleteMod(profileName, modPathName)
	if err != nil {
		log.Printf("Error deleting mod: %v", err)
		return "", err
	}

	return "Mod deleted", nil
}

func handleGetGlobalMods(ordering string, section string, query string, page int) ([]tsapi.GlobalModView, error) {
	mods, err := tsapi.GlobalListMods(tsapi.OrderingType(ordering), tsapi.SectionType(section), query, page)
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
	err := mod.InstallMod(profileName, types.ModName{
		Name:   modName,
		Author: authorName,
	}, func(current, total int, title string) {})

	if err != nil {
		log.Printf("Error installing mod: %v", err)
		return "", err
	}

	return "Mod installed", nil
}

func handleInstallModFromUrl(profileName, url string) (string, error) {
	err := mod.InstallModFromUrl(profileName, url, func(current, total int, title string) {})

	if err != nil {
		log.Printf("Error installing mod: %v", err)
		return "", err
	}

	return "Mod installed", nil
}
