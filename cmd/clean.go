package cmd

import (
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove all assets",
	Long:  `Remove all assets in database`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DeleteAssets()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
