package db

import "github.com/weaviate/weaviate-go-client/v4/weaviate"

func Init(dbClient *weaviate.Client, repositoryPath string) error {
	err := deleteSchema(dbClient)
	if err != nil {
		return err
	}

	err = createSchema(dbClient)
	if err != nil {
		return err
	}

	metadata, err := getMetadata(repositoryPath)
	if err != nil {
		return err
	}

	err = insertData(dbClient, metadata)

	return err
}
