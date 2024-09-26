package hellogame
import (
	"image/color"
	
	"example.raylib.gamedev/engine"
	engineenum "example.raylib.gamedev/engine/enum"

	"example.raylib.gamedev/game/asset"
	"example.raylib.gamedev/game/scene"
	"example.raylib.gamedev/game/character"
)

var HELLO string = "hello"

type hello struct {
	Assets asset.AssetsGroup
	name   string
	char   *character.MovableCharacter
	engine engine.Engine
}

func (s *hello) GetName() string {
	return s.name
}

func (s *hello) Start() {

}

func (s *hello) Update() string {
	kybrd := s.engine.Listener().Keyboard()
	s.triggerKeyboard(kybrd)
	if kybrd.Value == engineenum.KeyEvent_ENTER {
		return "world"
	}
	return "hello"
}
func (s *hello) Stop() {

}

func (s *hello) triggerKeyboard(evt engine.Event[engineenum.KeyEvent]) {
	switch evt.Value {
		case engineenum.KeyEvent_ArrowUp:
			s.char.Movable.Dest.POS.Y-=s.char.Movable.Velocity
		case engineenum.KeyEvent_ArrowDown:
			s.char.Movable.Dest.POS.Y+=s.char.Movable.Velocity
		case engineenum.KeyEvent_ArrowLeft: 
			s.char.Movable.Dest.POS.X-=s.char.Movable.Velocity
		case engineenum.KeyEvent_ArrowRight:
			s.char.Movable.Dest.POS.X+=s.char.Movable.Velocity
		case engineenum.KeyEvent_Q:
			
	}
}

func (s *hello) Render(loaded map[string]asset.Loaded, drawer engine.Drawer) {
	drawer.Text("This is an Hello World Game!", 20, engine.DrawerObject[int]{
		POS: engine.POS[int]{ X: 10, Y: 10},
		Tint: color.RGBA{0, 0, 128, 255},
	})
	// grass := loaded[grassAsset.Key]
	character := loaded[characterAsset.Key]
	// ngi.Drawer().Texture2D(grass.Loaded,  engine.DrawerObject[int]{
	// 	POS: engine.POS[int]{ X: 100, Y: 50},
	// 	Tint: color.RGBA{255, 255, 255, 255},
	// })
	drawer.Texture2D(character.Loaded,  engine.DrawerObject[int]{
		POS: engine.POS[int]{ X: 100, Y: 50},
		Tint: color.RGBA{255, 255, 255, 255},
	})
	s.char.Render(loaded, drawer)
	// rl.DrawTexture(grass.Loaded.(rl.Texture2D), 100, 50, rl.White)
	// rl.DrawTexture(character.Loaded.(rl.Texture2D), 300, 100, rl.White)

}

func (s *hello) GetAssetsGroup() asset.AssetsGroup {
	return s.Assets
}

func NewHelloScene(ngi engine.Engine, char *character.MovableCharacter) scene.Scene {
	return &hello{
		char: char,
		name: HELLO,
		Assets: asset.NewAssetGroup([]*asset.Asset{
			&grassAsset,
			&characterAsset,
		}),
		engine: ngi,
	}
}