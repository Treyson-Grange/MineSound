module example.com/mymodule

go 1.14

require (
	github.com/charmbracelet/bubbletea v1.1.1
	github.com/faiface/beep v1.1.0
	github.com/hajimehoshi/go-mp3 v0.3.4 // indirect
	github.com/hajimehoshi/oto v1.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
)

replace example.com/thatmodule => ../thatmodule

exclude example.com/thismodule v1.3.0
