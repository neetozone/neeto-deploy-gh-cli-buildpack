# GitHub CLI Sample Application

This is a sample Node.js application that demonstrates how to use the GitHub CLI buildpack.

## Features

- Express.js web server
- GitHub CLI integration
- REST API endpoints for GitHub operations

## Building with the GitHub CLI Buildpack

```bash
# Build the application with the GitHub CLI buildpack
pack build github-cli-sample \
  --buildpack paketo-buildpacks/nodejs \
  --buildpack paketo-buildpacks/github-cli \
  --path .

# Run the container
docker run -p 3000:3000 \
  -e GITHUB_TOKEN=your_github_token \
  github-cli-sample
```

## API Endpoints

- `GET /` - Application info and available endpoints
- `GET /version` - Get GitHub CLI version
- `GET /repos` - List your repositories (requires authentication)
- `GET /issues` - List issues from your repositories (requires authentication)

## Authentication

To use the GitHub CLI features, you need to authenticate:

1. Set the `GITHUB_TOKEN` environment variable with your GitHub Personal Access Token
2. Or run `gh auth login` inside the container

## Example Usage

```bash
# Check if GitHub CLI is available
curl http://localhost:3000/version

# List repositories (requires authentication)
curl http://localhost:3000/repos

# List issues (requires authentication)
curl http://localhost:3000/issues
```

## Development

To run locally without the buildpack:

```bash
# Install dependencies
npm install

# Install GitHub CLI manually
# Follow instructions at https://cli.github.com/

# Run the application
npm start
``` 