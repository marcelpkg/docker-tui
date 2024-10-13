package main

import (
	"fmt"

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
	// m.AddTask("put on skirt")
	// m.AddTask("put on thigh highs")
	// m.AddTask("install go")
	// m.AddTask("log on discord")
	tasks := []taskElement{createTaskElement("put on skirt"), createTaskElement("put on thigh highs"), createTaskElement("install go"), createTaskElement("log on discord")}

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
	Border(lipgloss.DoubleBorder(), true, true).
	PaddingTop(1).
	PaddingLeft(4).
	Width(48)

func (m model) View() string {
	s := ""

	for i, element := range m.tasks {

		selected := " "
		if i == m.cursor {
			selected = ">"
		}

		done := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF3344")).Render("✘")

		if element.done {
			done = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#33FF55")).Render("✓")
		}

		// > [✓] ↓ Task name
		s += fmt.Sprintf("%s %s %s\n", selected, done, element.name)
	}

	return style.Render(s)
}

func (m *model) AddTask(name string) {
	newTask := createTaskElement(name)

	m.tasks = append(m.tasks, newTask)
}

func main() {
	p := tea.NewProgram(initialModel())

	_, err := p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

// Order of events:
// tea.NewProgram(model)
// Error checked
// Initialise with .Init()
// On an event Update() is run (of the model)
// After an Update() finishes, View() is ran.
