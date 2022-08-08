package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TargetSprite struct {
	Sprite          []Target
	RemainingTarget int
}

type Target struct {
	C      Coordinate
	Width  float64
	Height float64
}

func (t *Target) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(t.C.X), float64(t.C.Y), t.Width, t.Height, color.RGBA{45, 90, 39, 0xff})

}
