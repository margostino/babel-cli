package editor

import (
	"github.com/margostino/babel-cli/pkg/common"
	"github.com/margostino/babel-cli/pkg/config"
	"os"
	"os/exec"
	"strings"
)

func Open(id *string, content string) string {
	var fileName string
	if id == nil {
		fileName = config.AssetPath
	} else {
		fileName = config.GetAssetPathById(*id)
	}
	createFile(content)
	dataPath := config.GetBabelDataPath()
	//cmd := exec.Command("vim", TempFileName)
	cmd := exec.Command("code", "-n", "-w", "-g", fileName, "-a", dataPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	common.Check(err, "Failed to run vim")
	updatedContent := readTempFile()
	removeTempFile()
	return updatedContent
}

func createFile(content string) {
	bytes := []byte(content)
	err := os.WriteFile(config.AssetPath, bytes, 0644)
	common.Check(err, "Failed to write to file")
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
