package content

import "fmt"

func (m Model) View() string {
	return fmt.Sprintf("%d", len(m.conflicts)) + m.Content.View()
}
