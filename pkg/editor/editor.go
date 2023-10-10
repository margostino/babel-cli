package editor

import (
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/config"
	"os"
	"os/exec"
	"strings"
)

func OpenBy(id *string, content string) string {
	var assetName string
	if id == nil {
		assetName = config.AssetPath
	} else {
		assetName = config.GetAssetPathById(*id)
	}
	if !fileExits(assetName) {
		createFile(assetName, content)
	}
	Open(&assetName)
	updatedContent := readTempFile()
	removeTempFile()
	return updatedContent
}

func Open(assetName *string) {
	var cmd *exec.Cmd
	dataPath := config.GetBabelDataPath()

	if assetName == nil {
		cmd = exec.Command("code", "-n", "-a", dataPath)
	} else {
		cmd = exec.Command("code", "-n", "-w", "-g", *assetName, "-a", dataPath)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	common.Check(err, "Failed to run vim")
}

func createFile(fileName string, content string) {
	bytes := []byte(content)
	err := os.WriteFile(fileName, bytes, 0644)
	common.Check(err, "Failed to write to file")
}

func fileExits(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func removeTempFile() {
	cmd := exec.Command("rm", config.AssetPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	common.Check(err, "Failed to run rm")
}

func readTempFile() string {
	// read text file and get text
	bytes, err := os.ReadFile(config.AssetPath)
	common.Check(err, "Failed to read file")
	content := string(bytes)
	return strings.TrimSpace(content)
}
