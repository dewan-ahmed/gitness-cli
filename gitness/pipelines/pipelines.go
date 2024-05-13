package pipelines

import "github.com/urfave/cli/v2"

// Command exports the pipeline command.
var Command = cli.Command{
	Name:  "pipelines",
	Usage: "manage pipelines",
	Subcommands: []*cli.Command{
		pipelinesListCmd,
	},
}
