package commands

import (
	"fmt"
	"gtmhub-cli/model"
	"gtmhub-cli/output"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
)

const focusedTextColor = "205"

var (
	color               = te.ColorProfile().Color
	focusedPrompt       = te.String("> ").Foreground(color("205")).String()
	blurredPrompt       = "> "
	focusedSubmitButton = "[ " + te.String("Submit").Foreground(color("205")).String() + " ]"
	blurredSubmitButton = "[ " + te.String("Submit").Foreground(color("240")).String() + " ]"
)

type interactiveMetrics struct {
	metrics      model.Metrics
	metricID     string
	metricChosen bool
	comment      string
	commentInput textinput.Model
	valueInput   textinput.Model
	value        *float64
	cursor       int
	err          error
	inputIndex   int
	submitButton string
}

func GetInteractiveMetricsModel(presentValue *float64, presentComment string, metricID string) interactiveMetrics{
	value := textinput.NewModel()
	value.Placeholder = "Value for this update? "
	value.Focus()
	value.Prompt = focusedPrompt
	value.TextColor = focusedTextColor
	if presentValue != nil {
		value.SetValue(strconv.FormatFloat(*presentValue, 'f', 1, 64))
	}

	comment := textinput.NewModel()
	comment.Placeholder = "Your metric update comment? (its ok if you leave it empty)"
	comment.Prompt = blurredPrompt
	comment.SetValue(presentComment)

	model := interactiveMetrics{}
	model.valueInput = value
	model.commentInput = comment

	model.metricID = metricID


	return model
}

func (im *interactiveMetrics) Init() tea.Cmd {
	if len(im.metricID) > 0 {
		im.metricChosen = true
		return nil
	}
	return getMetrics
}

func (im *interactiveMetrics) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if im.metricChosen == false {
		return chooseMetricUpdater(msg, im)
	}

	return commentAndValueUpdater(msg, im)
}

func commentAndValueUpdater(msg tea.Msg, im *interactiveMetrics) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return im, tea.Quit

		// Cycle between inputs
		case "tab", "shift+tab", "enter", "up", "down":

			inputs := []textinput.Model{
				im.valueInput,
				im.commentInput,
			}

			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && im.inputIndex == len(inputs) {
				return im, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				im.inputIndex--
			} else {
				im.inputIndex++
			}

			if im.inputIndex > len(inputs) {
				im.inputIndex = 0
			} else if im.inputIndex < 0 {
				im.inputIndex = len(inputs)
			}

			for i := 0; i <= len(inputs)-1; i++ {
				if i == im.inputIndex {
					// Set focused state
					inputs[i].Focus()
					inputs[i].Prompt = focusedPrompt
					continue
				}
				// Remove focused state
				inputs[i].Blur()
				inputs[i].Prompt = blurredPrompt
				inputs[i].TextColor = ""
			}

			im.valueInput = inputs[0]
			im.commentInput = inputs[1]

			if im.inputIndex == len(inputs) {
				im.submitButton = focusedSubmitButton
			} else {
				im.submitButton = blurredSubmitButton
			}

			return im, nil
		}
	}

	// Handle character input and blinks
	im, cmd = updateInputs(msg, im)
	return im, cmd
}

// Pass messages and models through to text input components. Only text inputs
// with Focus() set will respond, so it's safe to simply update all of them
// here without any further logic.
func updateInputs(msg tea.Msg, im *interactiveMetrics) (*interactiveMetrics, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	im.valueInput, cmd = im.valueInput.Update(msg)
	cmds = append(cmds, cmd)

	im.commentInput, cmd = im.commentInput.Update(msg)
	cmds = append(cmds, cmd)

	//m.emailInput, cmd = m.emailInput.Update(msg)
	//cmds = append(cmds, cmd)
	//
	//m.passwordInput, cmd = m.passwordInput.Update(msg)
	//cmds = append(cmds, cmd)

	return im, tea.Batch(cmds...)
}

func chooseMetricUpdater(msg tea.Msg, im *interactiveMetrics) (tea.Model, tea.Cmd) {
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
			selectedMetric := im.metrics[im.cursor]
			im.metricID = selectedMetric.ID
			if len(im.valueInput.Value()) > 0 && len(im.commentInput.Value()) >0 {
				return im, tea.Quit
			}

			im.metricChosen = true
			return im, nil
		}

	}
	return im, nil
}

func (im *interactiveMetrics) View() string {
	if im.metricChosen == false {
		return selectMetricView(im)
	}

	return commentAndValueView(im)
}

func selectMetricView(im *interactiveMetrics) string {
	if len(im.metrics) == 0 {
		return "Loading ..."
	}

	return output.GetMetricSelectionTable(im.metrics, im.cursor)
}

func commentAndValueView(im *interactiveMetrics) string {

	format := "The value you want to update with:\n%s\nThe comment you would like to leave with this update:\n%s\n\n\n%s"

	//s := "\n"

	//inputs := []string{
	//	im.valueInput.View(),
	//	im.commentInput.View(),
	//	//m.emailInput.View(),
	//	//m.passwordInput.View(),
	//}
	//
	//for i := 0; i < len(inputs); i++ {
	//	s += inputs[i]
	//	if i < len(inputs)-1 {
	//		s += "\n"
	//	}
	//}

	return fmt.Sprintf(format, im.valueInput.View(), im.commentInput.View(), im.submitButton)

	//s += "\n\n" + im.submitButton + "\n"
	//return s
}

func getMetrics() tea.Msg {
	metrics, err := client.GetMetricsInCurrentSession()
	if err != nil {
		return err
	}

	return metrics
}
