package cmd

import (
	"github.com/margostino/babel-cli/internal/tools"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Normalize all assets names",
	Long:  `Normalize all assets names in database`,
	Run: func(cmd *cobra.Command, args []string) {
		path := extractParam(args, 0)
		tools.CleanAssets(*path)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
