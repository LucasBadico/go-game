package engine

import (
	"example.raylib.gamedev/engine/enum"

	"image/color"
)

type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Definição de uma estrutura genérica restrita a inteiros e floats
type POS[T Numeric] struct {
    X T
    Y T
}

type Rectangle struct {
	POS POS[float32]
	Width float32
	Height float32
}

type Movable struct {
	Source   Rectangle
	Dest     Rectangle
	Velocity float32
}
// Definindo uma constraint que aceita tanto inteiros quanto floats


type DrawerObject[T Numeric] struct {
	POS POS[T]
	Tint color.RGBA
}

type Event[T any] struct {
	Source enum.EventSource
	Value T
}