package main

import (
	"github.com/margostino/babel-cli/cmd"
	"github.com/margostino/babel-cli/internal/config"
)

func main() {
	config.InitHome()
	//db.OpenDatabase()
	cmd.Execute()
}
