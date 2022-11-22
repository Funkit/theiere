package main

import (
	"fmt"
	"github.com/Funkit/crispy-engine/fancytext"
	"github.com/Funkit/crispy-engine/frame"
	"github.com/Funkit/crispy-engine/menu"
	"github.com/Funkit/crispy-engine/subframe"
	"github.com/Funkit/crispy-engine/tabs"
	"github.com/Funkit/crispy-engine/validation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

func main() {
	comm := make(chan struct{})
	items := []menu.ListItem{
		generateItem1(),
		generateItem2(comm),
		generateItem3(),
	}

	l, err := menu.New("This is a menu title", items)
	if err != nil {
		panic(err)
	}

	f, err := frame.New(frame.WithComponent(&l), frame.WithHorizontalAlignment(lipgloss.Left), frame.WithBorder())
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(f, tea.WithAltScreen())

	go func(ch <-chan struct{}) {
		select {
		case <-comm:
			p.Send(tea.Quit())
		}
	}(comm)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("Thank you for using this tool !")
}

func generateItem1() menu.ListItem {
	item1, err := fancytext.New("HELLO")
	if err != nil {
		panic(err)
	}

	return menu.ListItem{
		Item:      menu.NewItem("fancytext", "Full size text, default values"),
		Component: &item1,
	}
}

func generateItem2(comm chan struct{}) menu.ListItem {
	vld, err := validation.New(validation.WithChannel(comm))

	item2, err := subframe.New(subframe.WithComponent(&vld),
		subframe.WithHorizontalAlignment(lipgloss.Center),
		subframe.WithVerticalAlignment(lipgloss.Center))
	if err != nil {
		panic(err)
	}

	return menu.ListItem{
		Item:      menu.NewItem("validation", "Full size yes/no validation screen, default values"),
		Component: &item2,
	}
}

func generateItem3() menu.ListItem {
	tab1, err := tabs.NewTab("Tab 1")
	if err != nil {
		panic(err)
	}
	tab2, err := tabs.NewTab("Tab 2")
	if err != nil {
		panic(err)
	}
	tab3, err := tabs.NewTab("Tab 3")
	if err != nil {
		panic(err)
	}

	contentTabs := []tabs.Tab{
		tab1,
		tab2,
		tab3,
	}

	subTabs, err := tabs.New(contentTabs)
	if err != nil {
		panic(err)
	}

	subf, err := subframe.New(subframe.WithComponent(subTabs),
		subframe.WithHorizontalAlignment(lipgloss.Left),
		subframe.WithVerticalAlignment(lipgloss.Top))

	return menu.ListItem{
		Item:      menu.NewItem("tabs", "tabulations with content in each tab"),
		Component: &subf,
	}
}
