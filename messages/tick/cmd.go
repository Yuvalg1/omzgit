package tick

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Msg struct {
	RollOffset int
}

func Cmd(rollOffset int) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(10 * time.Second)
		return Msg{
			RollOffset: rollOffset,
		}
	}
}
