package main

import (
	"fmt"
	"os"

	"github.com/dewan-ahmed/gitness-cli/gitness/docker"
	"github.com/dewan-ahmed/gitness-cli/gitness/pipeline"
	"github.com/dewan-ahmed/gitness-cli/gitness/pipelines"
	"github.com/urfave/cli/v2"
)

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
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start Gitness Docker container",
				Action: func(c *cli.Context) error {
					if err := docker.StartGitness(); err != nil {
						return err
					}
					fmt.Println("🚀 Gitness started successfully!")
					return nil
				},
			},
			&pipeline.Command,
			&pipelines.Command,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
