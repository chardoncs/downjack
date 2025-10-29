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

	options []string

	filter          string
	filteredOptions []string

	offset int
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	switch msg := msg.(type) {
	case filterUpdateMsg:
		m.updateFilter(msg.text, false)

	case nextItemMsg:
		m.moveBy(1)
	case prevItemMsg:
		m.moveBy(-1)

	case selectionTriggerMsg:
		m.selected = true
		return m, tea.Quit
	}

	return m, nil
}

func (m listModel) View() string {
	renderedItems := make([]string, min(len(m.filteredOptions), listHeight))
	// i is display index
	for i := 0; i < listHeight && i < len(m.filteredOptions); i++ {
		item := m.filteredOptions[i+m.offset]

		if i+m.offset == m.index {
			renderedItems[i] = selectedListItemStyle.Render("> " + item)
		} else {
			renderedItems[i] = listItemStyle.Render("  " + item)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Right,
		listFrameStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left, renderedItems...),
		),
		listFooterStyle.Render(
			fmt.Sprintf("%d/%d", m.index+1, len(m.filteredOptions)),
		),
	)
}

func (m listModel) SelectedItem() (string, bool) {
	if !m.selected || m.index < 0 || m.index >= len(m.filteredOptions) {
		return "", false
	}

	return m.filteredOptions[m.index], true
}

func (m *listModel) updateFilter(filter string, force bool) {
	trimmedFilter := strings.TrimSpace(filter)
	if !force && trimmedFilter == m.filter {
		return
	}

	m.offset = 0
	m.index = 0

	m.filter = trimmedFilter
	m.filteredOptions = fuzzy.FindFold(trimmedFilter, m.options)
}

func (m *listModel) clampIndex() {
	m.index = min(max(m.index, 0), len(m.filteredOptions)-1)
}

func (m *listModel) moveBy(moves int) {
	m.index = m.index + moves
	m.clampIndex()

	topMostIndex := m.index - listHeight + 1
	bottomMostIndex := m.index + listHeight

	if m.offset < topMostIndex {
		m.offset = topMostIndex
	} else if m.offset+listHeight > bottomMostIndex {
		m.offset = bottomMostIndex - listHeight
	}
}

func initListModel(options []string, initialFilter string) listModel {
	m := listModel{
		options: options,
		filter:  strings.TrimSpace(initialFilter),
	}

	m.updateFilter(m.filter, true)

	return m
}
