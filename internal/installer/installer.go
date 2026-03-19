// Copyright (c) 2026 Alberto Tijunelis Neto. MIT License.

package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/albertotijunelis/mcp-manager/internal/registry"
)

// Installer is the interface for installing and removing MCP servers.
type Installer interface {
	Install(server registry.Server, baseDir string) error
	Remove(server registry.Server, baseDir string) error
}

// NewInstaller returns the appropriate installer for the given install type.
func NewInstaller(installType string) (Installer, error) {
	switch installType {
	case "go":
		return &GoInstaller{}, nil
	case "npm":
		return &NpmInstaller{}, nil
	case "pip":
		return &PipInstaller{}, nil
	case "git+build":
		return &GitInstaller{}, nil
	case "binary":
		return &BinaryInstaller{}, nil
	default:
		return nil, fmt.Errorf("unsupported install type: %s", installType)
	}
}

// ResolveBinaryPath determines the expected binary path after installation.
func ResolveBinaryPath(server registry.Server, baseDir string) string {
	switch server.InstallType {
	case "go":
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			home, _ := os.UserHomeDir()
			gopath = filepath.Join(home, "go")
		}
		binaryName := server.Binary
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}
		return filepath.Join(gopath, "bin", binaryName)

	case "npm":
		// npm global binaries are on PATH
		binaryName := server.Binary
		if path, err := exec.LookPath(binaryName); err == nil {
			return path
		}
		return binaryName

	case "pip":
		binaryName := server.Binary
		if path, err := exec.LookPath(binaryName); err == nil {
			return path
		}
		return binaryName

	case "git+build", "binary":
		binDir := filepath.Join(baseDir, "bin")
		binaryName := server.Binary
		if runtime.GOOS == "windows" {
			binaryName += ".exe"
		}
		return filepath.Join(binDir, binaryName)

	default:
		return server.Binary
	}
}

// GoInstaller installs MCP servers via `go install`.
type GoInstaller struct{}

func (g *GoInstaller) Install(server registry.Server, baseDir string) error {
	repo := server.Repo
	version := server.Version
	if version == "" || version == "latest" {
		version = "latest"
	}

	installTarget := repo + "@" + version
	cmd := exec.Command("go", "install", installTarget)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}

func (g *GoInstaller) Remove(server registry.Server, baseDir string) error {
	binaryPath := ResolveBinaryPath(server, baseDir)
	if binaryPath != "" {
		return os.Remove(binaryPath)
	}
	return nil
}

// NpmInstaller installs MCP servers via `npm install -g`.
type NpmInstaller struct{}

func (n *NpmInstaller) Install(server registry.Server, baseDir string) error {
	pkg := server.Package
	if pkg == "" {
		pkg = server.Binary
	}

	version := server.Version
	installTarget := pkg
	if version != "" && version != "latest" {
		installTarget = pkg + "@" + version
	}

	cmd := exec.Command("npm", "install", "-g", installTarget)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}

func (n *NpmInstaller) Remove(server registry.Server, baseDir string) error {
	pkg := server.Package
	if pkg == "" {
		pkg = server.Binary
	}

	cmd := exec.Command("npm", "uninstall", "-g", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// PipInstaller installs MCP servers via `pip install`.
type PipInstaller struct{}

func (p *PipInstaller) Install(server registry.Server, baseDir string) error {
	pkg := server.Package
	if pkg == "" {
		pkg = server.Binary
	}

	version := server.Version
	installTarget := pkg
	if version != "" && version != "latest" {
		installTarget = pkg + "==" + version
	}

	pythonCmd := "pip"
	if _, err := exec.LookPath("pip3"); err == nil {
		pythonCmd = "pip3"
	}

	cmd := exec.Command(pythonCmd, "install", installTarget)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}

func (p *PipInstaller) Remove(server registry.Server, baseDir string) error {
	pkg := server.Package
	if pkg == "" {
		pkg = server.Binary
	}

	pythonCmd := "pip"
	if _, err := exec.LookPath("pip3"); err == nil {
		pythonCmd = "pip3"
	}

	cmd := exec.Command(pythonCmd, "uninstall", "-y", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// GitInstaller clones a repo and runs a build command.
type GitInstaller struct{}

func (g *GitInstaller) Install(server registry.Server, baseDir string) error {
	serverDir := filepath.Join(baseDir, "servers", server.Name)
	binDir := filepath.Join(baseDir, "bin")

	// Create directories
	if err := os.MkdirAll(serverDir, 0750); err != nil {
		return fmt.Errorf("failed to create server directory: %w", err)
	}
	if err := os.MkdirAll(binDir, 0750); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	// Clone repository
	repoURL := "https://" + server.Repo
	if !strings.HasPrefix(server.Repo, "github.com") {
		repoURL = server.Repo
	}

	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, serverDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	// Build (try common build commands)
	buildCmds := [][]string{
		{"make", "build"},
		{"go", "build", "-o", filepath.Join(binDir, server.Binary), "."},
	}

	for _, buildCmd := range buildCmds {
		c := exec.Command(buildCmd[0], buildCmd[1:]...)
		c.Dir = serverDir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if c.Run() == nil {
			return nil
		}
	}

	return fmt.Errorf("no build command succeeded for %s", server.Name)
}

func (g *GitInstaller) Remove(server registry.Server, baseDir string) error {
	serverDir := filepath.Join(baseDir, "servers", server.Name)
	binPath := filepath.Join(baseDir, "bin", server.Binary)

	_ = os.RemoveAll(serverDir)
	_ = os.Remove(binPath)

	return nil
}

// BinaryInstaller downloads pre-built binaries from GitHub releases.
type BinaryInstaller struct{}

func (b *BinaryInstaller) Install(server registry.Server, baseDir string) error {
	binDir := filepath.Join(baseDir, "bin")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	// Build GitHub release download URL
	osName := runtime.GOOS
	archName := runtime.GOARCH

	ext := ".tar.gz"
	if runtime.GOOS == "windows" {
		ext = ".zip"
	}

	binaryName := server.Binary
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	version := server.Version
	if version == "" || version == "latest" {
		version = "latest"
	}

	downloadURL := fmt.Sprintf("https://github.com/%s/releases/%s/download/%s_%s_%s%s",
		strings.TrimPrefix(server.Repo, "github.com/"),
		version,
		server.Binary,
		osName,
		archName,
		ext,
	)

	client := &http.Client{
		Timeout: 120 * time.Second,
	}

	resp, err := client.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download binary: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d from %s", resp.StatusCode, downloadURL)
	}

	destPath := filepath.Join(binDir, binaryName)
	out, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0750)
	if err != nil {
		return fmt.Errorf("failed to create binary file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("failed to write binary: %w", err)
	}

	return nil
}

func (b *BinaryInstaller) Remove(server registry.Server, baseDir string) error {
	binaryName := server.Binary
	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	binPath := filepath.Join(baseDir, "bin", binaryName)
	return os.Remove(binPath)
}
