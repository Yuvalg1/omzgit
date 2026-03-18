package branches

type snapshot struct {
	remote bool
	total  int

	listNewSize        int
	listTextInputValue string

	width  int
	height int
}

func (m Model) getSnapshot() snapshot {
	return snapshot{
		remote: m.remote,
		total:  m.total,

		listNewSize:        m.list.NewSize(),
		listTextInputValue: m.list.TextInput.Value(),

		width:  m.width,
		height: m.height,
	}
}
