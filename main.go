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
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type rerenderMsg struct{}

func triggerRerender() tea.Msg {
	return rerenderMsg{}
}

type model struct {
	textInput textinput.Model
	viewport  viewport.Model
	history   []string
	width     int
	height    int
}

// Initialize the model with default values
func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type a command..."
	ti.Focus()
	ti.CharLimit = 156

	vp := viewport.New(80, 20) // Default viewport size

	return model{
		textInput: ti,
		viewport:  vp,
		history:   []string{},
		width:     80,
		height:    24,
	}
}

// Bubble Tea's init function (optional)
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update the model based on incoming messages
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle window resizing
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Resize viewport and text input based on new dimensions
		m.viewport.Width = msg.Width - 2
		m.viewport.Height = msg.Height - 5
		m.textInput.Width = msg.Width - 4
		return m, nil

	// Handle keyboard input
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Process the command when Enter is pressed
			command := strings.TrimSpace(m.textInput.Value())
			m.textInput.SetValue("") // Clear the input

			// Add the new command to the history or clear the history
			if command == "clear" {
				m.history = []string{}

			} else if command != "" {
				m.history = append(m.history, command)
			}
			m.viewport.SetContent(strings.Join(m.history, "\n"))

		case "esc":
			return m, tea.Quit

		}

	}

	// Update the text input and viewport state
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	m.viewport, _ = m.viewport.Update(msg)

	return m, cmd
}

// View renders the application's UI
func (m model) View() string {
	// Define a border style using lipgloss
	borderStyle := lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(1, 2)

	// Combine the viewport and input box into a layout
	return fmt.Sprintf(
		"%s\n\n%s",
		borderStyle.Render(m.viewport.View()), // Display history in a bordered box
		m.textInput.View(),                    // Show the input box
	)
}

func main() {
	// Create a new program with window size handling and full-screen mode
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
