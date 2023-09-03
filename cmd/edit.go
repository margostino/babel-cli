package cmd

import (
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/editor"
	"github.com/margostino/babel-cli/pkg/prompt"
	"github.com/spf13/cobra"
	"os"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "See a list of all notes you've added",
	Long:  `See a list of all notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		edit()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func edit() {
	assets := data.GetAll()

	items := []string{"Quit"}
	items = append(items, prompt.AssetsToSelector(assets)...)

	selector := prompt.Prompt{
		"",
	}

	choice := prompt.GetSelect(selector, items)
	if choice == 0 {
		os.Exit(0)
	}

	item := items[choice]
	prefix := common.NewString(item).Split(":").Get(0)
	id := common.NewString(prefix).ReplaceAll("(", "").ReplaceAll(")", "").Value()

	asset := data.GetBy(&id)

	asset.Content = editor.Open(asset.Content)
	data.Update(id, asset.Content)
}
