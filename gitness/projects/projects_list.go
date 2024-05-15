package projects

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/dewan-ahmed/gitness-cli/gitness/internal"

	"github.com/urfave/cli/v2"
)

var projectsListCmd = &cli.Command{
	Name:      "ls",
	Usage:     "list projects",
	ArgsUsage: "PROJECTNAME",
	Action:    projectsList,
}

type Project struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`
	IsPublic    bool   `json:"is_public"`
	Created     int64  `json:"created"`
}

func projectsList(ctx *cli.Context) error {
	baseURL := ctx.String("url")
	safeProjectName := url.QueryEscape(ctx.Args().First())
	body, err := internal.HttpRequest(ctx, baseURL+"api/v1/spaces/"+safeProjectName)
	if err != nil {
		return fmt.Errorf("failed to list projects: %w", err)
	}

	var project Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		return fmt.Errorf("failed to unmarshal project data: %w", err)
	}

	// Convert Unix timestamp to human-readable time
	createdTime := time.Unix(project.Created/1000, 0).Format("2006-01-02 15:04:05")

	// Print formatted project information
	fmt.Printf("Project ID: %s\n", project.Path)
	fmt.Printf("Project Description: %s\n", project.Description)
	visibility := "private"
	if project.IsPublic {
		visibility = "public"
	}
	fmt.Printf("Project Visibility: %s\n", visibility)
	fmt.Printf("Project Created: %s\n", createdTime)

	return nil
}
