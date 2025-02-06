package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Header Styling
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("63")).
			Padding(0, 2).
			BorderStyle(lipgloss.RoundedBorder())

	// Footer Styling
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("63")).
			Padding(0, 2).
			BorderStyle(lipgloss.RoundedBorder())

	// Scroll Percentage Styling
	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("63")).
			Padding(0, 1)
)

type model struct {
	content  string
	title    string
	ready    bool
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render(m.title)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	exitHelp := footerStyle.Render("Press Q / Esc to exit | ↑ / ↓ to scroll")
	scrollInfo := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	remaining := max(0, m.viewport.Width-lipgloss.Width(exitHelp)-lipgloss.Width(scrollInfo))
	line := strings.Repeat("─", remaining)
	return lipgloss.JoinHorizontal(lipgloss.Center, exitHelp, line, scrollInfo)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// StartPager initializes the pager with content and title.
func StartPager(content, title string) error {
	p := tea.NewProgram(
		model{content: content, title: title},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	_, err := p.Run()
	return err
}

func main() {
	var content, title string

	// Check if a file argument is provided
	if len(os.Args) > 1 {
		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Could not read file:", err)
			os.Exit(1)
		}
		content = string(data)
		title = filename // Set filename as the title
	} else {
		// Check if piped input exists
		fi, err := os.Stdin.Stat()
		if err != nil {
			fmt.Println("Could not stat STDIN:", err)
			os.Exit(1)
		}

		if fi.Mode()&os.ModeCharDevice == 0 {
			// Read from stdin
			bytes, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Println("Could not read from STDIN:", err)
				os.Exit(1)
			}
			content = string(bytes)
			title = "Piped Input" // Generic title for piped content
		} else {
			content = "Welcome to the Go Viewport Pager!"
			title = "Go View Pager"
		}
	}

	if err := StartPager(content, title); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
