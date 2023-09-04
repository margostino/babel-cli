package main

import (
	"github.com/margostino/babel-cli/cmd"
	"github.com/margostino/babel-cli/pkg/config"
	"github.com/margostino/babel-cli/pkg/data"
)

func main() {
	config.InitHome()
	data.OpenDatabase()
	cmd.Execute()
}
