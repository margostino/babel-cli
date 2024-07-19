package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genLimit int
var genPrompt string

var genSearchCmd = &cobra.Command{
	Use:   "gen-search",
	Short: "Generative Search",
	Long:  `Generative Search`,
	Run: func(cmd *cobra.Command, args []string) {
		query := extractParam(args, 0)
		openAiApiKey := viper.GetString("openai.apiKey")
		dbClient := db.NewDBClient(openAiApiKey)
		results, errs := db.GenSearch(dbClient, *query, genPrompt, genLimit)
		if errs != nil {
			jsonErrors, err := json.Marshal(errs)
			if err != nil {
				panic(err)
			}
			panic(string(jsonErrors))
		}
		for _, result := range results {
			fmt.Printf("Single Result: %s\n", result.SingleResult)
		}
	},
}

func init() {
	genSearchCmd.PersistentFlags().IntVarP(&genLimit, "limit", "l", 1, "limit for the search results")
	genSearchCmd.PersistentFlags().StringVarP(&genPrompt, "prompt", "p", "present the search results in a creative way", "prompt for the generative search results")
	rootCmd.AddCommand(genSearchCmd)
}
