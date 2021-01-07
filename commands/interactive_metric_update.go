package commands

import (
	"fmt"
	"gtmhub-cli/model"

	tea "github.com/charmbracelet/bubbletea"
)

type interactiveMetrics struct {
	metrics model.Metrics
	comment string
	value   *float64
	cursor  int
	err     error
}

func (im *interactiveMetrics) Init() tea.Cmd {
	return getMetrics
}

func (im *interactiveMetrics) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		im.err = msg
		return im, tea.Quit
	case model.Metrics:
		im.metrics = msg
		return im, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if im.cursor > 0 {
				im.cursor--
			}
		case "down", "j":
			if im.cursor < len(im.metrics)-1 {
				im.cursor++
			}
		case "q", "ctrl+c":
			im.cursor = -1
			return im, tea.Quit
		case "enter":
			return im, tea.Quit
		}

	}
	return im, nil
}

func (im *interactiveMetrics) View() string {
	s := "And to which metric does this value belong?\n\n"

	// Iterate over our choices
	for i, choice := range im.metrics {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if im.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.Name)
	}

	// The footer
	//s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func getMetrics() tea.Msg {
	metrics, err := client.GetMetricsInCurrentSession()
	if err != nil {
		return err
	}

	return metrics
}
