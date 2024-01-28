package handlers

import "github.com/zserge/lorca"

func SetupHandlers(ui lorca.UI) {

	ui.Bind("init", handleInit)

	ui.Bind("getProfiles", handleGetProfiles)
	ui.Bind("getMods", handleGetMods)
}
