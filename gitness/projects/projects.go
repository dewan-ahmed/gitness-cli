package projects

import "github.com/urfave/cli/v2"

// Command exports the project command.
var Command = cli.Command{
	Name:  "project",
	Usage: "manage projects",
	Subcommands: []*cli.Command{
		projectsListCmd,
		projectsCreateCmd,
		projectsDeleteCmd,
	},
}
