package pipeline

import "github.com/urfave/cli/v2"

// Command exports the repository command.
var Command = cli.Command{
	Name:  "pipeline",
	Usage: "manage pipelines",
	Subcommands: []*cli.Command{
		pipelineListCmd,
	},
}
