package pointers

import (
	"program/messages"
	"program/program"
	"program/program/files"
	"program/program/files/row"
)

var (
	_ messages.Ticker  = (*files.Model)(nil)
	_ messages.Deleter = (*program.Model)(nil)
)

var _ messages.Popuper = (*row.Model)(nil)
