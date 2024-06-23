package cmd

import (
	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new database",
	Long:  `Initialise a new database`,
	Run: func(cmd *cobra.Command, args []string) {
		db.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
