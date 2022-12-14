package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TargetSprite struct {
	Sprite    []Target
	HitSprite int
}

type Target struct {
	C      Coordinate
	Width  float64
	Height float64
	Color  color.RGBA
}

func (t *Target) Location() Coordinate {
	return t.C
}

func (t *Target) H() float64 {
	return t.Height
}

func (t *Target) W() float64 {
	return t.Width
}

func (t *Target) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(t.C.X), float64(t.C.Y), t.Width, t.Height, t.Color)

}

func NewTargetSprite() *TargetSprite {
	targetSprite := make([]Target, NUM_ROWS*NUM_TARGETS)
	// Five rows of target sprites
	for i := 0; i < NUM_ROWS; i++ {

		for j := 0; j < NUM_TARGETS; j++ {
			tmpCoordinate := &Coordinate{
				X: SCREEN_HORIZONTAL_PADDING + float64(j)*(SPRITE_HORIZONTAL_DISTANCE+TARGET_WIDTH),
				Y: SCREEN_VERTICAL_PADDING + float64(i)*(SPRITE_VERTICAL_DISTANCE+TARGET_HEIGHT),
			}
			targetSprite = append(targetSprite, Target{
				*tmpCoordinate,
				TARGET_WIDTH,
				TARGET_HEIGHT,
				PALETTE[i],
			})
		}
	}

	return &TargetSprite{
		targetSprite,
		0,
	}
}

func (t *TargetSprite) Draw(dst *ebiten.Image) {
	for _, v := range t.Sprite {
		v.Draw(dst)
	}
}

func (t *TargetSprite) Remove(index int) {
	t.Sprite[index] = t.Sprite[len(t.Sprite)-1]
	t.Sprite = t.Sprite[:len(t.Sprite)-1]
}
