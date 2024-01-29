package handlers

import (
	"os/exec"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
)

func handleOpenModDir(profile, modPathName string) (string, error) {

	profileMods, err := util.GetModsPath(profile)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("explorer", profileMods+"\\"+modPathName)
	err = cmd.Run()

	return "Mod directory opened", nil
}
