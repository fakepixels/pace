# Pace CLI

.------..------..------..------.
|P.--. ||A.--. ||C.--. ||E.--. |
| :/\: || (\/) || :/\: || (\/) |
| (__) || :\/: || :\/: || :\/: |
| '--'P|| '--'A|| '--'C|| '--'E|
'------''------''------''------'

A CLI tool for Pace Capital announcements and updates.

## Installation

### Option 1: Build from Source

If you have Go 1.24.4 or later installed:

```bash
go install github.com/fakepixels/pace@latest
```

### Option 2: Download Pre-built Binary

1. Visit the [releases page](https://github.com/fakepixels/pace/releases/latest)
2. Download the archive for your operating system:
   - macOS Intel: `pace_Darwin_x86_64.tar.gz`
   - macOS Apple Silicon: `pace_Darwin_arm64.tar.gz`
   - Linux x86_64: `pace_Linux_x86_64.tar.gz`
   - Linux ARM64: `pace_Linux_arm64.tar.gz`
   - Windows x86_64: `pace_Windows_x86_64.zip`
   - Windows ARM64: `pace_Windows_arm64.zip`
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

Simply run:

```bash
pace
```

Navigate through the menu using:
- ↑/↓ or k/j to move
- Enter to select
- b to go back
- q to quit

## Features

- Read announcement posts
- Send a hello email to Tina
- Discover random Substack posts
- Visit the announcement site

## License

MIT

Welcome to the **Pace CLI**. 

> "Hey! Only real G found their way here."

---

## 🌐 Announcement Site

Check out the full announcement online:

👉 [pace-announcement.vercel.app](https://pace-announcement.vercel.app/)

---

## 💡 About

Pace CLI is a playful, artful terminal app for the Pace community, built with 💙 using [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Lip Gloss](https://github.com/charmbracelet/lipgloss), and [Glamour](https://github.com/charmbracelet/glamour).

---

Made by [Tina](mailto:tina@pacecapital.com).