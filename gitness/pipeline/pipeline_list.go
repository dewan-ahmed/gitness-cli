package pipeline

import (
	"encoding/json"
	"html/template"
	"net/url"
	"os"

	"github.com/dewan-ahmed/gitness-cli/gitness/internal"

	"github.com/harness/gitness/types"
	"github.com/urfave/cli/v2"
)

var pipelineListCmd = &cli.Command{
	Name:      "ls",
	Usage:     "list pipelines",
	ArgsUsage: " ",
	Action:    pipelineList,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: `{{ .Identifier }}`,
		},
		&cli.StringFlag{
			Name:        "repo-ref",
			Usage:       "repo ref",
			Required:    true,
			DefaultText: "project/repo",
		},
	},
}

func pipelineList(ctx *cli.Context) error {
	base_url := ctx.String("url")
	safe_repo_ref := url.QueryEscape(ctx.String("repo-ref"))
	body, err := internal.HttpRequest(ctx, base_url+"api/v1/repos/"+safe_repo_ref+"/pipelines")
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var pipelines []types.Pipeline
	json.Unmarshal(body, &pipelines)

	for _, pipeline := range pipelines {
		tmpl.Execute(os.Stdout, pipeline)
	}

	return nil
}
