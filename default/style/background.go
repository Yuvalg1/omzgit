package style

import (
	"omzgit/default/colors/bg"

	"github.com/charmbracelet/lipgloss"
)

var Bg = lipgloss.NewStyle().
	BorderBackground(bg.C[0]).
	Background(bg.C[0]).
	BorderForeground(bg.C[4])
