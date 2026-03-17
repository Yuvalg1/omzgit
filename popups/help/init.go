package help

import (
	"reflect"
	"strings"

	"omzgit/env"
	"omzgit/lib/list"
	"omzgit/popups/help/option"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list           list.Model[option.Model]
	defaultOptions []env.Option
	visible        bool
	total          int

	width  int
	height int
}

func InitialModel(width int, height int) Model {
	initialList := list.InitialModel(getHeight(height), []option.Model{}, 0, "No Branches Found")

	m := Model{
		list: initialList,

		width:  getWidth(width),
		height: getHeight(height),
	}

	m.list.Children = []option.Model{option.EmptyInitialModel(m.width, getHeight(height))}
	m.list.SetCreateChild(func(name string) *option.Model {
		created := option.EmptyInitialModel(m.width, getHeight(height))
		return &created
	})

	return m
}

func (m Model) Init() tea.Cmd {
	return m.list.Children[m.list.ActiveRow].Init()
}

func getHeight(height int) int {
	return height - 4
}

func getWidth(width int) int {
	return width
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m *Model) getOptions() []option.Model {
	index := 0

	var options []option.Model
	for len(options) < m.list.NewSize() && index < len(m.defaultOptions) {
		option := option.InitialModel(m.width, m.defaultOptions[index])

		if filterFn(option, m.list.TextInput.Value()) {
			options = append(options, option)
		}

		index++
	}

	if m.list.TextInput.Value() != "" {
		m.total = len(options)
	}

	return options
}

func filterFn(option option.Model, text string) bool {
	return strings.Contains(strings.ToLower(option.Roller.Name), strings.ToLower(text)) ||
		strings.Contains(option.Msg, text)
}

func GetEnvOptions(configuration any) []env.Option {
	values := reflect.ValueOf(configuration)
	if values.Kind() == reflect.Pointer {
		values = values.Elem()
	}

	envOptions := []env.Option{}
	itemType := reflect.TypeOf(env.Option{})

	for i := 0; i < values.NumField(); i++ {
		field := values.Field(i)

		if field.Type() == itemType {
			envOptions = append(envOptions, field.Interface().(env.Option))
		}
	}

	return envOptions
}
