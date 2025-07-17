# Contributing to Pace CLI

Thank you for your interest in contributing to Pace CLI! This document provides guidelines and information for contributors.

## Development Setup

### Prerequisites

- [Go 1.24.4+](https://golang.org/dl/)
- [Git](https://git-scm.com/)
- [Make](https://www.gnu.org/software/make/) (usually pre-installed on macOS/Linux)

### Optional Tools

- [golangci-lint](https://golangci-lint.run/) for linting
- [GoReleaser](https://goreleaser.com/) for releases

### Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/fakepixels/pace.git
   cd pace
   ```

2. **Set up development environment:**
   ```bash
   make setup
   ```

3. **Run the application:**
   ```bash
   make dev
   ```

4. **Run tests:**
   ```bash
   make test
   ```

## Development Workflow

### Available Make Commands

- `make dev` - Run the application in development mode
- `make serve` - Run the SSH server for development
- `make build` - Build binary for current platform
- `make test` - Run tests
- `make fmt` - Format code
- `make lint` - Run linter (requires golangci-lint)
- `make clean` - Clean build artifacts
- `make help` - Show all available commands

### Code Style

- Run `make fmt` before committing
- Follow standard Go conventions
- Use meaningful variable and function names
- Add comments for exported functions and complex logic

### Testing

- Write tests for new functionality
- Run `make test` to ensure all tests pass
- Maintain good test coverage

## Project Structure

```
pace/
├── main.go              # Main application entry point
├── announcement.md      # Announcement content
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Development commands
├── .goreleaser.yml     # Release configuration
├── .github/workflows/  # CI/CD workflows
├── public/             # Static assets
└── README.md           # Project documentation
```

## Making Changes

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Your Changes

- Follow the coding standards
- Add tests for new functionality
- Update documentation if needed

### 3. Test Your Changes

```bash
make test
make build
./bin/pace  # Test the binary
```

### 4. Commit Your Changes

```bash
git add .
git commit -m "Add meaningful commit message"
```

### 5. Push and Create PR

```bash
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

## Pull Request Guidelines

- Provide a clear description of the changes
- Include steps to test the changes
- Ensure CI checks pass
- Link to any related issues

## Release Process

Releases are automated using GoReleaser and GitHub Actions:

1. Create and push a git tag: `git tag -a v1.0.0 -m "Release v1.0.0"`
2. Push the tag: `git push origin v1.0.0`
3. GitHub Actions will automatically create a release

## Bug Reports

When reporting bugs, please include:

- Go version (`go version`)
- Operating system and version
- Steps to reproduce
- Expected vs actual behavior
- Any error messages

## Feature Requests

For feature requests:

- Describe the feature and its use case
- Explain how it would benefit users
- Provide examples if possible

## Questions?

- Open an issue for questions about development
- Check existing issues for common questions

## Code of Conduct

Please be respectful and inclusive in all interactions. This project follows standard open source community guidelines.

Thank you for contributing! 🎉