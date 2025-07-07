package main

import (
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/unni/github-cli-buildpack/githubcli"
)

func main() {
	parser := githubcli.NewConfigParser()
	packit.Detect(githubcli.Detect(parser))
}
