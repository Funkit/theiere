package executor

import (
	"github.com/Funkit/theiere/subview"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Message struct {
	Success     bool
	Description string
}

type Model struct {
	spinner           spinner.Model
	successStyle      lipgloss.Style
	failStyle         lipgloss.Style
	container         lipgloss.Style
	GotResult         bool
	success           bool
	resultDescription string
}

type options struct {
	successStyle *lipgloss.Style
	failStyle    *lipgloss.Style
}

type Option func(option *options) error

func WithSuccessStyle(successStyle lipgloss.Style) Option {
	return func(options *options) error {
		options.successStyle = &successStyle

		return nil
	}
}

func WithFailStyle(failStyle lipgloss.Style) Option {
	return func(options *options) error {
		options.failStyle = &failStyle

		return nil
	}
}

func New(opts ...Option) (Model, error) {
	var options options
	for _, opt := range opts {
		if err := opt(&options); err != nil {
			return Model{}, err
		}
	}

	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#6AA84F"))
	if options.successStyle != nil {
		successStyle = *options.successStyle
	}

	failStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#F44336"))
	if options.failStyle != nil {
		failStyle = *options.failStyle
	}

	return Model{
		spinner:      spinner.New(),
		successStyle: successStyle,
		failStyle:    failStyle,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *Model) Update(msg tea.Msg) (subview.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, subview.GoUp
		}
	case Message:
		m.GotResult = true
		m.success = msg.Success
		m.resultDescription = msg.Description
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m *Model) View() string {
	if !m.GotResult {
		return m.spinner.View() + " Processing in progress..."
	}
	if m.success {
		return m.container.Render(lipgloss.JoinVertical(lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Center,
				"Status: ", m.successStyle.Render("SUCCESS")),
			m.resultDescription))
	}
	return m.container.Render(lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center,
			"Status: ", m.failStyle.Render("FAIL")),
		m.resultDescription))
}

func (m *Model) SetWidth(width int) {
	m.container.Width(width)
}

func (m *Model) SetHeight(height int) {
	m.container.Height(height)
}
func (m *Model) Reset() {
	m.success = false
}
