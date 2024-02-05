package main

import (
	"fmt"
	"strings"

	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/mod"
	"github.com/nsf/termbox-go"
)

/*
	uilogic.go

	Ui logic for the CLI.
*/

func CLI_SearchMods(profile string) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	_, mods, err := mod.EnumMods(profile)
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

func CLI_ListMods(profile string) {
	_, mNames, err := mod.EnumMods(profile)
	if err != nil {
		fmt.Println("Error listing mods:", err)
		return
	}

	fmt.Println("Mods for profile", profile, ":")
	for i, mName := range mNames {
		fmt.Printf("%d: %s\n", i+1, mName)
	}
}

func CLI_ZipMods(profile string) {
	err := mod.ZipMods(profile, UpdateProgressBar)
	if err != nil {
		fmt.Println("Error zipping mods:", err)
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
			termbox.SetCell(j, i+3, ch, termbox.ColorDefault, termbox.ColorDefault) // +2 to account for the query and instruction messages
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

func CLI_InstallMod(profile string) {
	modUrl, err := promptForMod()
	if err != nil {
		fmt.Println("Error prompting for mod:", err)
		return
	}
	err = mod.InstallModFromUrl(profile, modUrl, UpdateProgressBar)
	if err != nil {
		fmt.Println("Error installing mod:", err)
		return
	}
	fmt.Println("\nMod installed successfully")
}

func UpdateProgressBar(current, total int, title string) {
	printLoadingBar(current, total, title)
}

func printLoadingBar(current, total int, title string) {
	const barLength = 30
	progress := float64(current) / float64(total)
	filledLength := int(progress * float64(barLength))

	fmt.Printf("\r%s: [", title)
	for i := 0; i < filledLength; i++ {
		fmt.Print("=")
	}
	for i := filledLength; i < barLength; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %d%%", int(progress*100))
}
