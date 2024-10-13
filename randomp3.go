package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

/**
* Helper functions to play random MP3 files
* MP3 System changes? Here it is.
 */
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

func getRandomMP3(choice string, mp3Counts map[string]int) string {
	count := mp3Counts[choice]
	if count == 0 {
		return ""
	}
	fileIndex := rand.Intn(count) + 1
	return fmt.Sprintf("mp3/%s/%d.mp3", strings.ToLower(choice), fileIndex)
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
