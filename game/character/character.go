package character
import (
	"example.raylib.gamedev/engine"
	"example.raylib.gamedev/game/asset"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

type MovableCharacter struct {
	engine.Movable
	assetKey string
}

func (cc *MovableCharacter) Render(loaded map[string]asset.Loaded, drawer engine.Drawer) {
	characterSprt := loaded[cc.assetKey]
	drawer.RectangleCharacter(characterSprt.Loaded, cc.Movable)
}

func NewMovableCharacter(assetKey string, velocity float32) *MovableCharacter {
	return &MovableCharacter{
		assetKey: assetKey,
		Movable: engine.Movable{
			Velocity: velocity,
			Source: engine.Rectangle{
				POS: engine.POS[float32]{
					X: 0,
					Y: 0,
				},
				Width: 48,
				Height: 48,
			},
			Dest: engine.Rectangle{
				POS: engine.POS[float32]{
					X: 200,
					Y: 200,
				},
				Width: 100,
				Height: 100,
			},
		},
	}
}