package scene

import (
	"example.raylib.gamedev/game/asset"
	"example.raylib.gamedev/engine"
)

type SceneManager struct {
	scenes map[string]Scene
}

func (sm *SceneManager) GetSceneByName(name string) Scene {
	return sm.scenes[name]
}

type Scene interface {
	GetName() string
	Render(map[string]asset.Loaded, engine.Drawer)
	// Load(loaded map[string]asset.Loaded)
	GetAssetsGroup() asset.AssetsGroup
	Start()
	Update() string
	Stop()
}

func NewSceneManager(s map[string]Scene) *SceneManager {
	return &SceneManager{
		scenes: s,
	}
}