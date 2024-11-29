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
)

var metadataCollection sync.Map

type IndexData struct {
	Summary    string
	Highlights []string
}

type Index struct {
	IndexData []IndexData
}

func writeIndexFile(root string) {
	indexData := make(map[string]map[string]interface{})
	indexFilePath := filepath.Join(root, "z-metadata", "index.json")
	indexFileContent, err := os.ReadFile(indexFilePath)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Failed to read index file: %v\n", err)
		return
	}

	if len(indexFileContent) > 0 {
		err = json.Unmarshal(indexFileContent, &indexData)
		if err != nil {
			log.Printf("Failed to unmarshal index file: %v\n", err)
			return
		}
	}

	metadataCollection.Range(func(key, value interface{}) bool {
		relativePath := key.(string)
		metadata := value.(map[string]interface{})

		highlights, ok := metadata["highlights"].([]interface{})
		if !ok {
			log.Printf("Invalid metadata for file %s: highlights missing or wrong type\n", relativePath)
			return true
		}

		summary, ok := metadata["summary"].(string)
		if !ok {
			log.Printf("Invalid metadata for file %s: summary missing or wrong type\n", relativePath)
			return true
		}

		indexData[relativePath] = map[string]interface{}{
			"highlights": highlights,
			"summary":    summary,
		}

		return true
	})

	if len(indexData) > 0 {
		indexFile, err := os.Create(indexFilePath)
		common.Check(err, "Failed to create index file")
		defer indexFile.Close()

		log.Printf("Writing index data to file %s\n", indexFilePath)

		indexJSON, err := json.MarshalIndent(indexData, "", "  ")
		if err != nil {
			log.Printf("Failed to marshal index data to JSON: %v\n", err)
			return
		}

		err = os.WriteFile(indexFilePath, indexJSON, 0644)
		if err != nil {
			log.Printf("Failed to write JSON to index file %s: %v\n", indexFilePath, err)
		}
	} else {
		log.Printf("No data to write to index file %s\n", indexFilePath)
	}
}

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
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	common.Check(err, "Failed to unmarshal JSON")

	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	common.Check(err, "Failed to marshal JSON")

	dir := filepath.Dir(filePath)
	err = os.MkdirAll(dir, 0755)
	common.Check(err, "Failed to create directory")

	err = os.WriteFile(filePath, prettyJSON, 0644)
	common.Check(err, "Failed to write JSON to file")
}

func initMetadata(root string, relativeFilePath string, skipNamesMap map[string]struct{}, openAiAPIKey string, wg *sync.WaitGroup) {
	defer wg.Done()
	absoluteFilePath := fmt.Sprintf("%s/%s", root, relativeFilePath)
	absoluteMetadataFilePath := fmt.Sprintf("%s/metadata/%s.json", root, relativeFilePath)
	content, err := os.ReadFile(absoluteFilePath)
	common.Check(err, "Failed to read file content")

	log.Printf("Getting metadata for file %s\n", relativeFilePath)

	metadata, err := openai.GetChatCompletionForMetadata(openAiAPIKey, relativeFilePath, string(content))

	if err != nil {
		log.Printf("Error getting metadata for file %s: %v\n", relativeFilePath, err)
		return
	}

	var metadataMap map[string]interface{}
	err = json.Unmarshal([]byte(metadata), &metadataMap)
	common.Check(err, "Failed to unmarshal metadata JSON")
	metadataCollection.Store(relativeFilePath, metadataMap)
	writePrettyJSONToFile(metadata, absoluteMetadataFilePath)
	log.Printf("Metadata written for file %s\n", absoluteMetadataFilePath)
}

func syncMetadata(root string, relativeFilePath string, skipNamesMap map[string]struct{}, openAiAPIKey string, wg *sync.WaitGroup) {
	absoluteMetadataFilePath := fmt.Sprintf("%s/metadata/%s.json", root, relativeFilePath)
	if _, err := os.Stat(absoluteMetadataFilePath); os.IsNotExist(err) {
		initMetadata(root, relativeFilePath, skipNamesMap, openAiAPIKey, wg)
	} else {
		wg.Done()
	}
}

func processTaskResult(taskResultChan <-chan openai.LLMTaskResult, root string) {
	for task := range taskResultChan {
		common.Check(task.Error, "Failed to get content from task result")
		content := task.Content
		absoluteMetadataFilePath := fmt.Sprintf("%s/metadata/%s.json", root, task.RelativeFilePath)

		var metadataMap map[string]interface{}
		err := json.Unmarshal([]byte(content), &metadataMap)
		common.Check(err, "Failed to unmarshal metadata JSON")
		metadataCollection.Store(task.RelativeFilePath, metadataMap)
		writePrettyJSONToFile(content, absoluteMetadataFilePath)
		log.Printf("Metadata written for file %s\n", absoluteMetadataFilePath)
	}
}

func walkAndEnrichMetadata(root string, skipNamesMap map[string]struct{}, openAiAPIKey string, fn func(root string, relativeFilePath string, skipNamesMap map[string]struct{}, openAiAPIKey string, wg *sync.WaitGroup)) error {
	var wg sync.WaitGroup

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if shouldSkipPath(path, skipNamesMap) {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		relativePath, err := filepath.Rel(root, path)
		common.Check(err, "Failed to get relative path")
		metadataFilePath := filepath.Join(root, "z-metadata", relativePath)
		metadataDir := filepath.Dir(metadataFilePath)

		if _, err := os.Stat(metadataDir); os.IsNotExist(err) {
			if err := os.MkdirAll(metadataDir, os.ModePerm); err != nil {
				log.Printf("Error creating directory %s: %v\n", metadataDir, err)
				return nil
			}
		}

		wg.Add(1)
		relativeFilePath := common.NewString(path).ReplaceAll(fmt.Sprintf("%s/", root), "").Value()

		go fn(root, relativeFilePath, skipNamesMap, openAiAPIKey, &wg)

		return nil
	})

	wg.Wait()
	writeIndexFile(root)

	log.Println("Done!")

	return err
}

func EnrichMetadata(repositoryPath string, openAiAPIKey string) {
	log.Println(fmt.Sprintf("Initializing metadata..."))
	skipNames := []string{".git", "z-metadata", "0-description", "0-babel", "metadata_index"}
	skipNamesMap := utils.ListToMap(skipNames)

	if err := walkAndEnrichMetadata(repositoryPath, skipNamesMap, openAiAPIKey, initMetadata); err != nil {
		log.Printf("Error walking through the path: %v\n", err)
	}
}

func SyncMetadata(repositoryPath string, openAiAPIKey string) {
	log.Println(fmt.Sprintf("Syncing metadata..."))
	skipNames := []string{".git", "z-metadata", "0-description", "0-babel", "metadata_index"}
	skipNamesMap := utils.ListToMap(skipNames)

	if err := walkAndEnrichMetadata(repositoryPath, skipNamesMap, openAiAPIKey, syncMetadata); err != nil {
		log.Printf("Error walking through the path: %v\n", err)
	}
}
