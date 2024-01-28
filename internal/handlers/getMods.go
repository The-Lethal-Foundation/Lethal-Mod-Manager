package handlers

import (
	"log"

	mod "github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/mod"
)

type GetModsResponse struct {
	ModName        string `json:"mod_name"`
	ModAuthor      string `json:"mod_author"`
	ModVersion     string `json:"mod_version"`
	ModDescription string `json:"mod_description"`
	ModPathName    string `json:"mod_path_name"`
}

func handleGetMods(profileName string) []GetModsResponse {
	_, modNames, err := mod.EnumMods(profileName)
	if err != nil {
		log.Fatal(err)
	}

	// Go over each mod and grab manifest if it exists
	mods := []GetModsResponse{}
	for _, modName := range modNames {
		parsedModName, err := mod.LocalParseModName(modName)
		if err != nil {
			log.Fatal(err)
		}

		hasManifest, err := mod.LocalModHasManifest(profileName, modName)
		if err != nil {
			log.Fatal(err)
		}
		if hasManifest {
			manifest, err := mod.LocalGetModManifest(profileName, modName)
			if err != nil {
				log.Fatal(err)
			}

			mods = append(mods, GetModsResponse{
				ModName:        manifest.Name,
				ModAuthor:      parsedModName.Author,
				ModVersion:     manifest.VersionNumber,
				ModDescription: manifest.Description,
				ModPathName:    modName,
			})
		}
	}

	return mods
}
