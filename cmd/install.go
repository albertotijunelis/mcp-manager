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

var installVersion string
var installYes bool

func init() {
	installCmd.Flags().StringVar(&installVersion, "version", "", "Specific version to install (default: latest)")
	installCmd.Flags().BoolVar(&installYes, "yes", false, "Skip confirmation prompt")
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install <name>",
	Short: "Install an MCP server from the registry",
	Long:  "Install a Model Context Protocol server by name. Looks up the server in the registry and installs it using the appropriate method (go, npm, pip, git+build, or binary).",
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
			ui.PrintInfo("Run 'mcp search " + name + "' to search available servers")
			return err
		}

		cfg, err := config.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		if cfg.IsInstalled(name) {
			ui.PrintWarning(fmt.Sprintf("Server '%s' is already installed", name))
			ui.PrintInfo("Use 'mcp update " + name + "' to update it")
			return nil
		}

		// Show server info card
		fields := [][]string{
			{"Name", server.DisplayName},
			{"Description", server.Description},
			{"Install Type", server.InstallType},
			{"Repository", server.Repo},
			{"Version", server.Version},
		}
		if len(server.RequiresEnv) > 0 {
			fields = append(fields, []string{"Required Env", strings.Join(server.RequiresEnv, ", ")})
		}
		fmt.Println(ui.RenderInfoCard("Install Server", fields))
		fmt.Println()

		// Confirmation
		if !installYes {
			fmt.Print(ui.InfoStyle.Render("  Install this server? [Y/n] "))
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer != "" && answer != "y" && answer != "yes" {
				ui.PrintWarning("Installation cancelled")
				return nil
			}
		}

		fmt.Println()

		// Override version if specified
		if installVersion != "" {
			server.Version = installVersion
		}

		// Install
		inst, err := installer.NewInstaller(server.InstallType)
		if err != nil {
			return err
		}

		err = ui.RunWithSpinner(fmt.Sprintf("Installing %s...", server.DisplayName), func() error {
			return inst.Install(*server, mcpDir())
		})
		if err != nil {
			ui.PrintError(fmt.Sprintf("Installation failed: %s", err.Error()))
			return err
		}

		// Determine binary path
		binaryPath := installer.ResolveBinaryPath(*server, mcpDir())

		// Save to config
		if err := cfg.Install(*server, mcpDir(), binaryPath); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println()
		ui.PrintSuccess(fmt.Sprintf("Successfully installed %s!", server.DisplayName))

		// Show env var setup instructions
		if len(server.RequiresEnv) > 0 {
			fmt.Println()
			ui.PrintWarning("This server requires environment variables:")
			for _, env := range server.RequiresEnv {
				fmt.Printf("  %s\n", ui.InfoStyle.Render(fmt.Sprintf("export %s=<your-value>", env)))
			}
			fmt.Println()
			ui.PrintInfo(fmt.Sprintf("Run 'mcp run %s' to start the server", name))
		}

		return nil
	},
}
