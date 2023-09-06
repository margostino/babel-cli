package cmd

import (
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop table from database",
	Long:  `Drop table from database`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DropTable()
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
