package raylib
import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "example.raylib.gamedev/engine"

)

type MusicPlayer struct {

}

func (a *MusicPlayer) Play(loaded interface{}) {
	rl.PlayMusicStream(loaded.(rl.Music))
}

func (a *MusicPlayer) Update(loaded interface{}) {
	rl.UpdateMusicStream(loaded.(rl.Music))
}

func (a *MusicPlayer) Pause(loaded interface{}) {
	rl.PauseMusicStream(loaded.(rl.Music))
}

func (a *MusicPlayer) Resume(loaded interface{}) {
	rl.ResumeMusicStream(loaded.(rl.Music))
}

func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{}
}