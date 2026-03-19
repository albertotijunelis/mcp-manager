# INSTRUCTIONS FOR CLAUDE

You have been uploaded this file as your complete build specification.
Read everything below and immediately start building the project.
Do not ask questions. Do not ask for confirmation. Just build it.

---

# YOUR ROLE

You are an elite Go engineer and open-source maintainer with deep expertise in building viral developer tools. Your task is to build **mcp-manager** — a production-ready CLI that acts as a package manager for MCP (Model Context Protocol) servers, similar to Homebrew but for the AI agent toolchain.

---

# CONTEXT: What is MCP?

The Model Context Protocol is an open standard for connecting AI models to external tools and data sources. GitHub Copilot, Claude Code, Cursor, and Windsurf all support MCP servers. Developers currently must manually clone repos, configure processes, and manage dependencies by hand. There is no package manager. mcp-manager fills this gap.

---

# YOUR STANDARDS

- Every file is complete and production-ready. Zero TODOs, zero placeholders.
- Go idioms: handle all errors, use context, avoid global state.
- Beautiful terminal output: colors, spinners, progress bars, tables.
- The project must compile and run with `go run main.go install github` immediately after setup.
- README must be viral-optimised: ASCII logo, badges, demo section, quick-start in 30 seconds.
- Think like a maintainer who wants 1,000 GitHub stars on day one.

---

# AUTHOR & LICENSE

- Author: Alberto Tijunelis Neto
- Email: albertotijunelis@gmail.com
- GitHub: github.com/albertotijunelis
- License: MIT
- Copyright notice in every .go file header: `// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.`
- LICENSE file: full MIT text with "Copyright (c) 2026 Alberto Tijunelis Neto"
- README author badge linking to github.com/albertotijunelis
- All git commit author: `Alberto Tijunelis Neto <albertotijunelis@gmail.com>`
- go.mod: use `github.com/albertotijunelis/mcp-manager`

# CLAUDE AS CO-AUTHOR — THIS IS CRITICAL FOR VIRALITY

This project was built with Claude (Anthropic). This must be reflected everywhere:

## 1. Every git commit message must end with:
```
Co-authored-by: Claude <claude@anthropic.com>
```

## 2. README.md must include this badge in the badges row:
```
[![Built with Claude](https://img.shields.io/badge/built%20with-Claude%20Sonnet-CC785C?style=flat-square&logo=anthropic)](https://claude.ai)
```

## 3. README.md must include a dedicated section called "Built with Claude":
```markdown
## Built with Claude

This project was fully scaffolded and built with [Claude](https://claude.ai) by Anthropic — including all Go source code, the registry, CI/CD pipelines, and this README.

The entire codebase was generated via a single structured prompt and then refined by the author. This is an example of what's possible when a developer uses Claude as a serious engineering co-pilot.

If you're curious about AI-assisted open source, feel free to open an issue or reach out.
```

## 4. CONTRIBUTING.md must mention:
"This project was originally built with Claude (Anthropic). AI-assisted contributions are welcome and encouraged."

## 5. The initial commit message must be exactly:
```
feat: initial release of mcp-manager v0.1.0

The package manager for MCP servers — like Homebrew for your AI toolchain.

✦ Built entirely with Claude by Anthropic
✦ 27 files, 20+ MCP servers in registry on day one
✦ Beautiful TUI with Charm's Bubble Tea + Lip Gloss
✦ Cross-platform binaries via GoReleaser
✦ MIT License

Co-authored-by: Claude <claude@anthropic.com>
```

---

# TECH STACK

- CLI framework: github.com/spf13/cobra
- Terminal UI: github.com/charmbracelet/bubbletea + github.com/charmbracelet/lipgloss
- Progress bars: github.com/charmbracelet/bubbles/progress
- Spinner: github.com/charmbracelet/bubbles/spinner
- HTTP: standard net/http with retries
- JSON: standard encoding/json
- Go version: 1.22
- Module: github.com/albertotijunelis/mcp-manager

---

# COLOR PALETTE (Lip Gloss)

- Primary / cyan:  #00D4FF
- Success / green: #00FF9F
- Warning / amber: #FFB800
- Error / red:     #FF4757
- Muted / gray:    #64748B
- Background: default terminal bg (no hardcoded bg)

