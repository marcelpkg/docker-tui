package main

import (
	"fmt"

	"github.com/docker/docker/client"
	docker "github.com/marcelpkg/docker-tui/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	containers []docker.Container
	cursor     int
}

func initialModel() model {
	return model{
		containers: make([]docker.Container, 0),
		cursor:     0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.containers = docker.GetContainers()
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "j", "down":
			if m.cursor < len(m.containers)-1 {
				m.cursor++
			}
		case "t", " ":
			target := m.containers[m.cursor]
			if target.IsRunning() {
				target.Pause()
			} else {
				target.Resume()
			}
		}
	}

	return m, nil
}

var style = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F3F3F3")).
	Border(lipgloss.RoundedBorder(), true, true).
	PaddingTop(1).
	PaddingLeft(4).
	Width(48)

func (m model) View() string {
	display := ""

	d := docker.GetClient()
	defer func(d *client.Client) {
		err := d.Close()
		if err != nil {

		}
	}(d)

	containers := docker.GetContainers()

	if len(containers) == 0 {
		display = "No containers found!\n"
	} else {
		for i, element := range containers {
			selected := " "
			if i == m.cursor {
				selected = "→ "
			}

			status := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF3344")).Render("✘ Paused ")

			if element.IsRunning() {
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#33FF55")).Render("✓ Running ")
			}

			display += fmt.Sprintln(selected + status + element.Names[0])
		}
	}

	controlStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d9d9d9")).Render("\nt - pause/unpause")

	display += controlStyle

	return style.Render(display)
}

func main() {
	tea.ClearScreen()
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
