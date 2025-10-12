package list

import (
	"omzgit/messages"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = getHeight(msg.Height)

		var cmds []tea.Cmd
		for index, element := range m.Children {
			res, cmd := element.Update(msg)
			m.Children[index] = res.(T)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)

	case messages.RollerMsg:
		res, cmd := m.Children[m.ActiveRow].Update(msg)
		m.Children[m.ActiveRow] = res.(T)

		return m, cmd

	case tea.KeyMsg:
		if m.TextInput.Focused() {
			switch msg.String() {
			case "esc":
				m.TextInput.Blur()
				m.TextInput.SetValue("")

				return m, m.ModeCmd("goto")

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
		case "enter":
			return m, nil

		case "esc":
			m.TextInput.SetValue("")

			res, cmd := m.Children[m.ActiveRow].Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'g'}})
			m.Children[m.ActiveRow] = res.(T)

			return m, cmd

		case "g":
			curr := m.ActiveRow

			m.ActiveRow = 0
			cmd := move(m, msg, curr, 0)
			m.innerOffset = 0

			return m, cmd

		case "G":
			curr := m.ActiveRow

			m.ActiveRow = len(m.Children) - 1
			cmd := move(m, msg, curr, len(m.Children)-1)
			m.innerOffset = m.height - 2

			return m, cmd

		case "j", "down":
			curr := m.ActiveRow
			next := (m.ActiveRow + 1 + len(m.Children)) % len(m.Children)

			m.innerOffset = min(m.height-2, m.innerOffset+1)
			m.ActiveRow = next
			cmd := move(m, msg, curr, next)

			return m, cmd

		case "k", "up":
			curr := m.ActiveRow
			next := (m.ActiveRow - 1 + len(m.Children)) % len(m.Children)

			m.innerOffset = max(0, m.innerOffset-1)

			if next == len(m.Children)-1 {
				m.innerOffset = m.height - 2
			}

			m.ActiveRow = next
			cmd := move(m, msg, curr, next)

			return m, cmd

		case "/":
			curr := m.ActiveRow

			m.ActiveRow = 0
			cmd := move(m, msg, curr, 0)

			m.TextInput.Focus()
			return m, tea.Batch(cmd, m.ModeCmd("search"))

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

func move[T tea.Model](m Model[T], msg tea.Msg, curr int, next int) tea.Cmd {
	if len(m.Children) == 0 {
		return nil
	}

	res1, cmd1 := m.Children[curr].Update(msg)
	m.Children[curr] = res1.(T)

	res2, cmd2 := m.Children[next].Update(msg)
	m.Children[next] = res2.(T)

	return tea.Batch(cmd1, cmd2)
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
