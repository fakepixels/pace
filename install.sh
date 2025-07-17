#!/bin/bash

# Pace CLI Installer Script
# Usage: curl -fsSL https://raw.githubusercontent.com/fakepixels/pace/main/install.sh | bash

set -e

# Configuration
REPO="fakepixels/pace"
BINARY_NAME="pace"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
log_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

log_success() {
    echo -e "${GREEN}✅${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

log_error() {
    echo -e "${RED}❌${NC} $1"
}

# Detect OS and architecture
detect_platform() {
    local os
    local arch
    
    # Detect OS
    case "$(uname -s)" in
        Darwin*)  os="Darwin" ;;
        Linux*)   os="Linux" ;;
        MINGW*|CYGWIN*|MSYS*) os="Windows" ;;
        *)        log_error "Unsupported operating system: $(uname -s)" && exit 1 ;;
    esac
    
    # Detect architecture
    case "$(uname -m)" in
        x86_64|amd64) arch="x86_64" ;;
        arm64|aarch64) arch="arm64" ;;
        *)            log_error "Unsupported architecture: $(uname -m)" && exit 1 ;;
    esac
    
    echo "${os}_${arch}"
}

# Get latest release version
get_latest_version() {
    local latest_url="https://api.github.com/repos/${REPO}/releases/latest"
    
    if command -v curl >/dev/null 2>&1; then
        curl -s "$latest_url" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
    elif command -v wget >/dev/null 2>&1; then
        wget -qO- "$latest_url" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
    else
        log_error "curl or wget is required"
        exit 1
    fi
}

# Download and install binary
install_binary() {
    local version="$1"
    local platform="$2"
    local archive_name="${BINARY_NAME}_${platform}.tar.gz"
    local download_url="https://github.com/${REPO}/releases/download/${version}/${archive_name}"
    local temp_dir=$(mktemp -d)
    
    log_info "Downloading Pace CLI ${version} for ${platform}..."
    
    # Download
    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "$download_url" -o "${temp_dir}/${archive_name}"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "$download_url" -O "${temp_dir}/${archive_name}"
    else
        log_error "curl or wget is required"
        exit 1
    fi
    
    # Extract
    log_info "Extracting archive..."
    tar -xzf "${temp_dir}/${archive_name}" -C "$temp_dir"
    
    # Install
    log_info "Installing to ${INSTALL_DIR}..."
    
    # Check if we can write to install directory
    if [ -w "$INSTALL_DIR" ] || [ "$EUID" -eq 0 ]; then
        mv "${temp_dir}/${BINARY_NAME}" "$INSTALL_DIR/"
    else
        log_warning "Need sudo permissions to install to ${INSTALL_DIR}"
        sudo mv "${temp_dir}/${BINARY_NAME}" "$INSTALL_DIR/"
    fi
    
    # Make executable
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    
    # Cleanup
    rm -rf "$temp_dir"
    
    log_success "Pace CLI installed successfully!"
}

# Verify installation
verify_installation() {
    if command -v pace >/dev/null 2>&1; then
        local version=$(pace --version 2>/dev/null || echo "unknown")
        log_success "Installation verified. Version: ${version}"
        log_info "Run 'pace' to get started!"
    else
        log_warning "Installation completed but 'pace' command not found in PATH"
        log_info "You may need to add ${INSTALL_DIR} to your PATH or restart your shell"
    fi
}

# Check if already installed
check_existing_installation() {
    if command -v pace >/dev/null 2>&1; then
        local current_version=$(pace --version 2>/dev/null || echo "unknown")
        log_info "Found existing installation: ${current_version}"
        
        read -p "Do you want to continue with the installation? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            log_info "Installation cancelled"
            exit 0
        fi
    fi
}

# Main installation function
main() {
    echo
    log_info "🚀 Pace CLI Installer"
    echo
    
    # Check for existing installation
    check_existing_installation
    
    # Detect platform
    local platform=$(detect_platform)
    log_info "Detected platform: ${platform}"
    
    # Get latest version
    log_info "Fetching latest release information..."
    local version=$(get_latest_version)
    
    if [ -z "$version" ]; then
        log_error "Failed to get latest version"
        exit 1
    fi
    
    log_info "Latest version: ${version}"
    
    # Install
    install_binary "$version" "$platform"
    
    # Verify
    verify_installation
    
    echo
    log_success "🎉 Installation complete!"
    echo
    log_info "Get started with: pace"
    log_info "Run SSH server: pace --serve"
    log_info "For help: pace --help"
    echo
}

# Run the installer
main "$@"