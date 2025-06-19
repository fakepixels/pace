package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/glamour"
	"io/ioutil"
	"os/exec"
)

// App screens
const (
	screenWelcome = iota
	screenMenu
	screenAnnouncement
	screenHelloTina
	screenTryMyLuck
	screenAnnouncementSite
)

const announcementSiteURL = "https://pace-announcement.vercel.app/"

var (
	neonBlue = lipgloss.Color("#1e90ff") // Neon blue
	neonBlueLight = lipgloss.Color("#63aaff") // Lighter neon blue for selected

	announcementStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888")).
		Foreground(lipgloss.Color("#a259f7")).
		Padding(1, 2).
		Margin(1, 2).
		Width(80)
	announcementTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("236")).
		Padding(0, 1)

	menuBoxStyle = lipgloss.NewStyle().
		Padding(0, 2).
		Margin(0, 0).
		Border(lipgloss.NormalBorder()).
		BorderForeground(neonBlue).
		Foreground(neonBlue)

	menuBoxSelectedStyle = menuBoxStyle.Copy().
		BorderForeground(neonBlueLight).
		Background(lipgloss.Color("#24283b")).
		Foreground(neonBlueLight).
		Bold(true)

	menuTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(neonBlue).
		MarginBottom(1)

	menuFooterStyle = lipgloss.NewStyle().
		Foreground(neonBlue).
		MarginTop(1)

	welcomeLogo = `
.------..------..------..------.
|P.--. ||A.--. ||C.--. ||E.--. |
| :/\: || (\/) || :/\: || (\/) |
| (__) || :\/: || :\/: || :\/: |
| '--'P|| '--'A|| '--'C|| '--'E|
'------''------''------''------'
`
	welcomeLogoStyle = lipgloss.NewStyle().
		Foreground(neonBlue).
		Bold(true).
		Margin(1, 0, 0, 0).
		Align(lipgloss.Center)

	welcomeMsgStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a259f7")).
		Bold(true).
		Align(lipgloss.Center)

	welcomeBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(neonBlue).
		Background(lipgloss.Color("#181825")).
		Padding(1, 4).
		Width(60).
		Align(lipgloss.Center)
)

type model struct {
	screen      int
	menuCursor  int
	menuChoices []string
	// For try my luck
	substackPosts []string
	selectedPost string
	// For announcement
	viewport viewport.Model
	viewportReady bool
	announcementMD string // holds the loaded markdown
	announcementRendered string // holds the glamour-rendered output
}

