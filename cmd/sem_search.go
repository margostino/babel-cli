package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var limit int

var semSearchCmd = &cobra.Command{
	Use:   "sem-search",
	Short: "Semantic Search",
	Long:  `Semantic Search`,
	Run: func(cmd *cobra.Command, args []string) {
		query := extractParam(args, 0)
		openAiApiKey := viper.GetString("openai.apiKey")
		dbClient := db.NewDBClient(openAiApiKey)
		results, errs := db.SemSearch(dbClient, *query, limit)
		if errs != nil {
			jsonErrors, err := json.Marshal(errs)
			if err != nil {
				panic(err)
			}
			panic(string(jsonErrors))
		}
		for _, result := range results {
			message := fmt.Sprintf("Category: %s\nPath: %s\nSummary: %s\n", result.Category, result.Path, result.Summary)
			fmt.Println(message)
		}
	},
}

func init() {
	semSearchCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 1, "limit for the search results")
	rootCmd.AddCommand(semSearchCmd)
}
