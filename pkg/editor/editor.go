package editor

import (
	"github.com/margostino/babel-cli/pkg/common"
	"os"
	"os/exec"
	"strings"
)

const TempFileName = "asset.temp"

func Open(content string) string {
	createTempFile(content)
	cmd := exec.Command("vim", TempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	common.Check(err, "Failed to run vim")
	updatedContent := readTempFile()
	removeTempFile()
	return updatedContent
}

func createTempFile(content string) {
	bytes := []byte(content)
	err := os.WriteFile(TempFileName, bytes, 0644)
	common.Check(err, "Failed to write to file")
}

func removeTempFile() {
	cmd := exec.Command("rm", TempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	common.Check(err, "Failed to run rm")
}

func readTempFile() string {
	// read text file and get text
	bytes, err := os.ReadFile("asset.temp")
	common.Check(err, "Failed to read file")
	content := string(bytes)
	return strings.TrimSpace(content)
}
