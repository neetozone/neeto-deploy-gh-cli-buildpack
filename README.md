# GitHub CLI Cloud Native Buildpack

A Cloud Native Buildpack that automatically installs GitHub CLI (`gh`) in containerized applications.

## Features

- ✅ Always detects and installs GitHub CLI (version 2.40.1)
- ✅ Cross-platform support (Linux x86-64)
- ✅ Proper buildpack lifecycle implementation
- ✅ Environment variables and process types configured

## Quick Start

### Build and Deploy

```bash
# Build the buildpack
./scripts/build.sh

# Package and push to ECR
./scripts/package.sh
docker tag github-cli-buildpack:latest 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 348674388966.dkr.ecr.us-east-1.amazonaws.com
docker push 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
```

### Usage

```bash
# Include in any application build
pack build my-app --buildpack 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest

# GitHub CLI will be available in the container
docker run my-app gh --version
```

## How It Works

### Detection
Always detects and runs regardless of application content.

### Build Process
1. Creates a layer for GitHub CLI installation
2. Downloads and installs GitHub CLI 2.40.1 from GitHub releases
3. Sets up PATH and environment variables
4. Creates `github-cli` process type

### Environment Variables
- `PATH`: Updated to include GitHub CLI binary location
- `GITHUB_CLI_VERSION`: Set to "2.40.1"

## Examples

### Node.js with GitHub CLI
```bash
pack build my-app \
  --buildpack paketo-buildpacks/nodejs \
  --buildpack 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
```

### GitHub CLI Commands
```bash
gh --version
gh auth login
gh repo list
gh issue create --title "Bug report" --body "Description"
```

## Development

### Prerequisites
- Go 1.21+
- Docker
- AWS CLI

### Local Testing
```bash
./test_buildpack.sh  # May fail on macOS (Linux binaries)
```

## Troubleshooting

- **"exec format error"**: Fixed by proper buildpack structure with separate entry points
- **Authentication**: Set up GitHub Personal Access Token or use `gh auth login`
- **Binary not found**: Verify buildpack was applied and PATH is set

## License

Apache License 2.0 