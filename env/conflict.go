package env

type conflict struct {
	Back Option
	Both Option

	Our  Option
	Ours Option

	Their  Option
	Theirs Option

	Next     Option
	Previous Option

	Up   Option
	Down Option

	PgDown Option
	PgUp   Option

	Quit  Option
	CtrlC Option
}

var Conflict = conflict{
	Back: Option{
		Msg:         "esc",
		Description: "returns to files page",
	},

	Both: Option{
		Msg:         "b",
		Description: "saves both highlighted conflict changes in file",
	},

	Our: Option{
		Msg:         "o",
		Description: "saves our highlighted conflict in file",
	},

	Ours: Files.Ours,

	Their: Option{
		Msg:         "t",
		Description: "saves their highlighted conflict in file",
	},

	Theirs: Files.Theirs,

	Next: Option{
		Msg:         "n",
		Description: "scrolls to next conflict",
	},
	Previous: Option{
		Msg:         "N",
		Description: "scrolls to previous conflict",
	},

	Up: Option{
		Msg:         "up",
		Description: "highlights previous branch",
		AltMsg:      "k",
	},
	Down: Option{
		Msg:         "down",
		Description: "highlights next branch",
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
