# Contributing to mcp-manager

Welcome! I'm [Alberto Tijunelis Neto](https://github.com/albertotijunelis) and I appreciate your interest in contributing to mcp-manager.

This project was originally built with [Claude](https://claude.ai) (Anthropic). AI-assisted contributions are welcome and encouraged.

---

## Easiest Contribution: Add a New MCP Server

The simplest and most impactful contribution is to **add a new MCP server** to the registry. Just edit `registry/registry.json` and open a Pull Request.

### JSON Schema

Each server entry must include these fields:

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | string | ✅ | Unique short name (lowercase, hyphens OK) |
| `display_name` | string | ✅ | Human-readable display name |
| `description` | string | ✅ | Clear, one-line description of what the server does |
| `repo` | string | ✅ | Repository path (e.g. `github.com/owner/repo`) |
| `install_type` | string | ✅ | One of: `go`, `npm`, `pip`, `git+build`, `binary` |
| `binary` | string | ✅ | Name of the executable binary |
| `package` | string | ❌ | npm/pip package name (if different from binary) |
| `version` | string | ✅ | Version constraint or `latest` |
| `homepage` | string | ✅ | URL to the project homepage or docs |
| `topics` | string[] | ✅ | Array of relevant tags for search |
| `requires_env` | string[] | ❌ | Environment variables required to run |
| `stars` | int | ✅ | Approximate GitHub stars (we'll verify) |

### Example Entry

```json
{
  "name": "my-server",
  "display_name": "My Amazing MCP Server",
  "description": "An MCP server that does something incredibly useful",
  "repo": "github.com/myorg/my-mcp-server",
  "install_type": "npm",
  "binary": "my-mcp-server",
  "package": "@myorg/my-mcp-server",
  "version": "latest",
  "homepage": "https://github.com/myorg/my-mcp-server",
  "topics": ["utility", "productivity"],
  "requires_env": ["MY_API_KEY"],
  "stars": 150
}
```

### Review Criteria

For a server to be accepted into the registry, it must:

1. **Implement the Model Context Protocol** — Must be a valid MCP server, not just any CLI tool
2. **Be publicly available** — The repo must be public on GitHub (or a public npm/pip package)
3. **Be actively maintained** — Should have recent commits and respond to issues
4. **Have accurate metadata** — Description, topics, and env vars must be correct
5. **Be installable** — The install type must work as expected (we test this in CI)

---

## Code Contributions

### Prerequisites

- Go 1.22+
- git

### Running Locally

```bash
# Clone the repo
git clone https://github.com/albertotijunelis/mcp-manager.git
cd mcp-manager

# Build
make build

# Run
./bin/mcp --help

# Run tests
make test

# Lint (requires golangci-lint)
make lint
```

### Project Structure

```
mcp-manager/
├── cmd/                  # CLI commands (cobra)
│   ├── root.go           # Root command and global flags
│   ├── install.go        # mcp install
│   ├── remove.go         # mcp remove
│   ├── list.go           # mcp list
│   ├── run.go            # mcp run
│   ├── search.go         # mcp search
│   ├── update.go         # mcp update
│   ├── info.go           # mcp info
│   └── doctor.go         # mcp doctor
├── internal/
│   ├── registry/         # Registry loading, saving, searching
│   ├── config/           # User config management
│   ├── installer/        # Install/remove logic for all types
│   └── ui/               # Styles, spinner, TUI components
├── registry/
│   └── registry.json     # The server registry
├── main.go               # Entry point
├── go.mod
├── Makefile
└── README.md
```

### Code Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Handle all errors explicitly
- Use `context.Context` where appropriate
- Add the copyright header to every `.go` file:
  ```go
  // Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.
  ```

---

## Pull Request Guidelines

1. Fork the repo and create a branch from `main`
2. If you've added code, add or update tests
3. Ensure `go build ./...` and `go test ./...` pass
4. Write a clear PR description explaining what and why
5. Reference any relevant issues

---

## Code of Conduct

Be kind, be respectful. We follow the standard [Contributor Covenant](https://www.contributor-covenant.org/). Harassment, discrimination, or hostility of any kind will not be tolerated.

---

## Questions?

Open an issue or reach out to [@albertotijunelis](https://github.com/albertotijunelis). Happy hacking!
