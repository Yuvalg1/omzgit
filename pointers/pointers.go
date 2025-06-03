package pointers

import (
	"program/lib/list"
	"program/messages"
	"program/program"
	"program/program/files"
	"program/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ messages.Cokerer = (*files.Model)(nil)
	_ messages.Popuper = (*files.Model)(nil)
	_ messages.Ticker  = (*files.Model)(nil)

	_ messages.Deleter = (*program.Model)(nil)
	_ messages.Moderer = (*program.Model)(nil)

	_ messages.Popuper = (*row.Model)(nil)

	_ messages.Moderer = (*list.Model[tea.Model])(nil)
)
