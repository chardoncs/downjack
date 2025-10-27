package ui

type filterUpdateMsg struct {
	text string
}

type (
	nextItemMsg struct{}
	prevItemMsg struct{}
)

type itemSelectedMsg struct {}

type selectionTriggerMsg struct{}

type abortMsg struct{}
