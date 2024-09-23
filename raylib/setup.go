package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "example.raylib.gamedev/hello-world/engine"

)

type SetupEngine struct {

}

func (se *SetupEngine) TargetFPS(amount int) {
	rl.SetTargetFPS(int32(amount))
}

func (se *SetupEngine) ExitKey(keyNumber int) {
	rl.SetExitKey(int32(keyNumber))
}

func (se *SetupEngine) InitAudioDevice() {
	rl.InitAudioDevice()
}

func NewSetupEngine() *SetupEngine {
	return &SetupEngine{}
}