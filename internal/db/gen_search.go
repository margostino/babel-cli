package db

import (
	"context"
	"encoding/json"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
)

func GenSearch(dbClient *weaviate.Client, query string, prompt string, limit int) ([]*GenerativeSearchResult, []*models.GraphQLError) {
	nearTextConceptWithDistance := dbClient.GraphQL().
		NearTextArgBuilder().
		WithConcepts([]string{"query"}).
		WithDistance(0.95)

	generativeSearchPrompt := graphql.NewGenerativeSearch().SingleResult(prompt)

	result, err := dbClient.GraphQL().
		Get().
		WithClassName("Babel").
		WithNearText(nearTextConceptWithDistance).
		WithFields(
			graphql.Field{Name: "category"},
			graphql.Field{Name: "path"},
			graphql.Field{Name: "summary"},
		).
		WithGenerativeSearch(generativeSearchPrompt).
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

	var rawGenerativeSearchResults []*RawGenerativeSearchResult
	err = json.Unmarshal(jsonData, &rawGenerativeSearchResults)
	if err != nil {
		panic(err)
	}

	generativeSearchResults := make([]*GenerativeSearchResult, len(rawGenerativeSearchResults))
	for i, r := range rawGenerativeSearchResults {
		generativeSearchResults[i] = &r.Additional.Generate
	}

	return generativeSearchResults, nil
}
