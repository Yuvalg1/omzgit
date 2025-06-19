package commit

import (
	"omzgit/messages"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	commitMessageType string
	moreOptions       bool
	options           map[byte]string
	textinput         textinput.Model
	visible           bool

	width  int
	height int
}

func InitialModel(width int, height int, title string) Model {
	ti := textinput.New()
	ti.CharLimit = 50
	ti.Placeholder = "Message"
	ti.Width = getWidth(width)

	options := map[byte]string{}
	options['a'] = ""
	options['e'] = ""
	options['E'] = ""
	options['n'] = ""
	options['y'] = ""

	return Model{
		commitMessageType: "-m",
		moreOptions:       false,
		options:           options,
		textinput:         ti,
		visible:           false,

		width:  getWidth(width),
		height: getHeight(height),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) PopupCmd(pType string, placeholder string, title string, fn func() tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return messages.PopupMsg{
			Name: title,
			Type: pType,
			Verb: placeholder,
		}
	}
}

func (m Model) RefreshCmd() tea.Cmd {
	return func() tea.Msg {
		return messages.RefreshMsg{}
	}
}

func getWidth(width int) int {
	return min(width-4+width%2, 48)
}

func getHeight(height int) int {
	return 2
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m Model) getCommitString() []string {
	commitStrings := []string{"commit"}
	for _, element := range m.options {
		if element != "" {
			commitStrings = append(commitStrings, element)
		}
	}

	commitStrings = append(commitStrings, m.commitMessageType, "\""+m.textinput.Value()+"\"")

	return commitStrings
}

func (m Model) getCommitStringString() string {
	commitStrings := "commit"
	for _, element := range m.options {
		if element != "" {
			commitStrings += " " + element
		}
	}

	commitStrings += " " + m.commitMessageType + " " + "\"" + m.textinput.Value() + "\""

	return commitStrings
}
