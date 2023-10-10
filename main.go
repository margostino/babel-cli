package main

import (
	"github.com/margostino/babel-cli/cmd"
	"github.com/margostino/babel-cli/pkg/config"
)

func main() {
	config.InitHome()
	//db.OpenDatabase()
	cmd.Execute()
}
