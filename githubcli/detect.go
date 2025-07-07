package githubcli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
)

type BuildPlanMetadata struct {
	Launch bool `toml:"launch"`
}

func Detect(configParser ConfigParser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		// Check for common GitHub CLI configuration files
		configPaths := []string{
			filepath.Join(context.WorkingDir, ".github"),
			filepath.Join(context.WorkingDir, "gh.yml"),
			filepath.Join(context.WorkingDir, ".gh.yml"),
			filepath.Join(context.WorkingDir, "package.json"), // For Node.js projects that might use GitHub CLI
		}

		hasGhConfig := false
		for _, path := range configPaths {
			if _, err := os.Stat(path); err == nil {
				hasGhConfig = true
				break
			}
		}

		// Also check if there's a specific requirement file
		hasRequirement, err := configParser.Parse(filepath.Join(context.WorkingDir, "gh-requirements.txt"))
		if err != nil {
			return packit.DetectResult{}, fmt.Errorf("failed to parse gh-requirements.txt: %w", err)
		}

		if hasGhConfig || hasRequirement {
			fmt.Println("GitHub CLI configuration or requirements found")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "github-cli",
						Metadata: BuildPlanMetadata{
							Launch: true,
						},
					},
				},
			},
		}, nil
	}
} 