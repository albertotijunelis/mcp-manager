// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/albertotijunelis/mcp-manager/internal/registry"
)

// InstalledServer represents an installed MCP server.
type InstalledServer struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	InstallPath string            `json:"install_path"`
	BinaryPath  string            `json:"binary_path"`
	InstalledAt time.Time         `json:"installed_at"`
	Env         map[string]string `json:"env"`
}

// Config holds the user's mcp-manager configuration.
type Config struct {
	Version   string                     `json:"version"`
	Installed map[string]InstalledServer  `json:"installed"`
	path      string
}

const configFileName = "config.json"

// Load reads the config from disk at the given base directory.
func Load(baseDir string) (*Config, error) {
	configPath := filepath.Join(baseDir, configFileName)

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				Version:   "1",
				Installed: make(map[string]InstalledServer),
				path:      configPath,
			}, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Installed == nil {
		cfg.Installed = make(map[string]InstalledServer)
	}
	cfg.path = configPath

	return &cfg, nil
}

// Save writes the config to disk.
func (c *Config) Save() error {
	dir := filepath.Dir(c.path)
	if err := os.MkdirAll(dir, 0750); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	return os.WriteFile(c.path, data, 0600)
}

// Install records a server as installed.
func (c *Config) Install(server registry.Server, baseDir, binaryPath string) error {
	envMap := make(map[string]string)
	for _, envVar := range server.RequiresEnv {
		// Preserve existing value if already set
		existing := ""
		if prev, ok := c.Installed[server.Name]; ok {
			if v, exists := prev.Env[envVar]; exists {
				existing = v
			}
		}
		if existing == "" {
			existing = os.Getenv(envVar)
		}
		envMap[envVar] = existing
	}

	c.Installed[server.Name] = InstalledServer{
		Name:        server.Name,
		Version:     server.Version,
		InstallPath: filepath.Join(baseDir, "servers", server.Name),
		BinaryPath:  binaryPath,
		InstalledAt: time.Now().UTC(),
		Env:         envMap,
	}

	return c.Save()
}

// Remove deletes a server from the installed list.
func (c *Config) Remove(name string) error {
	delete(c.Installed, name)
	return c.Save()
}

// IsInstalled checks if a server is installed.
func (c *Config) IsInstalled(name string) bool {
	_, ok := c.Installed[name]
	return ok
}

// GetInstalled returns an installed server by name.
func (c *Config) GetInstalled(name string) (*InstalledServer, error) {
	server, ok := c.Installed[name]
	if !ok {
		return nil, fmt.Errorf("server not installed: %s", name)
	}
	return &server, nil
}
