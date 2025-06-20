# Pace CLI

![pace-cli.png](public/pace-cli.png)

Welcome to the **Pace CLI**.

> "Hey! Only real G found their way here."

---

A CLI tool for Pace Capital announcements and updates, accessible both locally and over SSH.

## Installation

### Option 1: Build from Source

If you have Go 1.24.4 or later installed. If you haven't, you can install Go by:

1. Visit [Go's official download page](https://go.dev/dl/)
2. Run the installer and follow the prompts
3. Verify installation by opening a terminal and running:
   ```bash
   go version
   ```

Once Go is installed:

```bash
go install github.com/fakepixels/pace@latest
```

### Option 2: Download Pre-built Binary

1. Visit the [releases page](https://github.com/fakepixels/pace/releases/latest)
2. Download the archive for your operating system.
3. Extract the archive:
   ```bash
   # For macOS/Linux:
   tar xzf pace_*_*.tar.gz
   
   # For Windows:
   # Use Windows Explorer to extract the .zip file
   ```
4. Move the binary to your PATH:
   ```bash
   # macOS/Linux:
   sudo mv pace /usr/local/bin/

   # Windows:
   # Move pace.exe to C:\Windows\System32\
   ```

## Usage

The Pace CLI can be run in two modes:

### 1. Local Mode (Default)

For a quick, local experience, simply run the command:
```bash
pace
```
This will launch the application directly in your current terminal session.

Navigate through the menu using:
- ↑/↓ or k/j to move
- Enter to select
- b to go back
- q to quit

### 2. SSH Server Mode

To run the app as a secure SSH server that others can connect to, use the `--serve` flag:
```bash
pace --serve
```
This will start an SSH server. You can then connect to it from another terminal window:
```bash
ssh localhost -p 23234
```
The first time you start the server, it will automatically generate and save an SSH key in a `.ssh` directory. This key is used to secure the connection.

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

## 💡 About

Pace CLI is a playful, artful terminal app for the Pace community, built with 💙 using [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lip Gloss](https://github.com/charmbracelet/lipgloss), [Glamour](https://github.com/charmbracelet/glamour), and [Wish](https://github.com/charmbracelet/wish).

---

Made by [Tina](https://x.com/fkpxls).