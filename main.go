package main

/*
 * Copyright (C) 2024  Nicholas Popovic
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <https://www.gnu.org/licenses/>.
 */

import (
	"fmt"
	"log"

	"main/helpers"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/statusbar"
)

var (
	// Tabs data & styling
	tabs            = []string{"Home", "Profile", "Settings"}
	tabStyle        = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#555"))
	activeTabStyle  = lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("#000")).Background(lipgloss.Color("#FFF")).Bold(true)
	highlightedText = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color("#444"))
)

// model represents the properties of the UI.
type model struct {
	selectedTab int

	loading bool
	spinner spinner.Model

	viewport viewport.Model

	textInput textinput.Model

	statusbar  statusbar.Model
	input_mode string
}

// New creates a new instance of the UI.
func New() model {

	ti := textinput.New()
	ti.Placeholder = "Type here..."
	ti.Focus()

	s := spinner.New()
	s.Spinner = spinner.Dot

	sb := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#A550DF", Dark: "#A550DF"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#6124DF", Dark: "#6124DF"},
		},
	)

	return model{
		loading: false,
		spinner: s,

		statusbar:  sb,
		viewport:   viewport.Model{},
		textInput:  ti,
		input_mode: "INSERT",
	}
}

// Init intializes the UI.
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles all UI interactions.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	// Handle window resize events
	case tea.WindowSizeMsg:

		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 3

		m.statusbar.SetSize(msg.Width)
		m.statusbar.SetContent(m.input_mode, "~/.config/nvim", "1/23", "SB")

	// Handle key presses
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
			cmds = append(cmds, tea.Quit)
		case "esc":
			if m.input_mode == "COMMAND" {
				m.input_mode = "INSERT"
				m.statusbar.SetContent(m.input_mode, "~/.config/nvim", "1/23", "SB")
				m.textInput.Focus()
			} else if m.input_mode == "INSERT" {
				m.input_mode = "COMMAND"
				m.statusbar.SetContent(m.input_mode, "~/.config/nvim", "1/23", "SB")
				m.textInput.Blur()
			}
		case "enter":
			// Handle enter key pressed
			if m.input_mode == "INSERT" {
				// Perform action on enter key pressed in INSERT mode
				//log.Println("Enter key pressed: ", m.textInput.Value())
				m.loading = true
				cmds = append(cmds, m.spinner.Tick, helpers.GetChatCompletion(m.textInput.Value()))
				m.textInput.SetValue("")
			}
		}

	case spinner.TickMsg:

		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

	case helpers.FetchedDataMsg:
		m.loading = false
		if msg.Err == nil {
			r, _ := glamour.NewTermRenderer(
				glamour.WithStandardStyle("dark"),
				glamour.WithWordWrap(m.viewport.Width),
			)
			out, _ := r.Render(msg.Data)
			m.viewport.SetContent(out)
		} else {
			m.viewport.SetContent(msg.Err.Error())
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View returns a string representation of the UI.
func (m model) View() string {

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

	output := fmt.Sprintf("%s", renderedTabs)

	if m.loading {
		m.viewport.SetContent(m.spinner.View() + " Loading...")
	}

	output += "\n" + m.viewport.View() + "\n" + m.textInput.View() + "\n" + m.statusbar.View()

	return output
}

func main() {
	b := New()
	p := tea.NewProgram(b, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
