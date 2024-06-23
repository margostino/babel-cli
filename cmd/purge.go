package cmd

import (
	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove all assets",
	Long:  `Remove all assets in database`,
	Run: func(cmd *cobra.Command, args []string) {
		db.DeleteAssets()
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
}
