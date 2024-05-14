package projects

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var projectsDeleteCmd = &cli.Command{
	Name:      "delete",
	Usage:     "delete a project",
	ArgsUsage: "PROJECTNAME",
	Action:    projectsDelete,
}

func projectsDelete(ctx *cli.Context) error {
	baseURL := ctx.String("url")
	token := os.Getenv("GITNESS_TOKEN") // Retrieve the authentication token from environment variable

	projectName := ctx.Args().First()
	if projectName == "" {
		return fmt.Errorf("missing PROJECTNAME argument")
	}

	req, err := http.NewRequest("DELETE", baseURL+"api/v1/spaces/"+projectName, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // Set the Authorization header with the token

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete project: %s", resp.Status)
	}

	fmt.Println("Project deleted successfully")
	return nil
}
