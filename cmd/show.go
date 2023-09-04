package cmd

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/prompt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "show",
	Short: "See a list of all assets",
	Long:  `See a list of all notes`,
	Run: func(cmd *cobra.Command, args []string) {
		var id *string
		if len(args) > 0 {
			id = &args[0]
		}
		show(id)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func show(id *string) {
	assets := make([]*data.Asset, 0)

	if id != nil {
		asset := data.GetBy(id)
		assets = append(assets, asset)
	} else {
		assets = data.GetAll()
	}

	if len(assets) == 0 {
		fmt.Println(prompt.Red, "No assets found")
		return
	}

	items := prompt.AssetsToItems(assets)

	for _, item := range items {
		fmt.Println(prompt.Cyan, item)
	}
}
