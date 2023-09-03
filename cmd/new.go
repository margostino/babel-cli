package cmd

import (
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/editor"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new asset",
	Long:  `Creates a new Babel Asset (quick note, idea, knowledge, resource, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewAsset()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createNewAsset() {
	content := editor.Open("")
	data.InsertNote(content)
}
