package tui

import (
	"github.com/JGugino/pbterm/pb"
	tea "github.com/charmbracelet/bubbletea"
)

type PocketBase struct {
	Auth        pb.PBAuth
	Records     pb.PBRecord
	Collections pb.PBCollection
}

type PBTerm struct {
	Id string
	PB []PocketBase
}

func CreateNewPBTerm() PBTerm {
	return PBTerm{
		PB: make([]PocketBase, 0),
	}
}

func (term PBTerm) Init() tea.Cmd {
	return nil
}

func (term PBTerm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return term, tea.Quit
		}
	}

	return term, nil
}
func (term PBTerm) View() string {
	return "Hello World"
}
