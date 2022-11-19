package validation

import (
	"github.com/Funkit/crispy-engine/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1).MarginRight(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginTop(1).
				Underline(true)
)

type Model struct {
	buttonPos     bool
	width, height int
	frameStyle    lipgloss.Style
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

	height := 40
	if options.height != nil {
		height = *options.height
	}

	return Model{
		width:  width,
		height: height,
		frameStyle: lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center),
	}, nil
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (common.SubView, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right":
			m.buttonPos = !m.buttonPos
		case "enter":
			return m, ReturnStatus(m.buttonPos)
		case "q", "esc":
			return m, common.GoUp
		}
	}
	return m, nil
}

func (m *Model) View() string {
	okButton := ""
	cancelButton := ""
	if m.buttonPos {
		okButton = activeButtonStyle.Render("Yes")
		cancelButton = buttonStyle.Render("No")
	} else {
		okButton = buttonStyle.Render("Yes")
		cancelButton = activeButtonStyle.Render("No")
	}

	question := lipgloss.NewStyle().Width(40).MarginBottom(1).Align(lipgloss.Center).Render("Do you confirm your choice ?")
	buttons := lipgloss.JoinHorizontal(lipgloss.Center, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

type Status struct {
	Validated bool
}

func ReturnStatus(proceed bool) tea.Cmd {
	if proceed {
		return func() tea.Msg {
			return Status{Validated: proceed}
		}
	}

	return common.GoUp
}

func (m *Model) SetWidth(width int) {
	m.width = width
}
func (m *Model) SetHeight(height int) {
	m.height = height
}
