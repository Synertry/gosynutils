# Contributing to GoSynUtils

Thank you for your interest in contributing to GoSynUtils!

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Development Workflow](#development-workflow)
    - [Branching Strategy](#branching-strategy)
    - [Coding Standards](#coding-standards)
    - [Testing](#testing)
    - [Documentation](#documentation)
- [Pull Request Process](#pull-request-process)
- [Issue Reporting](#issue-reporting)
- [Communication](#communication)

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Development Workflow

Always develop for the latest stable Go version available: [Latest Go Version](https://go.dev/dl/)
There is little to no use as a library to not support the latest version.
Apart from that, Go is very stable in itself, so it is not a problem to use the latest version.

### Branching Strategy

- `main` branch is always stable and deployable
- Create feature branches from `main`:
  ```bash
  git checkout -b feature/your-feature-name
  # or
  git checkout -b fix/issue-you-are-fixing
  ```

### Coding Standards

#### No Shadow Variables
Means a new inner variable should not have the same name as an outer variable.
This is common for error variables, which get shadowed in a lot of public snippets.

Bad example:
```go
func Example() {
    err := doSomething()
    if err != nil {
        // This shadows the outer 'err' variable
        err := doSomethingElse()
        if err != nil {
            // This is confusing and can lead to bugs
        }
    }
}
```

Good example:
```go
func Example() {
    err := doSomething()
    if err != nil {
        // Use a different name for the inner variable
        errInner := doSomethingElse()
        if errInner != nil {
            // This is clearer and avoids shadowing
		}
	}
```

### Testing

- All new code should be accompanied by tests
- Aim for high test coverage on new functionality

### Documentation

Document all exported functions, types, and packages



## Pull Request Process

1. Update your fork to include the latest changes from upstream:
   ```bash
   git fetch origin
   git rebase origin/main
   ```

2. Ensure your code adheres to our coding standards and passes all tests

3. Submit a pull request with a clear title and description:
    - What does this PR do?
    - Any specific issues or challenges to note?
    - References to related issues or discussions

4. Your PR will be reviewed by maintainers who may request changes

5. Once approved, a maintainer will merge your PR

## Issue Reporting

- Use the issue tracker to report bugs or propose features
- For bugs, include:
    - A clear title and description
    - Steps to reproduce
    - Expected vs. actual behavior
    - Version information (Go version, OS, etc.)
    - Screenshots or logs if applicable

- For feature requests, include:
    - The problem you're trying to solve
    - Proposed solution or ideas
    - Any alternatives you've considered

## Communication

- For quick questions, open a [Discussion](https://github.com/Synertry/gosynutils/discussions)
- For bugs and feature requests, use [Issues](https://github.com/Synertry/gosynutils/issues)
- For significant changes, consider discussing in an issue before implementing

---

Thank you for contributing!