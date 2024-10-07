package core

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/centrx/webcore/core/server"
)

func Run() error {
	app := &cli.App{
		Name:    "webcore",
		Usage:   "webcore",
		Version: "0.0.1",
		Flags:   []cli.Flag{},
		Commands: []*cli.Command{
			server.Command(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}

	return nil
}
