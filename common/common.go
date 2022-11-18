package common

import tea "github.com/charmbracelet/bubbletea"

type SubView interface {
	SetWidth(width int)
	SetHeight(height int)
	View() string
	Init() tea.Cmd
	Update(msg tea.Msg) (SubView, tea.Cmd)
}

type TreeUp struct {
	bool
}

func GoUp() tea.Msg {
	return TreeUp{}
}
