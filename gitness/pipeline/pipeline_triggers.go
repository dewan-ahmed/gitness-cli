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

var pipelineTriggerCmd = &cli.Command{
	Name:      "trigger",
	Usage:     "list pipeline triggers",
	ArgsUsage: " ",
	Action:    pipelineTrigger,
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
		&cli.StringFlag{
			Name:        "pipeline-id",
			Usage:       "pipeline id",
			Required:    true,
			DefaultText: "default",
		},
	},
}

func pipelineTrigger(ctx *cli.Context) error {
	base_url := ctx.String("url")
	safe_repo_ref := url.QueryEscape(ctx.String("repo-ref"))
	safe_pipeline_id := url.QueryEscape(ctx.String("pipeline-id"))
	body, err := internal.HttpGetRequest(ctx, base_url+"api/v1/repos/"+safe_repo_ref+"/pipelines/"+safe_pipeline_id+"/triggers")
	if err != nil {
		return fmt.Errorf("failed for trigger '%s': %w", ctx.String("pipeline-id"), err)
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var triggers []types.Trigger
	json.Unmarshal(body, &triggers)

	for _, trigger := range triggers {
		tmpl.Execute(os.Stdout, trigger)
	}

	return nil
}
