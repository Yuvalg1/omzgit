package roller

import (
	"time"

	"omzgit/messages"

	"github.com/bep/debounce"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Name       string
	Offset     int
	debounceFn func(f func())

	Width int
}

func InitialModel(width int, name string) Model {
	return Model{
		Name:       name,
		Offset:     0,
		debounceFn: debounce.New(400 * time.Millisecond),

		Width: getWidth(width),
	}
}

func (m Model) Init() tea.Cmd {
	return m.RollerCmd()
}

func (m Model) InitRollerCmd() tea.Cmd {
	return func() tea.Msg {
		channel := make(chan tea.Msg)
		m.debounceFn(func() {
			time.Sleep(2 * time.Second)
			if m.Width < len(m.Name) {
				channel <- m.RollerCmd()()
			}
		})
		return <-channel
	}
}

func (m Model) RollerCmd() tea.Cmd {
	return func() tea.Msg {
		channel := make(chan tea.Msg)
		m.debounceFn(func() {
			if m.Width < len(m.Name) {
				channel <- messages.RollerMsg{}
			}
		})
		return <-channel
	}
}

func getWidth(width int) int {
	return width - 3
}
