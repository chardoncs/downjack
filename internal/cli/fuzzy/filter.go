package fuzzy

import "github.com/charmbracelet/bubbles/v2/list"

type item string

func (self item) FilterValue() string {
	return string(self)
}

func toList(optionList []string) []list.Item {
	length := len(optionList)
	results := make([]list.Item, length, length)

	for i, value := range optionList {
		results[i] = item(value)
	}

	return results
}
