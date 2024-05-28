package projects

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/dewan-ahmed/gitness-cli/gitness/internal"

	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"
	"github.com/urfave/cli/v2"
)

var projectsCreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "create a project",
	ArgsUsage: "<project id>",
	Action:    projectsCreate,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "format output",
			Value: `created {{ .Identifier }}`,
		},
		&cli.StringFlag{
			Name:     "description",
			Usage:    "project description",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "public",
			Usage:    "make project public",
			Required: false,
		},
	},
}

type ProjectCreateRequest struct {
	Description string `json:"description,omitempty"`
	IsPublic    bool   `json:"is_public"`
	Identifier  string `json:"identifier"`
}

func projectsCreate(ctx *cli.Context) error {
	baseURL := ctx.String("url")

	projectId := ctx.Args().First()
	if projectId == "" {
		return fmt.Errorf("project id is required")
	}
	err := check.SpaceIdentifierDefault(projectId, true)
	if err != nil {
		return err
	}

	description := ctx.String("description")
	err = check.Description(description)
	if err != nil {
		return err
	}

	// TODO: why doesn't the space struct have IsPublic?
	// https://github.com/harness/gitness/blob/main/types/space.go#L33
	//
	// We shouldn't have to create our own struct for a space/project
	project := ProjectCreateRequest{
		Description: description,
		IsPublic:    ctx.Bool("public"),
		Identifier:  projectId,
	}

	reqBody, err := json.Marshal(project)
	if err != nil {
		return fmt.Errorf("failed to marshal project data: %w", err)
	}

	body, err := internal.HttpPostRequest(ctx, baseURL+"api/v1/spaces", reqBody)
	if err != nil {
		return fmt.Errorf("failed for project '%s': %w", projectId, err)
	}

	tmpl, err := template.New("_").Parse(ctx.String("format") + "\n")
	if err != nil {
		return err
	}

	var space types.Space
	json.Unmarshal(body, &space)

	tmpl.Execute(os.Stdout, space)

	return nil
}
