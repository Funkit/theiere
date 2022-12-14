package validation

import (
	"github.com/Funkit/theiere/executor"
	"github.com/Funkit/theiere/subview"
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
	exec          executor.Model
	execEnabled   bool
	//this channel must be initialized outside this model
	clientCom chan<- struct{}
}

type options struct {
	width     *int
	height    *int
	clientCom chan<- struct{}
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

func WithChannel(signalChan chan<- struct{}) Option {
	return func(options *options) error {
		options.clientCom = signalChan

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

	exec, err := executor.New()
	if err != nil {
		return Model{}, err
	}

	return Model{
		width:     width,
		height:    height,
		exec:      exec,
		clientCom: options.clientCom,
		frameStyle: lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center),
	}, nil
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (subview.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right":
			m.buttonPos = !m.buttonPos
		case "enter":
			if m.execEnabled {

				recv, cmd := m.exec.Update(msg)
				if ex, ok := recv.(*executor.Model); ok {
					m.exec = *ex
				}

				return m, cmd
			}
			if !m.buttonPos {
				return m, subview.GoUp
			}
			if m.clientCom != nil {
				m.clientCom <- struct{}{}
			}
			return m, Proceed()
		case "q", "esc":
			return m, subview.GoUp
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

func Proceed() tea.Cmd {
	return func() tea.Msg {
		return Status{Validated: true}
	}
}

func (m *Model) SetWidth(width int) {
	m.width = width
}

func (m *Model) SetHeight(height int) {
	m.height = height
}

func (m *Model) Reset() {
	m.buttonPos = false
	m.execEnabled = false
	m.exec.Reset()
}
