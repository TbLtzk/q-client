# Agent instructions

## GitHub operations

When modifying GitHub issues, pull requests, comments, or Projects from this repository, use the `GITHUB_PAT` environment secret as the GitHub token.

Examples:

- `GH_TOKEN="$GITHUB_PAT" gh issue view 2 --repo TbLtzk/q-client`
- `GH_TOKEN="$GITHUB_PAT" gh api graphql ...`

Do not rely on the default `gh` authentication for GitHub writes. It may only have read or metadata access.

## Completing issues

When an agent considers its assigned issue complete:

- Mark the pull request as ready for review.
- Comment on the issue with a link to the pull request, or explain why the issue is complete without a pull request.
- Move the issue to `In review` on the project board.
