package cmd

import (
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var editCmd = &cobra.Command{
	Use:   "list",
	Short: "See a list of all notes you've added",
	Long:  `See a list of all notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		edit()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func edit() *data.Asset {
	//assets := data.GetAll()

	//items = append(items, "Quit")
	//for _, asset := range assets {
	//	category := data.GetCategoryAsString(asset.Category)
	//	item := fmt.Sprintf("(%d:%s) %s", asset.Id, category, asset.Content)
	//	items = append(items, item)
	//	//fmt.Println(string(prompt.Cyan), asset.Content)
	//}
	//
	//selectPrompt := prompt.Prompt{
	//	"What asset you want to edit?",
	//}
	//
	//choice := prompt.GetSelect(selectPrompt, items)
	//
	//println(choice)
	return nil
}