---

# REGISTRY FORMAT (~/.mcp/registry.json)

```json
{
  "version": "1.0.0",
  "updated_at": "2026-03-19T00:00:00Z",
  "registry_url": "https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/registry/registry.json",
  "servers": [
    {
      "name": "github",
      "display_name": "GitHub MCP Server",
      "description": "Official GitHub MCP server for repository, PR, and issue management",
      "repo": "github.com/github/github-mcp-server",
      "install_type": "go",
      "binary": "github-mcp-server",
      "version": "latest",
      "homepage": "https://github.com/github/github-mcp-server",
      "topics": ["github", "api", "developer"],
      "requires_env": ["GITHUB_TOKEN"],
      "stars": 8420
    }
  ]
}
```

---

# CONFIG FORMAT (~/.mcp/config.json)

```json
{
  "version": "1",
  "installed": {
    "github": {
      "name": "github",
      "version": "v1.0.0",
      "install_path": "/Users/user/.mcp/servers/github",
      "binary_path": "/Users/user/.mcp/bin/github-mcp-server",
      "installed_at": "2026-03-19T12:00:00Z",
      "env": {"GITHUB_TOKEN": ""}
    }
  }
}
```

---

# INSTALL TYPES TO SUPPORT

- `go`        — run `go install {repo}@{version}`
- `npm`       — run `npm install -g {package}`
- `pip`       — run `pip install {package}`
- `git+build` — clone repo, then run build command
- `binary`    — download pre-built binary from GitHub releases

---

# COMMANDS TO IMPLEMENT

| Command | Description |
|---|---|
| `mcp install <name>` | Install from registry with spinner + progress |
| `mcp remove <name>` | Uninstall and clean up |
| `mcp list` | Table of installed servers with status |
| `mcp run <name> [args...]` | Launch the MCP server process |
| `mcp search <query>` | Fuzzy search the registry |
| `mcp update [name]` | Update one or all servers + sync registry |
| `mcp info <name>` | Detailed server info card |
| `mcp doctor` | Check Go, Node, Python, git are available |

---

# OUTPUT FORMAT FOR EVERY FILE

Use this exact format for each file. Never truncate. Never use "...":

```
=== FILE: path/to/filename.ext ===
[complete file content]
=== END FILE ===
```

---

# FILES TO GENERATE — BUILD ALL OF THEM IN ORDER

## FILE 01 — go.mod

Module: github.com/albertotijunelis/mcp-manager
Go version: 1.22
Dependencies: cobra, bubbletea, lipgloss, bubbles

## FILE 02 — main.go

Entry point. Calls cmd.Execute(). Sets version from ldflags.
Add copyright header: // Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

## FILE 03 — cmd/root.go

Root cobra command.
Global flags: --no-color, --debug, --json.
On first run, show this ASCII logo:

```
╔╦╗╔═╗╔═╗   ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗
║║║║  ╠═╝───║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝
╩ ╩╚═╝╩     ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═
The package manager for MCP servers
```

## FILE 04 — cmd/install.go

`mcp install <name> [--version v1.2.3]`
- Look up name in registry
- Show server info card (display name, description, install type, required env vars)
- Ask for confirmation (skip with --yes)
- Show spinner during install
- Show progress bar if downloading binary
- After install: print required env vars with setup instructions
- Update ~/.mcp/config.json

## FILE 05 — cmd/remove.go

`mcp remove <name> [--yes]`
- Show what will be removed
- Confirmation prompt (skip with --yes)
- Remove binary and config entry
- Clean up ~/.mcp/servers/{name} directory

## FILE 06 — cmd/list.go

`mcp list [--json]`
Beautiful lipgloss table with columns:
  Name | Version | Status | Install Type | Installed At
Green dot (●) for running, gray dot for stopped.
If no servers installed: show helpful empty state with install hint.

## FILE 07 — cmd/run.go

`mcp run <name> [-- args...]`
- Check server is installed (error if not)
- Load env vars from config
- Print startup message with server name and transport info
- exec the binary with any extra args passed after --

## FILE 08 — cmd/search.go

`mcp search <query> [--limit 10]`
- Fuzzy match against name, description, topics
- Table output: Name | Description | Stars | Install Type
- Highlight matching text in results
- Show total match count below table

