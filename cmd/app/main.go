package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/mod"
	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/profile"
	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/util"
	"github.com/zserge/lorca"
)

//go:embed www
var fs embed.FS

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("", "", 1080, 650, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	profilesPath, err := util.GetProfilesPath()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Serve files from the embedded file system under the /www path
	mux.Handle("/", http.StripPrefix("/www/", http.FileServer(http.FS(fs))))

	// Serve local files from the profilesPath under the /images path
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(profilesPath))))

	// Use the mux in the HTTP server
	go http.Serve(ln, mux)

	// go http.Serve(ln, http.FileServer(http.FS(fs)))

	// Setup UI
	ui.Load(fmt.Sprintf("http://%s/www/www", ln.Addr()))

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
					ModPicture:     fmt.Sprintf("http://%s/images/%s/BepInEx/plugins/%s/icon.png", ln.Addr(), profile, modName),
				})
			}
		}

		return mods
	})

	ui.Bind("installMod", func(profile string, modUrl string) string {
		err := mod.InstallModFromUrl(profile, modUrl, func(current, total int, title string) {})
		if err != nil {
			return err.Error()
		}
		return "Success"
	})

	ui.Bind("setProfile", func(profile string) {
		log.Println("Switching to profile", profile)
		ui.Load(fmt.Sprintf("http://%s/www/profile.html", ln.Addr()))
	})

	// Bind Ctr+R to reload current page
	ui.Bind("reload", func() {
		ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
	})

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
