package cmd

import (
	"log"

	"github.com/margostino/babel-cli/internal/editor"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "See a list of all notes you've added",
	Long:  `See a list of all notes you've added`,
	Run: func(cmd *cobra.Command, args []string) {
		id := extractParam(args, 0)
		edit(id)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func edit(id *string) {
	asset := getAssetByIdOrSelection(id)
	asset.Content = editor.OpenBy(id, asset.Content)
	//data.Update(asset.Id, asset.Content)
	log.Printf("asset [%d] updated successfully", asset.Id)
}
