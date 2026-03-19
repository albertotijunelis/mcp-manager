// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/registry"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system dependencies and configuration",
	Long:  "Verify that all required tools are installed and the mcp-manager configuration is healthy.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println()
		fmt.Println(ui.TitleStyle.Render("  MCP Doctor"))
		fmt.Println(ui.MutedStyle.Render("  Checking your system..."))
		fmt.Println()

		allOk := true

		// Check Go
		allOk = checkTool("go", "Go", "https://go.dev/dl/") && allOk

		// Check Node.js
		allOk = checkTool("node", "Node.js", "https://nodejs.org/") && allOk

		// Check Python
		allOk = checkToolAny([]string{"python3", "python"}, "Python", "https://python.org/") && allOk

		// Check git
		allOk = checkTool("git", "git", "https://git-scm.com/") && allOk

		fmt.Println()

		// Check ~/.mcp directory
		mcpPath := mcpDir()
		if info, err := os.Stat(mcpPath); err == nil && info.IsDir() {
			fmt.Println(ui.SuccessStyle.Render("  ✓ ") + fmt.Sprintf("~/.mcp directory exists (%s)", mcpPath))
		} else {
			fmt.Println(ui.ErrorStyle.Render("  ✗ ") + "~/.mcp directory not found")
			fmt.Println(ui.MutedStyle.Render("    Will be created on first install"))
			allOk = false
		}

		// Check registry file
		registryPath := filepath.Join(mcpPath, "registry.json")
		if _, err := os.Stat(registryPath); err == nil {
			reg, err := registry.Load(mcpPath)
			if err == nil {
				fmt.Println(ui.SuccessStyle.Render("  ✓ ") + fmt.Sprintf("Registry loaded (%d servers)", len(reg.Servers)))
			} else {
				fmt.Println(ui.ErrorStyle.Render("  ✗ ") + "Registry file is corrupted")
				fmt.Println(ui.MutedStyle.Render("    Run 'mcp update' to re-sync"))
				allOk = false
			}
		} else {
			fmt.Println(ui.WarningStyle.Render("  ⚠ ") + "Registry file not found")
			fmt.Println(ui.MutedStyle.Render("    Run 'mcp update' to sync"))
			allOk = false
		}

		// Check registry URL is reachable
		registryURL := "https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/registry/registry.json"
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(registryURL)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				fmt.Println(ui.SuccessStyle.Render("  ✓ ") + "Registry URL is reachable")
			} else {
				fmt.Println(ui.WarningStyle.Render("  ⚠ ") + fmt.Sprintf("Registry URL returned status %d", resp.StatusCode))
				allOk = false
			}
		} else {
			fmt.Println(ui.WarningStyle.Render("  ⚠ ") + "Registry URL is not reachable")
			fmt.Println(ui.MutedStyle.Render("    Offline mode will use cached registry"))
			allOk = false
		}

		// Check installed server binaries
		cfg, err := config.Load(mcpPath)
		if err == nil && len(cfg.Installed) > 0 {
			fmt.Println()
			fmt.Println(ui.SubtitleStyle.Render("  Installed Servers"))
			for name, server := range cfg.Installed {
				if _, err := os.Stat(server.BinaryPath); err == nil {
					fmt.Println(ui.SuccessStyle.Render("  ✓ ") + fmt.Sprintf("%s — binary OK (%s)", name, server.BinaryPath))
				} else {
					fmt.Println(ui.ErrorStyle.Render("  ✗ ") + fmt.Sprintf("%s — binary missing (%s)", name, server.BinaryPath))
					fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("    Reinstall with: mcp remove %s && mcp install %s", name, name)))
					allOk = false
				}
			}
		}

		fmt.Println()
		if allOk {
			ui.PrintSuccess("Everything looks good! You're ready to go.")
		} else {
			ui.PrintWarning("Some checks failed. See above for details.")
		}
		fmt.Println()

		return nil
	},
}

func checkTool(binary, displayName, installURL string) bool {
	path, err := exec.LookPath(binary)
	if err != nil {
		fmt.Println(ui.ErrorStyle.Render("  ✗ ") + fmt.Sprintf("%s not found", displayName))
		fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("    Install: %s", installURL)))
		return false
	}

	version := getToolVersion(binary)
	fmt.Println(ui.SuccessStyle.Render("  ✓ ") + fmt.Sprintf("%s %s (%s)", displayName, version, path))
	return true
}

func checkToolAny(binaries []string, displayName, installURL string) bool {
	for _, binary := range binaries {
		path, err := exec.LookPath(binary)
		if err == nil {
			version := getToolVersion(binary)
			fmt.Println(ui.SuccessStyle.Render("  ✓ ") + fmt.Sprintf("%s %s (%s)", displayName, version, path))
			return true
		}
	}

	fmt.Println(ui.ErrorStyle.Render("  ✗ ") + fmt.Sprintf("%s not found", displayName))
	fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("    Install: %s", installURL)))
	return false
}

func getToolVersion(binary string) string {
	out, err := exec.Command(binary, "--version").CombinedOutput()
	if err != nil {
		return "unknown"
	}
	version := string(out)
	// Take only the first line
	for i, c := range version {
		if c == '\n' || c == '\r' {
			version = version[:i]
			break
		}
	}
	return version
}
