// This code is not working
// ./gitness-cli project create PROJECTNAME
// returning
// Error: failed to create project: 401 Unauthorized
package projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli/v2"
)

var projectsCreateCmd = &cli.Command{
	Name:      "create",
	Usage:     "create a project",
	ArgsUsage: "PROJECTNAME",
	Action:    projectsCreate,
	Flags: []cli.Flag{
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
	projectName := ctx.Args().First()
	if projectName == "" {
		return fmt.Errorf("missing PROJECTNAME argument")
	}

	description := ctx.String("description")
	isPublic := ctx.Bool("public")

	project := ProjectCreateRequest{
		Description: description,
		IsPublic:    isPublic,
		Identifier:  projectName,
	}

	reqBody, err := json.Marshal(project)
	if err != nil {
		return fmt.Errorf("failed to marshal project data: %w", err)
	}

	resp, err := http.Post(baseURL+"api/v1/spaces", "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create project: %s", resp.Status)
	}

	fmt.Println("Project created successfully")
	return nil
}
