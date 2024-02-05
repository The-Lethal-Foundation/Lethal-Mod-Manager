package mod

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
	types "github.com/KonstantinBelenko/lethal-mod-manager/pkg/types"
)

/*
	modparse.go

	Parsing mod names / urls / manifests.
*/

type ModManifest struct {
	Name          string   `json:"name"`
	VersionNumber string   `json:"version_number"`
	WebsiteUrl    string   `json:"website_url"`
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
}

// Converts a mod name string into a ModName struct
func LocalParseModName(modName string) (types.ModName, error) {
	// name consists of <author> ? optional -<name> ? optional -<version>
	parts := strings.Split(modName, "-")
	if len(parts) < 1 {
		return types.ModName{}, fmt.Errorf("invalid mod name: %s", modName)
	}

	if len(parts) == 1 {
		return types.ModName{
			Name: parts[0],
		}, nil
	}

	if len(parts) == 2 {
		return types.ModName{
			Author: parts[0],
			Name:   parts[1],
		}, nil
	}

	return types.ModName{
		Author:  parts[0],
		Name:    parts[1],
		Version: parts[2],
	}, nil

}

// Reads the manifest file for the specified mod, if it exists
func LocalGetModManifest(profileName, exactModName string) (ModManifest, error) {
	modsPath, err := util.GetModsPath(profileName)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error getting profile path: %w", err)
	}

	modPath := filepath.Join(modsPath, exactModName)
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return ModManifest{}, fmt.Errorf("mod does not exist: %s", modPath)
	}

	manifestPath := filepath.Join(modPath, "manifest.json")
	manifestFile, err := os.Open(manifestPath)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error opening manifest file: %w", err)
	}
	defer manifestFile.Close()

	// Wrap the file reader in a bufio.Reader
	reader := bufio.NewReader(manifestFile)

	// Read the first few bytes for BOM
	bom := make([]byte, 3)
	_, err = reader.Read(bom)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error reading file: %w", err)
	}
	if bom[0] != 0xEF || bom[1] != 0xBB || bom[2] != 0xBF {
		// Not a BOM; reset the reader to the start of the file
		_, err = manifestFile.Seek(0, 0)
		if err != nil {
			return ModManifest{}, fmt.Errorf("error seeking file: %w", err)
		}
		reader = bufio.NewReader(manifestFile)
	}

	manifest := ModManifest{}
	err = json.NewDecoder(reader).Decode(&manifest)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error decoding manifest file: %w", err)
	}

	return manifest, nil
}

func LocalModHasManifest(profileName, exactModName string) (bool, error) {
	modsPath, err := util.GetModsPath(profileName)
	if err != nil {
		return false, fmt.Errorf("error getting profile path: %w", err)
	}

	modPath := filepath.Join(modsPath, exactModName)
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return false, fmt.Errorf("mod does not exist: %s", modPath)
	}

	manifestPath := filepath.Join(modPath, "manifest.json")
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}
