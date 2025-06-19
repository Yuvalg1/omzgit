package pointers

import (
	"omzgit/lib/list"
	"omzgit/messages"
	"omzgit/popups/async"
	"omzgit/popups/commit"
	"omzgit/popups/discard"
	"omzgit/popups/input"
	"omzgit/program"
	"omzgit/program/branches"
	"omzgit/program/files"
	"omzgit/program/files/row"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ messages.Apier = (*async.Model)(nil)

	_ messages.Cokerer      = (*branches.Model)(nil)
	_ messages.Popuper[any] = (*branches.Model)(nil)

	_ messages.Popuper[func()] = (*commit.Model)(nil)
	_ messages.Refresher       = (*commit.Model)(nil)

	_ messages.Popuper[func() tea.Cmd] = (*discard.Model)(nil)
	_ messages.Refresher               = (*discard.Model)(nil)

	_ messages.Cokerer                 = (*files.Model)(nil)
	_ messages.Popuper[func() tea.Cmd] = (*files.Model)(nil)
	_ messages.Ticker                  = (*files.Model)(nil)

	_ messages.Refresher = (*input.Model)(nil)

	_ messages.Moderer = (*list.Model[tea.Model])(nil)

	_ messages.Moderer                 = (*program.Model)(nil)
	_ messages.Refresher               = (*program.Model)(nil)
	_ messages.Popuper[func() tea.Cmd] = (*program.Model)(nil)

	_ messages.Popuper[func() bool] = (*row.Model)(nil)
)
