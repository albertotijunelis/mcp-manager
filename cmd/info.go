// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/registry"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info <name>",
	Short: "Show detailed information about an MCP server",
	Long:  "Display a detailed info card for an MCP server from the registry, including install status.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		reg, err := registry.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load registry: %w", err)
		}

		server, err := reg.FindByName(name)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Server '%s' not found in registry", name))
			return err
		}

		fields := [][]string{
			{"Name", server.DisplayName},
			{"Description", server.Description},
			{"Repository", server.Repo},
			{"Homepage", server.Homepage},
			{"Install Type", server.InstallType},
			{"Version", server.Version},
		}

		if server.Stars > 0 {
			fields = append(fields, []string{"Stars", fmt.Sprintf("⭐ %s", formatStars(server.Stars))})
		}

		if len(server.Topics) > 0 {
			var badges []string
			for _, topic := range server.Topics {
				badges = append(badges, ui.RenderBadge(topic))
			}
			fields = append(fields, []string{"Topics", strings.Join(badges, " ")})
		}

		if len(server.RequiresEnv) > 0 {
			fields = append(fields, []string{"Required Env", strings.Join(server.RequiresEnv, ", ")})
		}

		// Check if installed
		cfg, err := config.Load(mcpDir())
		if err == nil {
			if installed, instErr := cfg.GetInstalled(name); instErr == nil {
				fields = append(fields, []string{"", ""}) // separator
				fields = append(fields, []string{"Status", ui.SuccessStyle.Render("● Installed")})
				fields = append(fields, []string{"Installed Version", installed.Version})
				fields = append(fields, []string{"Install Path", installed.InstallPath})
				fields = append(fields, []string{"Binary Path", installed.BinaryPath})
				fields = append(fields, []string{"Installed At", installed.InstalledAt.Format("2006-01-02 15:04:05")})
			} else {
				fields = append(fields, []string{"Status", ui.MutedStyle.Render("○ Not installed")})
			}
		}

		fmt.Println()
		fmt.Println(ui.RenderInfoCard(server.DisplayName, fields))
		fmt.Println()

		return nil
	},
}
