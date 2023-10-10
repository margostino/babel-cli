package cmd

import (
	"github.com/margostino/babel-cli/pkg/editor"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open editor",
	Long:  `Open a new editor with Babel Workspace`,
	Run: func(cmd *cobra.Command, args []string) {
		open()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}

func open() {
	editor.Open(nil)
}
