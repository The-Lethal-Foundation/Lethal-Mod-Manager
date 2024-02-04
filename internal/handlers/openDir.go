package handlers

import (
	"os/exec"
	"path/filepath"

	"github.com/The-Lethal-Foundation/lethal-core/filesystem"
)

func handleOpenModDir(profile, modPathName string) (string, error) {

	modFullPath := filepath.Join(filesystem.GetDefaultPath(), "LethalCompany", "Profiles", profile, "BepInEx", "plugins", modPathName)

	cmd := exec.Command("explorer", modFullPath)
	_ = cmd.Run()

	return "Mod directory opened", nil
}
