package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

func insertData(client *weaviate.Client, assets []*Asset) error {
	var wg sync.WaitGroup
	errors := make(chan error)

	ready, err := client.Misc().ReadyChecker().Do(context.Background())
	if err != nil {
		fmt.Println("Weaviate is NOT ready yet")
	} else {
		fmt.Printf("Weaviate is ready: %t\n", ready)
	}

	for _, item := range assets {
		wg.Add(1)
		go func(item *Asset) {
			defer wg.Done()
			resp, err := client.Batch().ObjectsBatcher().
				WithObjects(
					&models.Object{
						Class: "babel",
						Properties: map[string]interface{}{
							"content":       item.Content,
							"category":      item.Metadata.Category,
							"highlights":    item.Metadata.Highlights,
							"keywords":      item.Metadata.Keywords,
							"path":          item.Metadata.Path,
							"references":    item.Metadata.References,
							"related_links": item.Metadata.RelatedLinks,
							"summary":       item.Metadata.Summary,
							"tags":          item.Metadata.Tags,
						},
					},
				).
				Do(context.Background())
			if err != nil {
				errors <- err
				return
			}
			fmt.Printf("Object '%s' inserted successfully\n", resp[0].ID)
		}(item)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	var firstError error
	done := false
	for !done {
		select {
		case err := <-errors:
			if firstError == nil {
				firstError = err
				done = true
			}
		}
	}

	if firstError != nil {
		return firstError
	}

	fmt.Println("Class 'babel' initialized successfully")
	return nil
}
