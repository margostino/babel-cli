package config

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/mitchellh/go-homedir"
	"os"
)

const (
	BabelHome         = ".babel"
	AssetFileName     = "asset.babel"
	MaxSelectorLength = 30
	BabelDatabase     = "db"
)

var AssetPath = GetAssetPath()
var BabelHomePath = GetBabelHomePath()
var BabelDataPath = GetBabelDataPath()

func InitHome() {
	if _, err := os.Stat(BabelHomePath); os.IsNotExist(err) {
		err = os.Mkdir(BabelHomePath, 0755)
		common.Check(err, "error creating .babel folder")
	}
}

func GetAssetPath() string {
	return fmt.Sprintf("%s/%s", BabelDataPath, AssetFileName)
}

func GetAssetPathById(id string) string {
	fileName := fmt.Sprintf("%s.babel", id)
	return fmt.Sprintf("%s/%s", BabelDataPath, fileName)
}

func GetBabelHomePath() string {
	home, err := homedir.Dir()
	common.Check(err, "error getting home directory")
	babelFolder := fmt.Sprintf("%s/%s", home, BabelHome)
	return babelFolder
}

func GetBabelDataPath() string {
	babelFolder := fmt.Sprintf("%s/%s", BabelHomePath, BabelDatabase)
	return babelFolder
}
