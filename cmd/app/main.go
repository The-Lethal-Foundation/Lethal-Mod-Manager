package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/mod"
	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/profile"
	"github.com/zserge/lorca"
)

//go:embed www
var fs embed.FS

type context struct {
	sync.Mutex
	selectedProfile string
}

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("", "", 1080, 650, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Init context
	ctx := &context{}

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))
	// ui.Load("file://C:\\Users\\Tractor\\Desktop\\sideprojects\\my-thunderstoremm\\www\\index.html")

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
	// ui.Load("data:text/html,file://C:\\Users\\Tractor\\Desktop\\sideprojects\\my-thunderstoremm\\www\\index.html")

	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	ui.Bind("getProfiles", func() []string {
		_, profiles, err := profile.EnumProfiles()
		if err != nil {
			log.Fatal(err)
		}
		return profiles
	})

	ui.Bind("getMods", func(profile string) []GetModsResponse {
		_, modNames, err := mod.EnumMods(profile)
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

			hasManifest, err := mod.LocalModHasManifest(profile, modName)
			if err != nil {
				log.Fatal(err)
			}
			if hasManifest {
				manifest, err := mod.LocalGetModManifest(profile, modName)
				if err != nil {
					log.Fatal(err)
				}

				mods = append(mods, GetModsResponse{
					ModName:        manifest.Name,
					ModAuthor:      parsedModName.Author,
					ModVersion:     manifest.VersionNumber,
					ModDescription: manifest.Description,
					ModPicture:     "placeholder.svg",
				})
			}
		}

		return mods
	})

	ui.Bind("setProfile", func(profile string) {
		log.Println("Switching to profile", profile)
		ctx.Lock()
		ctx.selectedProfile = profile
		ctx.Unlock()
		ui.Load(fmt.Sprintf("http://%s/www/profile.html", ln.Addr()))
	})

	ui.Bind("getSelectedProfile", func() string {
		ctx.Lock()
		defer ctx.Unlock()
		return ctx.selectedProfile
	})

	// Bind Ctr+R to reload current page
	ui.Bind("reload", func() {
		ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
	})

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	// ui.Eval(`
	// 	console.log("Hello, world!");
	// 	console.log('Multiple values:', [1, false, {"x":5}]);
	// `)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
