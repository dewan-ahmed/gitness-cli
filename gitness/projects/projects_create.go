// This code is not working
// ./gitness-cli project create PROJECTNAME
// returning
// Error: failed to create project: 401 Unauthorized
package projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	token := os.Getenv("GITNESS_TOKEN") // Retrieve the authentication token from environment variable

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

	req, err := http.NewRequest("POST", baseURL+"api/v1/spaces", strings.NewReader(string(reqBody)))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // Set the Authorization header with the token

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create project: %s", resp.Status)
	}

	fmt.Println("Project created successfully")
	return nil
}