## FILE 09 — cmd/update.go

`mcp update [name] [--all]`
- Without args: sync registry from upstream URL, show diff of new/updated servers
- With name: reinstall that server to latest version
- With --all: update every installed server

## FILE 10 — cmd/info.go

`mcp info <name>`
Full info card using lipgloss box:
  - Display name
  - Description
  - Repository URL
  - Homepage
  - Topics (as colored badges)
  - Stars
  - Install type
  - Required env vars (with descriptions)
  - If installed: version, install path, installed date

## FILE 11 — cmd/doctor.go

`mcp doctor`
Check each of these and print ✓ or ✗ with fix hints:
  - Go installed + show version
  - Node.js installed + show version
  - Python installed + show version  
  - git installed + show version
  - ~/.mcp directory exists and is writable
  - Registry file exists
  - Registry URL is reachable (HTTP GET)
  - Each installed server binary exists at expected path

## FILE 12 — internal/registry/registry.go

```go
type Server struct {
    Name        string   `json:"name"`
    DisplayName string   `json:"display_name"`
    Description string   `json:"description"`
    Repo        string   `json:"repo"`
    InstallType string   `json:"install_type"`
    Binary      string   `json:"binary"`
    Version     string   `json:"version"`
    Homepage    string   `json:"homepage"`
    Topics      []string `json:"topics"`
    RequiresEnv []string `json:"requires_env"`
    Stars       int      `json:"stars"`
}

type Registry struct {
    Version     string    `json:"version"`
    UpdatedAt   time.Time `json:"updated_at"`
    RegistryURL string    `json:"registry_url"`
    Servers     []Server  `json:"servers"`
}
```

Functions: Load(), Save(), Search(query string), FindByName(name string), Sync() error
Default path: ~/.mcp/registry.json
Sync() fetches from RegistryURL, merges, saves.

## FILE 13 — internal/config/config.go

```go
type InstalledServer struct {
    Name        string            `json:"name"`
    Version     string            `json:"version"`
    InstallPath string            `json:"install_path"`
    BinaryPath  string            `json:"binary_path"`
    InstalledAt time.Time         `json:"installed_at"`
    Env         map[string]string `json:"env"`
}

type Config struct {
    Version   string                     `json:"version"`
    Installed map[string]InstalledServer `json:"installed"`
}
```

Functions: Load(), Save(), Install(server), Remove(name), IsInstalled(name), GetInstalled(name)
Default path: ~/.mcp/config.json

## FILE 14 — internal/installer/installer.go

Installer interface:
```go
type Installer interface {
    Install(server registry.Server) error
    Remove(server registry.Server) error
}
```

Implement all five:
- GoInstaller   — runs `go install {repo}@{version}`, finds binary in GOPATH/bin
- NpmInstaller  — runs `npm install -g {package}`
- PipInstaller  — runs `pip install {package}`
- GitInstaller  — clones repo to ~/.mcp/servers/{name}, runs build command, links binary
- BinaryInstaller — downloads from GitHub releases API, verifies checksum, chmod +x

Factory function: NewInstaller(installType string) (Installer, error)
Stream all command output to the terminal in real time.

## FILE 15 — internal/ui/styles.go

All lipgloss styles as package-level vars.

Colors:
```go
const (
    ColorCyan   = "#00D4FF"
    ColorGreen  = "#00FF9F"
    ColorAmber  = "#FFB800"
    ColorRed    = "#FF4757"
    ColorMuted  = "#64748B"
)
```

Styles to define:
- TitleStyle, SubtitleStyle
- SuccessStyle, ErrorStyle, WarningStyle, InfoStyle
- TableHeaderStyle, TableRowStyle, TableAltRowStyle
- BadgeStyle (for topics)
- InfoCardStyle (lipgloss box)
- SeparatorStyle
- LogoStyle

Helper functions:
- RenderTable(headers []string, rows [][]string) string
- RenderInfoCard(title string, fields [][]string) string
- PrintSuccess(msg string)
- PrintError(msg string)
- PrintWarning(msg string)
- PrintInfo(msg string)
- RenderBadge(text string) string

## FILE 16 — internal/ui/spinner.go

