package engine

import (
	"image/color"
	"example.raylib.gamedev/hello-world/engine/enum"
	
)

type Window interface {
	ShouldClose() bool
	Init(int, int, string)
	Close()
}

type SetupEngine interface {
	TargetFPS(int)
	ExitKey(int)
	InitAudioDevice()
}

type Loader interface { 
	Texture2D(string) interface{}
	Music(string) interface{}
	Sound(string) interface{}
}

type Unloader interface { 
	Texture2D(interface{})
	Music(interface{})
	Sound(interface{})
}

type Drawer interface {
	Begin()
	Clear(color.RGBA)
	End()
	Text(string, int, DrawerObject[int])
	Texture2D(interface{}, DrawerObject[int])
	RectangleCharacter(interface{}, Movable)
}

type Listener interface {
	Keyboard() Event[enum.KeyEvent]
}

type MusicPlayer interface {
	Play(interface{})
	Update(interface{})
	Pause(interface{})
	Resume(interface{})
}

type Engine interface {
	// Renderer() <-chan interface{}
	// Start()
	Window()         Window
	Setup()          SetupEngine
	Loader()         Loader
	Unloader()       Unloader
	Drawer()         Drawer
	Listener()       Listener
	MusicPlayer()    MusicPlayer
}
