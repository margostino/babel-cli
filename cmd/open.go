package cmd

import (
	"fmt"

	"github.com/margostino/babel-cli/internal/editor"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open-babel",
	Short: "Open editor",
	Long:  `Open a new editor with Babel Workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		// open()
		fmt.Println("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func open() {
	editor.Open(nil)
}
