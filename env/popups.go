package env

type alert struct {
	Yank   Option
	Up     Option
	Down   Option
	PgUp   Option
	PgDown Option

	Quit  Option
	CtrlC Option
}

var Alert = alert{
	Yank: Option{
		Msg:         "y",
		Description: "yanks error message to clipboard",
	},
	Up: Option{
		Msg:         "up",
		Description: "scrolls up a row",
		AltMsg:      "k",
	},
	Down: Option{
		Msg:         "down",
		Description: "scrolls down a row",
		AltMsg:      "j",
	},
	PgUp: Option{
		Msg:         "pgup",
		Description: "scrolls up a page",
	},
	PgDown: Option{
		Msg:         "pgdown",
		Description: "scrolls down a page",
	},

	Quit:  Quit.Quit,
	CtrlC: Quit.CtrlC,
}

type commit struct {
	Back    Option
	File    Option
	Message Option
	Options Option
	Yank    Option

	Quit  Option
	CtrlC Option
}

var Commit = commit{
	Back: Option{
		Msg:         "esc",
		Description: "returns to files page",
	},

	File: Option{
		Msg:         "F",
		Description: "allows writing file name to take commit message from",
	},

	Message: Option{
		Msg:         "m",
		Description: "allows writing commit message",
	},

	Options: Option{
		Msg:         "o",
		Description: "opens more options for commit",
	},

	Yank: Option{
		Msg:         "y",
		Description: "yanks the message/file name to clipboard",
	},

	Quit:  Quit.Quit,
	CtrlC: Quit.CtrlC,
}
