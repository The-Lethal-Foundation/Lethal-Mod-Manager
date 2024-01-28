package handlers

import "github.com/zserge/lorca"

func SetupHandlers(ui lorca.UI) {

	ui.Bind("getProfiles", handleGetProfiles)

}
