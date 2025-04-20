package messages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TerminalMsg struct {
	Width  int
	Height int
}

type TerminalResizer interface {
	TerminalCmd() tea.Cmd
}
