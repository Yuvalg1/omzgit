package messages

import tea "github.com/charmbracelet/bubbletea"

type TitleMsg struct {
	Title string
}

type Titler interface {
	TitleCmd(title string) tea.Cmd
}
