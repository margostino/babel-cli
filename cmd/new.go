package cmd

import (
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/editor"
	"github.com/margostino/babel-cli/pkg/prompt"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new asset",
	Long:  `Creates a new Babel Asset (quick note, idea, knowledge, resource, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		id := concatAllParams(args)
		createNewAsset(id)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createNewAsset(id *string) {
	content := editor.Open(id, "")
	if len(content) == 0 {
		println(prompt.Red, "No content provided")
		return
	}
	data.Insert(content)
}
