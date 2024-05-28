package repo

import "github.com/urfave/cli/v2"

// Command exports the pipeline command.
var Command = cli.Command{
	Name:  "repo",
	Usage: "manage repo",
	Subcommands: []*cli.Command{
		repoImportCmd,
	},
}
