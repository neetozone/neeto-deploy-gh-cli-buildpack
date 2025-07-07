#!/usr/bin/env bash
set -euo pipefail

# Test script for the GitHub CLI buildpack
APP_DIR="/Users/unni/Work/test_apps/hello-world-rails-8"
BUILDPACK_DIR="/Users/unni/Work/Neeto/neeto-deploy-buildpacks/neeto-deploy-gh-cli-buildpack"

echo "Testing GitHub CLI Buildpack"
echo "============================"
echo "App directory: $APP_DIR"
echo "Buildpack directory: $BUILDPACK_DIR"
echo ""

# Check if app directory exists
if [ ! -d "$APP_DIR" ]; then
    echo "Error: App directory does not exist: $APP_DIR"
    exit 1
fi

# Check if buildpack directory exists
if [ ! -d "$BUILDPACK_DIR" ]; then
    echo "Error: Buildpack directory does not exist: $BUILDPACK_DIR"
    exit 1
fi

# Check for GitHub CLI related files in the app
echo "Checking for GitHub CLI configuration files..."
GITHUB_FILES=(
    ".github"
    "gh.yml"
    ".gh.yml"
    "package.json"
    "gh-requirements.txt"
)

FOUND_FILES=()
for file in "${GITHUB_FILES[@]}"; do
    if [ -e "$APP_DIR/$file" ]; then
        echo "✓ Found: $file"
        FOUND_FILES+=("$file")
    else
        echo "✗ Not found: $file"
    fi
done

echo ""
if [ ${#FOUND_FILES[@]} -gt 0 ]; then
    echo "✅ GitHub CLI configuration detected!"
    echo "Found files: ${FOUND_FILES[*]}"
else
    echo "❌ No GitHub CLI configuration found"
    exit 1
fi

echo ""
echo "Testing buildpack binaries..."

# Test detect binary
echo "Testing detect binary..."
cd "$APP_DIR"
if "$BUILDPACK_DIR/bin/detect" . > /dev/null 2>&1; then
    echo "✅ Detect binary executed successfully"
else
    echo "❌ Detect binary failed"
    echo "Running with verbose output:"
    "$BUILDPACK_DIR/bin/detect" .
fi

echo ""
echo "Buildpack test completed!" 