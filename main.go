package main

import (
	"fmt"
	"github.com/docker/docker/client"
	docker "github.com/marcelpkg/docker-tui/api"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type taskElement struct {
	name        string
	description string
	done        bool
}

type model struct {
	tasks  []taskElement // 0: task
	cursor int
}

func createTaskElement(name string) taskElement {
	return taskElement{
		name:        name,
		description: "",
		done:        false,
	}
}

func initialModel() model {
	tasks := []taskElement{createTaskElement("nothing here...")}

	return model{
		tasks:  tasks,
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "j", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "k", "down":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.tasks[m.cursor].done = !m.tasks[m.cursor].done
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
				selected = ">"
			}

			status := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF3344")).Render("✘ Stopped")

			if element.State == "running" {
				status = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#33FF55")).Render("✓ Running")
			}

			display += fmt.Sprintln(selected + status + element.Names[0])
		}
	}
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
