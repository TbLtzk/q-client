#!/bin/bash
# Script to run goimports in Docker using docker-compose
# Usage: ./docker-goimports.sh [file1.go] [file2.go] ... or ./docker-goimports.sh . to format all files

set -e

FILES="${@:-.}"

echo "Running goimports in Docker on: $FILES"

# Build the command with arguments - escape properly for shell
ARGS="$@"
if [ $# -eq 0 ]; then
    ARGS="."
fi

# Use docker-compose which handles Windows paths better
# Override the command to include the file arguments
# goimports dependencies are pre-installed in the image
docker-compose -f docker-compose.dev.yml --profile goimports run --rm \
    goimports sh -c "go run golang.org/x/tools/cmd/goimports@v0.19.0 -w $ARGS"

echo "goimports completed!"

