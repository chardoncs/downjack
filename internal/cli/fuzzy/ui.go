package fuzzy

import (
	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

const (
	listWidth  = 20
	listHeight = 14
)

type model struct {
	inputModel textinput.Model
	listModel  list.Model
}

func (self model) Init() tea.Cmd {
	return nil
}

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return self, tea.Quit
		}
	}

	var inputCmd tea.Cmd
	var listCmd tea.Cmd

	self.inputModel, inputCmd = self.inputModel.Update(msg)
	self.listModel, listCmd = self.listModel.Update(msg)

	return self, tea.Batch(inputCmd, listCmd)
}

func (self model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		self.inputModel.View(),
		self.listModel.View(),
	)
}

func (self model) SelectedIndex() int {
	return self.listModel.Index()
}

func initialModel(prompt string, options []string) model {
	im := textinput.New()

	lm := list.New(toList(options), itemDelegate{}, listWidth, listHeight)
	lm.SetShowTitle(false)

	return model{
		inputModel: im,
		listModel:  lm,
	}
}
