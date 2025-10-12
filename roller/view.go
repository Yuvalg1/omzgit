package roller

import (
	"strings"

	"omzgit/consts"
)

func (m Model) View() string {
	if len(m.Name) < m.Width {
		return m.Name
	}

	rolled := m.Name[m.Offset:] + strings.Repeat(" ", max(1, m.Width-len(m.Name))) + m.Name[:m.Offset]
	return consts.TrimRight(rolled, m.Width)
}
