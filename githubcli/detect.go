package githubcli

import (
	"fmt"

	"github.com/paketo-buildpacks/packit/v2"
)

type BuildPlanMetadata struct {
	Launch bool `toml:"launch"`
}

func Detect(configParser ConfigParser) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		// Always detect and provide GitHub CLI, regardless of app content
		fmt.Println("GitHub CLI buildpack detected - will provide GitHub CLI for the application")

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{
						Name: "github-cli",
					},
				},
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
