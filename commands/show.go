package commands

import (
	"fmt"
	"gtmhub-cli/output"

	"github.com/urfave/cli/v2"
)

var (
	ShowCommand = &cli.Command{
		Name:  "show",
		Usage: "shows all your items from current sessions",
		//Action: ShowKRSAction,
		Subcommands: []*cli.Command{
			{
				Name: "krs",
				Usage: "shows all your krs that are in a current session",
				Action: ShowKRSAction,
			},
			{
				Usage: "shows all your objectives that are in a current session",
				Action: ShowGoalsAction,
				Name: "goals",
			},
		},
	}
)

func ShowGoalsAction(c *cli.Context) error {
	return fmt.Errorf("not implemented")
}

func ShowKRSAction(c *cli.Context) error {
	krs, err := client.GetMetricsInCurrentSession()
	if err != nil {
		return err
	}

	output.PrintMetrics(krs)

	return nil
}
