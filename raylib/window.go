package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
	// "example.raylib.gamedev/hello-world/engine"

)

type Window struct {
	screen Screen
	title  string
}


func (w *Window) ShouldClose() bool {
	return rl.WindowShouldClose()
}

func (w *Window) Init(width, height int, title string) {
	fmt.Println("window#20")
	w.screen = Screen{
		Width: int32(width),
		Height: int32(height),
	}
	fmt.Println("window#25")
	w.title = title
	rl.InitWindow(w.screen.Width, w.screen.Height, "Hello a Golang Game")
	fmt.Println("window#28")
}

func (w *Window) Close() {
 rl.CloseWindow()
}
func NewEmptyWindow() *Window {
	return &Window{}
}