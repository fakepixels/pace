package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// App screens
const (
	screenWelcome = iota
	screenMenu
	screenAnnouncement
	screenTeam
	screenTryMyLuck
	screenPaceDesktop
	screenSignup
	screenSignupConfirm
)

const announcementSiteURL = "https://pace-announcement.vercel.app/"
const paceDesktopURL = "https://desktop.pacecapital.com/"

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
 _______  _______  _______  _______ 
|       ||   _   ||       ||       |
|    _  ||  |_|  ||       ||    ___|
|   |_| ||       ||       ||   |___ 
|    ___||       ||      _||    ___|
|   |    |   _   ||     |_ |   |___ 
|___|    |__| |__||_______||_______|
`
	welcomeLogoStyle = lipgloss.NewStyle().
		Foreground(neonBlue).
		Bold(true)

	welcomeMsgStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a259f7")).
		Bold(true)

	welcomeBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(neonBlue).
		Background(lipgloss.Color("#181825")).
		Padding(1, 4).
		Width(60)

	teamCardStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Background(lipgloss.Color("236")).
		Padding(1, 2).
		Width(40)

	teamCardSelectedStyle = teamCardStyle.Copy().
		BorderForeground(lipgloss.Color("205")).
		Background(lipgloss.Color("17")).
		Bold(true)
)

type teamMember struct {
	name  string
	email string
	emoji string
}

type model struct {
	screen       int
	menuCursor   int
	menuChoices  []string
	teamMembers  []teamMember
	teamCursor   int
	// For try my luck
	substackPosts []string
	selectedPost  string
	// For announcement
	viewport      viewport.Model
	viewportReady bool
	announcementMD string // holds the loaded markdown
	announcementRendered string // holds the glamour-rendered output
	// For signup form
	signupForm *huh.Form
	signupMsg  string
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

	// Create the signup form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("What's your name?"),
			huh.NewInput().
				Key("email").
				Title("Your email address").
				Validate(func(str string) error {
					if !strings.Contains(str, "@") || !strings.Contains(str, ".") {
						return fmt.Errorf("Please enter a valid email address.")
					}
					return nil
				}),
		),
	)

	return model{
		screen:      screenWelcome,
		menuChoices: []string{"Read latest announcement ✨", "Pace Team Members", "Try my luck", "Go to Pace Desktop", "Sign up to stay up to date"},
		teamMembers: []teamMember{
			{name: "Aryan Naik", email: "aryan@pacecapital.com", emoji: "🦄"},
			{name: "Chris Peck", email: "chris@pacecapital.com", emoji: "🦉"},
			{name: "Grace Kasten", email: "grace@pacecapital.com", emoji: "🦋"},
			{name: "Jordan Cooper", email: "jordan@pacecapital.com", emoji: "🦁"},
			{name: "Tina He", email: "tina@pacecapital.com", emoji: "🐉"},
		},
		teamCursor: 0,
		substackPosts: []string{
			"https://fakepixels.substack.com/p/context-is-all-you-need",
			"https://fakepixels.substack.com/p/the-art-of-understanding-whats-going",
			"https://fakepixels.substack.com/p/ai-heidegger-and-evangelion",
			"https://fakepixels.substack.com/p/infinite-playgrounds-the-building",
			"https://fakepixels.substack.com/p/the-art-of-understanding-whats-going",
			"https://fakepixels.substack.com/p/the-spectacle-of-building",
			"https://fakepixels.substack.com/p/bring-the-mind-home",
		},
		viewport:             vp,
		announcementMD:       md,
		announcementRendered: styled,
		signupForm:           form,
	}
}

func (m model) Init() tea.Cmd {
	return nil // No initial command needed
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle the signup form screen separately
	if m.screen == screenSignup {
		form, cmd := m.signupForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.signupForm = f
		}

		if m.signupForm.State == huh.StateCompleted {
			// Form is done, process the data.
			name := m.signupForm.GetString("name")
			email := m.signupForm.GetString("email")

			// Send to Google Sheet or webhook here
			go func(name, email string) {
				// TODO: Replace with your Google Sheet endpoint
				endpoint := "https://your-google-sheet-endpoint.example.com" // <-- Replace this
				data := fmt.Sprintf("name=%s&email=%s", name, email)
				_, err := http.Post(endpoint, "application/x-www-form-urlencoded", strings.NewReader(data))
				if err != nil {
					log.Printf("Failed to send signup: %v", err)
				}
			}(name, email)

			m.signupMsg = "Thank you for signing up! You'll stay up to date."
			m.screen = screenSignupConfirm
		}

		return m, cmd
	}

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
					m.screen = screenTeam
				case 2:
					// Pick a random post
					rand.Seed(time.Now().UnixNano())
					m.selectedPost = m.substackPosts[rand.Intn(len(m.substackPosts))]
					exec.Command("open", m.selectedPost).Start()
					m.screen = screenTryMyLuck
				case 3:
					exec.Command("open", paceDesktopURL).Start()
					m.screen = screenPaceDesktop
				case 4:
					m.screen = screenSignup
					return m, m.signupForm.Init() // Initialize/reset the form
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
			case "o":
				exec.Command("open", announcementSiteURL).Start()
			case "e":
				mailto := "mailto:tina@pacecapital.com?subject=Hi%20from%20the%20dark&body=I%20saw%20your%20Pace%20CLI%20announcement%20and%20wanted%20to%20say%20hi."
				exec.Command("open", mailto).Start()
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		case screenTeam:
			switch msg.String() {
			case "up", "k":
				if m.teamCursor > 0 {
					m.teamCursor--
				}
			case "down", "j":
				if m.teamCursor < len(m.teamMembers)-1 {
					m.teamCursor++
				}
			case "enter":
				selectedMember := m.teamMembers[m.teamCursor]
				if selectedMember.email != "" {
					mailto := fmt.Sprintf("mailto:%s?subject=Hi%%20from%%20the%%20dark&body=I%%20saw%%20your%%20Pace%%20CLI%%20and%%20wanted%%20to%%20say%%20hi.", selectedMember.email)
					exec.Command("open", mailto).Start()
				}
			case "b":
				m.screen = screenMenu
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		case screenPaceDesktop:
			if msg.String() == "b" {
				m.screen = screenMenu
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		case screenTryMyLuck:
			if msg.String() == "b" {
				m.screen = screenMenu
			}
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
		case screenSignupConfirm:
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
	if m.screen == screenSignup {
		return m.signupForm.View()
	}

	switch m.screen {
	case screenWelcome:
		logo := welcomeLogo
		msg := "Welcome to Pace CLI.\nThis is a terminal app to learn about us\nand stay updated for events and surprise drops."
		prompt := "Press Enter to continue..."

		logoStyled := welcomeLogoStyle.Render(logo)
		msgStyled := welcomeMsgStyle.Render(msg)

		content := lipgloss.JoinVertical(
			lipgloss.Center,
			logoStyled,
			"\n",
			msgStyled,
			"\n\n",
			prompt,
		)

		return welcomeBoxStyle.Render(content)
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
		controls := "\n[↑/↓/PgUp/PgDn scroll, o: open in browser, e: email Tina, b: back, q: quit]"
		return title + "\n" + announcementStyle.Render(content) + controls
	case screenTeam:
		s := menuTitleStyle.Render("Pace Team Members") + "\n"
		for i, member := range m.teamMembers {
			cardStyle := teamCardStyle
			if m.teamCursor == i {
				cardStyle = teamCardSelectedStyle
			}

			nameLine := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%s %s", member.emoji, member.name))
			emailLine := ""
			if member.email != "" {
				emailLine = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Underline(true).Render(member.email)
			} else {
				emailLine = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true).Render("No email provided")
			}

			card := fmt.Sprintf("%s\n%s", nameLine, emailLine)
			s += cardStyle.Render(card) + "\n"
		}
		s += menuFooterStyle.Render("\nUse ↑/↓ or k/j to move, Enter to email. Press b to go back.")
		return s
	case screenTryMyLuck:
		return fmt.Sprintf("Try My Luck\n\nHere's a random Substack post for you:\n%s\n\nPress 'b' to go back, 'q' to quit.", m.selectedPost)
	case screenPaceDesktop:
		title := announcementTitleStyle.Render("Pace Desktop")
		url := lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Underline(true).Render(paceDesktopURL)
		msg := "Opening Pace Desktop in your browser...\n" + url + "\n\nPress 'b' to go back, 'q' to quit."
		return title + "\n" + announcementStyle.Render(msg)
	case screenSignupConfirm:
		return fmt.Sprintf("%s\n\nPress 'b' to go back, 'q' to quit.", m.signupMsg)
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