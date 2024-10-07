package server

import "github.com/urfave/cli/v2"

func Command() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start webcore server",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			Run()
			return nil
		},
	}
}
