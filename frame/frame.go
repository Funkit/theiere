package frame

import (
	"github.com/Funkit/crispy-engine/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	borderColor lipgloss.AdaptiveColor
	Style       lipgloss.Style
	Content     common.SubView
	hasContent  bool
}

type options struct {
	height      *int
	width       *int
	component   *common.SubView
	border      bool
	borderColor *lipgloss.AdaptiveColor
	alignment   *lipgloss.Position
}

type Options func(options *options) error

func WithWidth(width int) Options {
	return func(options *options) error {
		options.width = &width

		return nil
	}
}
func WithHeight(height int) Options {
	return func(options *options) error {
		options.height = &height

		return nil
	}
}
func WithComponent(content common.SubView) Options {
	return func(options *options) error {
		options.component = &content

		return nil
	}
}

func WithBorder() Options {
	return func(options *options) error {
		options.border = true

		return nil
	}
}

func WithAlignment(alignment lipgloss.Position) Options {
	return func(options *options) error {
		options.alignment = &alignment

		return nil
	}
}

func New(opts ...Options) (Model, error) {

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

	alignment := lipgloss.Center
	if options.alignment != nil {
		alignment = *options.alignment
	}

	style := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(alignment)
	if options.border {
		color := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
		if options.borderColor != nil {
			color = *options.borderColor
		}
		style = style.BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(color)
	}

	m := Model{
		Style: style,
	}

	if options.component != nil {
		m.hasContent = true
		m.Content = *options.component
		m.Content.SetWidth(width)
		m.Content.SetHeight(height)
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return m.Content.Init()
}

func (m Model) View() string {
	if m.hasContent {
		return m.Style.Render(m.Content.View())
	}
	return m.Style.Render("")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch recv := msg.(type) {
	case tea.WindowSizeMsg:
		m.Style.Width(recv.Width)
		m.Style.Height(recv.Height)
		m.Content.SetHeight(recv.Height)
		m.Content.SetWidth(recv.Width)
		return m, nil
	case tea.KeyMsg:
		cmd := recv.String()
		switch cmd {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	if m.hasContent {
		var cmd tea.Cmd
		m.Content, cmd = m.Content.Update(msg)

		return m, cmd
	}

	return m, nil
}
