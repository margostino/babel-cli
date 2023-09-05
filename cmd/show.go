package cmd

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/prompt"
	"github.com/spf13/cobra"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "show",
	Short: "See a list of all assets",
	Long:  `See a list of all notes`,
	Run: func(cmd *cobra.Command, args []string) {
		id := extractParam(args, 0)
		show(id)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func show(id *string) {
	assets := data.GetAssetsBy(id)

	if len(assets) == 0 {
		fmt.Println(prompt.Red, "No assets found")
		return
	}

	items := prompt.AssetsToItems(assets)

	for _, item := range items {
		prefix := common.NewString(item).Split(" ").Get(0)
		summary := strings.ReplaceAll(item, prefix, "")
		fmt.Println(prompt.Yellow, prefix, prompt.Cyan, summary)
	}
}
