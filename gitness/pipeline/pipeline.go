package pipeline

import "github.com/urfave/cli/v2"

// Command exports the pipeline command.
var Command = cli.Command{
	Name:  "pipeline",
	Usage: "manage pipeline",
	Subcommands: []*cli.Command{
		pipelineExecutionsCmd,
		pipelineListCmd,
	},
}
