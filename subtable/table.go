package subtable

import (
	"errors"
	"github.com/Funkit/crispy-engine/subview"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	helpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type Model struct {
	Table         table.Model
	KeyMap        KeyMap
	Help          help.Model
	helpEnabled   bool
	initCmd       func() tea.Cmd
	columns       []table.Column
	rows          []table.Row
	tableStyle    table.Styles
	height, width int
}

type options struct {
	columns     []table.Column
	rows        []table.Row
	width       *int
	focusColor  *lipgloss.Color
	helpEnabled bool
	initCmd     func() tea.Cmd
	keyMap      *KeyMap
}

type Option func(options *options) error

func WithColumns(cols []table.Column) Option {
	return func(options *options) error {
		options.columns = cols
		return nil
	}
}

func WithRows(rows []table.Row) Option {
	return func(options *options) error {
		options.rows = rows
		return nil
	}
}

func WithWidth(width int) Option {
	return func(options *options) error {
		if width <= 0 {
			return errors.New("invalid table width")
		}
		options.width = &width
		return nil
	}
}

func WithColor(color lipgloss.Color) Option {
	return func(options *options) error {
		options.focusColor = &color
		return nil
	}
}

func WithHelpDisplayed() Option {
	return func(options *options) error {
		options.helpEnabled = true
		return nil
	}
}

func WithKeyMap(km KeyMap) Option {
	return func(options *options) error {
		options.keyMap = &km
		return nil
	}
}

func WithInitCmd(initCmd func() tea.Cmd) Option {
	return func(options *options) error {
		options.initCmd = initCmd
		return nil
	}
}

func New(opts ...Option) (Model, error) {

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

	adjustColumnWidth(options.columns, width)

	t := table.New(
		table.WithColumns(options.columns),
		table.WithRows(options.rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)

	selectColor := lipgloss.Color("212")
	if options.focusColor != nil {
		selectColor = *options.focusColor
	}

	s := table.Styles{
		Selected: lipgloss.NewStyle().Bold(true).Foreground(selectColor),
		Header:   lipgloss.NewStyle().Bold(true).Padding(0, 1),
		Cell:     lipgloss.NewStyle().Padding(0, 1),
	}
	t.SetStyles(s)

	var km KeyMap

	if options.keyMap == nil {
		km = DefaultKeyMap()

	} else {
		km = *options.keyMap
	}
	t.KeyMap = km.asInternalTableMap()

	return Model{
		Table:       t,
		KeyMap:      km,
		Help:        help.New(),
		helpEnabled: options.helpEnabled,
		initCmd:     options.initCmd,
		columns:     options.columns,
		rows:        options.rows,
		tableStyle:  s,
		height:      height,
		width:       width,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	if m.initCmd != nil {
		return m.initCmd()
	}
	return nil
}

func (m *Model) Update(msg tea.Msg) (subview.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, subview.GoUp
		}
	}
	var cmd tea.Cmd
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.helpEnabled {
		return lipgloss.JoinVertical(lipgloss.Left, baseStyle.Render(m.Table.View()), helpStyle.Render(m.Help.View(m.KeyMap)))
	}

	return baseStyle.Render(m.Table.View()) + "\n"
}

func (m *Model) SetHeight(height int) {
	m.height = height - 4
	//m.height = 25

	adjustColumnWidth(m.columns, m.width)

	t := table.New(
		table.WithColumns(m.columns),
		table.WithRows(m.rows),
		table.WithFocused(true),
		table.WithHeight(m.height),
	)

	t.SetStyles(m.tableStyle)

	m.Table = t
}

func (m *Model) SetWidth(width int) {
	m.width = width - 10
	//m.width = 90

	adjustColumnWidth(m.columns, m.width)

	t := table.New(
		table.WithColumns(m.columns),
		table.WithRows(m.rows),
		table.WithFocused(true),
		table.WithHeight(m.height),
	)

	t.SetStyles(m.tableStyle)

	m.Table = t
}

func (m *Model) Reset() {
	m.Table.SetCursor(0)
}

func adjustColumnWidth(col []table.Column, maxWidth int) {
	sumWidth := 0
	for i := 0; i < len(col); i++ {
		sumWidth += col[i].Width
	}

	for i := 0; i < len(col); i++ {
		var f float64
		f = float64(col[i].Width) / float64(sumWidth) * float64(maxWidth)
		col[i].Width = int(f)
	}
}
