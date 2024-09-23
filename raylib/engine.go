package raylib
import (
	"example.raylib.gamedev/hello-world/engine"
)


type RaylibEngine struct {
	// Renderer() <-chan interface{}
	window   *Window
	setup    *SetupEngine
    loader   *Loader
	unloader *Unloader
	drawer   *Drawer
	listener *Listener
	music    *MusicPlayer
}	

func (r *RaylibEngine) Renderer() {

}

func (r *RaylibEngine) Window() engine.Window {
	return r.window
}

func (r *RaylibEngine) Setup() engine.SetupEngine {
	return r.setup
}

func (r *RaylibEngine) Loader() engine.Loader {
	return r.loader
}

func (r *RaylibEngine) Unloader() engine.Unloader {
	return r.unloader
}

func (r *RaylibEngine) Drawer() engine.Drawer {
	return r.drawer
}


func (r *RaylibEngine) Listener() engine.Listener {
	return r.listener
}

func (r *RaylibEngine) MusicPlayer() engine.MusicPlayer {
	return r.music
}

func NewRaylib() *RaylibEngine {
	return &RaylibEngine{
		window:   NewEmptyWindow(),
		unloader: NewUnloader(),
		loader:   NewLoader(),
		setup:    NewSetupEngine(),
		drawer:   NewDrawer(),
		listener: NewListener(),
		music: NewMusicPlayer(),
	}
}