package raylib
import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"example.raylib.gamedev/engine"
	"example.raylib.gamedev/engine/enum"
"fmt"
	// "image/color"
)

type Listener struct {

}

func buildKeyEvent(v enum.KeyEvent) engine.Event[enum.KeyEvent] {
	return engine.Event[enum.KeyEvent]{
		Source: enum.EventSource_KeyDown,
		Value: v,
	}
}

func (l *Listener) Keyboard() engine.Event[enum.KeyEvent] {
	if rl.IsKeyDown(rl.KeyUp) {
		fmt.Println("UP____")
		return buildKeyEvent(enum.KeyEvent_ArrowUp) 
	}
	if rl.IsKeyDown(rl.KeyDown) {
		fmt.Println("DOWN____")
		return buildKeyEvent(enum.KeyEvent_ArrowDown) 
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		fmt.Println("LEFT____")
		return buildKeyEvent(enum.KeyEvent_ArrowLeft) 
	}
	if rl.IsKeyDown(rl.KeyRight) {
		fmt.Println("RIGHT___")
		return buildKeyEvent(enum.KeyEvent_ArrowRight) 
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		fmt.Println("SPC_____")
		return buildKeyEvent(enum.KeyEvent_SPACE) 
	}
	if rl.IsKeyPressed(rl.KeyEnter) {
		fmt.Println("ENTER_____")
		return buildKeyEvent(enum.KeyEvent_ENTER) 
	}
	return buildKeyEvent(enum.KeyEvent_NoKey)
}



func NewListener() *Listener {
	return &Listener{}
}