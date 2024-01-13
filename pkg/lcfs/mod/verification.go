package mod

import (
	"fmt"
	"strings"
)

func LocalModExists(profile string, mod ModName) (bool, error) {

	// Ignore the mod and return true in case the mod name contains "BepInExPack"
	// We do this because BepInExPack is the host for all BepInEx plugins
	if strings.Contains(mod.Name, "BepInExPack") {
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
			return true, nil
		}
	}

	return false, nil
}

func LocalModOutdated(profile string, mod ModName) (bool, error) {

	// Ignore the mod and return true in case the mod name contains "BepInExPack"
	// We do this because BepInExPack is the host for all BepInEx plugins
	if strings.Contains(mod.Name, "BepInExPack") {
		return true, nil
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
