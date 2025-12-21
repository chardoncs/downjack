package ui

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainModel struct {
	inputModel textinput.Model
	listModel  listModel

	title string
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var listCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		default:
			var handled bool
			handled, m.listModel, listCmd = m.handleKeyBindings(msg)
			if handled {
				return m, listCmd
			}

		case "ctrl+c":
			return m, tea.Quit
		}
	}

	var inputCmd tea.Cmd

	m.inputModel, inputCmd = m.inputModel.Update(msg)
	m.listModel, listCmd = m.listModel.Update(filterUpdateMsg{text: m.inputModel.Value()})

	return m, tea.Batch(inputCmd, listCmd)
}

func (m MainModel) View() tea.View {
	var v tea.View

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(m.title),
		m.inputModel.View(),
		m.listModel.String(),
	)

	v.SetContent(content)
	return v
}

func (m MainModel) SelectedItem() (string, bool) {
	return m.listModel.SelectedItem()
}

// Handle key bindings
//
// returns: if message handled, the model, and the command
func (m MainModel) handleKeyBindings(msg tea.KeyPressMsg) (bool, listModel, tea.Cmd) {
	handled := true
	var model listModel
	var cmd tea.Cmd

	var downMsg tea.Msg

	switch msg.String() {
	case "enter":
		downMsg = selectionTriggerMsg{}

	case "ctrl+c":
		downMsg = abortMsg{}

	case "ctrl+n", "down":
		downMsg = nextItemMsg{}
	case "ctrl+p", "up":
		downMsg = prevItemMsg{}

	case "ctrl+v":
		downMsg = nextPageMsg{}
	case "meta+v":
		downMsg = prevPageMsg{}
	}

	if downMsg != nil {
		model, cmd = m.listModel.Update(downMsg)
	} else {
		handled = false
		model = m.listModel
	}

	return handled, model, cmd
}

func InitialModel(title string, options []string, initialInput string) MainModel {
	im := textinput.New()
	im.Focus()

	im.ShowSuggestions = true
	im.SetSuggestions(options)

	im.SetVirtualCursor(true)

	im.SetValue(initialInput)

	return MainModel{
		inputModel: im,
		listModel:  initListModel(options, initialInput),
		title:      title,
	}
}
