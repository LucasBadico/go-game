package hellogame

import (
	"example.raylib.gamedev/engine"
	"example.raylib.gamedev/game/character"
	"example.raylib.gamedev/game/scene"
)

func NewScenes(ngi engine.Engine) *scene.SceneManager {
	scenes := map[string]scene.Scene{}
	char := character.NewMovableCharacter("character.basic", 1.0)
	scenes[HELLO] = NewHelloScene(ngi, char)
	scenes[WORLD] = NewWorldScene(ngi, char)
	return scene.NewSceneManager(scenes)
	// return &scene.SceneManager{
	// 	scenes: scenes,
	// }
}
