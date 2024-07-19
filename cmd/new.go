package cmd

import (
	"fmt"

	"github.com/margostino/babel-cli/internal/editor"
	"github.com/margostino/babel-cli/internal/prompt"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new asset",
	Long:  `Creates a new Babel Asset (quick note, idea, knowledge, resource, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		// id := concatAllParams(args)
		// createNewAsset(id)
		fmt.Println("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createNewAsset(id *string) {
	content := editor.OpenBy(id, "")
	if len(content) == 0 {
		println(prompt.Red, "No content provided")
		return
	}
	//data.Insert(content)
}
