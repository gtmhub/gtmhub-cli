package main

import (
	"fmt"
	"gtmhub-cli/commands"
	config2 "gtmhub-cli/config"
	"gtmhub-cli/output"
	"os"

	"github.com/urfave/cli/v2"
)

var major, minor, patch string

func main() {

	config2.InitConfig()
	if major == "" && minor == "" && patch == "" {
		major = "0"
		minor = "1"
		patch = "0"
	}

	app := &cli.App{
		Name: "gtmhub",
		Version: fmt.Sprintf("%s.%s.%s", major, minor, patch),
		Usage: "The best cli tool in the world. Helps you stay focused while still keeping up with your okrs.",
		Commands: []*cli.Command{
			commands.LoginCommand,
			commands.LogoutCommand,
			commands.StatusCommand,
			commands.UpdateCommand,
			commands.GetCommand,
		},
		Flags: []cli.Flag{
			cli.VersionFlag,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		output.PrintErrorAndExit(err)
	}
}
