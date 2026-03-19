// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/albertotijunelis/mcp-manager/internal/registry"
	"github.com/albertotijunelis/mcp-manager/internal/ui"
	"github.com/spf13/cobra"
)

var searchLimit int

func init() {
	searchCmd.Flags().IntVar(&searchLimit, "limit", 10, "Maximum number of results to show")
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search the MCP server registry",
	Long:  "Fuzzy search the registry by name, description, or topics. Results are sorted by relevance.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		reg, err := registry.Load(mcpDir())
		if err != nil {
			return fmt.Errorf("failed to load registry: %w", err)
		}

		results := reg.Search(query)

		if len(results) == 0 {
			fmt.Println()
			ui.PrintWarning(fmt.Sprintf("No servers found matching '%s'", query))
			fmt.Println()
			ui.PrintInfo("Try a broader search or browse the full registry:")
			fmt.Println(ui.MutedStyle.Render("  https://github.com/albertotijunelis/mcp-manager/blob/main/registry/registry.json"))
			fmt.Println()
			return nil
		}

		if jsonOut {
			data, err := json.MarshalIndent(results, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		}

		// Limit results
		displayed := results
		if len(displayed) > searchLimit {
			displayed = displayed[:searchLimit]
		}

		headers := []string{"Name", "Description", "Stars", "Type"}
		var rows [][]string

		for _, server := range displayed {
			desc := server.Description
			if len(desc) > 60 {
				desc = desc[:57] + "..."
			}

			stars := ""
			if server.Stars > 0 {
				stars = formatStars(server.Stars)
			}

			rows = append(rows, []string{
				ui.InfoStyle.Render(server.Name),
				desc,
				stars,
				server.InstallType,
			})
		}

		fmt.Println()
		fmt.Println(ui.TitleStyle.Render(fmt.Sprintf("  Search results for '%s'", query)))
		fmt.Println()
		fmt.Println(ui.RenderTable(headers, rows))
		fmt.Println()
		fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("  %d result(s) found", len(results))))
		if len(results) > searchLimit {
			fmt.Println(ui.MutedStyle.Render(fmt.Sprintf("  Showing top %d — use --limit to show more", searchLimit)))
		}
		fmt.Println()
		ui.PrintInfo(fmt.Sprintf("Install a server: mcp install <name>"))

		return nil
	},
}

func formatStars(n int) string {
	if n >= 1000 {
		return fmt.Sprintf("%.1fk", float64(n)/1000.0)
	}
	return strconv.Itoa(n)
}
