package pointers

import (
	"program/lib/list"
	"program/messages"
	"program/program"
	"program/program/branches"
	"program/program/files"
	"program/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ messages.Popuper[any] = (*branches.Model)(nil)

	_ messages.Cokerer         = (*files.Model)(nil)
	_ messages.Popuper[func()] = (*files.Model)(nil)
	_ messages.Ticker          = (*files.Model)(nil)

	_ messages.Moderer = (*list.Model[tea.Model])(nil)

	_ messages.Deleter = (*program.Model)(nil)
	_ messages.Moderer = (*program.Model)(nil)

	_ messages.Popuper[func()] = (*row.Model)(nil)
)
