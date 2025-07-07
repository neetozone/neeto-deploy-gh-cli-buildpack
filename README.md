# GitHub CLI Cloud Native Buildpack

## `gcr.io/paketo-buildpacks/github-cli`

The GitHub CLI CNB installs and configures the GitHub CLI (`gh`) for applications that need to interact with GitHub from the command line.

## Integration

This CNB installs GitHub CLI and makes it available in the container environment. It can be used as a dependency for other buildpacks or applications that need GitHub CLI functionality.

To package this buildpack for consumption:
```
$ ./scripts/package.sh
```
This builds the buildpack's source using GOOS=linux by default. You can supply another value as the first argument to package.sh.

## Detection

The buildpack will detect and run when it finds any of the following:

- `.github/` directory (indicating GitHub workflows or configuration)
- `gh.yml` or `.gh.yml` files (GitHub CLI configuration)
- `package.json` (for Node.js projects that might use GitHub CLI)
- `gh-requirements.txt` file containing GitHub CLI related requirements

## Build

During the build phase, the buildpack:

1. Creates a layer for GitHub CLI installation
2. Downloads and installs the latest GitHub CLI binary
3. Sets up environment variables and PATH
4. Creates a launch process for running GitHub CLI commands

## Usage

### Basic Usage

Include this buildpack in your buildpack group:

```bash
pack build my-app --buildpack paketo-buildpacks/github-cli
```

### With Other Buildpacks

This buildpack can be combined with other buildpacks:

```bash
pack build my-app \
  --buildpack paketo-buildpacks/nodejs \
  --buildpack paketo-buildpacks/github-cli
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

- `GITHUB_CLI_VERSION`: Set to specify a particular version of GitHub CLI (defaults to "latest")
- `PATH`: Automatically updated to include the GitHub CLI binary location

### Authentication

To use GitHub CLI with authentication, you'll need to:

1. Set up a GitHub Personal Access Token
2. Configure authentication in your application or container environment
3. Use `gh auth login` or set the `GITHUB_TOKEN` environment variable

## `buildpack.yml` Configurations

There are no extra configurations for this buildpack based on `buildpack.yml`.

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

### GitHub Actions Workflow

The presence of a `.github/workflows/` directory will trigger this buildpack, making GitHub CLI available for CI/CD operations.

## Contributing

To contribute to this buildpack:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This buildpack is licensed under the Apache License 2.0. 