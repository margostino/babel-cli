package cmd

import (
	"github.com/margostino/babel-cli/internal/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Enrich assets metadata",
	Long:  "Enrich assets metadata in database based on output from LLM",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := viper.GetString("repository.path")
		openAiAPIKey := viper.GetString("openai.apikey")
		tools.EnrichMetadata(repositoryPath, openAiAPIKey)
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
}
