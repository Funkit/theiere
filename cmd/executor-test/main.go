package main

import (
	"fmt"
	"github.com/Funkit/theiere/executor"
	"github.com/Funkit/theiere/frame"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"time"
)

func main() {
	exec, err := executor.New()
	if err != nil {
		log.Fatal(err)
	}

	f, err := frame.New(frame.WithComponent(&exec))
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(f, tea.WithAltScreen())

	go func() {
		time.Sleep(time.Second * 5)
		p.Send(executor.Message{
			Success:     false,
			Description: "This function has failed",
		})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	fmt.Println("Thank you for using this tool !")
}
