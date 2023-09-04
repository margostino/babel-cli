package config

import (
	"fmt"
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/mitchellh/go-homedir"
	"os"
)

const (
	BabelHome    = ".babel"
	TempFileName = "asset.temp"
)

var AssetTempPath = GetAssetTempPath()
var BabelHomePath = GetBabelHomePath()

func InitBabelHome() {
	if _, err := os.Stat(BabelHomePath); os.IsNotExist(err) {
		err = os.Mkdir(BabelHomePath, 0755)
		common.Check(err, "error creating .babel folder")
	}
}

func GetAssetTempPath() string {
	home, err := homedir.Dir()
	common.Check(err, "error getting home directory")
	return fmt.Sprintf("%s/%s/%s", home, BabelHome, TempFileName)
}

func GetBabelHomePath() string {
	home, err := homedir.Dir()
	common.Check(err, "error getting home directory")
	babelFolder := fmt.Sprintf("%s/%s", home, BabelHome)
	return babelFolder
}
