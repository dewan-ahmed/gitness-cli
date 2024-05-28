package pipeline

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"os"

	"github.com/dewan-ahmed/gitness-cli/gitness/internal"

	"github.com/harness/gitness/types"
	"github.com/urfave/cli/v2"
)

var pipelineListCmd = &cli.Command{
	Name:      "list",
	Usage:     "list pipelines",
	ArgsUsage: "<pipeline id>",
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
	pipeline_id := ctx.Args().First()
	var api_path string
	if pipeline_id != "" {
		safe_pipeline_id := url.QueryEscape(pipeline_id)
		api_path = "api/v1/repos/" + safe_repo_ref + "/pipelines/" + safe_pipeline_id
	} else {
		api_path = "api/v1/repos/" + safe_repo_ref + "/pipelines"
	}

	body, err := internal.HttpGetRequest(ctx, base_url+api_path)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var pipeline types.Pipeline
	var pipelines []types.Pipeline

	if pipeline_id != "" {
		json.Unmarshal(body, &pipeline)
		tmpl.Execute(os.Stdout, pipeline)
	} else {
		json.Unmarshal(body, &pipelines)
		for _, pipeline := range pipelines {
			tmpl.Execute(os.Stdout, pipeline)
		}
	}

	return nil
}
