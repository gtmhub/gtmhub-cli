package commands

import (
	"gtmhub-cli/config"
	"gtmhub-cli/output"

	"github.com/urfave/cli/v2"
)

var (
	LogoutCommand = &cli.Command{
		Name:  "logout",
		Usage: "logs you out of gtmhub",
		Action: LogoutAction,
	}
)

func LogoutAction(c *cli.Context) error {
	config.Clear()
	output.Print(logoutMsg, output.Green)
	return nil
}
