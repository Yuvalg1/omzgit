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
