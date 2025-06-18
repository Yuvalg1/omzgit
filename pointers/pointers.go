package pointers

import (
	"program/lib/list"
	"program/messages"
	"program/popups/async"
	"program/popups/discard"
	"program/popups/input"
	"program/program"
	"program/program/branches"
	"program/program/files"
	"program/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ messages.Apier = (*async.Model)(nil)

	_ messages.Cokerer      = (*branches.Model)(nil)
	_ messages.Popuper[any] = (*branches.Model)(nil)

	_ messages.Popuper[func() bool] = (*discard.Model)(nil)
	_ messages.Refresher            = (*discard.Model)(nil)

	_ messages.Cokerer              = (*files.Model)(nil)
	_ messages.Popuper[func() bool] = (*files.Model)(nil)
	_ messages.Ticker               = (*files.Model)(nil)

	_ messages.Refresher = (*input.Model)(nil)

	_ messages.Moderer = (*list.Model[tea.Model])(nil)

	_ messages.Moderer      = (*program.Model)(nil)
	_ messages.Refresher    = (*program.Model)(nil)
	_ messages.Popuper[any] = (*program.Model)(nil)

	_ messages.Popuper[func() bool] = (*row.Model)(nil)
)
