// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"fmt"
	"os"

	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

var (
	noColor bool
	debug   bool
	jsonOut bool

	appVersion = "dev"
	appCommit  = "none"
	appDate    = "unknown"
)

const asciiLogo = `
╔╦╗╔═╗╔═╗   ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗
║║║║  ╠═╝───║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝
╩ ╩╚═╝╩     ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═
The package manager for MCP servers
`

func SetVersionInfo(version, commit, date string) {
	appVersion = version
	appCommit = commit
	appDate = date
}

var rootCmd = &cobra.Command{
	Use:   "mcp",
	Short: "The package manager for MCP servers",
	Long:  "mcp-manager — Install, manage and run Model Context Protocol servers like Homebrew for your AI toolchain.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if noColor {
			ui.DisableColors()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ui.LogoStyle.Render(asciiLogo))
		fmt.Println(ui.InfoStyle.Render("  Use 'mcp --help' to see available commands"))
		fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("  Version %s (%s) built %s", appVersion, appCommit, appDate)))
		fmt.Println()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable colored output")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVar(&jsonOut, "json", false, "Output in JSON format")

	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mcp-manager %s\n", appVersion)
		fmt.Printf("  commit: %s\n", appCommit)
		fmt.Printf("  built:  %s\n", appDate)
	},
}

func mcpDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		ui.PrintError("Cannot determine home directory: " + err.Error())
		os.Exit(1)
	}
	return home + string(os.PathSeparator) + ".mcp"
}