Bubble Tea spinner model.

Provide:
- A SpinnerModel struct (implements tea.Model)
- Start(label string), Stop()
- RunWithSpinner(label string, fn func() error) error — runs fn in a goroutine, shows spinner until done, returns error

Use charmbracelet/bubbles/spinner with the Dot spinner style in cyan.

## FILE 17 — registry/registry.json

Include EXACTLY these 20 MCP servers with accurate real-world data:

1.  github         — github.com/github/github-mcp-server          — go
2.  filesystem     — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-filesystem)
3.  sqlite         — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-sqlite)
4.  playwright     — github.com/microsoft/playwright-mcp           — npm (package: @playwright/mcp)
5.  brave-search   — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-brave-search)
6.  slack          — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-slack)
7.  postgres       — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-postgres)
8.  google-drive   — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-gdrive)
9.  memory         — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-memory)
10. time           — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-time)
11. fetch          — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-fetch)
12. puppeteer      — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-puppeteer)
13. aws-kb         — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-aws-kb-retrieval)
14. google-maps    — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-google-maps)
15. gitlab         — github.com/modelcontextprotocol/servers       — npm (package: @modelcontextprotocol/server-gitlab)
16. sentry         — github.com/getsentry/sentry-mcp               — go
17. linear         — github.com/linear/linear-mcp-server           — npm
18. notion         — github.com/makenotion/notion-mcp-server       — npm
19. stripe         — github.com/stripe/agent-toolkit               — npm (package: @stripe/agent-toolkit)
20. cloudflare     — github.com/cloudflare/mcp-server-cloudflare   — npm

For each server write a detailed, accurate description, correct topics array, realistic stars count, and real required_env values.

## FILE 18 — README.md

Use this exact ASCII logo at the top:

```
╔╦╗╔═╗╔═╗   ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗
║║║║  ╠═╝───║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝
╩ ╩╚═╝╩     ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═
```

Tagline: **The package manager for MCP servers**
Subtitle: *Install, manage and run Model Context Protocol servers — like Homebrew for your AI toolchain*

