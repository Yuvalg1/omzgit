package commit

import (
	"omzgit/clipboard"
	"omzgit/git"
	"omzgit/messages/refresh"
	"omzgit/program/popups"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textinput.Width = getWidth(msg.Width) - 5

		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)
		return m, nil

	case popups.Msg:
		m.visible = true
		return m, nil

	case tea.KeyMsg:
		if m.textinput.Focused() {
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit

			case "esc":
				m.textinput.Blur()
				return m, nil

			case "down":
				m.textinput.SetCursor(len(m.textinput.Value()))
				return m, nil

			case "up":
				m.textinput.SetCursor(0)
				return m, nil

			case "enter":
				output, err := git.Exec(m.getCommitString()...)
				if err == nil {
					m = InitialModel(m.width, m.height, "commit")
					return m, refresh.Cmd()
				}

				return m, popups.Cmd("alert", "Commit Error!", output, func() tea.Cmd { return nil })

			default:
				res, cmd := m.textinput.Update(msg)
				m.textinput = res
				return m, cmd
			}
		}

		if m.moreOptions {
			switch msg.String() {
			case "a":
				m.writeOption('a', "--amend")
				return m, nil

			case "ctrl+c", "q":
				return m, tea.Quit

			case "e":
				m.writeOption('e', "--edit")
				return m, nil

			case "esc", "o":
				m.moreOptions = false
				return m, nil

			case "E":
				m.writeOption('E', "--no-edit")
				return m, nil

			case "n":
				m.writeOption('n', "--no-verify")
				return m, nil

			case "y":
				m.writeOption('y', "--allow-empty")
				return m, nil

			default:
				return m, nil
			}
		}

		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			output, err := git.Exec(m.getCommitString()...)
			if err == nil {
				m = InitialModel(m.width, m.height, "commit")
				return m, refresh.Cmd()
			}

			return m, popups.Cmd("alert", "Commit Error!", output, func() tea.Cmd { return nil })

		case "esc":
			m.visible = false
			return m, nil

		case "F":
			if m.commitMessageType != "-F" {
				m.commitMessageType = "-F"
				m.textinput.Placeholder = "File"
				m.textinput.SetValue("")
			}

			m.textinput.Focus()
			return m, nil

		case "m":
			if m.commitMessageType != "-m" {
				m.commitMessageType = "-m"
				m.textinput.Placeholder = "Message"
				m.textinput.SetValue("")
			}

			m.textinput.Focus()
			return m, nil

		case "o":
			m.moreOptions = true
			return m, nil

		case "y":
			clipboard.Copy(m.textinput.Value())
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}

func (m *Model) writeOption(letter byte, option string) {
	if len(m.options[letter]) == 0 {
		m.options[letter] = option
	} else {
		m.options[letter] = ""
	}
}
