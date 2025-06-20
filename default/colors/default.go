package colors

import "github.com/charmbracelet/lipgloss"

var (
	Red    = lipgloss.Color("#FF7B72")
	Yellow = lipgloss.Color("#F2CC60")
	Green  = lipgloss.Color("#56D364")
	Blue   = lipgloss.Color("#79C0FF")
	Orange = lipgloss.Color("#F0883E")
	Aqua   = lipgloss.Color("#A5D6FF")
	Pink   = lipgloss.Color("#FF9BCE")
)

func GetColor(cond bool, color1 lipgloss.Color, color2 lipgloss.Color) lipgloss.Color {
	if cond {
		return color1
	}

	return color2
}
