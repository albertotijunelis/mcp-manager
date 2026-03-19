// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/albertotijunelis/mcp-manager/internal/config"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run <name> [-- args...]",
	Short: "Run an installed MCP server",
	Long:  "Launch an installed MCP server process. Any arguments after -- are passed directly to the server binary.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		var serverArgs []string
		if cmd.ArgsLenAtDash() > 0 {
			serverArgs = args[cmd.ArgsLenAtDash():]
		} else if len(args) > 1 {
			serverArgs = args[1:]
		}

		cfg, err := config.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		installed, err := cfg.GetInstalled(name)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Server '%s' is not installed", name))
			ui.PrintInfo(fmt.Sprintf("Install it with: mcp install %s", name))
			return err
		}

		binaryPath := installed.BinaryPath
		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			ui.PrintError(fmt.Sprintf("Binary not found at %s", binaryPath))
			ui.PrintInfo(fmt.Sprintf("Try reinstalling: mcp remove %s && mcp install %s", name, name))
			return fmt.Errorf("binary not found: %s", binaryPath)
		}

		// Set environment variables from config
		for k, v := range installed.Env {
			if v != "" {
				os.Setenv(k, v)
			}
		}

		// Show startup info
		fmt.Println()
		ui.PrintInfo(fmt.Sprintf("Starting %s (v%s)", name, installed.Version))
		if len(serverArgs) > 0 {
			ui.PrintInfo(fmt.Sprintf("Args: %s", strings.Join(serverArgs, " ")))
		}

		// Check for missing env vars
		missingEnv := false
		for k, v := range installed.Env {
			if v == "" && os.Getenv(k) == "" {
				if !missingEnv {
					fmt.Println()
					ui.PrintWarning("Missing environment variables:")
					missingEnv = true
				}
				fmt.Printf("  %s\n", ui.WarningStyle.Render(k))
			}
		}
		if missingEnv {
			fmt.Println()
		}

		fmt.Println(ui.MutedStyle.Render("  Press Ctrl+C to stop"))
		fmt.Println()

		// Exec the binary
		binary, err := exec.LookPath(binaryPath)
		if err != nil {
			binary = binaryPath
		}

		execErr := syscall.Exec(binary, append([]string{binary}, serverArgs...), os.Environ())
		if execErr != nil {
			// Fallback: use exec.Command if syscall.Exec fails (e.g., on Windows)
			command := exec.Command(binaryPath, serverArgs...)
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			return command.Run()
		}

		return nil
	},
}
