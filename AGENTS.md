# AGENTS.md

## Tooling

This project uses [mise](https://mise.jdx.dev/) for task management and development environment configuration.

### Setup

1.  Install `mise`.
2.  Run `mise install` to install dependencies (Go, linters).

### Standard Tasks

*   `mise run lint`: Runs all linters (golangci-lint, actionlint).
*   `mise run fmt`: Runs formatters (go fmt).
*   `mise run test`: Runs tests.
*   `mise run ci`: Runs the CI pipeline (lint + test).
*   `mise run codegen`: Updates generated code.
*   `mise run install`: Installs dependencies.

## CI/CD

The project uses GitHub Actions with a single workflow `.github/workflows/autorelease.yml`.
This workflow follows the standard flow: Install -> Codegen -> PR (if changes) -> CI -> Release.

## Conventions

*   **Go Version**: 1.24.0 (managed by mise).
*   **Linters**: golangci-lint (Go), actionlint (GitHub Actions).
*   **Formatting**: go fmt.
*   **Error Handling**: Check all errors. Do not use `_` to ignore errors unless absolutely necessary and documented.
*   **Context**: Use `context.Context` for cancellation and timeouts.
