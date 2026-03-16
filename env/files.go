package env

type files struct {
	Add        Option
	AddAll     Option
	Commit     Option
	Discard    Option
	DiscardAll Option
	Enter      Option
	Reset      Option
	ResetAll   Option
	Yank       Option
	Ours       Option
	Theirs     Option

	Up     Option
	Down   Option
	Bottom Option
	PgUp   Option
	PgDown Option

	Refresh Option
	Search  Option
}

var Files = files{
	Add: Option{
		Msg:         "a",
		Description: "adds highlighted file",
	},
	AddAll: Option{
		Msg:         "A",
		Description: "adds all changed files",
	},
	Commit: Option{
		Msg:         "c",
		Description: "opens commit menu",
	},
	Discard: Option{
		Msg:         "d",
		Description: "opens discard popup to discard highlighted file",
	},
	DiscardAll: Option{
		Msg:         "D",
		Description: "opens discard popup to discard all files",
	},
	Enter: Option{
		Msg:         "enter",
		Description: "a magic key that opens conflict menu and toggle staging of highlighted file",
	},
	Reset: Option{
		Msg:         "r",
		Description: "resets highlighted file",
	},
	ResetAll: Option{
		Msg:         "R",
		Description: "resets all files",
	},
	Yank: Option{
		Msg:         "y",
		Description: "yanks highlighted file path",
	},
	Ours: Option{
		Msg:         "O",
		Description: "checkouts --ours changes in case of conflict",
	},
	Theirs: Option{
		Msg:         "T",
		Description: "checkouts --theirs changes in case of conflict",
	},

	PgUp: Option{
		Msg:         "pgup",
		Description: "scrolls up one line in file diff",
	},
	PgDown: Option{
		Msg:         "pgdown",
		Description: "scrolls down one line in file diff",
	},
	Up: Option{
		Msg:         "up",
		Description: "highlights previous file",
		AltMsg:      "k",
	},
	Down: Option{
		Msg:         "down",
		Description: "highlights next file",
		AltMsg:      "j",
	},
	Bottom: Option{
		Msg:         "G",
		Description: "highlights last file",
	},

	Refresh: Refresh,
	Search:  Search,
}
