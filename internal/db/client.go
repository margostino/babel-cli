package db

import (
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func NewDBClient(openAiApiKey string, port int) *weaviate.Client {
	cfg := weaviate.Config{
		Host:   fmt.Sprintf("%s:%d", "localhost", port),
		Scheme: "http",
		Headers: map[string]string{
			"X-OpenAI-Api-Key": openAiApiKey,
		},
	}
	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}
