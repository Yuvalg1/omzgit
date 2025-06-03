package list

import (
	"program/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TerminalMsg:
		m.width = getWidth(msg.Width)
		m.height = getHeight(msg.Height)

		var cmds []tea.Cmd
		for index, element := range m.Children {
			res, cmd := element.Update(msg)
			m.Children[index] = res.(T)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		if m.TextInput.Focused() {
			switch msg.String() {
			case "esc":
				m.TextInput.Blur()
				m.TextInput.SetValue("")

				res, cmd := m.Children[m.ActiveRow].Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'g'}})
				m.Children[m.ActiveRow] = res.(T)

				return m, tea.Batch(cmd, m.ModeCmd(""))

			case "enter":
				m.SetContent(m.Children)
				m.TextInput.Blur()

				res, cmd := m.Children[0].Update(msg)
				m.Children[0] = res.(T)

				return m, tea.Batch(cmd, m.ModeCmd(""))
			}

			res, cmd := m.TextInput.Update(msg)
			m.TextInput = res

			return m, cmd
		}

		switch keypress := msg.String(); keypress {
		case "esc":
			m.TextInput.SetValue("")

			res, cmd := m.Children[m.ActiveRow].Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'g'}})
			m.Children[m.ActiveRow] = res.(T)

			return m, cmd

		case "g":
			curr := m.ActiveRow

			m.ActiveRow = 0
			cmds := move(m, msg, curr, 0)

			return m, tea.Batch(cmds...)

		case "G":
			curr := m.ActiveRow

			m.ActiveRow = len(m.Children) - 1
			cmds := move(m, msg, curr, len(m.Children)-1)

			return m, tea.Batch(cmds...)

		case "j", "down":
			curr := m.ActiveRow
			next := (m.ActiveRow + 1 + len(m.Children)) % len(m.Children)

			m.ActiveRow = next
			cmds := move(m, msg, curr, next)

			return m, tea.Batch(cmds...)

		case "k", "up":
			curr := m.ActiveRow
			next := (m.ActiveRow - 1 + len(m.Children)) % len(m.Children)

			m.ActiveRow = next
			cmds := move(m, msg, curr, next)

			return m, tea.Batch(cmds...)

		case "/":
			curr := m.ActiveRow

			m.ActiveRow = 0
			cmds := move(m, msg, curr, 0)

			cmds = append(cmds, m.ModeCmd("search"))
			m.TextInput.Focus()
			return m, tea.Batch(cmds...)

		default:
			var cmds []tea.Cmd

			for index, element := range m.Children {
				res, cmd := element.Update(msg)
				m.Children[index] = res.(T)
				cmds = append(cmds, cmd)
			}

			return m, tea.Batch(cmds...)
		}
	}

	return m, nil
}

func move[T tea.Model](m Model[T], msg tea.Msg, curr int, next int) []tea.Cmd {
	if len(m.Children) == 0 {
		return []tea.Cmd{}
	}

	res1, cmd1 := m.Children[curr].Update(msg)
	m.Children[curr] = res1.(T)

	res2, cmd2 := m.Children[next].Update(msg)
	m.Children[next] = res2.(T)

	return []tea.Cmd{cmd1, cmd2}
}

func (m Model[T]) getFilteredChildren() []T {
	filteredChildren := []T{}

	for _, element := range m.Children {
		if m.filterFn(element, m.TextInput.Value()) {
			filteredChildren = append(filteredChildren, element)
		}
	}

	return filteredChildren
}
