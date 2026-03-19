// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package main

import (
	"os"

	"github.com/albertotijunelis/mcp-manager/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
