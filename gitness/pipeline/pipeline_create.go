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

var pipelineCreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "create pipeline",
	ArgsUsage: "<pipeline id>",
	Action:    pipelineCreate,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: `created {{ .Identifier }}`,
		},
		&cli.StringFlag{
			Name:        "repo-ref",
			Usage:       "repo ref",
			Required:    true,
			DefaultText: "project/repo",
		},
		&cli.StringFlag{
			Name:        "config-path",
			Usage:       "config path",
			Required:    false,
			DefaultText: ".harness/<pipeline id>.yml",
		},
		&cli.StringFlag{
			Name:     "default-branch",
			Usage:    "default branch",
			Required: false,
			Value:    "main",
		},
		&cli.StringFlag{
			Name:     "description",
			Usage:    "description",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "disabled",
			Usage:    "disabled",
			Required: false,
		},
	},
}

func pipelineCreate(ctx *cli.Context) error {
	base_url := ctx.String("url")
	safe_repo_ref := url.QueryEscape(ctx.String("repo-ref"))
	pipeline_id := ctx.Args().First()
	if pipeline_id == "" {
		return fmt.Errorf("pipeline id is required")
	}
	safe_pipeline_id := url.QueryEscape(pipeline_id)

	new_pipeline := &types.Pipeline{
		ConfigPath:    ctx.String("config-path"),
		DefaultBranch: ctx.String("default-branch"),
		Description:   ctx.String("description"),
		Disabled:      ctx.Bool("disabled"),
		Identifier:    safe_pipeline_id,
	}

	if new_pipeline.ConfigPath == "" {
		new_pipeline.ConfigPath = ".harness/" + pipeline_id + ".yml"
	}

	reqBody, err := json.Marshal(&new_pipeline)
	if err != nil {
		return fmt.Errorf("failed to marshal pipeline: %w", err)
	}

	body, err := internal.HttpPostRequest(ctx, base_url+"api/v1/repos/"+safe_repo_ref+"/pipelines", reqBody)
	if err != nil {
		return fmt.Errorf("failed for pipeline '%s': %w", pipeline_id, err)
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
