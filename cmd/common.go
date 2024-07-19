package cmd

import (
	"strings"
)

func extractParam(args []string, pos int) *string {
	var param *string
	if len(args) > pos {
		param = &args[pos]
	}
	return param
}

func concatAllParams(args []string) *string {
	if len(args) == 0 {
		return nil
	}
	joinedArgs := strings.TrimSpace(strings.Join(args[0:], "_"))
	return &joinedArgs
}
