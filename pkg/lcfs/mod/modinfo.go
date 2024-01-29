package mod

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/tsapi"
	types "github.com/KonstantinBelenko/lethal-mod-manager/pkg/types"
)

/*
	modinfo.go

	Structs & functions that provide information about mods.
*/

func EnumMods(ProfileName string) ([]string, []string, error) {

	profilePath, err := util.GetProfilePath(ProfileName)

	// 1. Check if the path exists
	path := profilePath + "\\BepInEx\\plugins"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("path does not exist: %s", path)
	}

	// 2. Enumerate the profiles in folder
	mod_names := []string{}
	mod_paths := []string{}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to enumerate profiles: %s", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			mod_paths = append(mod_paths, filepath.Join(path, entry.Name()))
			mod_names = append(mod_names, entry.Name())
		}
	}

	return mod_paths, mod_names, nil
}

func LocalModExists(profile string, mod types.ModName) (bool, error) {

	// Ignore the mod and return true in case the mod name contains "BepInExPack"
	// We do this because BepInExPack is the host for all BepInEx plugins
	if strings.Contains(mod.Name, "BepInExPack") {
		return true, nil
	}

	// Same for MMHOOK
	if strings.Contains(mod.Author, "MMHOOK") {
		return true, nil
	}

	// Enumerate the mods
	_, mods, err := EnumMods(profile)
	if err != nil {
		return false, fmt.Errorf("error enumerating mods: %w", err)
	}

	// Check if the mod exists
	for _, m := range mods {
		if strings.Contains(m, mod.Name) {

			// Check if mod dir is empty, if so, remove it and return false
			profilePath, err := util.GetProfilePath(profile)
			if err != nil {
				return false, fmt.Errorf("error getting profile path: %w", err)
			}

			modPath := filepath.Join(profilePath, "BepInEx", "plugins", m)
			isEmpty, err := util.IsDirEmpty(modPath)
			if err != nil {
				return false, fmt.Errorf("error checking if mod dir is empty: %w", err)
			}

			if isEmpty {
				err := os.Remove(modPath)
				if err != nil {
					return false, fmt.Errorf("error removing mod dir: %w", err)
				}
				return false, nil
			}

			return true, nil
		}
	}

	return false, nil
}

// Checks if the mod is outdated in the specified profile
// Requires the provided mod to have a version
func LocalModOutdated(profile string, mod types.ModName) (bool, error) {

	// Ignore the mod and return true in case the mod name contains "BepInExPack"
	// We do this because BepInExPack is the host for all BepInEx plugins
	if strings.Contains(mod.Name, "BepInExPack") {
		return false, nil
	}

	// Same for MMHOOK
	if strings.Contains(mod.Author, "MMHOOK") {
		return false, nil
	}

	_, mods, err := EnumMods(profile)
	if err != nil {
		return false, fmt.Errorf("error enumerating mods: %w", err)
	}

	// Check if the mod is outdated
	for _, m := range mods {
		otherMod, err := LocalParseModName(m)
		if err != nil {
			return false, fmt.Errorf("error parsing mod name: %w", err)
		}

		if strings.Contains(otherMod.Name, "BepInExPack") {
			continue
		}

		// Same for MMHOOK
		if strings.Contains(otherMod.Name, "MMHOOK") {
			continue
		}

		otherModManifest, err := LocalGetModManifest(profile, m)
		if err != nil {
			continue
		}
		otherMod.Version = otherModManifest.VersionNumber

		if mod.Name == otherMod.Name {
			if otherMod.Version != "" && mod.Version != "" {
				if otherMod.Version != mod.Version {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func InstallModFromUrl(profile, modUrl string, installDepCallback types.ProgressCallback) error {
	// 1. Parse mod url into namespace and name
	mod, err := ParseModName(modUrl)
	if err != nil {
		return fmt.Errorf("error parsing mod name: %w", err)
	}

	err = InstallMod(profile, mod, installDepCallback)
	if err != nil {
		return fmt.Errorf("error installing mod: %w", err)
	}

	return nil
}

// Installs / updates the specified mod in the specified profile
func InstallMod(profile string, mod types.ModName, installDepCallback types.ProgressCallback) error {

	// 1. Retrieve latest mod version
	modInfo, err := tsapi.GetModInfo(mod)
	if err != nil {
		return fmt.Errorf("error getting mod info: %w", err)
	}
	mod.Version = modInfo.LatestVersion

	// 2. Check if such mod already exists or the version is outdated
	exists, err := LocalModExists(profile, mod)
	if err != nil {
		return fmt.Errorf("error checking if mod exists: %w", err)
	}

	outdated, err := LocalModOutdated(profile, mod)
	if err != nil {
		return fmt.Errorf("error checking if mod is outdated: %w", err)
	}

	if exists && !outdated {
		return nil
	}

	// 3. Download mod to temp folder
	zipName, err := tsapi.DownloadMod(mod)
	if err != nil {
		return fmt.Errorf("error downloading mod: %w", err)
	}

	// 4. Unzip mod to profile folder
	newModName, err := UnzipMod(profile, zipName, mod)
	if err != nil {
		return fmt.Errorf("error unzipping mod: %w", err)
	}

	// 5. Go over the dependencies and install / update them
	fmt.Printf("Installing dependencies for mod %s\n", mod.Name)
	manifest, err := LocalGetModManifest(profile, newModName)
	if err != nil {
		return fmt.Errorf("error getting mod manifest: %w", err)
	}

	numberOfDeps := len(manifest.Dependencies)
	numberOfDepsInstalled := 0
	installDepCallback(0, numberOfDeps, "Installing dependencies")
	for _, dep := range manifest.Dependencies {

		depMod, err := LocalParseModName(dep)
		if err != nil {
			return fmt.Errorf("error parsing mod name: %w", err)
		}
		err = InstallMod(profile, depMod, installDepCallback)
		if err != nil {
			return fmt.Errorf("error installing dependency: %w", err)
		}

		numberOfDepsInstalled++
		installDepCallback(numberOfDepsInstalled, numberOfDeps, "Installing dependencies")
	}

	return nil
}

func ParseModName(modUrl string) (types.ModName, error) {

	// Remove the trailing slash if exists
	modUrl = strings.TrimSuffix(modUrl, "/")

	// Parse the URL
	parsedUrl, err := url.Parse(modUrl)
	if err != nil {
		return types.ModName{}, fmt.Errorf("error parsing URL: %w", err)
	}

	// Split the path into segments
	segments := strings.Split(path.Clean(parsedUrl.Path), "/")[1:]

	// Assuming the URL format is like https://thunderstore.io/c/lethal-company/p/namespace/modname
	// and that there are at least 5 segments ("/c/lethal-company/p/namespace/modname")
	if len(segments) < 5 {
		return types.ModName{}, fmt.Errorf("invalid mod URL format")
	}

	return types.ModName{
		Author: segments[len(segments)-2],
		Name:   segments[len(segments)-1],
	}, nil
}
