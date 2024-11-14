package cmd

import (
	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new vector collection",
	Long:  `Initialise a new vector collection`,
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := viper.GetString("repository.path")
		openAiApiKey := viper.GetString("openai.apiKey")
		dbPort := viper.GetInt("db.port")
		dbClient := db.NewDBClient(openAiApiKey, dbPort)
		err := db.Init(dbClient, repositoryPath)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
