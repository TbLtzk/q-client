#!/bin/bash
# Script to run golangci-lint in Docker using docker-compose
# Usage: ./docker-lint.sh [file or directory paths]
# Examples:
#   ./docker-lint.sh                    # Run on all files
#   ./docker-lint.sh ./params/version.go  # Run on specific file
#   ./docker-lint.sh ./params            # Run on specific directory

set -e

# Build arguments for golangci-lint
if [ $# -eq 0 ]; then
    LINT_ARGS="./..."
else
    LINT_ARGS="$*"
fi

echo "Running golangci-lint in Docker on: $*"

# Use docker-compose which handles Windows paths better
# Override the command to pass lint arguments
# golangci-lint is pre-installed in the image
docker-compose -f docker-compose.dev.yml --profile lint run --rm \
    golangci-lint sh -c 'golangci-lint run --config .golangci.yml '"$*"

echo "Linting completed successfully!"