Badges row (shields.io, flat-square style):
- CI: ![CI](https://img.shields.io/github/actions/workflow/status/albertotijunelis/mcp-manager/ci.yml?style=flat-square&label=CI)
- Release: ![Release](https://img.shields.io/github/v/release/albertotijunelis/mcp-manager?style=flat-square)
- Go version: ![Go](https://img.shields.io/badge/go-1.22-00D4FF?style=flat-square)
- License: ![License](https://img.shields.io/badge/license-MIT-00FF9F?style=flat-square)
- Stars: ![Stars](https://img.shields.io/github/stars/albertotijunelis/mcp-manager?style=flat-square)
- Author: ![Author](https://img.shields.io/badge/author-Alberto%20Tijunelis%20Neto-a855f7?style=flat-square&logo=github)
- Built with: ![Bubble Tea](https://img.shields.io/badge/built%20with-Bubble%20Tea-FF75B7?style=flat-square)

Sections in order:
1. Logo + badges
2. **What is this?** — 3-sentence explanation of MCP and the problem
3. **Demo** — realistic terminal output using Unicode box-drawing chars for these 4 commands:
   - $ mcp install github
   - $ mcp list
   - $ mcp search database
   - $ mcp run github
4. **Install** — three methods:
   ```bash
   # Homebrew (recommended)
   brew install albertotijunelis/tap/mcp-manager

   # curl one-liner
   curl -sSfL https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/install.sh | sh

   # Go install
   go install github.com/albertotijunelis/mcp-manager@latest
   ```
5. **Quick start** — 3 steps, get a running MCP server in 60 seconds
6. **Commands** — full table: Command | Description | Example
7. **Registry** — what it is, how to browse, link to registry/registry.json
8. **Adding a server** — JSON snippet + "open a PR" call to action
9. **Supported install types** — table: Type | How it works | Example
10. **Contributing** — link to CONTRIBUTING.md
11. **License** — MIT, link to LICENSE
12. Footer: `Built with ❤️ by [Alberto Tijunelis Neto](https://github.com/albertotijunelis) · Powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)`

## FILE 19 — LICENSE

Full MIT License text:

```
MIT License

Copyright (c) 2026 Alberto Tijunelis Neto

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## FILE 20 — CONTRIBUTING.md

Contents:
- Welcome message from Alberto Tijunelis Neto
- **Easiest contribution: add a new MCP server** — just edit registry/registry.json and open a PR
- Full JSON schema with field descriptions and examples
- Review criteria (what makes a server eligible)
- How to run locally: go build, go test
- Code of conduct reference
- PR template guidance

## FILE 21 — .goreleaser.yml

Full GoReleaser v2 config:
- Builds: darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, windows/amd64
- Archives: .tar.gz for unix, .zip for windows
- Binary name: `mcp` (short, memorable — `mcp install github` beats `mcp-manager install github`)
- ldflags: -X main.version={{.Version}} -X main.commit={{.Commit}} -X main.date={{.Date}}
- Homebrew tap: albertotijunelis/homebrew-tap
- Homebrew formula description and install test
- Checksums file: checksums.txt (sha256)
- Changelog: group by feat/fix/chore

## FILE 22 — .github/workflows/release.yml

GitHub Actions workflow:
- Trigger: on push tag v*.*.*
- Job: release
- Steps: checkout, setup-go@v5 (go 1.22), run GoReleaser action
- Env: GITHUB_TOKEN from secrets
- GoReleaser: distribution: goreleaser, version: latest, args: release --clean

## FILE 23 — .github/workflows/ci.yml

CI workflow:
- Trigger: push to main, pull_request to main
- Matrix: os: [ubuntu-latest, macos-latest], go: [1.22]
- Steps: checkout, setup-go, go mod download, go build ./..., go vet ./..., go test ./...
- Job name: "build"

## FILE 24 — Makefile

Targets:
```makefile
build:         go build -o bin/mcp .
install:       go install .
test:          go test ./...
lint:          golangci-lint run
clean:         rm -rf bin/ dist/
release-dry:   goreleaser release --snapshot --clean
run:           go run main.go
```

## FILE 25 — install.sh

A curl-installable shell script:
- Detects OS (darwin/linux) and ARCH (amd64/arm64)
- Downloads the correct binary from the latest GitHub release
- Verifies SHA256 checksum
- Installs to /usr/local/bin/mcp
- Prints success message with first command to run
- Handles errors gracefully with helpful messages
- Works with: `curl -sSfL https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/install.sh | sh`

## FILE 26 — SECURITY.md

Simple security policy:
- Supported versions table (latest release = supported)
- How to report a vulnerability: email albertotijunelis@gmail.com with subject "mcp-manager security"
- Response time commitment: 48 hours
- This file makes GitHub show a "Security policy" badge on the repo — adds credibility

## FILE 27 — demo.tape

A VHS tape file (by Charm) to auto-generate the README demo GIF:
```
# demo.tape — run: vhs demo.tape
Output demo.gif
Set FontSize 14
Set Width 1200
Set Height 600
Set Theme "Catppuccin Mocha"

Type "mcp doctor"
Enter
Sleep 2s

Type "mcp search database"  
Enter
Sleep 2s

Type "mcp install sqlite"
Enter
Sleep 3s

Type "mcp list"
Enter
Sleep 2s
```

---

# CRITICAL RULES

1. Output EVERY file using the `=== FILE: === / === END FILE ===` format
2. NEVER truncate any file. If a file is long, keep going until it is complete
3. NEVER use "..." or "[rest of implementation]" or "[similar to above]"
4. NEVER skip a file. All 27 files must be present in your output
5. If you reach your output limit mid-file, stop at a clean line and I will ask you to continue
6. Every .go file must have the copyright header: `// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.`
7. The project must compile with `go build ./...` after all files are in place
8. Use `albertotijunelis` everywhere — no placeholder text remains
9. The binary is called `mcp`, not `mcp-manager` — every reference to the CLI command uses `mcp`
10. NOTE: Alberto must also create a second GitHub repo called `homebrew-mcp` (or `homebrew-tap`) for the Homebrew formula. The .goreleaser.yml points to `albertotijunelis/homebrew-mcp`. Mention this in the README install section.

---

# START NOW

Generate all 27 files in order. Begin with FILE 01 (go.mod) immediately.
