package commits

type snapshot struct {
	total int

	listNewSize        int
	listTextInputValue string

	width  int
	height int
}

func (m Model) getSnapshot() snapshot {
	return snapshot{
		total: m.total,

		listNewSize:        m.list.NewSize(),
		listTextInputValue: m.list.TextInput.Value(),

		width:  m.width,
		height: m.height,
	}
}
