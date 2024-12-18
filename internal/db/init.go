package db

import (
	"github.com/margostino/babel-cli/internal/common"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func Init(dbClient *weaviate.Client, repositoryPath string) {
	err := deleteSchema(dbClient)
	common.CheckPanic(err, "Error deleting schema")

	err = createSchema(dbClient)
	common.CheckPanic(err, "Error creating schema")

	assets, err := getAssetsWithMetadata(repositoryPath)
	common.CheckPanic(err, "Error getting assets")

	err = insertData(dbClient, assets)
	common.CheckPanic(err, "Error inserting data in new schema")
}
