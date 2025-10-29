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
		switch msg.String() {
		default:
			var handled bool
			handled, self.listModel, listCmd = self.handleKeyBindings(msg)
			if handled {
				return self, listCmd
			}

		case "ctrl+c":
			return self, tea.Quit
		}
	}

	var inputCmd tea.Cmd

	self.inputModel, inputCmd = self.inputModel.Update(msg)
	self.listModel, listCmd = self.listModel.Update(filterUpdateMsg{text: self.inputModel.Value()})

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

// Handle key bindings
//
// returns: if message handled, the model, and the command
func (self MainModel) handleKeyBindings(msg tea.KeyPressMsg) (bool, listModel, tea.Cmd) {
	handled := true
	var model listModel
	var cmd tea.Cmd

	switch msg.String() {
	case "enter":
		model, cmd = self.listModel.Update(selectionTriggerMsg{})

	case "ctrl+c":
		model, cmd = self.listModel.Update(abortMsg{})

	case "ctrl+n", "down":
		model, cmd = self.listModel.Update(nextItemMsg{})
	case "ctrl+p", "up":
		model, cmd = self.listModel.Update(prevItemMsg{})
	// TODO: More

	default:
		handled = false
		model = self.listModel
	}

	return handled, model, cmd
}

func InitialModel(title string, options []string, initialInput string) MainModel {
	im := textinput.New()
	im.Focus()

	im.ShowSuggestions = true
	im.SetSuggestions(options)

	im.VirtualCursor = true

	im.SetValue(initialInput)

	return MainModel{
		inputModel: im,
		listModel:  initListModel(options, initialInput),
		title:      title,
	}
}
