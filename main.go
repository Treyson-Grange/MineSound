package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

// Color Vars
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[38;5;22m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// Timing Variables | default 1-10 seconds
var Minimum = 1
var Maximum = 10

// Spacing Vars
var Padding = "\n\n\n\n\n\n\n\n\n\n" // can you tell I don't know golang?

// Global channel for stopping sound playback
var stopChannel chan bool

/**
* Main function to start the program
* big ole init function
 */
func main() {
	if len(os.Args) > 2 {
		min, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid minimum value: %v", err)
		} else {
			println("Minimum: ", min)
		}
		max, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid maximum value: %v", err)
		}
		Minimum = min
		Maximum = max
	}

	stopChannel = make(chan bool)
	p := tea.NewProgram(initialModel())

	if err, _ := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

/**
* Model struct to hold the state of the program
* State changes? Change here.
 */
type model struct {
	choices     []string
	choicesText map[string]string
	cursor      int
	selected    string
	isPlaying   bool
	mp3Counts   map[string]int
}

func initialModel() model {
	return model{
		choices: []string{"Cave", "Mobs", "Thunder", "Basaltdelta"},
		choicesText: map[string]string{
			"Cave":        "Classic minecraft cave scares!",
			"Mobs":        "Scariest mob sounds (spooky)",
			"Thunder":     "Thunderstorm/Rain sounds",
			"Basaltdelta": "Sounds from the Basalt Delta",
		},
		cursor:    0,
		isPlaying: false,
		mp3Counts: map[string]int{
			"Cave":        18,
			"Mobs":        5,
			"Thunder":     3,
			"Basaltdelta": 9,
		},
	}
}

/**
* Update function to handle user input
* User input? Change here.
 */
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			if m.isPlaying {
				stopChannel <- true
			}
			m.selected = m.choices[m.cursor]
			m.isPlaying = true
			return m, playMP3Cmd(m.selected, m.mp3Counts)
		}
	}

	return m, nil
}

/**
* View function to display the menu, and overall UI
* Visual changes? Change here.
 */
func (m model) View() string {
	var menu string

	content, err := os.ReadFile("ascii/logo.txt")
	if err != nil {
		log.Fatal(err)
	}
	menu += fmt.Sprintf(Padding+"\n\n%s\n\n", Green+string(content)+Reset)
	menu += "\nBy Treyson Grange\n\nTo start the program, choose a sound stage to play.\n"

	for i, choice := range m.choices {
		cursor := " "
		postfix := " "
		if m.cursor == i {
			cursor = Red + ">" + Reset + Green
		}
		if m.selected == choice {
			postfix = Green + "✔" + Reset
		}
		menu += fmt.Sprintf("%s %s %s\n", cursor, choice+Reset, postfix)
	}

	menu += "\n" + m.choicesText[m.choices[m.cursor]] + "\n\n"

	return fmt.Sprint(menu + Padding)
}
