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
	callbackFn     func() tea.Cmd
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
	return width - 2
}

func (m Model) GetVisible() bool {
	return m.visible
}

func (m *Model) getOptions() []option.Model {
	index := 0

	defaultOptions := []string{}
	for _, element := range m.defaultOptions {
		defaultOptions = append(defaultOptions, element.Msg+"\n"+element.AltMsg+"\n"+element.Description)
	}

	input := list.FormatInput(m.list.TextInput.Value())
	finds := defaultOptions

	if len(input) != 0 {
		finds = list.Filter(defaultOptions, input)
	}

	var options []option.Model
	for len(options) < m.list.NewSize() && index < len(finds) {
		parts := strings.Split(finds[index], "\n")
		option := option.InitialModel(m.width, env.Option{Msg: parts[0], AltMsg: parts[1], Description: parts[2]})
		options = append(options, option)

		index++
	}

	if m.list.TextInput.Value() != "" {
		m.total = len(options)
	}

	return options
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
