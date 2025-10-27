package ui

import tea "github.com/charmbracelet/bubbletea/v2"

type listModel struct {
	index    int
	selected bool

	filter string

	options         []string
	filteredOptions []string

	offset int

	clearUiAndExit bool
}

func (self listModel) Init() tea.Cmd {
	return nil
}

func (self listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	if self.clearUiAndExit {
		// UI must have been cleared already
		return self, tea.Quit
	}

	switch msg := msg.(type) {
	case filterUpdateMsg:
		self.updateFilter(msg.text)

	case nextItemMsg:
		self.moveBy(1)
	case prevItemMsg:
		self.moveBy(-1)

	case selectionTriggerMsg:
		self.selected = true
		self.clearUiAndExit = true
	case abortMsg:
		self.clearUiAndExit = true
	}

	return self, nil
}

func (self listModel) View() string {
	if self.clearUiAndExit {
		return ""
	}

	return ""
}

func (self listModel) SelectedItem() (string, bool) {
	if !self.selected || self.index < 0 || self.index >= len(self.filteredOptions) {
		return "", false
	}

	return self.filteredOptions[self.index], true
}

// Handle key bindings
//
// returns: if message handled, the model, and the command
func (self listModel) HandleKeyBindings(msg tea.KeyPressMsg) (bool, listModel, tea.Cmd) {
	handled := true
	var model listModel
	var cmd tea.Cmd

	switch msg.String() {
	case "enter":
		model, cmd = self.Update(selectionTriggerMsg{})

	case "ctrl+c":
		model, cmd = self.Update(abortMsg{})

	case "ctrl+n", "down":
		model, cmd = self.Update(nextItemMsg{})
	case "ctrl+p", "up":
		model, cmd = self.Update(prevItemMsg{})
	// TODO: More

	default:
		handled = false
	}

	return handled, model, cmd
}

func (self *listModel) updateFilter(filter string) {
	self.filter = filter
	// TODO: filter logic
}

func (self *listModel) clampIndex() {
	self.index = min(max(self.index, 0), len(self.filteredOptions)-1)
}

func (self *listModel) moveBy(moves int) {
	self.index = self.index + moves
	self.clampIndex()

	topMostIndex := self.index - listHeight + 1
	bottomMostIndex := self.index + listHeight - 1

	if self.offset < topMostIndex {
		self.offset = topMostIndex
	} else if self.offset > bottomMostIndex {
		self.offset = bottomMostIndex
	}
}

func initListModel(options []string) listModel {
	return listModel{
		options: options,
	}
}
