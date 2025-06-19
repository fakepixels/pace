package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
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

var announcementText = `Every morning, we wake into histories unfolding, currents of possibility quietly reshaping the future. We stand together at the edge of a great transformation, sensing the gathering force of change. This moment is thrilling and unsettling, stripping away distractions and revealing the one thing that truly matters: our infinite capacity for creation.

Creation has never been this intimate or explosive. Agentic systems, 24/7 onchain markets, open-source chips, neural interfaces, and generative media engines have scattered the seeds of genesis into every corner of our lives. We have compressed the gap between imagination and instantiation down to milliseconds.
Such profound compression demands a new kind of partner. Partners who can navigate entropy, sense the right moment to intervene, and place contrarian bets before the faint signal becomes consensus. Partners possessed by the same obsession that keeps founders awake at 3 a.m.

Obsession isn't a lifestyle, it's the necessary cost of conviction. I learned that building my startup Station Labs, hammering late‑night builds with my co-founders until we finally felt like the product was pushing the boundaries of onchain DevX and UX. We later got acquired by Coinbase to become part of Base, where I led the 0 to 1 of some of the most well-loved developer tools including OnchainKit, MiniKit, and Base MCP. Obsession compels us to stake a claim long before others even realize there's a map.
The most transformative opportunities arise where deep obsession meets potent societal shifts. The migration from traditional finance to crypto rails, the transition from fossil fuels to renewables, the decentralization of inference to the edge, and the evolution from passive media to personalized experiences—all begin subtly, then swiftly remake our world.

Building at these turning points, just before their magnitude becomes undeniable, fuels my day and keeps me awake at night.

I've known the Pace team and witnessed its growth for many years, first as a junior investor, then as a founder they backed. As much as rejoining the team feels like homecoming, it feels more like joining something brand new: bigger, sharper, and even more assured of what it represents in the ecosystem.

Pace is built for the sliver of time after impossible and before obvious; we practice pre‑consensus investing: go deep, decide quickly, ignore the crowd. We trade broad opinions for sharp theses, comfort for honesty, gradients for stark clarity. Exceptional investing at Pace transcends identifying great companies; it demands precision of thought, rapid learning curves, and dedication to founders.
Central to our philosophy is "double-loop learning", a perpetual refinement of frameworks that illuminate not only the essence of transformative businesses and technologies but also our deepest intuitions and biases. This discipline compels us beyond the question, "Why does this company matter?" and into deeper territory: "Why does it resonate uniquely with me, and what makes me the most valuable partner to this founder?"
Double-loop learning inherently rejects superficial analysis or detached market maps that miss the essence of extraordinary, category-defining companies. Instead, we strive for profoundly personal, fiercely introspective clarity, identifying the sparks destined to ignite new markets before they're visible to the world.

While others construct empires, Pace deliberately wants to build a temple dedicated to clarity, originality, and conviction. Our ambition isn't measured merely by the size of the funds we raise, but by the magnitude of returns our clarity unlocks, the audacity of our challenges to orthodoxy, and our courage to act decisively before consensus emerges.
Authentic partnerships here with founders are born not from comfort but from relentless honesty. We rigorously interrogate assumptions, sharpen intuitions, and demand intentionality at every juncture. We commit ourselves to visionary founders whose obsessions align with the most potent undercurrents reshaping society, whose work defines not just markets but entire paradigms.

If you are tugging at the seams of reality, writing the first line of code, a half‑finished memo, a prototype repo, a voice note scribbled at 3 a.m.—I will meet you there, in the dark, before the world wakes up. I publish ideas, theses, and code frequently, and will be a skin-in-the-game partner in crime. Together, we can make inevitability arrive sooner.`

var (
	neonBlue = lipgloss.Color("#1e90ff") // Neon blue
	neonBlueLight = lipgloss.Color("#63aaff") // Lighter neon blue for selected

	announcementStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(neonBlue).
		Foreground(lipgloss.Color("#0a1a2f")).
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
}

func initialModel() model {
	viewportWidth := 80
	viewportHeight := 20
	paragraphs := strings.Split(announcementText, "\n\n")
	wrappedParagraphs := make([]string, len(paragraphs))
	for i, p := range paragraphs {
		wrappedParagraphs[i] = lipgloss.NewStyle().Width(viewportWidth - 4).Render(p)
	}
	wrapped := strings.Join(wrappedParagraphs, "\n\n")
	vp := viewport.New(viewportWidth, viewportHeight)
	vp.SetContent(wrapped)
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
					m.screen = screenHelloTina
				case 2:
					// Pick a random post
					rand.Seed(time.Now().UnixNano())
					m.selectedPost = m.substackPosts[rand.Intn(len(m.substackPosts))]
					m.screen = screenTryMyLuck
				case 3:
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
			// Re-wrap the text to the new width, preserving paragraphs
			paragraphs := strings.Split(announcementText, "\n\n")
			wrappedParagraphs := make([]string, len(paragraphs))
			for i, p := range paragraphs {
				wrappedParagraphs[i] = lipgloss.NewStyle().Width(m.viewport.Width - 4).Render(p)
			}
			wrapped := strings.Join(wrappedParagraphs, "\n\n")
			m.viewport.SetContent(wrapped)
			m.viewportReady = true
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.screen {
	case screenWelcome:
		return "Hello! Welcome to the CLI app.\n\nPress Enter to continue..."
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
		return "Hello, Tina! 👋\n\nHope you have a wonderful day!\n\nPress 'b' to go back, 'q' to quit."
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

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
} 