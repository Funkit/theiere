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

	text1, err := fancytext.New("HELLO")
	if err != nil {
		panic(err)
	}
	text2, err := fancytext.New("WORLD")
	if err != nil {
		panic(err)
	}
	text3, err := fancytext.New("HOW ARE YOU")
	if err != nil {
		panic(err)
	}

	items := []menu.ListItem{
		{
			Item:      menu.NewItem("Book a Pass", "Select GS and RF channel to book"),
			Component: text1,
		},
		{
			Item:      menu.NewItem("Manage Plans", "List plans by status"),
			Component: text2,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
		{
			Item:      menu.NewItem("Cancel Plan", "Cancel an upcoming plan"),
			Component: text3,
		},
	}

	l, err := menu.New("This is a menu", items)
	if err != nil {
		panic(err)
	}

	f, err := frame.New(frame.WithComponent(l), frame.WithAlignment(lipgloss.Left), frame.WithBorder())
	if err != nil {
		panic(err)
	}

	if _, err := tea.NewProgram(f, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("Thank you for using nano-MCS !")
}
