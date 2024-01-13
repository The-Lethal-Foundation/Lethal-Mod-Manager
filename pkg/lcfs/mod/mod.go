package mod

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nsf/termbox-go"
)

type ProgressCallback func(current, total int)

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

func ZipMods(profile string, progressCallback ProgressCallback) error {
	modPaths, _, err := EnumMods(profile)
	if err != nil {
		return fmt.Errorf("error enumerating mods: %w", err)
	}

	totalFiles, err := countTotalFiles(modPaths)
	if err != nil {
		return fmt.Errorf("error counting files: %w", err)
	}

	// Determine the path to the user's desktop
	desktopPath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home directory: %w", err)
	}
	desktopPath = filepath.Join(desktopPath, "Desktop", "LethalCompanyMods.zip")

	// Create a zip file
	zipFile, err := os.Create(desktopPath)
	if err != nil {
		return fmt.Errorf("error creating zip file: %w", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add each mod file to the zip
	var filesProcessed int
	for _, modPath := range modPaths {
		if err := addFileToZip(zipWriter, modPath, progressCallback, &filesProcessed, totalFiles); err != nil {
			return fmt.Errorf("error adding file to zip: %w", err)
		}
	}

	fmt.Printf("\nMods zipped successfully at: %s\n", desktopPath)
	return nil
}

func countTotalFiles(paths []string) (int, error) {
	var total int
	for _, path := range paths {
		err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				total++
			}
			return nil
		})
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func UpdateProgressBar(current, total int) {
	printLoadingBar(current, total)
}

func addFileToZip(zipWriter *zip.Writer, filePath string, callback ProgressCallback, filesProcessed *int, totalFiles int) error {
	info, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(filePath)
	}

	// Function to handle each file or directory
	fileWalkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}

		relPath := path
		if baseDir != "" {
			relPath, err = filepath.Rel(filePath, path)
			if err != nil {
				return err
			}
			relPath = filepath.Join(baseDir, relPath)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		*filesProcessed++
		if callback != nil {
			callback(*filesProcessed, totalFiles)
		}
		return nil
	}

	if info.IsDir() {
		return filepath.Walk(filePath, fileWalkFunc)
	} else {
		return fileWalkFunc(filePath, info, nil)
	}
}

func printLoadingBar(current, total int) {
	const barLength = 30
	progress := float64(current) / float64(total)
	filledLength := int(progress * float64(barLength))

	fmt.Printf("\rZipping mods: [")
	for i := 0; i < filledLength; i++ {
		fmt.Print("=")
	}
	for i := filledLength; i < barLength; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %d%%", int(progress*100))
}
