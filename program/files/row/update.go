package row

import (
	"os/exec"
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "a":
			if !m.Staged {
				m.Staged = gitAdd(m)
			}
			return m, nil

		case "A":
			m.Staged = true
			return m, nil

		case "g", "G", "/", "esc":
			m.Active = !m.Active
			return m, nil

		case "j", "k", "down", "up":
			m.Active = !m.Active
			return m, nil

		case "d":
			return m, m.PopupCmd(m.Path, func() {
				if m.Staged {
					m.Staged = !gitReset(m)
				}
				gitRestore(m.Path)
			})

		case "r":
			if m.Staged {
				m.Staged = !gitReset(m)
			}
			return m, nil

		case "R":
			m.Staged = false
			return m, nil

		default:
			return m, nil
		}
	}

	return m, nil
}

func gitAdd(m Model) bool {
	cmd := exec.Command("git", "add", m.Path)

	_, err := cmd.Output()

	return err == nil
}

func gitReset(m Model) bool {
	cmd := exec.Command("git", "reset", "--", m.Path)

	_, err := cmd.Output()

	return err == nil
}

func gitRestore(path string) {
	cmd := exec.Command("git", "restore", path)

	cmd.Output()
}
