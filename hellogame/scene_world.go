package hellogame
import (
	"image/color"

	"example.raylib.gamedev/engine"
	engineenum "example.raylib.gamedev/engine/enum"
	"example.raylib.gamedev/game/asset"
	"example.raylib.gamedev/game/character"
	"example.raylib.gamedev/game/scene"
)

const WORLD = "world"

type world struct {
	Assets asset.AssetsGroup
	name   string
	char   *character.MovableCharacter
	engine engine.Engine
}

func (s *world) GetName() string {
	return s.name
}

func (s *world) Start() {

}

func (s *world) Update() string {
	kybrd := s.engine.Listener().Keyboard() // down
	s.triggerKeyboard(kybrd)
	if kybrd.Value == engineenum.KeyEvent_ENTER {
		return "hello"
	}
	return "world"
}
func (s *world) Stop() {

}

func (s *world) triggerKeyboard(evt engine.Event[engineenum.KeyEvent]) {
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

func (s *world) Render(loaded map[string]asset.Loaded, drawer engine.Drawer) {
	drawer.Text("You changed Scene!", 40, engine.DrawerObject[int]{
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
		Tint: color.RGBA{0, 0, 255, 255},
	})
	s.char.Render(loaded, drawer)
	// rl.DrawTexture(grass.Loaded.(rl.Texture2D), 100, 50, rl.White)
	// rl.DrawTexture(character.Loaded.(rl.Texture2D), 300, 100, rl.White)

}

func (s *world) GetAssetsGroup() asset.AssetsGroup {
	return s.Assets
}

func NewWorldScene(ngi engine.Engine, char *character.MovableCharacter) scene.Scene {
	return &world{
		char: char,
		name: WORLD,
		Assets: asset.NewAssetGroup([]*asset.Asset{
			&characterAsset,
		}),
		engine: ngi,
	}
}