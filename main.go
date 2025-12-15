package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JGugino/pbterm/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	term := tui.CreateNewPBTerm()

	p := tea.NewProgram(term)
	if _, err := p.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
