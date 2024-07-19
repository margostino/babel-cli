package db

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/entities/schema"
)

func deleteSchema(client *weaviate.Client) error {
	err := client.Schema().AllDeleter().Do(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Schema deleted successfully")
	return nil
}

func createSchema(client *weaviate.Client) error {
	err := client.Schema().ClassCreator().WithClass(&models.Class{
		Class:      "babel",
		Vectorizer: "text2vec-openai",
		VectorIndexConfig: &map[string]interface{}{
			"distance": "dot",
		},
		ModuleConfig: &map[string]interface{}{
			"generative-openai": map[string]interface{}{
				"model": "gpt-4-1106-preview",
			},
		},
		Properties: []*models.Property{
			{
				Name:     "category",
				DataType: schema.DataTypeText.PropString(),
				// DataType: []string{"Category"},
			},
			{
				Name:     "highlights",
				DataType: schema.DataTypeStringArray.PropString(),
			},
			{
				Name:     "keywords",
				DataType: schema.DataTypeStringArray.PropString(),
			},
			{
				Name:     "path",
				DataType: schema.DataTypeText.PropString(),
			},
			{
				Name:     "references",
				DataType: schema.DataTypeStringArray.PropString(),
			},
			{
				Name:     "related_links",
				DataType: schema.DataTypeStringArray.PropString(),
			},
			{
				Name:     "summary",
				DataType: schema.DataTypeText.PropString(),
			},
			{
				Name:     "tags",
				DataType: schema.DataTypeStringArray.PropString(),
			},
		},
	}).Do(context.Background())

	if err != nil {
		return err
	}

	fmt.Println("Class 'babel' created successfully")
	return nil
}
