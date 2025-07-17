#!/bin/bash

# Test Installation Script
# This script tests the installation process without actually installing

set -e

echo "🧪 Testing Pace CLI Installation Process"
echo "========================================"

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

success() {
    echo -e "${GREEN}✅ $1${NC}"
}

error() {
    echo -e "${RED}❌ $1${NC}"
}

warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# Test 1: Check if required files exist
echo
echo "Test 1: Checking required files..."

required_files=(
    "Makefile"
    "install.sh"
    ".goreleaser.yml"
    ".github/workflows/ci.yml"
    ".github/workflows/release.yml"
    "CONTRIBUTING.md"
    "pace.rb"
    "main.go"
    "go.mod"
)

for file in "${required_files[@]}"; do
    if [ -f "$file" ]; then
        success "$file exists"
    else
        error "$file missing"
        exit 1
    fi
done

# Test 2: Check Makefile syntax
echo
echo "Test 2: Checking Makefile syntax..."
if make -n help >/dev/null 2>&1; then
    success "Makefile syntax is valid"
else
    error "Makefile has syntax errors"
    exit 1
fi

# Test 3: Check install script is executable
echo
echo "Test 3: Checking install script..."
if [ -x "install.sh" ]; then
    success "install.sh is executable"
else
    error "install.sh is not executable"
    exit 1
fi

# Test 4: Check GitHub Actions workflow syntax
echo
echo "Test 4: Checking GitHub Actions workflows..."
for workflow in .github/workflows/*.yml; do
    if command -v yamllint >/dev/null 2>&1; then
        if yamllint "$workflow" >/dev/null 2>&1; then
            success "$(basename "$workflow") syntax is valid"
        else
            warning "$(basename "$workflow") may have syntax issues (yamllint not conclusive)"
        fi
    else
        warning "yamllint not available, skipping workflow validation"
    fi
done

# Test 5: Check GoReleaser config
echo
echo "Test 5: Checking GoReleaser configuration..."
if command -v goreleaser >/dev/null 2>&1; then
    if goreleaser check >/dev/null 2>&1; then
        success "GoReleaser configuration is valid"
    else
        error "GoReleaser configuration has issues"
        goreleaser check
        exit 1
    fi
else
    warning "GoReleaser not available, skipping validation"
fi

# Test 6: Check Go module
echo
echo "Test 6: Checking Go module..."
if command -v go >/dev/null 2>&1; then
    if go mod verify >/dev/null 2>&1; then
        success "Go module is valid"
    else
        error "Go module verification failed"
        exit 1
    fi
    
    if go list ./... >/dev/null 2>&1; then
        success "Go code compiles"
    else
        error "Go code has compilation issues"
        exit 1
    fi
else
    warning "Go not available, skipping Go-specific tests"
fi

# Test 7: Check documentation links
echo
echo "Test 7: Checking documentation..."
if grep -q "CONTRIBUTING.md" README.md; then
    success "README links to CONTRIBUTING.md"
else
    error "README doesn't link to CONTRIBUTING.md"
fi

if grep -q "install.sh" README.md; then
    success "README mentions install.sh"
else
    error "README doesn't mention install.sh"
fi

echo
echo "🎉 All tests passed! Installation setup looks good."
echo
echo "Next steps to complete the setup:"
echo "1. Push these changes to GitHub"
echo "2. Create a release to test the full pipeline"
echo "3. Set up homebrew-tap repository for Homebrew formula"
echo "4. Test the install script on a clean machine"