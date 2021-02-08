package commands

import (
	"fmt"
	"gtmhub-cli/model"
	"gtmhub-cli/output"
	"strconv"
	"time"

	//"gtmhub-cli/output"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

var (
	UpdateCommand = &cli.Command{
		Name:   "update",
		Usage:  "lets you update your krs. Specify a specific kr using --id or select from the list instead!",
		Action: UpdateAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "specifies the id of the metric to be updated. Use gtmhub show krs to get a list of krs.",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "comment",
				Aliases:  []string{"c"},
				Usage:    "specifies a comment to be associated with this checkin",
				Required: false,
			},
			&cli.Float64Flag{
				Name:     "value",
				Aliases:  []string{"v"},
				Usage:    "specifies the value for this update",
				Required: false,
			},
		},
	}
)

func UpdateAction(c *cli.Context) error {
	now := time.Now()
	id := c.String("id")
	var value *float64
	if c.IsSet("value") {
		inputValue := c.Float64("value")
		value = &inputValue
	}
	comment := c.String("comment")

	interactiveId, valueFromInteractive, commentFromInteractive, err := getIdInteractively(value, comment, id)
	if err != nil {
		return err
	}

	id = interactiveId
	value = valueFromInteractive
	comment = commentFromInteractive

	if len(id) == 0 {
		return fmt.Errorf("no id provided with the --id option and non selected from the list")
	}

	req := model.CheckInMetricRequest{
		Actual:      *value,
		CheckInDate: &now,
		Comment:     comment,
	}

	err = client.UpdateMetric(req, id)
	if err != nil {
		return err
	}

	output.Print(metricUpdatedMsg, output.Green)

	return nil
}

func getIdInteractively(value *float64, comment, metricID string) (string, *float64, string, error) {
	if value != nil && len(metricID) > 0 && len(comment) > 0 {
		return metricID, value, comment, nil
	}
	model := GetInteractiveMetricsModel(value, comment, metricID)
	p := tea.NewProgram(&model)

	if err := p.Start(); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)

	}
	if model.err != nil {
		return "", value, "", model.err
	}

	if model.cursor < 0 {
		// no choice was made
		return "", value, comment, nil
	}

	valueEntered := model.valueInput.Value()
	parsed, err := getFloat(valueEntered)
	if err != nil {
		return "", value, comment, err
	}

	return model.metricID, &parsed, model.commentInput.Value(), nil
}

func getFloat(entered string) (float64, error) {
	return strconv.ParseFloat(entered, 64)
}
