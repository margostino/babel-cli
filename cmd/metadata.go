package cmd

import (
	"github.com/margostino/babel-cli/internal/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var metadataInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Enrich assets metadata",
	Long:  "Initialize assets metadata in database based on output from LLM",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := viper.GetString("repository.path")
		openAiAPIKey := viper.GetString("openai.apikey")
		tools.EnrichMetadata(repositoryPath, openAiAPIKey)
	},
}

var metadataSyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync assets metadata",
	Long:  "Sync assets metadata in database based on output from LLM",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryPath := viper.GetString("repository.path")
		openAiAPIKey := viper.GetString("openai.apikey")
		tools.SyncMetadata(repositoryPath, openAiAPIKey)
	},
}

var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Enrich assets metadata",
	Long:  "Enrich assets metadata in database based on output from LLM",
}

func init() {
	metadataCmd.AddCommand(metadataInitCmd)
	metadataCmd.AddCommand(metadataSyncCmd)
	rootCmd.AddCommand(metadataCmd)
}
