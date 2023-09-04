/*
Copyright © 2023 margostino (maj.dagostino@gmail.com)

*/
package main

import (
	"github.com/margostino/babel-cli/cmd"
	"github.com/margostino/babel-cli/pkg/config"
	"github.com/margostino/babel-cli/pkg/data"
)

func main() {
	config.InitBabelHome()
	data.OpenDatabase()
	cmd.Execute()
}
