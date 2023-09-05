package cmd

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/data"
	"github.com/margostino/babel-cli/pkg/prompt"
	"os"
)

func extractParam(args []string, pos int) *string {
	var param *string
	if len(args) > pos {
		param = &args[pos]
	}
	return param
}

func getAssetByIdOrSelection(id *string) *data.Asset {
	var asset *data.Asset
	assets := data.GetAssetsBy(id)

	if len(assets) == 0 {
		fmt.Println(prompt.Red, "No assets found")
		return nil
	}

	if id != nil {
		asset = assets[0]
	} else {

		items := []string{"Quit"}
		items = append(items, prompt.AssetsToItems(assets)...)

		selector := prompt.Prompt{
			"",
		}

		choice := prompt.GetSelect(selector, items)
		if choice == 0 {
			os.Exit(0)
		}

		item := items[choice]
		prefix := common.NewString(item).Split(":").Get(0)
		selectedId := common.NewString(prefix).ReplaceAll("(", "").ReplaceAll(")", "").Value()
		idAsInt := data.ToInt(&selectedId)
		asset = data.GetBy(idAsInt)
	}
	return asset
}