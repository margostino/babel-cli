package cmd

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/prompt"
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
	input := prompt.Prompt{
		"",
	}
	fmt.Printf("Start writing...\n\n")
	content := prompt.GetInput(input)
	println(content)
	data.InsertNote(content)
}
