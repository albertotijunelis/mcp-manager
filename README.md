<p align="center">
  <pre align="center">
╔╦╗╔═╗╔═╗   ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗
║║║║  ╠═╝───║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝
╩ ╩╚═╝╩     ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═
  </pre>
</p>

<h3 align="center"><strong>The package manager for MCP servers</strong></h3>
<p align="center"><em>Install, manage and run Model Context Protocol servers — like Homebrew for your AI toolchain</em></p>

<p align="center">
  <a href="https://github.com/albertotijunelis/mcp-manager/actions"><img src="https://img.shields.io/github/actions/workflow/status/albertotijunelis/mcp-manager/ci.yml?style=flat-square&label=CI" alt="CI"></a>
  <a href="https://github.com/albertotijunelis/mcp-manager/releases"><img src="https://img.shields.io/github/v/release/albertotijunelis/mcp-manager?style=flat-square" alt="Release"></a>
  <img src="https://img.shields.io/badge/go-1.22-00D4FF?style=flat-square" alt="Go">
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-00FF9F?style=flat-square" alt="License"></a>
  <a href="https://github.com/albertotijunelis/mcp-manager/stargazers"><img src="https://img.shields.io/github/stars/albertotijunelis/mcp-manager?style=flat-square" alt="Stars"></a>
  <a href="https://github.com/albertotijunelis"><img src="https://img.shields.io/badge/author-Alberto%20Tijunelis%20Neto-a855f7?style=flat-square&logo=github" alt="Author"></a>
  <a href="https://github.com/charmbracelet/bubbletea"><img src="https://img.shields.io/badge/built%20with-Bubble%20Tea-FF75B7?style=flat-square" alt="Bubble Tea"></a>
  <a href="https://claude.ai"><img src="https://img.shields.io/badge/built%20with-Claude%20Sonnet-CC785C?style=flat-square&logo=anthropic" alt="Built with Claude"></a>
</p>

---

## What is this?

The **Model Context Protocol (MCP)** is an open standard for connecting AI models to external tools and data sources. GitHub Copilot, Claude Code, Cursor, and Windsurf all support MCP servers — but there's no easy way to discover, install, and manage them.

**mcp-manager** fixes that. It's a single CLI that acts as a package manager for the entire MCP ecosystem. Think Homebrew, but for your AI toolchain.

Install any MCP server in seconds, keep them updated, and run them with a single command.

---

## Demo

```
$ mcp install github

  ╭──────────────────────────────────────────────────────────────╮
  │  Install Server                                              │
  │                                                              │
  │  Name         │ GitHub MCP Server                            │
  │  Description  │ Official GitHub MCP server for repository,   │
  │               │ PR, issue management, code search, and       │
  │               │ GitHub Actions workflows                     │
  │  Install Type │ go                                           │
  │  Repository   │ github.com/github/github-mcp-server          │
  │  Version      │ latest                                       │
  │  Required Env │ GITHUB_TOKEN                                 │
  ╰──────────────────────────────────────────────────────────────╯

  Install this server? [Y/n] y

  ⠋ Installing GitHub MCP Server...
  ✓ Installing GitHub MCP Server...

  ✓ Successfully installed GitHub MCP Server!

  ⚠ This server requires environment variables:
    export GITHUB_TOKEN=<your-value>
```

```
$ mcp list

  Installed MCP Servers

  Name     Version   Status      Install Path              Installed At
  ──────   ───────   ──────      ────────────              ────────────
  github   latest    ● stopped   ~/.mcp/servers/github     2026-03-19

  1 server(s) installed
```

```
$ mcp search database

  Search results for 'database'

  Name       Description                                              Stars   Type
  ────       ───────────                                              ─────   ────
  sqlite     MCP server for querying and managing SQLite databases..  12.5k   npm
  postgres   MCP server for querying PostgreSQL databases with...     12.5k   npm

  2 result(s) found
```

```
$ mcp run github

  → Starting github (vlatest)
  Press Ctrl+C to stop
```

---

## Install

### Homebrew (recommended)

```bash
brew install albertotijunelis/tap/mcp-manager
```

