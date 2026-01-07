---
applyTo: "**/*.md"
---

# Markdown Instructions

## Linting

We use **markdownlint** for markdown linting (via markdownlint-cli2-action in CI).

## Rules

- Headings must have 1 blank line above
- Multiple headings with same content allowed (siblings_only)
- No line length limit (MD013 disabled)
- First line heading not required (MD041 disabled)

## Configuration

See `.markdownlint.json` for linting rules.

## Excluded Files

- `CLAUDE.md` is excluded from linting

## Commands

```bash
# Lint markdown files (if markdownlint-cli installed locally)
markdownlint '**/*.md' --ignore CLAUDE.md
```
