package githubcli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func Build(logger scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		ghLayer, err := context.Layers.Get("github-cli")
		if err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to get github-cli layer: %w", err)
		}

		// Set up GitHub CLI installation
		ghLayer.Launch = true
		ghLayer.Cache = true

		// Create bin directory in the layer
		binDir := filepath.Join(ghLayer.Path, "bin")
		if err := os.MkdirAll(binDir, 0755); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to create bin directory: %w", err)
		}

		// Install GitHub CLI (this would typically download and install the binary)
		// For now, we'll create a placeholder script that would be replaced with actual installation
		installScript := `#!/bin/bash
# Install GitHub CLI
curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
sudo apt update && sudo apt install gh -y
`

		installPath := filepath.Join(binDir, "install-gh.sh")
		if err := os.WriteFile(installPath, []byte(installScript), 0755); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to write install script: %w", err)
		}

		// Set up environment variables
		ghLayer.SharedEnv.Default("PATH", filepath.Join(ghLayer.Path, "bin"))
		ghLayer.SharedEnv.Default("GITHUB_CLI_VERSION", "latest")

		// Create a process to run GitHub CLI
		processes := []packit.Process{
			{
				Type:    "github-cli",
				Command: "gh",
				Args:    []string{"--version"},
				Default: false,
				Direct:  true,
			},
		}

		logger.LaunchProcesses(processes)

		return packit.BuildResult{
			Layers: []packit.Layer{ghLayer},
			Launch: packit.LaunchMetadata{
				Processes: processes,
			},
		}, nil
	}
} 