package commands

import "github.com/urfave/cli/v2"

var (
	GetCommand = &cli.Command{
		Name:  "get",
		Usage: "gets you items from gtmhub",
		Subcommands: []*cli.Command{
			ListCommand,
			KRsCommand,
		},
	}
)
