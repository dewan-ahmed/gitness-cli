package docker

import (
	"fmt"
	"os/exec"
	"strings"
)

// StartGitness starts the Gitness Docker container
func StartGitness() error {
	if !isDockerRunning() {
		return fmt.Errorf("docker is not running. Please start Docker before running Gitness")
	}

	// Define Docker command
	dockerCmd := exec.Command("docker", "run", "-d",
		"-p", "3000:3000",
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-v", "/tmp/gitness:/data",
		"--name", "gitness",
		"--restart", "always",
		"harness/gitness",
	)

	// Execute Docker command
	output, err := dockerCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "Conflict") {
			return fmt.Errorf("a container with the name 'gitness' already exists. Please remove the existing container or choose a different name")

		}
		return fmt.Errorf("failed to start Gitness: %v. Output: %s", err, output)
	}

	return nil
}

// isDockerRunning checks if Docker is running
func isDockerRunning() bool {
	// Check if Docker info can be retrieved
	dockerCmd := exec.Command("docker", "info")
	err := dockerCmd.Run()
	return err == nil
}
