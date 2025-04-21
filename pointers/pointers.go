package pointers

import (
	"program/messages"
	"program/program"
	"program/program/files"
	"program/program/files/row"
)

var (
	_ messages.Popuper = (*files.Model)(nil)
	_ messages.Ticker  = (*files.Model)(nil)

	_ messages.Deleter = (*program.Model)(nil)
	_ messages.Titler  = (*files.Model)(nil)

	_ messages.Popuper = (*row.Model)(nil)
)
