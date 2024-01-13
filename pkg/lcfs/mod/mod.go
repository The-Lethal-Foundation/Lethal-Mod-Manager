package mod

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nsf/termbox-go"
)

func EnumMods(ProfileName string) ([]string, []string, error) {

	// path = %AppData% / Roaming / Thunderstore Mod Manager / DataFolder / LethalCompany / profiles / ...
	appDataPath := os.Getenv("APPDATA")
	if appDataPath == "" {
		return nil, nil, fmt.Errorf("APPDATA environment variable is not set")
	}

	// 1. Check if the path exists
	path := appDataPath + "\\Thunderstore Mod Manager\\DataFolder\\LethalCompany\\profiles\\" + ProfileName + "\\BepInEx\\plugins"
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

func SearchMods(profile string) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	_, mods, err := EnumMods(profile)
	if err != nil {
		fmt.Println("Error retrieving mods:", err)
		return
	}

	query := ""
	for {
		printMods(mods, query)

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				if len(query) > 0 {
					query = query[:len(query)-1]
				}
			default:
				if ev.Ch != 0 {
					query += string(ev.Ch)
				}
			}
		}
	}
}

func ListMods(profile string) {
	_, mNames, err := EnumMods(profile)
	if err != nil {
		fmt.Println("Error listing mods:", err)
		return
	}

	fmt.Println("Mods for profile", profile, ":")
	for i, mName := range mNames {
		fmt.Printf("%d: %s\n", i+1, mName)
	}
}

func printMods(mods []string, query string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Display the instruction message
	instMsg := "Press 'Escape' to exit search"
	for i, ch := range instMsg {
		termbox.SetCell(i, 0, ch, termbox.ColorYellow, termbox.ColorDefault)
	}

	// Display the user's current query
	queryMsg := fmt.Sprintf("Search: %s", query)
	for i, ch := range queryMsg {
		termbox.SetCell(i, 1, ch, termbox.ColorWhite, termbox.ColorDefault)
	}

	// Display filtered mods
	filteredMods := filterMods(mods, query)
	for i, mod := range filteredMods {
		for j, ch := range mod {
			termbox.SetCell(j, i+2, ch, termbox.ColorDefault, termbox.ColorDefault) // +2 to account for the query and instruction messages
		}
	}

	termbox.Flush()
}

func filterMods(mods []string, query string) []string {
	var filtered []string
	for _, mod := range mods {
		if strings.Contains(strings.ToLower(mod), strings.ToLower(query)) {
			filtered = append(filtered, mod)
		}
	}
	return filtered
}
