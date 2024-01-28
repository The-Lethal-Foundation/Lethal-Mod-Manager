package mod

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/KonstantinBelenko/lethal-mod-manager/pkg/lcfs/util"
	types "github.com/KonstantinBelenko/lethal-mod-manager/pkg/types"
)

/*
	modmanagement.go

	Contains functions for mod operations such as installing, updating, checking mods and more.
*/

func ZipMods(profile string, progressCallback types.ProgressCallback) error {
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

func addFileToZip(zipWriter *zip.Writer, filePath string, callback types.ProgressCallback, filesProcessed *int, totalFiles int) error {
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
			callback(*filesProcessed, totalFiles, "Zipping mods")
		}
		return nil
	}

	if info.IsDir() {
		return filepath.Walk(filePath, fileWalkFunc)
	} else {
		return fileWalkFunc(filePath, info, nil)
	}
}

// UnzipMod unzips the mod file into the specified profile folder and removes the zip file.
// Returns mod name and error
func UnzipMod(profileName, zipPath string, modName types.ModName) (string, error) {
	// Open the zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return "", fmt.Errorf("error opening zip file: %w", err)
	}
	defer r.Close()

	modsPath, err := util.GetModsPath(profileName)
	fmt.Println(modsPath)
	if err != nil {
		return "", fmt.Errorf("error getting profile path: %w", err)
	}

	// Create a target directory for the mod
	newModDirName := fmt.Sprintf("%s-%s-%s", modName.Author, modName.Name, modName.Version)
	modPath := filepath.Join(modsPath, newModDirName)
	if err := os.MkdirAll(modPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creating mod directory: %w", err)
	}

	modsPath = modPath

	// Iterate through the files in the zip archive
	for _, f := range r.File {
		// Check and remove previous version if exists
		targetPath := filepath.Join(modsPath, f.Name)
		if err := os.RemoveAll(targetPath); err != nil {
			return "", fmt.Errorf("error removing previous version of the mod: %w", err)
		}

		// Create necessary directories
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return "", fmt.Errorf("error creating directory: %w", err)
			}
			continue
		}

		// Open the file inside the zip
		fileInZip, err := f.Open()
		if err != nil {
			return "", fmt.Errorf("error opening file in zip: %w", err)
		}

		// Create the file in the target directory
		targetFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fileInZip.Close()
			return "", fmt.Errorf("error creating file: %w", err)
		}

		// Copy the contents of the file
		if _, err := io.Copy(targetFile, fileInZip); err != nil {
			fileInZip.Close()
			targetFile.Close()
			return "", fmt.Errorf("error copying file contents: %w", err)
		}

		fileInZip.Close()
		targetFile.Close()
	}

	return newModDirName, nil
}

func DeleteMod(profileName, modName string) error {
	modsPath, err := util.GetModsPath(profileName)
	if err != nil {
		return fmt.Errorf("error getting profile path: %w", err)
	}

	modPath := filepath.Join(modsPath, modName)
	if _, err := os.Stat(modPath); os.IsNotExist(err) {
		return fmt.Errorf("mod does not exist: %s", modPath)
	}

	if err := os.RemoveAll(modPath); err != nil {
		return fmt.Errorf("error deleting mod: %w", err)
	}

	return nil
}
