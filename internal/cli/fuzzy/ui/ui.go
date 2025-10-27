package ui

import (
	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

type MainModel struct {
	inputModel textinput.Model
	listModel  listModel

	title string
}

func (self MainModel) Init() tea.Cmd {
	return nil
}

func (self MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var listCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		var handled bool
		handled, self.listModel, listCmd = self.listModel.HandleKeyBindings(msg)
		if handled {
			return self, listCmd
		}

	case itemSelectedMsg:
		return self, tea.Quit
	}

	var inputCmd tea.Cmd

	self.inputModel, inputCmd = self.inputModel.Update(msg)
	self.listModel, listCmd = self.listModel.Update(msg)

	return self, tea.Batch(inputCmd, listCmd)
}

func (self MainModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(self.title),
		self.inputModel.View(),
		self.listModel.View(),
	)
}

func (self MainModel) SelectedItem() (string, bool) {
	return self.listModel.SelectedItem()
}

func InitialModel(title string, options []string) MainModel {
	return MainModel{
		inputModel: textinput.New(),
		listModel:  initListModel(options),
		title:      title,
	}
}
