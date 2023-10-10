package cmd

import (
	"github.com/margostino/babel-cli/pkg/db"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop table",
	Long:  `Drop table`,
	Run: func(cmd *cobra.Command, args []string) {
		db.DropTable()
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
