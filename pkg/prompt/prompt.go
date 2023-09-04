package prompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/margostino/babel-cli/pkg/data"
	"os"
	"strings"
)

func GetInput(pc Prompt) string {
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
		Label:     pc.Label,
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

func GetSelect(pc Prompt, items []string) int {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }} ",
		Active:   "{{ . | bold | yellow }} ",
		Selected: "{{ . | green }} ",
		Inactive: "{{ . | cyan }} ",
	}

	prompt := promptui.Select{
		Label:     pc.Label,
		Items:     items,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return index
}

func AssetsToItems(assets []*data.Asset) []string {
	items := make([]string, 0)
	for _, asset := range assets {
		category := data.GetCategoryAsString(asset.Category)
		content := strings.ReplaceAll(asset.Content, "\n", " ")
		if len(content) > 30 {
			content = content[:20] + "..."
		}
		item := fmt.Sprintf("(%d:%s) %s", asset.Id, category, content)
		item = strings.ReplaceAll(item, "\n", "")
		items = append(items, item)
	}
	return items
}
