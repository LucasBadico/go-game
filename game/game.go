package game

import (
	"errors"
	"fmt"
	"image/color"

	"example.raylib.gamedev/game/asset"
	"example.raylib.gamedev/game/asset/enum"
	"example.raylib.gamedev/game/scene"
	"example.raylib.gamedev/engine"

)

type Game interface {
	Start()
	Stop()

	Unload([]*asset.Asset) error
	Load([]*asset.Asset) error

	Update()
	Render(scene.Scene)

	IsRunning() bool
}
type Screen struct {
	Height int
	Width  int
}
type game struct {
	screen        Screen
	bgColor       color.RGBA
	fps           int
	scenes       *scene.SceneManager
	loaded        map[string]asset.Loaded
	currentScene  scene.Scene
	engine 		  engine.Engine
	nextSceneName string
}

type GameConfig struct {
	Screen          Screen
	BackgroundColor color.RGBA
	FPS             int
	Engine          engine.Engine
}

func (g *game) Render(s scene.Scene) {
	g.engine.Drawer().Begin()
	g.engine.Drawer().Clear(g.bgColor)
	s.Render(g.loaded, g.engine.Drawer())
	g.engine.Drawer().End()
}

func (g *game) Start() {
	g.engine.Window().Init(g.screen.Width, g.screen.Height, "Hello a Golang Game")
	g.engine.Setup().TargetFPS(g.fps)
	g.engine.Setup().ExitKey(0)
	g.engine.Setup().InitAudioDevice()
}

func (g *game) Stop() {
	assetsToUnload := make([]*asset.Asset, 0)
	for _, asst := range g.loaded {
		assetsToUnload = append(assetsToUnload, &asst.Asset)
	}
	g.Unload(assetsToUnload)
	g.engine.Window().Close()
}

func (g *game) IsRunning() bool {
	return !g.engine.Window().ShouldClose()
}

func (g *game) Load(ag []*asset.Asset) error {
	for _, asst := range ag {
		switch asst.Type {
		case enum.AssetType_2D_TEXTURE:
			loaded := asset.Loaded{
				Asset:  *asst,
				Loaded: g.engine.Loader().Texture2D(asst.Path), 
			}
			g.loaded[asst.Key] = loaded
		case enum.AssetType_AUDIO_MUSIC:
			loaded := asset.Loaded{
				Asset:  *asst,
				Loaded: g.engine.Loader().Music(asst.Path), 
			}
			g.loaded[asst.Key] = loaded
		case enum.AssetType_AUDIO_SOUND:
			loaded := asset.Loaded{
				Asset:  *asst,
				Loaded: g.engine.Loader().Sound(asst.Path), 
			}
			g.loaded[asst.Key] = loaded
		default:
			return errors.New("Asset type not found, please verify")
		}

	}
	return nil
}

func (g *game) SetNextSceneName(name string) {
	g.nextSceneName = name
}

func (g *game) Unload(ag []*asset.Asset) error {
	for _, asst := range ag {
		switch asst.Type {
		case enum.AssetType_2D_TEXTURE:
			loaded := g.loaded[asst.Key]
			// unload
			g.engine.Unloader().Texture2D(loaded.Loaded)
			delete(g.loaded, asst.Key)
		case enum.AssetType_AUDIO_MUSIC:
			loaded := g.loaded[asst.Key]
			// unload
			g.engine.Unloader().Music(loaded.Loaded)
			delete(g.loaded, asst.Key)
		case enum.AssetType_AUDIO_SOUND:
			loaded := g.loaded[asst.Key]
			// unload
			g.engine.Unloader().Sound(loaded.Loaded)
			delete(g.loaded, asst.Key)
		default:
			return errors.New("Asset type not found, please verify")
		}

	}
	return nil
}

func (g *game) loadFirstScene(firstScene scene.Scene) {
	firstAssetsGroup := firstScene.GetAssetsGroup().GetAssets()
	fmt.Println(g.loaded)
	g.Load(firstAssetsGroup)
	fmt.Println(g.loaded)

	g.currentScene = firstScene
	return
}

func (g *game) loadNextScene(nextScene scene.Scene) {
	nextAssetsGroup := nextScene.GetAssetsGroup()
	currAssetsGroup := g.currentScene.GetAssetsGroup()

	stayTheSame := currAssetsGroup.Intersects(&nextAssetsGroup)
	toRemove := currAssetsGroup.Diffs(&stayTheSame).GetAssets()
	toAdd := nextAssetsGroup.Diffs(&stayTheSame).GetAssets()

	g.currentScene.Stop()
	if len(toRemove) > 0 {
		g.Unload(toRemove)
	}
	if len(toAdd) > 0 {
		g.Load(toAdd)
	}
	g.currentScene = nextScene
	g.currentScene.Start()
	return
}

func (g *game) Update() {
	nextSceneName := g.nextSceneName
	nextScene := g.scenes.GetSceneByName(nextSceneName)

	if g.currentScene == nil {
		g.loadFirstScene(nextScene)
		return
	}

	if g.currentScene.GetName() != nextSceneName {
		g.loadNextScene(nextScene)
	}
	g.nextSceneName = g.currentScene.Update()
	g.Render(g.currentScene)
	return
	
}

func NewGame(config GameConfig, scenes *scene.SceneManager) Game {
	return &game{
		screen:       config.Screen,
		bgColor:      config.BackgroundColor,
		fps:          config.FPS,
		scenes:       scenes,
		loaded:       map[string]asset.Loaded{},
		engine:       config.Engine,
		nextSceneName: "hello",
	}
}
