package githubcli

import (
	"fmt"
	"os"
	"os/exec"
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

		// Download and install GitHub CLI
		logger.Process("Installing GitHub CLI...")

		// Download GitHub CLI
		ghVersion := "2.40.1"
		arch := "amd64" // Default to amd64 for the builder

		// Download the binary
		downloadURL := fmt.Sprintf("https://github.com/cli/cli/releases/download/v%s/gh_%s_linux_%s.tar.gz", ghVersion, ghVersion, arch)
		downloadCmd := exec.Command("curl", "-fsSL", downloadURL, "-o", "/tmp/gh.tar.gz")
		if err := downloadCmd.Run(); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to download GitHub CLI: %w", err)
		}

		// Extract the binary
		extractCmd := exec.Command("tar", "-xzf", "/tmp/gh.tar.gz", "-C", "/tmp")
		if err := extractCmd.Run(); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to extract GitHub CLI: %w", err)
		}

		// Copy the binary to the layer
		ghBinary := fmt.Sprintf("/tmp/gh_%s_linux_%s/bin/gh", ghVersion, arch)
		destPath := filepath.Join(binDir, "gh")
		if err := exec.Command("cp", ghBinary, destPath).Run(); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to copy GitHub CLI binary: %w", err)
		}

		// Make it executable
		if err := os.Chmod(destPath, 0755); err != nil {
			return packit.BuildResult{}, fmt.Errorf("failed to make GitHub CLI executable: %w", err)
		}

		logger.Process("GitHub CLI installed successfully")

		// Set up environment variables
		ghLayer.SharedEnv.Default("PATH", filepath.Join(ghLayer.Path, "bin"))
		ghLayer.SharedEnv.Default("GITHUB_CLI_VERSION", ghVersion)

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
