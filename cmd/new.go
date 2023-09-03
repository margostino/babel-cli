package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/margostino/babel-cli/data"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new asset",
	Long:  `Creates a new Babel Asset (quick note, idea, knowledge, resource, etc.)`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewAsset()
	},
}

type promptContent struct {
	label string
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func promptGetInput(pc promptContent) string {
	//validate := func(input string) error {
	//	if len(input) <= 0 {
	//		return errors.New(pc.errorMsg)
	//	}
	//	return nil
	//}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold | cyan }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		//Validate:  validate,
	}

	fullResult := ""
	// while last 2 characters are not a /n, keep prompting
	for fullResult == "" || fullResult[len(fullResult)-2:] != ":q" {
		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		fullResult += result
	}

	fullResult = strings.Replace(fullResult, ":q", "", -1)
	return fullResult
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewAsset() {
	input := promptContent{
		"",
	}

	fmt.Printf("Start writing...\n\n")

	content := promptGetInput(input)
	println(content)
	//definitionPromptContent := promptContent{
	//	"Please provide a definition.",
	//	fmt.Sprintf("What is the definition of the %s?", word),
	//}
	//definition := promptGetInput(definitionPromptContent)
	//
	//categoryPromptContent := promptContent{
	//	"Please provide a category.",
	//	fmt.Sprintf("What category does %s belong to?", word),
	//}
	//category := promptGetSelect(categoryPromptContent)
	//println(content, definition, category)
	data.InsertNote(content)
}
