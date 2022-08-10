package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Bullet struct {
	C      Coordinate
	Height float64
	Width  float64
}

func NewBullet() *Bullet {
	bullet_coordinate := &Coordinate{
		X: BULLET_X,
		Y: BULLET_Y,
	}
	return &Bullet{
		C:      *bullet_coordinate,
		Height: BULLET_HEIGHT,
		Width:  BULLET_WIDTH,
	}
}

func (b *Bullet) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(b.C.X), float64(b.C.Y), b.Width, b.Height, color.RGBA{255, 0, 0, 0xff})

}

func (b *Bullet) Move() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {

	}
	return nil
}
