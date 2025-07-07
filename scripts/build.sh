#!/usr/bin/env bash
set -euo pipefail

GOOS=${1:-linux}
GOARCH=${2:-amd64}

echo "Building for GOOS=${GOOS} GOARCH=${GOARCH}"

# Build the detect binary
GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o bin/detect ./cmd/detect

# Build the build binary
GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o bin/build ./cmd/build

# Build the run binary
cd run
GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o ../bin/run main.go
cd ..

echo "Build complete!" 