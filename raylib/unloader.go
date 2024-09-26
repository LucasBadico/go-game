package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "example.raylib.gamedev/engine"

)

type Unloader struct {

}

func (l *Unloader) Texture2D(loaded interface{}) {
	rl.UnloadTexture(loaded.(rl.Texture2D))
}

func (l *Unloader) Music(loaded interface{}) {
	rl.UnloadAudioStream(loaded.(rl.Music).Stream)
}

func (l *Unloader) Sound(loaded interface{}) {
	rl.UnloadSound(loaded.(rl.Sound))
}


func NewUnloader() *Unloader {
	return &Unloader{}
}
