package main

import (
	"image/color"

	"example.raylib.gamedev/game"
	"example.raylib.gamedev/hellogame"
	engine "example.raylib.gamedev/raylib"
)

func main() {
	ngi := engine.NewRaylib()
	s := hellogame.NewScenes(ngi)
	// start renderer routine

	game := game.NewGame(
		game.GameConfig{
			Screen: game.Screen{Width: 800, Height: 450},
			FPS: 60,
			BackgroundColor: color.RGBA{147, 211, 196, 255},
			Engine: ngi,
		},
		s)

	game.Start()
	defer game.Stop()

	for game.IsRunning() {
		game.Update()
	}
}
