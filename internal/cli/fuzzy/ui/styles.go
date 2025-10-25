package ui

import "github.com/charmbracelet/lipgloss/v2"

const (
	listWidth  = 20
	listHeight = 14
)

var (
	titleStyle            = lipgloss.NewStyle().PaddingLeft(2)
	listItemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	selectedListItemStyle = listItemStyle.Foreground(lipgloss.Color("170"))
)
