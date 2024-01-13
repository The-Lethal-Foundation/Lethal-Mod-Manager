package mod

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KonstantinBelenko/lethal_mod_manager/pkg/lcfs/profile"
	"github.com/nsf/termbox-go"
)

type ProgressCallback func(current, total int, title string)

type ModName struct {
	Author  string
	Name    string
	Version string
}

type ModManifest struct {
	Name          string   `json:"name"`
	VersionNumber string   `json:"version_number"`
	WebsiteUrl    string   `json:"website_url"`
	Description   string   `json:"description"`
	Dependencies  []string `json:"dependencies"`
}

func EnumMods(ProfileName string) ([]string, []string, error) {

	profilePath, err := profile.GetProfilePath(ProfileName)

	// 1. Check if the path exists
	path := profilePath + "\\BepInEx\\plugins"
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

// Converts a mod name string into a ModName struct
func LocalParseModName(modName string) (ModName, error) {
	// name consists of <author> ? optional -<name> ? optional -<version>
	parts := strings.Split(modName, "-")
	if len(parts) < 1 {
		return ModName{}, fmt.Errorf("invalid mod name: %s", modName)
	}

	if len(parts) == 1 {
		return ModName{
			Name: parts[0],
		}, nil
	}

	if len(parts) == 2 {
		return ModName{
			Author: parts[0],
			Name:   parts[1],
		}, nil
	}

	return ModName{
		Author:  parts[0],
		Name:    parts[1],
		Version: parts[2],
	}, nil

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

func GetModsPath(profileName string) (string, error) {
	profilePath, err := profile.GetProfilePath(profileName)
	if err != nil {
		return "", err
	}

	return profilePath + "\\BepInEx\\plugins", nil
}

func LocalGetModManifest(profileName, exactModName string) (ModManifest, error) {
	modsPath, err := GetModsPath(profileName)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error getting profile path: %w", err)
	}

	modPath := filepath.Join(modsPath, exactModName)
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return ModManifest{}, fmt.Errorf("mod does not exist: %s", modPath)
	}

	manifestPath := filepath.Join(modPath, "manifest.json")
	manifestFile, err := os.Open(manifestPath)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error opening manifest file: %w", err)
	}
	defer manifestFile.Close()

	// Wrap the file reader in a bufio.Reader
	reader := bufio.NewReader(manifestFile)

	// Read the first few bytes for BOM
	bom := make([]byte, 3)
	_, err = reader.Read(bom)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error reading file: %w", err)
	}
	if bom[0] != 0xEF || bom[1] != 0xBB || bom[2] != 0xBF {
		// Not a BOM; reset the reader to the start of the file
		_, err = manifestFile.Seek(0, 0)
		if err != nil {
			return ModManifest{}, fmt.Errorf("error seeking file: %w", err)
		}
		reader = bufio.NewReader(manifestFile)
	}

	manifest := ModManifest{}
	err = json.NewDecoder(reader).Decode(&manifest)
	if err != nil {
		return ModManifest{}, fmt.Errorf("error decoding manifest file: %w", err)
	}

	return manifest, nil
}
