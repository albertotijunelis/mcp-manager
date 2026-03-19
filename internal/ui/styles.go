// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Color palette
const (
	ColorCyan  = "#00D4FF"
	ColorGreen = "#00FF9F"
	ColorAmber = "#FFB800"
	ColorRed   = "#FF4757"
	ColorMuted = "#64748B"
)

var colorsEnabled = true

// DisableColors turns off colored output.
func DisableColors() {
	colorsEnabled = false
}

// TitleStyle renders section titles.
var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(ColorCyan)).
	MarginBottom(0)

// SubtitleStyle renders subtitles.
var SubtitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(ColorMuted))

// SuccessStyle renders success messages.
var SuccessStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorGreen))

// ErrorStyle renders error messages.
var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorRed))

// WarningStyle renders warning messages.
var WarningStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorAmber))

// InfoStyle renders informational messages.
var InfoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorCyan))

// MutedStyle renders muted/secondary text.
var MutedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorMuted))

// TableHeaderStyle renders table headers.
var TableHeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color(ColorCyan)).
	PaddingRight(2)

// TableRowStyle renders normal table rows.
var TableRowStyle = lipgloss.NewStyle().
	PaddingRight(2)

// TableAltRowStyle renders alternate table rows.
var TableAltRowStyle = lipgloss.NewStyle().
	PaddingRight(2).
	Foreground(lipgloss.Color("#cccccc"))

// BadgeStyle renders topic badges.
var BadgeStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Background(lipgloss.Color(ColorCyan)).
	Padding(0, 1).
	Bold(true)

// InfoCardStyle renders information cards.
var InfoCardStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color(ColorCyan)).
	Padding(1, 2).
	MarginLeft(2)

// SeparatorStyle renders horizontal separators.
var SeparatorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorMuted))

// LogoStyle renders the ASCII logo.
var LogoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorCyan)).
	Bold(true)

// RenderTable renders a formatted table with headers and rows.
func RenderTable(headers []string, rows [][]string) string {
	if len(rows) == 0 {
		return ""
	}

	// Calculate column widths
	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = lipgloss.Width(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) {
				w := lipgloss.Width(cell)
				if w > colWidths[i] {
					colWidths[i] = w
				}
			}
		}
	}

	var sb strings.Builder

	// Header
	sb.WriteString("  ")
	for i, h := range headers {
		padded := h + strings.Repeat(" ", max(0, colWidths[i]-lipgloss.Width(h)))
		sb.WriteString(TableHeaderStyle.Render(padded))
	}
	sb.WriteString("\n")

	// Separator
	sb.WriteString("  ")
	for i, w := range colWidths {
		sb.WriteString(SeparatorStyle.Render(strings.Repeat("─", w)))
		if i < len(colWidths)-1 {
			sb.WriteString("  ")
		}
	}
	sb.WriteString("\n")

	// Rows
	for _, row := range rows {
		sb.WriteString("  ")
		for i, cell := range row {
			if i < len(colWidths) {
				cellWidth := lipgloss.Width(cell)
				padded := cell + strings.Repeat(" ", max(0, colWidths[i]-cellWidth))
				sb.WriteString(TableRowStyle.Render(padded))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// RenderInfoCard renders a key-value info card in a lipgloss box.
func RenderInfoCard(title string, fields [][]string) string {
	var sb strings.Builder

	sb.WriteString(TitleStyle.Render(title))
	sb.WriteString("\n\n")

	labelWidth := 0
	for _, field := range fields {
		if len(field) >= 2 && len(field[0]) > labelWidth {
			labelWidth = len(field[0])
		}
	}

	for _, field := range fields {
		if len(field) < 2 {
			continue
		}
		label := field[0]
		value := field[1]

		if label == "" && value == "" {
			sb.WriteString(SeparatorStyle.Render(strings.Repeat("─", labelWidth+20)))
			sb.WriteString("\n")
			continue
		}

		paddedLabel := label + strings.Repeat(" ", max(0, labelWidth-len(label)))
		sb.WriteString(MutedStyle.Render(paddedLabel+" │ ") + value + "\n")
	}

	return InfoCardStyle.Render(sb.String())
}

// PrintSuccess prints a success message.
func PrintSuccess(msg string) {
	fmt.Println(SuccessStyle.Render("  ✓ " + msg))
}

// PrintError prints an error message.
func PrintError(msg string) {
	fmt.Println(ErrorStyle.Render("  ✗ " + msg))
}

// PrintWarning prints a warning message.
func PrintWarning(msg string) {
	fmt.Println(WarningStyle.Render("  ⚠ " + msg))
}

// PrintInfo prints an informational message.
func PrintInfo(msg string) {
	fmt.Println(InfoStyle.Render("  → " + msg))
}

// RenderBadge renders a topic badge.
func RenderBadge(text string) string {
	return BadgeStyle.Render(text)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
