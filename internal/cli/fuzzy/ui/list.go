package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type listModel struct {
	index    int
	selected bool

	options         []string
	filteredOptions []string

	offset int
}

func (self listModel) Init() tea.Cmd {
	return nil
}

func (self listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	switch msg := msg.(type) {
	case filterUpdateMsg:
		self.updateFilter(msg.text)

	case nextItemMsg:
		self.moveBy(1)
	case prevItemMsg:
		self.moveBy(-1)

	case selectionTriggerMsg:
		self.selected = true
		return self, tea.Quit
	}

	return self, nil
}

func (self listModel) View() string {
	renderedItems := make([]string, min(len(self.filteredOptions), listHeight))
	// i is display index
	for i := 0; i < listHeight && i < len(self.filteredOptions); i++ {
		item := self.filteredOptions[i + self.offset]

		if i + self.offset == self.index {
			renderedItems[i] = selectedListItemStyle.Render("> "+item)
		} else {
			renderedItems[i] = listItemStyle.Render("  "+item)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Right,
		listFrameStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left, renderedItems...),
		),
		listFooterStyle.Render(
			fmt.Sprintf("%d/%d", self.index + 1, len(self.filteredOptions)),
		),
	)
}

func (self listModel) SelectedItem() (string, bool) {
	if !self.selected || self.index < 0 || self.index >= len(self.filteredOptions) {
		return "", false
	}

	return self.filteredOptions[self.index], true
}

func (self *listModel) updateFilter(filter string) {
	self.offset = 0
	self.index = 0

	trimmedFilter := strings.TrimSpace(filter)
	if trimmedFilter == "" {
		self.filteredOptions = self.options
		return
	}

	self.filteredOptions = fuzzy.FindFold(trimmedFilter, self.options)
}

func (self *listModel) clampIndex() {
	self.index = min(max(self.index, 0), len(self.filteredOptions)-1)
}

func (self *listModel) moveBy(moves int) {
	self.index = self.index + moves
	self.clampIndex()

	topMostIndex := self.index - listHeight + 1
	bottomMostIndex := self.index + listHeight

	if self.offset < topMostIndex {
		self.offset = topMostIndex
	} else if self.offset+listHeight > bottomMostIndex {
		self.offset = bottomMostIndex - listHeight
	}
}

func initListModel(options []string) listModel {
	return listModel{
		options:         options,
		filteredOptions: options,
	}
}