func initialModel() model {
	viewportWidth := 80
	viewportHeight := 20
	// Load announcement markdown from file
	mdBytes, err := ioutil.ReadFile("announcement.md")
	md := ""
	if err == nil {
		md = string(mdBytes)
	}
	// Render with Glamour
	styled := ""
	if md != "" {
		r, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(viewportWidth-4),
		)
		styled, _ = r.Render(md)
	}
	vp := viewport.New(viewportWidth, viewportHeight)
	vp.SetContent(styled)
	return model{
		screen:      screenWelcome,
		menuChoices: []string{"Read announcement post", "Say hello to Tina", "Try my luck", "Go to announcement site"},
		substackPosts: []string{
			"https://fakepixels.substack.com/p/context-is-all-you-need",
			"https://fakepixels.substack.com/p/the-art-of-understanding-whats-going",
			"https://fakepixels.substack.com/p/ai-heidegger-and-evangelion",
			"https://fakepixels.substack.com/p/infinite-playgrounds-the-building",
			"https://fakepixels.substack.com/p/the-art-of-understanding-whats-going",
			"https://fakepixels.substack.com/p/the-spectacle-of-building",
			"https://fakepixels.substack.com/p/bring-the-mind-home",
		},
		viewport: vp,
		announcementMD: md,
		announcementRendered: styled,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.screen {
		case screenWelcome:
			if msg.String() == "enter" {
				m.screen = screenMenu
			}
		case screenMenu:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.menuCursor > 0 {
					m.menuCursor--
				}
			case "down", "j":
				if m.menuCursor < len(m.menuChoices)-1 {
					m.menuCursor++
				}
			case "enter":
				switch m.menuCursor {
				case 0:
					m.screen = screenAnnouncement
				case 1:
					// Open mailto link
					mailto := "mailto:tina@pacecapital.com?subject=Hi%20from%20the%20dark&body=I%20saw%20your%20Pace%20CLI%20and%20wanted%20to%20say%20hi."
					exec.Command("open", mailto).Start()
					m.screen = screenHelloTina
				case 2:
					// Pick a random post
					rand.Seed(time.Now().UnixNano())
					m.selectedPost = m.substackPosts[rand.Intn(len(m.substackPosts))]
					exec.Command("open", m.selectedPost).Start()
					m.screen = screenTryMyLuck
				case 3:
					exec.Command("open", announcementSiteURL).Start()
					m.screen = screenAnnouncementSite
				}
			}
		case screenAnnouncement:
			// Scroll viewport
			switch msg.String() {
			case "up", "k":
				m.viewport.LineUp(1)
			case "down", "j":
				m.viewport.LineDown(1)
			case "pgup":
				m.viewport.SetYOffset(m.viewport.YOffset - m.viewport.Height)
			case "pgdown":
				m.viewport.SetYOffset(m.viewport.YOffset + m.viewport.Height)
			case "b":
				m.screen = screenMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		case screenAnnouncementSite:
			if msg.String() == "b" {
				m.screen = screenMenu
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		case screenHelloTina, screenTryMyLuck:
			if msg.String() == "b" {
				m.screen = screenMenu
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		if m.screen == screenAnnouncement {
			m.viewport.Width = msg.Width - 8 // account for border/margin
			m.viewport.Height = msg.Height - 6
			if m.announcementMD != "" {
				r, _ := glamour.NewTermRenderer(
					glamour.WithAutoStyle(),
					glamour.WithWordWrap(m.viewport.Width-4),
				)
				styled, _ := r.Render(m.announcementMD)
				m.announcementRendered = styled
				m.viewport.SetContent(styled)
			}
			m.viewportReady = true
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.screen {
	case screenWelcome:
		logo := welcomeLogoStyle.Render(welcomeLogo)
		msg := welcomeMsgStyle.Render("Hey! Only real G found their way here.")
		prompt := lipgloss.NewStyle().Align(lipgloss.Center).Render("Press Enter to continue...")
		return "\n" + welcomeBoxStyle.Render(logo+"\n"+msg+"\n\n"+prompt)
	case screenMenu:
		s := menuTitleStyle.Render("What would you like to do?") + "\n"
		for i, choice := range m.menuChoices {
			if m.menuCursor == i {
				s += menuBoxSelectedStyle.Render(choice) + "\n"
			} else {
				s += menuBoxStyle.Render(choice) + "\n"
			}
		}
		s += menuFooterStyle.Render("Use ↑/↓ or k/j to move, Enter to select. Press q to quit.")
		return s
	case screenAnnouncement:
		title := announcementTitleStyle.Render("Announcement Post")
		content := m.viewport.View()
		return title + "\n" + announcementStyle.Render(content) + "\n[↑/↓/PgUp/PgDn scroll, b: back, q: quit]"
	case screenHelloTina:
		return "Mail client opened! Say hi to Tina at tina@pacecapital.com\n\nPress 'b' to go back, 'q' to quit."
	case screenTryMyLuck:
		return fmt.Sprintf("Try My Luck\n\nHere's a random Substack post for you:\n%s\n\nPress 'b' to go back, 'q' to quit.", m.selectedPost)
	case screenAnnouncementSite:
		title := announcementTitleStyle.Render("Announcement Site")
		url := lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Underline(true).Render(announcementSiteURL)
		msg := "Open this link in your browser (copy-paste):\n" + url + "\n\nPress 'b' to go back, 'q' to quit."
		return title + "\n" + announcementStyle.Render(msg)
	}
	return ""
}

var version = "dev" // will be set by goreleaser

func main() {
	flagVersion := false
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagSet.BoolVar(&flagVersion, "version", false, "print version and exit")
	flagSet.BoolVar(&flagVersion, "v", false, "print version and exit (shorthand)")
	flagSet.Parse(os.Args[1:])
	if flagVersion {
		fmt.Println(version)
		return
	}
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
} 