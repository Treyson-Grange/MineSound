package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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
	choices   []string
	cursor    int
	selected  string
	isPlaying bool
	mp3Counts map[string]int
}

func initialModel() model {
	return model{
		choices:   []string{"Cave", "Mobs", "Thunder"},
		cursor:    0,
		isPlaying: false,
		mp3Counts: map[string]int{
			"Cave":    18,
			"Mobs":    2,
			"Thunder": 3,
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
			postfix = Green + "✔" + Reset + Green
		}
		menu += fmt.Sprintf("%s %s %s\n", cursor, choice+Reset, postfix)
	}

	if m.cursor == 0 {
		menu += fmt.Sprintf("\n%s\n", "Classic minecraft cave scares!\n")
	} else if m.cursor == 1 {
		menu += fmt.Sprintf("\n%s\n", "Scariest mob sounds (spooky)\n")
	} else if m.cursor == 2 {
		menu += fmt.Sprintf("\n%s\n", "Thunderstorm/Rain sounds\n")
	}
	// More coming, finding more sounds is rough.

	return fmt.Sprintf(menu + Padding)
}

/**
* Helper functions to play random MP3 files
* MP3 System changes? Here it is.
 */
func getRandomMP3(choice string, mp3Counts map[string]int) string {
	count := mp3Counts[choice]
	if count == 0 {
		return ""
	}
	fileIndex := rand.Intn(count) + 1
	return fmt.Sprintf("mp3/%s/%d.mp3", strings.ToLower(choice), fileIndex)
}

func playMP3Cmd(selectedChoice string, mp3Counts map[string]int) tea.Cmd {
	return func() tea.Msg {
		if stopChannel != nil {
			select {
			case stopChannel <- true:
			default:
			}
		}

		stopChannel = make(chan bool)

		go func() {
			for {
				filePath := getRandomMP3(selectedChoice, mp3Counts)
				if filePath == "" {
					return
				}

				err := playMP3(filePath)
				if err != nil {
					log.Fatal(err)
				}

				select {
				case <-stopChannel:
					return
				case <-time.After(time.Duration(rand.Intn(Maximum)+Minimum) * time.Second):
				}
			}
		}()
		return nil
	}
}

func playMP3(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	select {
	case <-done:
		return nil
	case <-stopChannel:
		speaker.Clear()
		return nil
	}
}
