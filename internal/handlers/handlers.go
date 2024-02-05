package handlers

import "github.com/zserge/lorca"

func SetupHandlers(ui lorca.UI, addr string) {

	handleGetAddr := func() string {
		return addr
	}

	// setup
	ui.Bind("init", handleInit)
	ui.Bind("getAddr", handleGetAddr)
	ui.Bind("runGame", handleRunGame)

	// mods
	ui.Bind("getMods", handleGetMods)
	ui.Bind("openModDir", handleOpenModDir)
	ui.Bind("deleteMod", handleDeleteMod)
	ui.Bind("getGlobalMods", handleGetGlobalMods)
	ui.Bind("installMod", handleInstallMod)
	ui.Bind("installModFromUrl", handleInstallModFromUrl)

	// profile
	ui.Bind("getProfiles", handleGetProfiles)
	ui.Bind("saveLastUsedProfile", handleSaveLastUsedProfile)
	ui.Bind("loadLastUsedProfile", handleLoadLastUsedProfile)
	ui.Bind("renameProfile", handleRenameProfile)
}
