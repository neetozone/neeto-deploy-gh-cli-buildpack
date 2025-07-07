# GitHub CLI Cloud Native Buildpack

A Cloud Native Buildpack for installing and configuring GitHub CLI (`gh`) in containerized applications.

## Overview

This buildpack automatically installs and configures GitHub CLI for applications that need to interact with GitHub from the command line. It's designed to work with Cloud Native Buildpacks and can be used in various deployment environments.

## Features

- ✅ Automatic GitHub CLI installation
- ✅ Cross-platform support (Linux x86-64)
- ✅ Proper buildpack lifecycle implementation
- ✅ Environment variable configuration
- ✅ Process type definitions for container orchestration

## Buildpack Structure

The buildpack follows the standard Cloud Native Buildpack structure:

```
├── bin/
│   ├── detect    # Detection phase binary
│   ├── build     # Build phase binary
│   └── run       # Run phase binary
├── cmd/
│   ├── detect/   # Detect command entry point
│   └── build/    # Build command entry point
├── githubcli/    # Core buildpack logic
├── run/          # Run command entry point
├── scripts/      # Build and packaging scripts
└── buildpack.toml # Buildpack configuration
```

## Building the Buildpack

### Prerequisites

- Go 1.21 or later
- Docker
- AWS CLI (for ECR deployment)

### Local Build

To build the buildpack locally:

```bash
# Build for Linux (default)
./scripts/build.sh

# Build for specific platform
./scripts/build.sh linux amd64
```

### Packaging

To package the buildpack as a Docker image:

```bash
./scripts/package.sh
```

This creates a Docker image tagged as `github-cli-buildpack:latest`.

## Deployment to ECR

To deploy the buildpack to AWS ECR:

```bash
# Build and package
./scripts/package.sh

# Tag for ECR
docker tag github-cli-buildpack:latest 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest

# Authenticate with ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 348674388966.dkr.ecr.us-east-1.amazonaws.com

# Push to ECR
docker push 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
```

## Detection

The buildpack **always detects and runs** regardless of the application content. This means GitHub CLI will be available in any application that includes this buildpack, making it useful for:

- Applications that need GitHub CLI for automation
- CI/CD pipelines that require GitHub API access
- Development environments that need GitHub CLI tools
- Any containerized application that might need GitHub functionality

The buildpack provides the `github-cli` dependency and requires it for launch, ensuring GitHub CLI is available in the final container.

## Build Process

During the build phase, the buildpack:

1. **Creates a layer** for GitHub CLI installation with launch and cache enabled
2. **Downloads GitHub CLI** binary (version 2.40.1) from GitHub releases
3. **Extracts and installs** the binary in the layer's bin directory
4. **Sets up environment variables** including PATH and GITHUB_CLI_VERSION
5. **Creates process types** for container orchestration (github-cli process)
6. **Makes the binary executable** with proper permissions

## Usage

### Basic Usage

Include this buildpack in your buildpack group:

```bash
pack build my-app --buildpack 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
```

### With Other Buildpacks

This buildpack can be combined with other buildpacks:

```bash
pack build my-app \
  --buildpack paketo-buildpacks/nodejs \
  --buildpack 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest
```

### Running GitHub CLI Commands

Once the buildpack is applied, you can run GitHub CLI commands in your container:

```bash
# Check GitHub CLI version
gh --version

# Authenticate with GitHub
gh auth login

# List repositories
gh repo list

# Create issues, pull requests, etc.
gh issue create --title "Bug report" --body "Description"
```

## Configuration

### Environment Variables

- `PATH`: Automatically updated to include the GitHub CLI binary location
- `GITHUB_CLI_VERSION`: Set to "2.40.1" (hardcoded in the buildpack)
- `BP_LOG_LEVEL`: Set buildpack log level (optional)

### Authentication

To use GitHub CLI with authentication, you'll need to:

1. Set up a GitHub Personal Access Token
2. Configure authentication in your application or container environment
3. Use `gh auth login` or set the `GITHUB_TOKEN` environment variable

## Process Types

The buildpack defines the following process types:

- `github-cli`: Process for running GitHub CLI commands (non-default, direct execution)

## Examples

### Node.js Application with GitHub CLI

Create a `package.json` with GitHub CLI scripts:

```json
{
  "name": "my-github-app",
  "scripts": {
    "release": "gh release create v1.0.0 --generate-notes",
    "pr": "gh pr create --title 'Update dependencies' --body 'Automated PR'"
  }
}
```

### Container with GitHub CLI

Since this buildpack always detects, you can include it in any container build to make GitHub CLI available:

```bash
# Build any application with GitHub CLI available
pack build my-app \
  --buildpack 348674388966.dkr.ecr.us-east-1.amazonaws.com/neeto-deploy/buildpacks/gh-cli:latest

# GitHub CLI will be available in the container
docker run my-app gh --version
```

## Testing

To test the buildpack locally:

```bash
# Run the test script
./test_buildpack.sh
```

Note: 
- The test script may fail on macOS since the binaries are built for Linux
- The test script checks for GitHub CLI configuration files, but the buildpack always detects regardless of application content

## Troubleshooting

### Common Issues

1. **"exec format error"**: This was fixed by properly structuring the buildpack with separate entry points for each lifecycle phase.

2. **Authentication issues**: Ensure you have proper GitHub authentication configured.

3. **Binary not found**: Verify that the buildpack was applied correctly and the PATH is set.

## Contributing

To contribute to this buildpack:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This buildpack is licensed under the Apache License 2.0. 