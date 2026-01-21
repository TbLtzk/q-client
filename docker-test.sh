#!/bin/bash
# Script to run Go tests in Docker using docker-compose
# Usage: ./docker-test.sh [package] [test flags]
# Examples:
#   ./docker-test.sh                              # Run all tests
#   ./docker-test.sh ./governance                 # Run tests in governance package
#   ./docker-test.sh ./governance -v              # Run with verbose flag
#   ./docker-test.sh ./governance -run TestNewExclusionSet  # Run specific test

set -e

# Build arguments for go test
if [ $# -eq 0 ]; then
    TEST_ARGS="./..."
else
    TEST_ARGS="$@"
fi

echo "Running Go tests in Docker: $*"

# Use docker-compose which handles Windows paths better
# Override the command to pass test arguments
# Build dependencies are pre-installed in the image
docker-compose -f docker-compose.dev.yml --profile test run --rm \
    test sh -c 'go test '"$*"

echo "Tests completed!"

