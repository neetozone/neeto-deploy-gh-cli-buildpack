#!/usr/bin/env bash
set -euo pipefail

GOOS=${1:-linux}
GOARCH=${2:-amd64}

echo "Packaging for GOOS=${GOOS} GOARCH=${GOARCH}"

# Build the binaries
./scripts/build.sh ${GOOS} ${GOARCH}

# Create the buildpack
pack buildpack package github-cli-buildpack:latest \
  --config package.toml

echo "Package complete!" 