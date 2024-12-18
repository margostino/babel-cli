package db

import (
	"context"
	"encoding/json"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func SemSearch(dbClient *weaviate.Client, query string, limit int) ([]*SemanticSearchResult, []*models.GraphQLError) {
	nearTextConceptWithDistance := dbClient.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{query}).
		WithDistance(0.95)

	result, err := dbClient.GraphQL().
		Get().
		WithClassName("Babel").
		WithNearText(nearTextConceptWithDistance).
		WithFields(
			graphql.Field{Name: "category"},
			graphql.Field{Name: "path"},
			graphql.Field{Name: "summary"},
			graphql.Field{Name: "content"},
		).
		WithLimit(limit).
		Do(context.Background())

	if len(result.Errors) > 0 {
		return nil, result.Errors
	}
	if err != nil {
		return nil, []*models.GraphQLError{{Message: err.Error()}}
	}

	jsonData, err := json.Marshal(result.Data["Get"].(map[string]interface{})["Babel"])
	if err != nil {
		panic(err)
	}

	var semanticSearchResults []*SemanticSearchResult
	err = json.Unmarshal(jsonData, &semanticSearchResults)
	if err != nil {
		panic(err)
	}

	return semanticSearchResults, nil
}
