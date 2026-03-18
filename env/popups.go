package env

type alert struct {
	Yank Option

	Quit  Option
	CtrlC Option
}

var Alert = alert{
	Yank: Option{
		Msg:         "y",
		Description: "yanks error message to clipboard",
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
