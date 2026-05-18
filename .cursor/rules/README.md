# Cursor project rules

Committed rules live as `*.mdc` files in this directory (tracked in git). They apply to everyone using the repo, including cloud agents.

## Machine-specific rules (not in git)

Use either layout below; both are listed in `.gitignore`:

| Location | Purpose |
|----------|---------|
| `.cursor/rules/local/*.mdc` | Organize local overrides in a subfolder |
| `.cursor/rules/*.local.mdc` | Flat files at this level (see loading note below) |

Example: `local/windows-docker-go.mdc` — Windows host without a local Go toolchain; use `./docker-test.sh`, `./docker-lint.sh`, `./docker-goimports.sh`.

### Does Cursor load rules in `local/`?

Support for **nested** `.mdc` files under `.cursor/rules/` is not documented as reliably as **flat** files in this directory. Some Cursor versions and multi-root workspaces have had issues with subdirectory rules.

If a rule in `local/` does not seem to apply:

1. Copy or rename it to `.cursor/rules/<name>.local.mdc` (same directory as this README), or
2. Move the content to **Cursor → Settings → Rules → User** for a machine-wide rule.

Cloud agents and CI do not read gitignored local rules.
