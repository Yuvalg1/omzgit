package files

type snapshot struct {
	total int

	listNewSize        int
	listTextInputValue string

	width  int
	height int
}

func (m Model) getSnapshot() snapshot {
	width, height := m.getFilesAxis()
	return snapshot{
		total: m.total,

		listNewSize:        m.list.NewSize(),
		listTextInputValue: m.list.TextInput.Value(),

		width:  width,
		height: height,
	}
}
