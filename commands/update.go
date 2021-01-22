package commands

import (
	"fmt"
	"gtmhub-cli/input"
	"gtmhub-cli/model"
	"gtmhub-cli/output"
	"os"
	"time"

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
	value := c.Float64("value")
	comment := c.String("comment")

	req := model.CheckInMetricRequest{
		Actual:      value,
		CheckInDate: &now,
		Comment:     comment,
	}

	if c.IsSet("value") == false {
		val, err := input.GetFloat("You should provide a value for your update: ")
		if err != nil {
			return err
		}

		req.Actual = val
	}

	if len(id) == 0 {
		interactiveId, err := getIdInteractively()
		if err != nil {
			return err
		}

		id = interactiveId
	}

	if len(id) == 0 {
		return fmt.Errorf("no id provided with the --id option and non selected from the list")
	}

	err := client.UpdateMetric(req, id)
	if err != nil {
		return err
	}

	output.Print(metricUpdatedMsg, output.Green)

	return nil
}

func getIdInteractively() (string, error) {
	model := interactiveMetrics{}
	p := tea.NewProgram(&model)

	if err := p.Start(); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)

	}
	if model.err != nil {
		return "", model.err
	}

	if model.cursor < 0 {
		// no choice was made
		return "", nil
	}

	metric := model.metrics[model.cursor]

	return metric.ID, nil
}

