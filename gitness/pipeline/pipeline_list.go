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
	ArgsUsage: " ",
	Action:    pipelineList,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: `{{ .ConfigPath }}`,
		},
		&cli.StringFlag{
			Name:        "repo-ref",
			Usage:       "repo ref",
			Required:    true,
			DefaultText: "project/repo",
		},
		&cli.StringFlag{
			Name:        "pipeline-id",
			Usage:       "pipeline id",
			Required:    true,
			DefaultText: "default",
		},
	},
}

func pipelineList(ctx *cli.Context) error {
	base_url := ctx.String("url")
	safe_repo_ref := url.QueryEscape(ctx.String("repo-ref"))
	safe_pipeline_id := url.QueryEscape(ctx.String("pipeline-id"))
	body, err := internal.HttpRequest(ctx, base_url+"api/v1/repos/"+safe_repo_ref+"/pipelines/"+safe_pipeline_id)
	if err != nil {
		return fmt.Errorf("failed for pipeline '%s': %w", ctx.String("pipeline-id"), err)
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var pipeline types.Pipeline
	json.Unmarshal(body, &pipeline)

	tmpl.Execute(os.Stdout, pipeline)

	return nil
}
