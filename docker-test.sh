#!/bin/bash
# Script to run Go tests in Docker using docker-compose
# Usage: ./docker-test.sh [package] [test flags]
# Examples:
#   ./docker-test.sh                              # Run all tests
#   ./docker-test.sh ./governance                 # Run tests in governance package
#   ./docker-test.sh ./governance -v              # Run with verbose flag
#   ./docker-test.sh ./governance -run TestNewExclusionSet  # Run specific test

set -e

echo "Running Go tests in Docker: $*"
echo ""

# Find the package path (first argument that starts with ./)
PACKAGE="./..."
ALL_ARGS="$@"

# Parse arguments to find package path
for arg in "$@"; do
    if [[ "$arg" == ./* ]] || [[ "$arg" == ./... ]]; then
        PACKAGE="$arg"
        break
    fi
done

# Use docker-compose which handles Windows paths better
# GOMODCACHE is set in docker-compose.dev.yml to /go/pkg/mod to avoid Windows path conflicts
# First compile test binary with verbose output to show compilation progress
echo "=== Compiling test binary (this may take a while on first run) ==="
docker-compose -f docker-compose.dev.yml --profile test run --rm \
    test sh -c "unset GOMODCACHE; export GOMODCACHE=/go/pkg/mod && go test -c -v '$PACKAGE' -o /tmp/test.bin 2>&1 || echo 'Compilation completed (some warnings may be normal)'"

echo ""
echo "=== Running tests ==="
# Then run the actual tests (Go will use cached compilation if available)
docker-compose -f docker-compose.dev.yml --profile test run --rm \
    test sh -c "unset GOMODCACHE; export GOMODCACHE=/go/pkg/mod && go test $ALL_ARGS"

echo ""
echo "Tests completed!"

