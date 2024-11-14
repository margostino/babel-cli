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
var genQuery string

var genSearchCmd = &cobra.Command{
	Use:   "gen-search",
	Short: "Generative Search",
	Long:  `Generative Search`,
	Run: func(cmd *cobra.Command, args []string) {
		// query := extractParam(args, 0)
		if genQuery == "" {
			fmt.Println("Please provide a query!")
			return
		}
		openAiApiKey := viper.GetString("openai.apiKey")
		dbPort := viper.GetInt("db.port")
		dbClient := db.NewDBClient(openAiApiKey, dbPort)
		results, errs := db.GenSearch(dbClient, genQuery, genPrompt, genLimit)
		if errs != nil {
			jsonErrors, err := json.Marshal(errs)
			if err != nil {
				panic(err)
			}
			panic(string(jsonErrors))
		}
		for _, result := range results {
			if result.SingleResult != "" {
				fmt.Printf("%s\n", result.SingleResult)
			}
			if result.Error != "" {
				fmt.Printf("Error: %s\n", result.Error)
			}
		}
	},
}

func init() {
	genSearchCmd.PersistentFlags().IntVarP(&genLimit, "limit", "l", 1, "limit for the search results")
	genSearchCmd.PersistentFlags().StringVarP(&genPrompt, "prompt", "p", "re-write this **{summary}** in a creative way", "prompt for the generative search results")
	genSearchCmd.PersistentFlags().StringVarP(&genQuery, "query", "q", "", "query for the generative search")
	rootCmd.AddCommand(genSearchCmd)
}
