package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/dewan-ahmed/gitness-cli/gitness/internal"

	"github.com/harness/gitness/app/api/controller/repo"
	"github.com/harness/gitness/app/services/importer"
	"github.com/harness/gitness/types"
	"github.com/urfave/cli/v2"
)

var repoImportCmd = &cli.Command{
	Name:      "import",
	Usage:     "import a repo",
	ArgsUsage: "<org/repo>",
	Action:    repoImport,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: `imported {{ .Identifier }}`,
		},
		&cli.StringFlag{
			Name:     "description",
			Usage:    "repo description",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "project-id",
			Usage:    "project id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "uid",
			Usage:    "unique id",
			Required: true,
		},
	},
}

func repoImport(ctx *cli.Context) error {
	base_url := ctx.String("url")
	provider_repo := ctx.Args().First()
	if provider_repo == "" {
		return fmt.Errorf("repo is required")
	}

	// TODO: does it make sense to reuse these Gitness structs?
	importReq := &repo.ImportInput{
		Description: ctx.String("description"),
		ParentRef:   ctx.String("project-id"),
		// TODO: support pipeline conversion
		Pipelines:    "ignore",
		ProviderRepo: provider_repo,
		UID:          ctx.String("uid"),
		// TODO: build out provider support
		Provider: importer.Provider{
			Type: "github",
		},
	}

	reqBody, err := json.Marshal(importReq)
	if err != nil {
		return fmt.Errorf("failed to marshal import data: %w", err)
	}

	body, err := internal.HttpPostRequest(ctx, base_url+"api/v1/repos/import", reqBody)
	if err != nil {
		return fmt.Errorf("failed for repo '%s': %w", provider_repo, err)
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var repo types.Repository
	json.Unmarshal(body, &repo)

	tmpl.Execute(os.Stdout, repo)

	return nil
}
