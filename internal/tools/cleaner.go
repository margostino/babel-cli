package tools

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/margostino/babel-cli/internal/utils"

	"github.com/margostino/babel-cli/internal/common"
)

func walkAndNormalizeFiles(root string, skipNamesMap map[string]struct{}) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if _, found := skipNamesMap[info.Name()]; found {
				return filepath.SkipDir
			}
		} else {
			if _, found := skipNamesMap[info.Name()]; !found {
				normalizedFileName := normalizeFileName(info.Name())
				if normalizedFileName != info.Name() {
					newPath := filepath.Join(filepath.Dir(path), normalizedFileName)
					if err := os.Rename(path, newPath); err != nil {
						log.Fatalf("Error renaming file: %v\n", err)
					} else {
						log.Printf("Renamed file: %s to %s\n", path, newPath)
					}
				} else {
					log.Printf("File already normalized: %s\n", path)
				}
			}
		}
		return nil
	})
}

func normalizeFileName(name string) string {
	ext := filepath.Ext(name)
	base := strings.TrimSuffix(name, ext)
	normalized := common.NewString(base).
		ToLower().
		TrimSpace().
		ReplaceAll(" ", "_").
		ReplaceAll(".", "_").
		Value()
	return normalized
}

func CleanAssets(path string) {
	log.Println("Running AssetsCleaner tool...")
	skipNames := []string{".git", "0-description", "0-babel", "metadata_index"}
	skipNamesMap := utils.ListToMap(skipNames)

	if err := walkAndNormalizeFiles(path, skipNamesMap); err != nil {
		log.Printf("Error walking through the path: %v\n", err)
	}
}
