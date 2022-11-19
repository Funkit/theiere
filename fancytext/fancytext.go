package fancytext

import (
	"github.com/Funkit/crispy-engine/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Content   string
	Style     lipgloss.Style
	fixedSize bool
}

type options struct {
	height    *int
	width     *int
	style     *lipgloss.Style
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
		options.width = &height

		return nil
	}
}

func WithStyle(style lipgloss.Style) Option {
	return func(options *options) error {
		options.style = &style

		return nil
	}
}

func WithFixedSize() Option {
	return func(options *options) error {
		options.fixedSize = true

		return nil
	}
}

func New(content string, opts ...Option) (Model, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return Model{}, err
		}
	}

	width := 30
	if options.width != nil {
		width = *options.width
	}

	height := 30
	if options.height != nil {
		height = *options.height
	}

	style := lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))
	if options.style != nil {
		style = *options.style
	}

	style.Height(height)
	style.Width(width)

	return Model{
		Content:   content,
		Style:     style,
		fixedSize: options.fixedSize,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.Style.Render(m.Content)
}

func (m Model) Update(msg tea.Msg) (common.SubView, tea.Cmd) {
	switch val := msg.(type) {
	case tea.KeyMsg:
		switch val.String() {
		case "esc", "q":
			return m, common.GoUp
		}
	}

	return m, nil
}

func (m Model) SetWidth(width int) {
	if !m.fixedSize {
		m.Style.Width(width)
	}

}

func (m Model) SetHeight(height int) {
	if !m.fixedSize {
		m.Style.Height(height)
	}
}
