package cmd

import (
	"github.com/margostino/babel-cli/internal/db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete assets from database",
	Long:  `Delete assets from database by ID or selection`,
	Run: func(cmd *cobra.Command, args []string) {
		id := extractParam(args, 0)
		deleteAsset(*id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteAsset(id string) {
	db.Delete(id)
	//asset := getAssetByIdOrSelection(id)
	//asset.Content = editor.Open(id, asset.Content)
	//data.Delete(asset.Id)
}
