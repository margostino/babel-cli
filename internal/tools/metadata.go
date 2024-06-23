package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/margostino/babel-cli/internal/common"
	"github.com/margostino/babel-cli/internal/openai"
	"github.com/margostino/babel-cli/internal/utils"
	//
)

func shouldSkipPath(path string, skipNamesMap map[string]struct{}) bool {
	for _, part := range strings.Split(path, string(filepath.Separator)) {
		if _, found := skipNamesMap[part]; found {
			return true
		}
	}
	return false
}

func getRelativePath(absolutePath string) (string, error) {
	const pattern = ".babel/db"
	absolutePath = filepath.Clean(absolutePath)
	index := strings.Index(absolutePath, pattern)
	if index == -1 {
		return "", fmt.Errorf("pattern %s not found in path %s", pattern, absolutePath)
	}
	relativePath := absolutePath[index+len(pattern)+1:] // +1 to remove the trailing separator
	return relativePath, nil
}

func writePrettyJSONToFile(jsonString, filePath string) {
	// Unmarshal the JSON string into a map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	common.Check(err, "Failed to unmarshal JSON")

	// Marshal the map into a pretty-formatted JSON string
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	common.Check(err, "Failed to marshal JSON")

	dir := filepath.Dir(filePath)
	err = os.MkdirAll(dir, 0755)
	common.Check(err, "Failed to create directory")

	err = os.WriteFile(fmt.Sprintf("%s.json", filePath), prettyJSON, 0644)
	common.Check(err, "Failed to write JSON to file")
}

func processFile(path string, root string, skipNamesMap map[string]struct{}, openAiAPIKey string, wg *sync.WaitGroup) {
	defer wg.Done()

	if shouldSkipPath(path, skipNamesMap) {
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		log.Printf("Error stating file %s: %v\n", path, err)
		return
	}
	if info.IsDir() {
		return
	}

	relativePath, err := filepath.Rel(root, path)
	common.Check(err, "Failed to get relative path")
	metadataFilePath := filepath.Join(root, "metadata", relativePath)
	metadataDir := filepath.Dir(metadataFilePath)

	// Ensure the metadata directory exists
	if _, err := os.Stat(metadataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(metadataDir, os.ModePerm); err != nil {
			log.Printf("Error creating directory %s: %v\n", metadataDir, err)
			return
		}
	}

	content, err := os.ReadFile(path)
	common.Check(err, "Failed to read file content")

	relativeFilePath, err := getRelativePath(path)
	common.Check(err, "Failed to get relative path")

	metadata, err := openai.GetChatCompletionForMetadata(openAiAPIKey, relativeFilePath, string(content))
	if err != nil {
		log.Printf("Error getting metadata for file %s: %v\n", path, err)
		return
	}

	if _, err := os.Stat(metadataFilePath); os.IsNotExist(err) {
		log.Printf("Created metadata for %s\n", path)
	} else {
		log.Printf("Updated Metadata for %s\n", path)
	}

	writePrettyJSONToFile(metadata, metadataFilePath)
}

func walkAndEnrichMetadata(root string, skipNamesMap map[string]struct{}, openAiAPIKey string) error {
	var wg sync.WaitGroup
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		wg.Add(1)
		go processFile(path, root, skipNamesMap, openAiAPIKey, &wg)
		return nil
	})

	wg.Wait()
	return err

	// metadataPath := filepath.Join(root, "metadata")
	// return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if info.IsDir() {
	// 		if _, found := skipNamesMap[info.Name()]; found {
	// 			return filepath.SkipDir
	// 		}
	// 	} else {
	// 		if _, found := skipNamesMap[info.Name()]; !found {
	// 			relativePath, err := filepath.Rel(root, path)
	// 			common.Check(err, "Failed to get relative path")
	// 			metadataFilePath := filepath.Join(metadataPath, relativePath)
	// 			metadataDir := filepath.Dir(metadataFilePath)

	// 			// Ensure the metadata directory exists
	// 			if _, err := os.Stat(metadataDir); os.IsNotExist(err) {
	// 				if err := os.MkdirAll(metadataDir, os.ModePerm); err != nil {
	// 					return err
	// 				}
	// 			}

	// 			content, err := os.ReadFile(path)
	// 			common.Check(err, "Failed to read file content")

	// 			relativeFilePath, err := getRelativePath(path)
	// 			common.Check(err, "Failed to get relative path")

	// 			metadata, err := openai.GetChatCompletionForMetadata(openAiAPIKey, relativeFilePath, string(content))
	// 			common.Check(err, "Failed to get metadata")

	// 			if _, err := os.Stat(metadataFilePath); os.IsNotExist(err) {
	// 				log.Printf("Created metadata for %s\n", path)
	// 			} else {
	// 				log.Printf("Updated Metadata for %s\n", path)
	// 			}

	// 			writePrettyJSONToFile(metadata, metadataFilePath)
	// 		}
	// 	}
	// 	return nil
	// })
}

func EnrichMetadata(repositoryPath string, openAiAPIKey string) {
	log.Println(fmt.Sprintf("Running MetadataEnrichment tool in bulk..."))
	skipNames := []string{".git", "metadata", "0-description", "0-babel", "metadata_index"}
	skipNamesMap := utils.ListToMap(skipNames)

	if err := walkAndEnrichMetadata(repositoryPath, skipNamesMap, openAiAPIKey); err != nil {
		log.Printf("Error walking through the path: %v\n", err)
	}
}
