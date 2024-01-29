package handlers

import "github.com/zserge/lorca"

func SetupHandlers(ui lorca.UI) {

	ui.Bind("init", handleInit)

	// mods
	ui.Bind("getMods", handleGetMods)
	ui.Bind("openModDir", handleOpenModDir)
	ui.Bind("deleteMod", handleDeleteMod)
	ui.Bind("getGlobalMods", handleGetGlobalMods)

	// profile
	ui.Bind("getProfiles", handleGetProfiles)
	ui.Bind("saveLastUsedProfile", handleSaveLastUsedProfile)
	ui.Bind("loadLastUsedProfile", handleLoadLastUsedProfile)
}
