# Contributing to SIG‑Agro Backend

Thank you for your interest in contributing to SIG‑Agro! This document outlines the guidelines for contributing to the project. By participating, you agree to abide by the terms of the [Apache License 2.0](LICENSE) under which this project is distributed.

## Code of Conduct

We are committed to providing a welcoming and harassment‑free experience for everyone. Please be respectful and constructive in all interactions. Unacceptable behavior will not be tolerated.

## How to Contribute

### Reporting Bugs or Requesting Features

- Use the [issue tracker](https://github.com/your-org/sig-agro-backend/issues) to report bugs or suggest enhancements.
- Before creating a new issue, please search existing issues to avoid duplicates.
- For bugs, include:
  - A clear description of the problem.
  - Steps to reproduce the behavior.
  - Expected vs. actual behavior.
  - Environment details (OS, Go version, etc.).
- For features, describe the use case and the value it would bring.

### Submitting Changes (Pull Requests)

1. **Fork the repository** and create a new branch from `main`.
2. **Follow the coding standards**:
   - Go code should be formatted with `go fmt`.
   - Follow the existing project structure (Clean Architecture).
   - Write tests for new functionality.
   - Update documentation as needed (README, comments).
3. **Commit messages** should be clear and descriptive. Use the imperative mood ("Add feature", not "Added feature").
4. **Pull request (PR) guidelines**:
   - Provide a descriptive title and summary.
   - Reference any related issues (e.g., `Fixes #123`).
   - Ensure all CI checks pass (lint, tests).
   - Keep PRs focused; avoid mixing unrelated changes.
5. **Sign your commits** (optional but encouraged).

### Development Setup

Please refer to the main [README.md](README.md) for instructions on setting up the development environment with Docker Compose.

### Code Style

- **Go**: Use `gofmt` and `golangci-lint`. The repository includes a `.golangci.yml` configuration.
- **Protocol Buffers**: Use the official style guide.
- **SQL**: Use lowercase keywords, consistent indentation.

### Testing

- Write unit tests for business logic (usecases) and integration tests for repositories.
- Run all tests with `go test ./...` before submitting a PR.
- For gRPC services, consider using `grpcurl` or a mock client for testing.

### Documentation

- Update the README or service‑specific documentation when changing user‑facing behavior.
- Document new environment variables in the service’s README or config section.

## Review Process

- Maintainers will review PRs and may request changes.
- Once approved, the PR will be merged by a maintainer.
- All contributions will be acknowledged.

## Licensing

By contributing to this project, you agree that your contributions will be licensed under the [Apache License 2.0](LICENSE). You retain the copyright of your contributions.

---

Thank you for helping build SIG‑Agro!