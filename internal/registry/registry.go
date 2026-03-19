// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Server represents an MCP server in the registry.
type Server struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
	Repo        string   `json:"repo"`
	InstallType string   `json:"install_type"`
	Binary      string   `json:"binary"`
	Package     string   `json:"package,omitempty"`
	Version     string   `json:"version"`
	Homepage    string   `json:"homepage"`
	Topics      []string `json:"topics"`
	RequiresEnv []string `json:"requires_env,omitempty"`
	Stars       int      `json:"stars"`
}

// Registry contains the full list of available MCP servers.
type Registry struct {
	Version     string   `json:"version"`
	UpdatedAt   string   `json:"updated_at"`
	RegistryURL string   `json:"registry_url"`
	Servers     []Server `json:"servers"`
}

const registryFileName = "registry.json"

// Load reads the registry from disk at the given base directory.
// If the file does not exist, it loads the embedded default registry.
func Load(baseDir string) (*Registry, error) {
	registryPath := filepath.Join(baseDir, registryFileName)

	data, err := os.ReadFile(registryPath)
	if err != nil {
		// Try loading from the bundled registry
		return loadBundled()
	}

	var reg Registry
	if err := json.Unmarshal(data, &reg); err != nil {
		return nil, fmt.Errorf("failed to parse registry: %w", err)
	}

	return &reg, nil
}

// loadBundled returns a minimal default registry when no file is found.
func loadBundled() (*Registry, error) {
	// Look for the bundled registry in the executable's directory
	execPath, err := os.Executable()
	if err == nil {
		bundledPath := filepath.Join(filepath.Dir(execPath), "registry", registryFileName)
		if data, err := os.ReadFile(bundledPath); err == nil {
			var reg Registry
			if err := json.Unmarshal(data, &reg); err == nil {
				return &reg, nil
			}
		}
	}

	// Look in the current working directory
	if data, err := os.ReadFile(filepath.Join("registry", registryFileName)); err == nil {
		var reg Registry
		if err := json.Unmarshal(data, &reg); err == nil {
			return &reg, nil
		}
	}

	return &Registry{
		Version:     "1.0.0",
		UpdatedAt:   time.Now().UTC().Format(time.RFC3339),
		RegistryURL: "https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/registry/registry.json",
		Servers:     []Server{},
	}, nil
}

// Save writes the registry to disk at the given base directory.
func (r *Registry) Save(baseDir string) error {
	if err := os.MkdirAll(baseDir, 0750); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	registryPath := filepath.Join(baseDir, registryFileName)

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal registry: %w", err)
	}

	return os.WriteFile(registryPath, data, 0600)
}

// FindByName returns the server with the given name, or an error if not found.
func (r *Registry) FindByName(name string) (*Server, error) {
	for i := range r.Servers {
		if strings.EqualFold(r.Servers[i].Name, name) {
			return &r.Servers[i], nil
		}
	}
	return nil, fmt.Errorf("server not found: %s", name)
}

// Search performs a fuzzy search across name, description, and topics.
func (r *Registry) Search(query string) []Server {
	query = strings.ToLower(query)
	var results []Server

	for _, server := range r.Servers {
		score := 0

		// Exact name match
		if strings.EqualFold(server.Name, query) {
			score += 100
		} else if strings.Contains(strings.ToLower(server.Name), query) {
			score += 50
		}

		// Description match
		if strings.Contains(strings.ToLower(server.Description), query) {
			score += 30
		}

		// Display name match
		if strings.Contains(strings.ToLower(server.DisplayName), query) {
			score += 40
		}

		// Topic match
		for _, topic := range server.Topics {
			if strings.Contains(strings.ToLower(topic), query) {
				score += 20
				break
			}
		}

		if score > 0 {
			results = append(results, server)
		}
	}

	// Sort by relevance (simple bubble sort since lists are small)
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if scoreServer(results[j], query) > scoreServer(results[i], query) {
				results[i], results[j] = results[j], results[i]
			}
		}
	}

	return results
}

func scoreServer(s Server, query string) int {
	score := 0
	if strings.EqualFold(s.Name, query) {
		score += 100
	} else if strings.Contains(strings.ToLower(s.Name), query) {
		score += 50
	}
	if strings.Contains(strings.ToLower(s.Description), query) {
		score += 30
	}
	if strings.Contains(strings.ToLower(s.DisplayName), query) {
		score += 40
	}
	for _, topic := range s.Topics {
		if strings.Contains(strings.ToLower(topic), query) {
			score += 20
			break
		}
	}
	return score
}

// Sync fetches the latest registry from the upstream URL and saves it.
func (r *Registry) Sync(baseDir string) error {
	url := r.RegistryURL
	if url == "" {
		url = "https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/registry/registry.json"
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch registry: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registry returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	var remote Registry
	if err := json.Unmarshal(body, &remote); err != nil {
		return fmt.Errorf("failed to parse remote registry: %w", err)
	}

	r.Version = remote.Version
	r.UpdatedAt = remote.UpdatedAt
	r.Servers = remote.Servers

	return r.Save(baseDir)
}