> **Note:** This requires the Homebrew tap at [albertotijunelis/homebrew-tap](https://github.com/albertotijunelis/homebrew-tap). The tap is auto-published by GoReleaser on each release.

### curl one-liner

```bash
curl -sSfL https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/install.sh | sh
```

### Go install

```bash
go install github.com/albertotijunelis/mcp-manager@latest
```

### From source

```bash
git clone https://github.com/albertotijunelis/mcp-manager.git
cd mcp-manager
make build
./bin/mcp --help
```

---

## Quick Start

Get a running MCP server in 60 seconds:

**1. Check your system is ready**

```bash
mcp doctor
```

**2. Install an MCP server**

```bash
mcp install github
```

**3. Run it**

```bash
export GITHUB_TOKEN=ghp_your_token_here
mcp run github
```

That's it. Your AI tools (Copilot, Claude, Cursor) can now connect to the GitHub MCP server.

---

## Commands

| Command | Description | Example |
|---|---|---|
| `mcp install <name>` | Install an MCP server from the registry | `mcp install github` |
| `mcp remove <name>` | Remove an installed server and clean up | `mcp remove github` |
| `mcp list` | List all installed servers with status | `mcp list` |
| `mcp run <name>` | Launch an MCP server process | `mcp run github` |
| `mcp search <query>` | Search the server registry | `mcp search database` |
| `mcp update [name]` | Update servers or sync the registry | `mcp update --all` |
| `mcp info <name>` | Show detailed server information | `mcp info github` |
| `mcp doctor` | Check system dependencies and health | `mcp doctor` |
| `mcp version` | Print version information | `mcp version` |

### Global flags

| Flag | Description |
|---|---|
| `--no-color` | Disable colored output |
| `--debug` | Enable debug logging |
| `--json` | Output in JSON format |

---

## Registry

The registry is a curated JSON file containing metadata for all supported MCP servers. It ships with **20+ servers** out of the box and is synced from the upstream repository.

Browse the full registry:
[`registry/registry.json`](registry/registry.json)

### Sync the registry

```bash
mcp update
```

This fetches the latest server list from GitHub so you always have access to newly added servers.

---

## Adding a Server

Want to add an MCP server to the registry? Edit `registry/registry.json` and open a PR!

Each server entry looks like this:

```json
{
  "name": "my-server",
  "display_name": "My MCP Server",
  "description": "A clear description of what this server does",
  "repo": "github.com/owner/repo",
  "install_type": "npm",
  "binary": "my-mcp-server",
  "package": "@scope/my-mcp-server",
  "version": "latest",
  "homepage": "https://github.com/owner/repo",
  "topics": ["relevant", "tags"],
  "requires_env": ["API_KEY"],
  "stars": 100
}
```

### Review criteria

- The server must implement the Model Context Protocol
- The repository must be public and actively maintained
- The description must be accurate and helpful
- The install type must be correct and tested

---

## Supported Install Types

| Type | How it works | Example |
|---|---|---|
| `go` | Runs `go install repo@version` | GitHub MCP Server |
| `npm` | Runs `npm install -g package` | Filesystem, SQLite, Slack |
| `pip` | Runs `pip install package` | Python-based servers |
| `git+build` | Clones repo and runs build command | Custom servers |
| `binary` | Downloads pre-built binary from GitHub Releases | Platform-specific servers |

---

## Built with Claude

This project was fully scaffolded and built with [Claude](https://claude.ai) by Anthropic — including all Go source code, the registry, CI/CD pipelines, and this README.

The entire codebase was generated via a single structured prompt and then refined by the author. This is an example of what's possible when a developer uses Claude as a serious engineering co-pilot.

If you're curious about AI-assisted open source, feel free to open an issue or reach out.

---

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

The easiest way to contribute is to **add a new MCP server** to the registry — just edit `registry/registry.json` and open a PR.

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

<p align="center">
  Built with ❤️ by <a href="https://github.com/albertotijunelis">Alberto Tijunelis Neto</a> · Powered by <a href="https://github.com/charmbracelet/bubbletea">Bubble Tea</a>
</p>
