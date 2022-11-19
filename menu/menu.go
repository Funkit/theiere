package menu

import (
	"github.com/Funkit/crispy-engine/subview"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Model This is a list built for selecting various subviews.
// Once a view is selected,the selected subview can trigger coming back
// to the menu by returning common.TreeUp as a tea.Msg when updating.
type Model struct {
	list      list.Model
	choice    string
	SubViews  map[string]subview.Model
	fixedSize bool
}

type ListItem struct {
	Item      Item
	Component subview.Model
}

type options struct {
	width     *int
	height    *int
	fixedSize bool
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

func WithFixedSize() Option {
	return func(options *options) error {
		options.fixedSize = true

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

	width := 50
	if options.width != nil {
		width = *options.width
	}

	height := 20
	if options.height != nil {
		height = *options.height
	}

	subViews := make(map[string]subview.Model)

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
		list:      l,
		SubViews:  subViews,
		fixedSize: options.fixedSize,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	var commands []tea.Cmd
	for key := range m.SubViews {
		commands = append(commands, m.SubViews[key].Init())
	}

	return tea.Batch(commands...)
}

func (m *Model) Update(msg tea.Msg) (subview.Model, tea.Cmd) {
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
			case "q", "esc", "ctrl+c":
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	if m.choice != "" {
		m.SubViews[m.choice], cmd = m.SubViews[m.choice].Update(msg)
		switch msg.(type) {
		case subview.TreeUp:
			m.choice = ""
			for key := range m.SubViews {
				m.SubViews[key].Reset()
			}
			return m, nil
		}
		return m, cmd
	}

	return m, cmd
}

func (m *Model) View() string {
	if m.choice != "" {
		return m.SubViews[m.choice].View()
	}

	return m.list.View()
}

func (m *Model) SetWidth(width int) {
	if !m.fixedSize {
		m.list.SetWidth(width)
		for key := range m.SubViews {
			m.SubViews[key].SetWidth(width)
		}
	}
}

func (m *Model) SetHeight(height int) {
	if !m.fixedSize {
		m.list.SetHeight(height)
		for key := range m.SubViews {
			m.SubViews[key].SetHeight(height)
		}
	}
}

func (m *Model) Reset() {
	m.choice = ""
	for key := range m.SubViews {
		m.SubViews[key].Reset()
	}
}
