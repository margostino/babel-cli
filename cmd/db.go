package cmd

import (
	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new vector collection and schema",
	Long:  `Initialise a new vector collection and schema`,
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := viper.GetString("repository.path")
		openAiApiKey := viper.GetString("openai.apiKey")
		dbPort := viper.GetInt("db.port")
		dbClient := db.NewDBClient(openAiApiKey, dbPort)
		db.Init(dbClient, repositoryPath)
	},
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Perform actions on database",
	Long:  "Perform actions on database such as initialise schema",
}

func init() {
	dbCmd.AddCommand(dbInitCmd)
	rootCmd.AddCommand(dbCmd)
}
