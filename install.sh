#!/bin/sh
# install.sh — Install mcp-manager
# Usage: curl -sSfL https://raw.githubusercontent.com/albertotijunelis/mcp-manager/main/install.sh | sh

set -e

REPO="albertotijunelis/mcp-manager"
BINARY_NAME="mcp"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

info() {
    printf "${CYAN}→ %s${NC}\n" "$1"
}

success() {
    printf "${GREEN}✓ %s${NC}\n" "$1"
}

error() {
    printf "${RED}✗ %s${NC}\n" "$1"
    exit 1
}

# Detect OS
detect_os() {
    OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
    case "$OS" in
        linux)  OS="linux" ;;
        darwin) OS="darwin" ;;
        *)      error "Unsupported operating system: $OS. Use 'go install' instead." ;;
    esac
    echo "$OS"
}

# Detect architecture
detect_arch() {
    ARCH="$(uname -m)"
    case "$ARCH" in
        x86_64|amd64)  ARCH="amd64" ;;
        arm64|aarch64) ARCH="arm64" ;;
        *)             error "Unsupported architecture: $ARCH" ;;
    esac
    echo "$ARCH"
}

# Get latest release tag from GitHub API
get_latest_version() {
    VERSION=$(curl -sSf "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
    if [ -z "$VERSION" ]; then
        error "Failed to fetch latest version. Check your internet connection."
    fi
    echo "$VERSION"
}

main() {
    printf "\n"
    printf "${CYAN}╔╦╗╔═╗╔═╗   ╔╦╗╔═╗╔╗╔╔═╗╔═╗╔═╗╦═╗${NC}\n"
    printf "${CYAN}║║║║  ╠═╝───║║║╠═╣║║║╠═╣║ ╦║╣ ╠╦╝${NC}\n"
    printf "${CYAN}╩ ╩╚═╝╩     ╩ ╩╩ ╩╝╚╝╩ ╩╚═╝╚═╝╩╚═${NC}\n"
    printf "  Installer for mcp-manager\n\n"

    OS=$(detect_os)
    ARCH=$(detect_arch)
    info "Detected platform: ${OS}/${ARCH}"

    VERSION=$(get_latest_version)
    info "Latest version: ${VERSION}"

    # Strip the 'v' prefix for the archive name
    VERSION_NUM="${VERSION#v}"

    ARCHIVE_NAME="mcp-manager_${VERSION_NUM}_${OS}_${ARCH}.tar.gz"
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARCHIVE_NAME}"
    CHECKSUMS_URL="https://github.com/${REPO}/releases/download/${VERSION}/checksums.txt"

    TMPDIR=$(mktemp -d)
    trap 'rm -rf "$TMPDIR"' EXIT

    info "Downloading ${ARCHIVE_NAME}..."
    curl -sSfL -o "${TMPDIR}/${ARCHIVE_NAME}" "$DOWNLOAD_URL" || error "Failed to download binary. Check if the release exists for your platform."

    info "Downloading checksums..."
    curl -sSfL -o "${TMPDIR}/checksums.txt" "$CHECKSUMS_URL" || error "Failed to download checksums."

    info "Verifying SHA256 checksum..."
    cd "$TMPDIR"
    EXPECTED_SUM=$(grep "$ARCHIVE_NAME" checksums.txt | awk '{print $1}')
    if [ -z "$EXPECTED_SUM" ]; then
        error "Checksum not found for ${ARCHIVE_NAME}"
    fi

    if command -v sha256sum > /dev/null 2>&1; then
        ACTUAL_SUM=$(sha256sum "$ARCHIVE_NAME" | awk '{print $1}')
    elif command -v shasum > /dev/null 2>&1; then
        ACTUAL_SUM=$(shasum -a 256 "$ARCHIVE_NAME" | awk '{print $1}')
    else
        error "Neither sha256sum nor shasum found. Cannot verify checksum."
    fi

    if [ "$EXPECTED_SUM" != "$ACTUAL_SUM" ]; then
        error "Checksum verification failed!\n  Expected: ${EXPECTED_SUM}\n  Got:      ${ACTUAL_SUM}"
    fi
    success "Checksum verified"

    info "Extracting binary..."
    tar xzf "$ARCHIVE_NAME"

    info "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_NAME" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        sudo mv "$BINARY_NAME" "${INSTALL_DIR}/${BINARY_NAME}"
    fi
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"

    success "mcp-manager ${VERSION} installed successfully!"
    printf "\n"
    info "Get started:"
    printf "  ${CYAN}mcp doctor${NC}          Check your system\n"
    printf "  ${CYAN}mcp search github${NC}   Search for servers\n"
    printf "  ${CYAN}mcp install github${NC}  Install your first server\n"
    printf "\n"
}

main
