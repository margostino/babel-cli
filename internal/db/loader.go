package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/margostino/babel-cli/internal/common"
)

func getAssetsWithMetadata(repositoryPath string) ([]*Asset, error) {
	var jsonFiles []string
	var allAssets []*Asset
	var wg sync.WaitGroup
	var mu sync.Mutex

	metadataPath := filepath.Join(repositoryPath, "z-metadata")

	err := filepath.Walk(metadataPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" && !strings.HasSuffix(path, "index.json") {
			jsonFiles = append(jsonFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	results := make(chan *Asset)
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
				fmt.Printf("Corrupted file %s: %s\n", filePath, err)
				errors <- err
				return
			}
			assetPath := common.NewString(filePath).ExtractPrefixUntil("/z-metadata/")
			assetContent, err := os.ReadFile(fmt.Sprintf("%s/%s", assetPath, assetMetadata.Path))
			if err != nil {
				errors <- err
				return
			}

			asset := &Asset{
				Content:  string(assetContent),
				Metadata: *assetMetadata,
			}

			results <- asset
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
				allAssets = append(allAssets, result)
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

	return allAssets, nil
}
