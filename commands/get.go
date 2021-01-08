package commands

import "github.com/urfave/cli/v2"

var (
	GetCommand = &cli.Command{
		Name:  "get",
		Usage: "gets you items from gtmhub",
		Action: ListAction,
		Subcommands: []*cli.Command{
			ListCommand,
			KRsCommand,
		},
	}
)
