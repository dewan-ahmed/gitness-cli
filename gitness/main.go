package main

import (
	"fmt"
	"os"

	"github.com/dewan-ahmed/gitness-cli/gitness/pipeline"
	"github.com/dewan-ahmed/gitness-cli/gitness/pipelines"

	"github.com/urfave/cli/v2"
)

// The code defines a CLI application using the urfave/cli package.
// It sets up flags for token and URL, assigns default values, and includes a command for pipelines.
// The application runs with error handling for any issues that may arise.
func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Usage:   "gitness personal access token",
				EnvVars: []string{"GITNESS_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "url",
				Usage:   "gitness server url",
				EnvVars: []string{"GITNESS_URL"},
				Value:   "http://localhost:3000/",
			},
		},
	}
	app.Commands = []*cli.Command{
		&pipeline.Command,
		&pipelines.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
