package ui

type filterUpdateMsg struct {
	text string
}

type (
	nextItemMsg struct{}
	prevItemMsg struct{}

	nextPageMsg struct{}
	prevPageMsg struct{}

	nextHalfPageMsg struct{}
	prevHalfPageMsg struct{}
)

type selectionTriggerMsg struct{}

type abortMsg struct{}
