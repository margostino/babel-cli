/*
Copyright © 2023 margostino (maj.dagostino@gmail.com)

*/
package main

import (
	"github.com/margostino/babel-cli/cmd"
	"github.com/margostino/babel-cli/pkg/data"
	"log"
)

func main() {
	err := data.OpenDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	cmd.Execute()
}
