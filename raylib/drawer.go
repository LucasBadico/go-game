package raylib
import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"example.raylib.gamedev/hello-world/engine"

	"image/color"
)

type Drawer struct {

}

func (d *Drawer) Begin() {
	rl.BeginDrawing()
}

func (d *Drawer) End() {
	rl.EndDrawing()
}

func (d *Drawer) Clear(c color.RGBA) {
	rl.ClearBackground(c)
}

func (d *Drawer) Text(text string, size int, obj engine.DrawerObject[int]) {
	rl.DrawText(text, int32(obj.POS.X), int32(obj.POS.Y), int32(size), obj.Tint)
}

func (d *Drawer) Texture2D(loaded interface{}, obj engine.DrawerObject[int]) {
	rl.DrawTexture(loaded.(rl.Texture2D), int32(obj.POS.X),  int32(obj.POS.Y), obj.Tint)
	// rl.DrawText(text, int32(obj.Pos.X), int32(obj.Pos.Y), int32(size), obj.Tint)
}

func (d *Drawer) RectangleCharacter(loaded interface{}, m engine.Movable) {
	plyrSrc := rl.NewRectangle(
		m.Source.POS.X,
		m.Source.POS.Y,
		m.Source.Width,
		m.Source.Height,
	)
	plyrDst := rl.NewRectangle(
		m.Dest.POS.X,
		m.Dest.POS.Y,
		m.Dest.Width,
		m.Dest.Height,
	)
	rl.DrawTexturePro(
		loaded.(rl.Texture2D),
		plyrSrc,
		plyrDst,
		rl.NewVector2(plyrDst.Width, plyrDst.Height),
		0,
		rl.White,
	)
}


func NewDrawer() *Drawer {
	return &Drawer{}
}