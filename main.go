package main

import (
	"fmt"
	"github.com/docker/docker/client"
	docker "github.com/marcelpkg/docker-tui/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model

type model struct {
	containers []docker.Container
	cursor     int
	window     int
	showAll    bool
}

func initialModel() model {
	return model{
		containers: make([]docker.Container, 0),
		cursor:     0,
		window:     0,
		showAll:    false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Updater

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.containers = docker.GetContainers(m.showAll)
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			// if m.GetView != "none" quit
			// else m.SetView none
			return m, tea.Quit

		case "tab":
			if m.window < 1 {
				m.window++
			} else {
				m.window = 0
			}

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "s":
			if m.containers[m.cursor].IsRunning() {
				go m.containers[m.cursor].Stop()
			}
			go m.containers[m.cursor].Start()

		case "j", "down":
			if m.cursor < len(m.containers)-1 {
				m.cursor++
			}

		case "r":
			target := m.containers[m.cursor]
			go target.Restart()
			fmt.Println(target.State)

		case " ":
			if m.showAll == true {
				if m.cursor > len(docker.GetContainers(false)) {
					m.cursor = 0
				}
			}
			m.showAll = !m.showAll

		case "t":
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
	Width(120)

func (m model) View() string {

	display := ""
	infoDisplay := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).Margin(0, 4).PaddingLeft(4)
	var infoText string

	d := docker.GetClient()
	defer func(d *client.Client) {
		err := d.Close()
		if err != nil {

		}
	}(d)

	containers := docker.GetContainers(m.showAll)

	if len(containers) == 0 {
		display = "No containers found!\n"
	} else {
		currentContainer := containers[m.cursor]
		for i, element := range containers {
			selected := " "
			if i == m.cursor {
				selected = "→ "
			}

			status := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF3344")).Render("✘ Stopped ")

			switch element.State {

			case "running":
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#33FF55")).Render("✓ Running ")

			case "paused":
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFFF00")).Render("… Paused ")

			case "restarting":
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFA500")).Render("↻ Restarting ")

			case "exited":
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FF0000")).Render("! Exited ")

			}

			display += fmt.Sprintln(selected + status + element.Names[0])
		}

		controlStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#d9d9d9")).Render("\nt - pause/unpause | r - restart | space - show all\n")

		display += controlStyle

		infoText += "Name: " + currentContainer.Names[0] + " | " + currentContainer.State + "\n"
		infoText += "Uptime: " + currentContainer.Status + "\n"
		infoText += "Image: " + currentContainer.Image + "\n"

	}

	return style.Render(lipgloss.JoinHorizontal(lipgloss.Left, display, infoDisplay.Render(infoText)))
}

func main() {
	tea.ClearScreen()

	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}
