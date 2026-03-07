package env

type quit struct {
	CtrlC Option
	Quit  Option
}

var Quit = quit{
	CtrlC: Option{
		Msg:         "ctrl+c",
		Description: "quits application",
	},
	Quit: Option{
		Msg:         "q",
		Description: "quits application",
	},
}
