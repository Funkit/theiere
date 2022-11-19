package subview

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	SetWidth(width int)
	SetHeight(height int)
	View() string
	Init() tea.Cmd
	Update(msg tea.Msg) (Model, tea.Cmd)
	Reset() // Reset the view when going up in the component tree
}

type TreeUp struct {
	bool
}

func GoUp() tea.Msg {
	return TreeUp{}
}
