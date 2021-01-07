package commands

import (
	"gtmhub-cli/output"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	StatusCommand = &cli.Command{
		Name:  "status",
		Usage: "shows you the items that currently need your attention",
		Action: StatusAction,
	}
)

func StatusAction(c *cli.Context) error {
	krs, err := client.GetMetricsInCurrentSession()
	if err != nil {
		return err
	}

	filteredKrs := krs.FilterMetrics(time.Now().AddDate(0,0,-7))
	if len(filteredKrs) == 0 {
		output.Print(nothingToUpdateCongratsMsg)
		return nil
	}

	output.PrintMetrics(filteredKrs)

	return nil
}
