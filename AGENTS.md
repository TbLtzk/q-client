# Agent instructions

## GitHub operations

When modifying GitHub issues, pull requests, comments, or Projects from this repository, use the `GITHUB_PAT` environment secret as the GitHub token.

Examples:

- `GH_TOKEN="$GITHUB_PAT" gh issue view 2 --repo TbLtzk/q-client`
- `GH_TOKEN="$GITHUB_PAT" gh api graphql ...`

Do not rely on the default `gh` authentication for GitHub writes. It may only have read or metadata access.

## Starting issues

When an agent is told to work on an issue:

- If the issue is not in `Ready`, explain that the issue is not marked as ready and ask whether to start anyway.
- If the issue is already assigned to someone else, explain who is working on it and ask whether to reassign it and start anyway.
- If the issue is in `Ready`, acknowledge it, assign it to the agent where possible, move it to `In progress`, and start working on the issue.

### Issue description **and** comments

Treat the **entire GitHub issue thread** as part of the work spec, not only the opening description:

- Before planning or coding, read the **issue body and all existing comments** (design sketches, naming, constraints, acceptance tweaks, and replies). Those comments are **binding scope** alongside the description—especially when **starting** an issue, but also whenever **new comments** appear before merge.
- If newer comments supersede older text, follow the latest agreed direction; if two parts conflict and it is unclear which wins, ask the maintainer instead of guessing.
- If the implementation must **diverge** from something written in the issue or comments (for example RPC or Go symbol names), call that out explicitly in the PR or an issue comment so reviewers can reconcile it.

When using automation, fetch the full comment thread (for example `gh issue view <n> --comments` or the issue comments API), not just the body field.

## Completing issues

When an agent considers its assigned issue complete:

- Mark the pull request as ready for review.
- Comment on the issue with a link to the pull request, or explain why the issue is complete without a pull request.
- Move the issue to `In review` on the project board.

## Pre-merge QA

Before considering work ready for review, run the narrow checks that match the changed packages and include CI-like coverage for the affected area. When the host has no working Go toolchain, use `./docker-test.sh`, `./docker-lint.sh`, and `./docker-goimports.sh` from the repo root instead of bare `go test` / `go run build/ci.go lint`.

For external L0 governance changes, use at least:

- `go test -timeout=20m -p 1 ./governance`
- `go test ./governance ./ethclient ./internal/web3ext ./cmd/utils`
- `go test ./governance -run '^$'`
- `go run build/ci.go lint` when the local environment has the same native dependencies as CI, including `blst` headers.

Automated tests must not depend on live testnets, public RPC endpoints, or other external networks. Use in-process RPC servers, generated chains, fixtures, and mocks instead.
