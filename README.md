# Pace CLI

![pace-cli.png](public/pace-cli.png)

Welcome to the **Pace CLI**.

> "Hey! Only real G found their way here."

---

A CLI tool for Pace Capital announcements and updates, accessible both locally and over SSH.

### ✨ What's New
- **One-line installer** for instant setup
- **Enhanced developer experience** with comprehensive tooling
- **Automated releases** with cross-platform binaries
- **Multiple installation methods** to fit your workflow

## 🚀 Installation

### Quick Install (Recommended)

**One command to install on macOS/Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/fakepixels/pace/main/install.sh | bash
```

### Package Managers

**Homebrew (macOS/Linux):**
```bash
# Coming soon - Homebrew tap in progress
brew install fakepixels/tap/pace
```

**Go users:**
```bash
go install github.com/fakepixels/pace@latest
```

### Other Options

<details>
<summary>📦 Download pre-built binaries</summary>

1. Go to [releases](https://github.com/fakepixels/pace/releases/latest)
2. Download for your platform
3. Extract and move to PATH:
   ```bash
   # macOS/Linux:
   tar xzf pace_*_*.tar.gz && sudo mv pace /usr/local/bin/
   
   # Windows: Extract .zip and add pace.exe to your PATH
   ```
</details>

<details>
<summary>🛠️ Build from source</summary>

Requires [Go 1.24.4+](https://golang.org/dl/)

```bash
git clone https://github.com/fakepixels/pace.git
cd pace
make install    # Builds and installs in one step
```
</details>

### ✅ Verify Installation

```bash
pace --version
```

## 🎮 Usage

### Quick Start
```bash
# Run the app locally
pace

# Check version
pace --version

# Get help
pace --help
```

### Two Ways to Run

**🖥️ Local Mode (Default)**
```bash
pace
```
Launches the interactive terminal app instantly in your current session.

**🌐 SSH Server Mode**
```bash
pace --serve
```
Starts a secure SSH server that others can connect to:
```bash
ssh localhost -p 23234
```

### Navigation
- **↑/↓** or **k/j** - Move through options
- **Enter** - Select
- **b** - Go back
- **q** - Quit

## Features

- Run as a local TUI or a secure, shareable SSH server
- Read announcement posts
- Send a hello email to Tina
- Discover random Substack posts
- Visit [Pace Desktop](https://desktop.pacecapital.com/)
- Secret drops, events, swag

## 🌐 Announcement Site

Check out the full announcement online:

👉 [pace-announcement.vercel.app](https://pace-announcement.vercel.app/)

---

## 🛠️ Development

**Quick Setup:**
```bash
git clone https://github.com/fakepixels/pace.git
cd pace
make setup
make dev
```

**Available Commands:**
- `make dev` - Run in development mode
- `make build` - Build binary
- `make test` - Run tests
- `make install` - Install locally
- `make help` - Show all commands

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## 💡 About

Pace CLI is a playful, artful terminal app for the Pace community, built with 💙 using [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lip Gloss](https://github.com/charmbracelet/lipgloss), [Glamour](https://github.com/charmbracelet/glamour), and [Wish](https://github.com/charmbracelet/wish).

---

Made by [Tina](https://x.com/fkpxls).