package mod

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

	err = Unzip(zipPath, modPath)
	if err != nil {
		return "", fmt.Errorf("error unzipping mod: %w", err)
	}

	if err := MovePlugins(modPath); err != nil {
		return "", fmt.Errorf("error moving plugins: %w", err)
	}

	return newModDirName, nil
}

// MovePlugins searches for a 'plugins' directory inside the modPath or
// inside 'modPath/BepInEx'. If found, it moves all contents from the 'plugins' directory
// to the modPath.
func MovePlugins(modPath string) error {
	// Potential plugins directories
	pluginsDir := filepath.Join(modPath, "plugins")
	bepInExDir := filepath.Join(modPath, "BepInEx")
	bepInExPluginsDir := filepath.Join(bepInExDir, "plugins")

	// Attempt to move contents from both possible plugins directories
	for _, pd := range []string{pluginsDir, bepInExPluginsDir} {
		err := moveContentsFromPluginDir(pd, modPath)
		if err != nil {
			fmt.Printf("Failed to move contents from: %s to: %s, error: %v\n", pd, modPath, err)
			return err
		}
	}

	// Check if BepInEx directory is now empty and remove it if so
	return removeDirIfEmpty(bepInExDir)
}

// Move contents from one plugin directory to the modPath
func moveContentsFromPluginDir(pluginsDir, modPath string) error {
	contents, err := ioutil.ReadDir(pluginsDir)
	if err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist, nothing to do
			fmt.Printf("Directory does not exist, skipping: %s\n", pluginsDir)
			return nil
		}
		// Some other error occurred while accessing the directory
		return err
	}

	for _, content := range contents {
		srcPath := filepath.Join(pluginsDir, content.Name())
		dstPath := filepath.Join(modPath, content.Name())

		// If the destination is a directory, merge contents, otherwise move
		if info, err := os.Stat(dstPath); err == nil && info.IsDir() {
			if err := mergeDirs(srcPath, dstPath); err != nil {
				fmt.Printf("Failed to merge directories from: %s to: %s, error: %v\n", srcPath, dstPath, err)
				return err
			}
		} else {
			if err := os.Rename(srcPath, dstPath); err != nil {
				fmt.Printf("Failed to move from: %s to: %s, error: %v\n", srcPath, dstPath, err)
				return err
			}
		}
	}

	// Remove the plugins directory if it's now empty
	return removeDirIfEmpty(pluginsDir)
}

// Removes a directory if it is empty
func removeDirIfEmpty(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist, nothing to do
			fmt.Printf("Directory does not exist, skipping removal: %s\n", dirPath)
			return nil
		}
		// Some other error occurred while opening the directory
		return err
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1) // Attempt to read at least one entry
	if err == io.EOF {
		// Directory is empty
		dir.Close() // Close before removing
		if err := os.Remove(dirPath); err != nil {
			fmt.Printf("Failed to remove directory: %s, error: %v\n", dirPath, err)
			return err
		}
		fmt.Printf("Removed empty directory: %s\n", dirPath)
	} else if err != nil {
		// Some other error occurred while reading the directory
		return err
	}

	// Directory is not empty or could not be read
	return nil
}

func mergeDirs(srcDir, dstDir string) error {
	// Check if the source directory exists
	srcInfo, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		// Source directory doesn't exist, nothing to merge
		return nil
	} else if err != nil {
		// Some other error occurred while accessing the source path
		return err
	}

	// Ensure the source is actually a directory
	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory: %s", srcDir)
	}

	// Ensure the destination directory exists
	dstInfo, err := os.Stat(dstDir)
	if os.IsNotExist(err) {
		// Destination directory does not exist, create it
		if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create destination directory: %s, error: %w", dstDir, err)
		}
	} else if err != nil {
		// Some other error occurred while accessing the destination path
		return err
	} else if !dstInfo.IsDir() {
		// Destination exists but is not a directory, which is a problem
		return fmt.Errorf("destination exists but is not a directory: %s", dstDir)
	}

	// Now we can safely move the contents
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			// Recursively merge subdirectories
			if err := mergeDirs(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Move files
			if err := os.Rename(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	// Remove the now-empty source directory
	return os.Remove(srcDir)
}

// removeIfEmpty removes the specified directory if it is empty.
func removeIfEmpty(dirPath string) error {
	dirEmpty, err := isDirEmpty(dirPath)
	if err != nil {
		return err
	}
	if dirEmpty {
		if err := os.Remove(dirPath); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

// isDirEmpty checks if a directory is empty.
func isDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
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
