// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"fmt"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/installer"
	"github.com/albertotijunelis/mcp-manager/internal/registry"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

var updateAll bool

func init() {
	updateCmd.Flags().BoolVar(&updateAll, "all", false, "Update all installed servers")
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update [name]",
	Short: "Update MCP servers or sync the registry",
	Long:  "Without arguments: sync the registry from upstream. With a server name: update that server to the latest version. With --all: update every installed server.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reg, err := registry.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load registry: %w", err)
		}

		// No args: sync registry
		if len(args) == 0 && !updateAll {
			fmt.Println()
			ui.PrintInfo("Syncing registry from upstream...")
			fmt.Println()

			err = ui.RunWithSpinner("Syncing registry...", func() error {
				return reg.Sync(mcpDir())
			})
			if err != nil {
				ui.PrintError(fmt.Sprintf("Failed to sync registry: %s", err.Error()))
				return err
			}

			ui.PrintSuccess("Registry synced successfully!")
			fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("  %d servers available", len(reg.Servers))))
			fmt.Println()
			return nil
		}

		cfg, err := config.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// --all: update every installed server
		if updateAll {
			if len(cfg.Installed) == 0 {
				ui.PrintWarning("No servers installed to update")
				return nil
			}

			fmt.Println()
			ui.PrintInfo(fmt.Sprintf("Updating %d installed server(s)...", len(cfg.Installed)))
			fmt.Println()

			var errors []string
			for name := range cfg.Installed {
				if err := updateServer(name, reg, cfg); err != nil {
					errors = append(errors, fmt.Sprintf("%s: %s", name, err.Error()))
				}
			}

			if len(errors) > 0 {
				fmt.Println()
				ui.PrintWarning("Some updates failed:")
				for _, e := range errors {
					fmt.Println(ui.ErrorStyle.Render("  ✗ " + e))
				}
			}

			fmt.Println()
			ui.PrintSuccess("Update complete!")
			return nil
		}

		// Update specific server
		name := args[0]
		if !cfg.IsInstalled(name) {
			ui.PrintError(fmt.Sprintf("Server '%s' is not installed", name))
			return fmt.Errorf("server not installed: %s", name)
		}

		fmt.Println()
		if err := updateServer(name, reg, cfg); err != nil {
			return err
		}

		fmt.Println()
		ui.PrintSuccess(fmt.Sprintf("Successfully updated %s!", name))
		return nil
	},
}

func updateServer(name string, reg *registry.Registry, cfg *config.Config) error {
	server, err := reg.FindByName(name)
	if err != nil {
		ui.PrintWarning(fmt.Sprintf("Server '%s' not found in registry, skipping", name))
		return err
	}

	inst, err := installer.NewInstaller(server.InstallType)
	if err != nil {
		return err
	}

	err = ui.RunWithSpinner(fmt.Sprintf("Updating %s...", server.DisplayName), func() error {
		return inst.Install(*server, mcpDir())
	})
	if err != nil {
		ui.PrintError(fmt.Sprintf("Failed to update %s: %s", name, err.Error()))
		return err
	}

	binaryPath := installer.ResolveBinaryPath(*server, mcpDir())
	return cfg.Install(*server, mcpDir(), binaryPath)
}
