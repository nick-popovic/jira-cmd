package uicomponents

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model to keep track of the current selected tab
type TabModel struct {
	selectedTab int
}

// Define styles using Lip Gloss
var (
	tabs            = []string{"AI 󰧑", "Search ", "Console 󰆍", "Active Sprint(s) ⚪", "Backlog(s) ", "Risks ⚠", "Dependencies "}
	activeTabStyle  = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("#000")).Background(lipgloss.Color("#FFF")).Bold(true)
	contentStyle    = lipgloss.NewStyle().Padding(1, 2).Foreground(lipgloss.Color("#888"))
	highlightedText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#444"))
	tabStyle        = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("#888"))
)

// Initialize the application model
//func initialModel() model {
//	return model{selectedTab: 0}
//}

// Init function for Bubble Tea
func (m TabModel) Init() tea.Cmd {
	return nil
}

// Message types for tab switching
type tabMsg int

// Update function for Bubble Tea
func (m TabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.selectedTab > 0 {
				m.selectedTab--
			}
		case "right":
			if m.selectedTab < len(tabs)-1 {
				m.selectedTab++
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View function for rendering the tabs and content
func (m TabModel) View() string {
	// Render Tabs
	var renderedTabs string
	for i, tab := range tabs {
		if i == m.selectedTab {
			renderedTabs += activeTabStyle.Render(tab)
		} else {
			renderedTabs += tabStyle.Render(tab)
		}
		if i < len(tabs)-1 {
			renderedTabs += " | "
		}
	}

	// Display content based on selected tab

	/*
		var content string
		switch m.selectedTab {
		case 0:
			content = "Welcome to the Home tab!"
		case 1:
			content = "This is your Profile tab."
		case 2:
			content = "Adjust your Settings here."
		}
	*/

	// Final view
	return fmt.Sprintf("%s", renderedTabs)
}
