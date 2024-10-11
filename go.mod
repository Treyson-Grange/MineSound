module example.com/mymodule

go 1.14

require (
	github.com/charmbracelet/bubbles v0.20.0 // indirect
	github.com/charmbracelet/bubbletea v1.1.1 // indirect
	github.com/charmbracelet/lipgloss v0.13.0 // indirect
	github.com/faiface/beep v1.1.0 // indirect
	github.com/gordonklaus/portaudio v0.0.0-20230709114228-aafa478834f5 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	github.com/hajimehoshi/oto v1.0.1 // indirect
)

replace example.com/thatmodule => ../thatmodule

exclude example.com/thismodule v1.3.0
