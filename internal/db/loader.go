package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func getMetadata(repositoryPath string) ([]*Metadata, error) {
	var jsonFiles []string
	var allMetadata []*Metadata
	var wg sync.WaitGroup
	var mu sync.Mutex

	metadataPath := filepath.Join(repositoryPath, "metadata")

	err := filepath.Walk(metadataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			jsonFiles = append(jsonFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	results := make(chan *Metadata)
	errors := make(chan error)

	for _, jsonFile := range jsonFiles {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				errors <- err
				return
			}
			fmt.Printf("Reading file %s\n", filePath)
			var assetMetadata *Metadata
			err = json.Unmarshal(fileContent, &assetMetadata)
			if err != nil {
				fmt.Printf("Corrupted file %s\n", filePath)
				errors <- err
				return
			}
			results <- assetMetadata
		}(jsonFile)

	}

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	var firstError error
	done := false
	for !done {
		select {
		case result, ok := <-results:
			if !ok {
				results = nil
			} else {
				mu.Lock()
				allMetadata = append(allMetadata, result)
				mu.Unlock()
			}
		case err := <-errors:
			if firstError == nil {
				firstError = err
			}
		}
		if results == nil && len(errors) == 0 {
			done = true
		}
	}

	if firstError != nil {
		return nil, firstError
	}

	return allMetadata, nil
}
