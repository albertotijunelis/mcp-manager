// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List installed MCP servers",
	Long:    "Display a table of all installed MCP servers with their status, version, and install type.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if len(cfg.Installed) == 0 {
			fmt.Println()
			ui.PrintInfo("No MCP servers installed yet")
			fmt.Println()
			ui.PrintInfo("Get started:")
			fmt.Println(ui.InfoStyle.Render("  mcp search <query>    ") + ui.MutedStyle.Render("Search available servers"))
			fmt.Println(ui.InfoStyle.Render("  mcp install <name>    ") + ui.MutedStyle.Render("Install a server"))
			fmt.Println()
			return nil
		}

		if jsonOut {
			data, err := json.MarshalIndent(cfg.Installed, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		}

		headers := []string{"Name", "Version", "Status", "Install Path", "Installed At"}
		var rows [][]string

		for _, server := range cfg.Installed {
			status := ui.MutedStyle.Render("● stopped")

			installedAt := server.InstalledAt.Format(time.DateOnly)

			rows = append(rows, []string{
				ui.InfoStyle.Render(server.Name),
				server.Version,
				status,
				server.InstallPath,
				installedAt,
			})
		}

		fmt.Println()
		fmt.Println(ui.TitleStyle.Render("  Installed MCP Servers"))
		fmt.Println()
		fmt.Println(ui.RenderTable(headers, rows))
		fmt.Println()
		fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("  %d server(s) installed", len(cfg.Installed))))
		fmt.Println()

		return nil
	},
}
