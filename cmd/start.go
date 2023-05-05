package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start DKIM service",
		Action: func(ctx *cli.Context) error {
			fmt.Println("hello")
			return nil
		},
	}
}
