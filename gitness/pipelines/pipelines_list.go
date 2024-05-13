package pipelines

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

var pipelinesListCmd = &cli.Command{
	Name:      "list",
	Usage:     "list pipelines",
	ArgsUsage: " ",
	Action:    pipelinesList,
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

// The code defines a function to list pipelines by making an HTTP request to a specified URL and parsing the response JSON data.
// It then uses a template to format and display the pipeline information to the standard output.
func pipelinesList(ctx *cli.Context) error {
	base_url := ctx.String("url")
	safe_repo_ref := url.QueryEscape(ctx.String("repo-ref"))
	body, err := internal.HttpRequest(ctx, base_url+"api/v1/repos/"+safe_repo_ref+"/pipelines")
	if err != nil {
		return fmt.Errorf("failed to list pipelines for '%s': %w", ctx.String("repo-ref"), err)
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
