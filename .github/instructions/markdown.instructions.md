---
applyTo: "**/*.md"
---

# Markdown Instructions

## Linting

We use **markdownlint** for markdown linting (via markdownlint-cli2-action in CI).

## Rules

- Headings must have 1 blank line above
- Multiple headings with same content allowed at different levels
- Standard markdownlint rules apply

## Excluded Files

- `CLAUDE.md` is excluded from linting

## Commands

```bash
# Lint markdown files (if markdownlint-cli installed locally)
markdownlint '**/*.md' --ignore CLAUDE.md
```
