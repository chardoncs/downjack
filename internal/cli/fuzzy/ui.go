package fuzzy

import (
	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/bubbles/v2/textinput"
	tea "github.com/charmbracelet/bubbletea/v2"
)

const (
	listWidth  = 20
	listHeight = 14
)

type model struct {
	selected int

	inputModel textinput.Model
	listModel  list.Model
}

func (self model) Init() tea.Cmd {
	return nil
}

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return self, nil
}

func (self model) View() string {
	return ""
}

func (self model) SelectedIndex() int {
	return self.selected
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
