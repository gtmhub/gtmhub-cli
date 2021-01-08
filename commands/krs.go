package commands

import (
	"gtmhub-cli/output"

	"github.com/urfave/cli/v2"
)

var (
	KRsCommand = &cli.Command{
		Name:  "krs",
		Usage: "shows all your krs that are in the current session",
		Action: ShowKRSAction,
	}
)

func ShowKRSAction(c *cli.Context) error {
	krs, err := client.GetMetricsInCurrentSession()
	if err != nil {
		return err
	}

	output.PrintMetrics(krs)

	return nil
}
