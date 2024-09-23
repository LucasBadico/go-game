package raylib
import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "example.raylib.gamedev/hello-world/engine"

)

type Loader struct {

}

func (l *Loader) Texture2D(path string) interface{} {
	return rl.LoadTexture(path)
}

func (l *Loader) Music(path string) interface{} {
	return rl.LoadMusicStream(path)
}

func (l *Loader) Sound(path string) interface{} {
	return rl.LoadSound(path)
}

func NewLoader() *Loader {
	return &Loader{}
}