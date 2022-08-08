package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	C      Coordinate
	Height float64
	Width  float64
}

func (p *Player) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(p.C.X), float64(p.C.Y), p.Width, p.Height, color.RGBA{45, 90, 39, 0xff})
}

// Update the player's position if there's any keystroke
func (p *Player) Update() error {
	// TODO: if the player hits the wall, stop moving in that direction
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		new_pos := p.C.X + 5
		p.C.SetX(new_pos)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		new_pos := p.C.X - 5
		p.C.SetX(new_pos)
	}

	return nil
}
