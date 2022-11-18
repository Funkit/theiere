package menu

import (
	"github.com/Funkit/crispy-engine/common"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model This is a list
type Model struct {
	list     list.Model
	choice   string
	quitting bool
	SubViews map[string]common.SubView
}

type ListItem struct {
	Item      Item
	Component common.SubView
}

type options struct {
	width  *int
	height *int
}

type Option func(options *options) error

func WithWidth(width int) Option {
	return func(options *options) error {
		options.width = &width

		return nil
	}
}

func WithHeight(height int) Option {
	return func(options *options) error {
		options.height = &height

		return nil
	}
}

func New(title string, items []ListItem, opts ...Option) (Model, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return Model{}, err
		}
	}

	width := 80
	if options.width != nil {
		width = *options.width
	}

	height := 20
	if options.height != nil {
		height = *options.height
	}

	subViews := make(map[string]common.SubView)

	var teaList []list.Item
	for _, item := range items {
		subViews[item.Item.title] = item.Component
		teaList = append(teaList, item.Item)
	}

	l := list.New(teaList, list.NewDefaultDelegate(), width, height)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return Model{
		list:     l,
		SubViews: subViews,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (common.SubView, tea.Cmd) {
	if m.choice == "" {
		m.list, _ = m.list.Update(msg)
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
			case "enter":
				i, ok := m.list.SelectedItem().(Item)
				if ok {
					m.choice = i.Title()
				}

				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	if m.choice != "" {
		m.SubViews[m.choice], cmd = m.SubViews[m.choice].Update(msg)
		switch msg.(type) {
		case common.TreeUp:
			m.choice = ""
			return m, nil
		}
		return m, cmd
	}

	return m, cmd
}

func (m Model) View() string {
	if m.choice != "" {
		return m.SubViews[m.choice].View()
	}

	return m.list.View()
}

func (m Model) SetWidth(width int) {
	m.list.SetWidth(width)
	for key := range m.SubViews {
		m.SubViews[key].SetWidth(width)
	}
}

func (m Model) SetHeight(height int) {
	m.list.SetHeight(height)
	for key := range m.SubViews {
		m.SubViews[key].SetHeight(height)
	}
}
