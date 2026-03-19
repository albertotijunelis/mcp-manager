// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/installer"
	"github.com/albertotijunelis/mcp-manager/internal/registry"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

var removeYes bool

func init() {
	removeCmd.Flags().BoolVar(&removeYes, "yes", false, "Skip confirmation prompt")
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:     "remove <name>",
	Aliases: []string{"uninstall", "rm"},
	Short:   "Remove an installed MCP server",
	Long:    "Remove an installed MCP server, its binary, and clean up the configuration.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := config.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		installed, err := cfg.GetInstalled(name)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Server '%s' is not installed", name))
			return err
		}

		// Show what will be removed
		fields := [][]string{
			{"Name", installed.Name},
			{"Version", installed.Version},
			{"Binary", installed.BinaryPath},
			{"Install Path", installed.InstallPath},
		}
		fmt.Println(ui.RenderInfoCard("Remove Server", fields))
		fmt.Println()

		// Confirmation
		if !removeYes {
			fmt.Print(ui.WarningStyle.Render("  Remove this server? [y/N] "))
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer != "y" && answer != "yes" {
				ui.PrintInfo("Removal cancelled")
				return nil
			}
		}

		fmt.Println()

		// Load registry to get install type
		reg, err := registry.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load registry: %w", err)
		}

		server, _ := reg.FindByName(name)

		err = ui.RunWithSpinner(fmt.Sprintf("Removing %s...", name), func() error {
			if server != nil {
				inst, instErr := installer.NewInstaller(server.InstallType)
				if instErr == nil {
					_ = inst.Remove(*server, mcpDir())
				}
			}

			// Remove server directory
			serverDir := mcpDir() + string(os.PathSeparator) + "servers" + string(os.PathSeparator) + name
			_ = os.RemoveAll(serverDir)

			// Remove binary if it exists
			if installed.BinaryPath != "" {
				_ = os.Remove(installed.BinaryPath)
			}

			return nil
		})
		if err != nil {
			ui.PrintError(fmt.Sprintf("Removal failed: %s", err.Error()))
			return err
		}

		// Remove from config
		if err := cfg.Remove(name); err != nil {
			return fmt.Errorf("failed to update config: %w", err)
		}

		fmt.Println()
		ui.PrintSuccess(fmt.Sprintf("Successfully removed %s", name))

		return nil
	},
}
