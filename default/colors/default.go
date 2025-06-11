package colors

import "github.com/charmbracelet/lipgloss"

var (
	Red    = lipgloss.Color("#E67E80")
	Orange = lipgloss.Color("#E69875")
	Yellow = lipgloss.Color("#DBBC7F")
	Green  = lipgloss.Color("#A7C080")
	Blue   = lipgloss.Color("#7FBBB3")
	Aqua   = lipgloss.Color("#83C092")
	Purple = lipgloss.Color("#D699B6")

	Fg = lipgloss.Color("#D3C6AA")
)

func GetColor(cond bool, color1 lipgloss.Color, color2 lipgloss.Color) lipgloss.Color {
	if cond {
		return color1
	}

	return color2
}
