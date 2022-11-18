package fancytext

import (
	"github.com/Funkit/crispy-engine/common"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Content string
	Style   lipgloss.Style
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	//if style was not set at creation, use default value
	if m.Style.Value() == "" {
		defaultStyle := lipgloss.NewStyle().
			Bold(true).
			Padding(1).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4"))
		return defaultStyle.Render(m.Content)
	} else {
		return m.Style.Render(m.Content)
	}
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
	m.Style.Width(width)
}

func (m Model) SetHeight(height int) {

}
