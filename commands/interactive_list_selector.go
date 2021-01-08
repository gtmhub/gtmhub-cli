package commands

import (
	"fmt"
	"gtmhub-cli/model"

	tea "github.com/charmbracelet/bubbletea"
)

type interactiveListSelector struct {
	lists        model.FullListResponse
	selectedList *model.ListResponse
	cursor       int
	err          error
}

func (il *interactiveListSelector) Init() tea.Cmd {
	return getLists
}

func (il *interactiveListSelector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		il.err = msg
		return il, tea.Quit
	case model.FullListResponse:
		il.lists = msg
		return il, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if il.cursor > 0 {
				il.cursor--
			}
		case "down", "j":
			if il.cursor < len(il.lists.Items)-1 {
				il.cursor++
			}
		case "q", "ctrl+c":
			il.cursor = -1
			return il, tea.Quit
		case "enter":
			il.selectedList = &il.lists.Items[il.cursor]
			return il, tea.Quit
		}

	}
	return il, nil
}

func (il *interactiveListSelector) View() string {
	s := "Which list would you like to see?\n"
	for i, choice := range il.lists.Items {

		cursor := " " // no cursor
		if il.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	s += "\nPress q to quit.\n"

	return s
}

func getLists() tea.Msg {
	metrics, err := client.GetAllLists()
	if err != nil {
		return err
	}

	return metrics
}
