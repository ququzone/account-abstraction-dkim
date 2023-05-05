package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:    "aa-dkim",
		Version: "v0.1.0",
		Authors: []*cli.Author{
			{
				Name:  "ququzone",
				Email: "xueping.yang@gmail.com",
			},
		},
		HelpName:  "aa-dkim",
		Usage:     "dkim service for account abstraction",
		UsageText: "aa-dkim <SUBCOMMAND>",
		Commands: []*cli.Command{
			start(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}
}
