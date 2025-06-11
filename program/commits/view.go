package commits

import (
	"program/default/style"
)

func (m Model) View() string {
	return style.Bg.Width(m.width).Height(m.height).Render(m.Title)
}
