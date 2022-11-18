package main

import (
	"fmt"
	"github.com/Funkit/crispy-engine/fancytext"
	"github.com/Funkit/crispy-engine/frame"
	"github.com/Funkit/crispy-engine/menu"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

func main() {

	items := []menu.ListItem{
		{
			Item: menu.NewItem("Book a Pass", "Select GS and RF channel to book"),
			Component: fancytext.Model{
				Content: "HELLO",
			},
		},
		{
			Item: menu.NewItem("Manage Plans", "List plans by status"),
			Component: fancytext.Model{
				Content: "WORLD",
			},
		},
		{
			Item: menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: fancytext.Model{
				Content: "TOTO",
			},
		},
	}

	l, err := menu.New("This is a menu", items)
	if err != nil {
		panic(err)
	}

	f, err := frame.New(frame.WithComponent(l), frame.WithAlignment(lipgloss.Left))
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(f, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("Thank you for using nano-MCS !")
}
