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
	fixedSize   bool
}

type options struct {
	height              *int
	width               *int
	component           *common.SubView
	border              bool
	borderColor         *lipgloss.AdaptiveColor
	horizontalAlignment *lipgloss.Position
	verticalAlignment   *lipgloss.Position
	fixedSize           *bool
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
func WithComponent(content common.SubView) Option {
	return func(options *options) error {
		options.component = &content

		return nil
	}
}

func WithBorder() Option {
	return func(options *options) error {
		options.border = true

		return nil
	}
}

func WithBorderColor(color lipgloss.AdaptiveColor) Option {
	return func(options *options) error {
		options.borderColor = &color

		return nil
	}
}

func WithHorizontalAlignment(alignment lipgloss.Position) Option {
	return func(options *options) error {
		options.horizontalAlignment = &alignment

		return nil
	}
}

func WithVerticalAlignment(alignment lipgloss.Position) Option {
	return func(options *options) error {
		options.verticalAlignment = &alignment

		return nil
	}
}

func WithFixedSize() Option {
	return func(options *options) error {
		isFixed := true
		options.fixedSize = &isFixed

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

	width := 30
	if options.width != nil {
		width = *options.width
	}

	height := 30
	if options.height != nil {
		height = *options.height
	}

	fixedSize := false
	if options.fixedSize != nil {
		fixedSize = *options.fixedSize
	}

	horizontalAlignment := lipgloss.Center
	if options.horizontalAlignment != nil {
		horizontalAlignment = *options.horizontalAlignment
	}

	verticalAlignment := lipgloss.Center
	if options.verticalAlignment != nil {
		verticalAlignment = *options.verticalAlignment
	}

	style := lipgloss.NewStyle().
		Width(width).
		Height(height).
		AlignHorizontal(horizontalAlignment).AlignVertical(verticalAlignment)
	if options.border {
		color := lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
		if options.borderColor != nil {
			color = *options.borderColor
		}
		style = style.BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(color)
	}

	m := Model{
		Style:     style,
		fixedSize: fixedSize,
	}

	if options.component != nil {
		m.hasContent = true
		m.Content = *options.component
		m.Content.SetWidth(width - 2)
		m.Content.SetHeight(height - 2)
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
		if !m.fixedSize {
			m.Style.Width(recv.Width - 2)
			m.Style.Height(recv.Height - 2)
			m.Content.SetHeight(recv.Height - 2)
			m.Content.SetWidth(recv.Width - 2)
		} else {
			m.Content.SetHeight(m.Style.GetHeight())
			m.Content.SetWidth(m.Style.GetWidth())
		}
		return m, nil
	case tea.KeyMsg:
		switch recv.String() {
		case "ctrl+c":
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
