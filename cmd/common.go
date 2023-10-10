package cmd

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/db"
	"github.com/margostino/babel-cli/pkg/prompt"
	"os"
	"strings"
)

func extractParam(args []string, pos int) *string {
	var param *string
	if len(args) > pos {
		param = &args[pos]
	}
	return param
}

func concatAllParams(args []string) *string {
	if len(args) == 0 {
		return nil
	}
	joinedArgs := strings.TrimSpace(strings.Join(args[0:], "_"))
	return &joinedArgs
}

func getAssetByIdOrSelection(id *string) *db.Asset {
	var asset *db.Asset
	assets := db.GetAssetsBy(id)

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
		idAsInt := db.ToInt(&selectedId)
		asset = db.GetBy(idAsInt)
	}
	return asset
}
