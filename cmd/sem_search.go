package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var limit int
var query string
var includeSemContent bool

var semSearchCmd = &cobra.Command{
	Use:   "sem-search",
	Short: "Semantic Search",
	Long:  `Semantic Search`,
	Run: func(cmd *cobra.Command, args []string) {
		// query := extractParam(args, 0)
		if query == "" {
			fmt.Println("Please provide a query!")
			return
		}
		openAiApiKey := viper.GetString("openai.apiKey")
		dbPort := viper.GetInt("db.port")
		dbClient := db.NewDBClient(openAiApiKey, dbPort)
		results, errs := db.SemSearch(dbClient, query, limit)
		if errs != nil {
			jsonErrors, err := json.Marshal(errs)
			if err != nil {
				panic(err)
			}
			panic(string(jsonErrors))
		}
		for _, result := range results {
			message := fmt.Sprintf("Category: %s\nPath: %s\nSummary: %s\n", result.Category, result.Path, result.Summary)
			if includeSemContent {
				message += fmt.Sprintf("\nContent: %s", result.Content)
			}
			fmt.Println(message)
		}
	},
}

func init() {
	semSearchCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 1, "limit for the search results")
	semSearchCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "query for the semantic search")
	semSearchCmd.PersistentFlags().BoolVarP(&includeSemContent, "with-content", "w", false, "display full content in results")
	rootCmd.AddCommand(semSearchCmd)
}
