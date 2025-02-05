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
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type model struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

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
			// Initialize the viewport once we know the window size.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.SetContent(m.content)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	// Update the viewport with the current message (for scrolling, etc.)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	// The view consists of a header, the viewport content, and a footer.
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	// Exit help text to show the user how to exit the pager.
	exitHelp := infoStyle.Render("Press q, esc, or ctrl+c to exit")
	// Show the scroll percentage.
	scrollInfo := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	// Fill the remaining space with a horizontal line.
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

func main() {
	var content []byte
	var err error

	// Check if there is piped input by examining os.Stdin.
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("could not stat STDIN:", err)
		os.Exit(1)
	}

	// Read from STDIN if data is being piped in.
	if fi.Mode()&os.ModeCharDevice == 0 {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("could not read from STDIN:", err)
			os.Exit(1)
		}
	}

	p := tea.NewProgram(
		model{content: string(content)},
		tea.WithAltScreen(),       // use the full terminal (alternate screen buffer)
		tea.WithMouseCellMotion(), // enable mouse support
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
