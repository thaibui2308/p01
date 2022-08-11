package breakout

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Bullet struct {
	C             Coordinate
	Height        float64
	Width         float64
	Moving        bool
	CollisionType int
}

func NewBullet() *Bullet {
	bullet_coordinate := &Coordinate{
		X: BULLET_X,
		Y: BULLET_Y,
	}
	return &Bullet{
		C:             *bullet_coordinate,
		Height:        BULLET_HEIGHT,
		Width:         BULLET_WIDTH,
		Moving:        false,
		CollisionType: NO_COLLISION,
	}
}

func (b *Bullet) SetMoving() {
	if b.Moving {
		return
	} else {
		b.Moving = true
	}
}

func (b *Bullet) SetStop() {
	if !b.Moving {
		return
	} else {
		b.Moving = false
	}
}

func (b *Bullet) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, float64(b.C.X), float64(b.C.Y), b.Width, b.Height, color.RGBA{255, 0, 0, 0xff})

}

func (b *Bullet) CollisionDetect(sprite *TargetSprite) (index []int) {
	index = make([]int, len(sprite.Sprite))

	// four cases
	for i, v := range sprite.Sprite {
		// Collision at the top and bottom of the Target's block
		if b.C.X+b.Width >= v.C.X && b.C.X <= v.C.X+v.Width {
			if b.C.Y <= v.C.Y+v.Height && b.C.Y >= v.C.Y {
				index = append(index, i)
				fmt.Println("Buzz")
				continue
			}
		} else if b.C.Y+b.Height >= v.C.Y && b.C.Y <= v.C.Y+v.Height {
			if b.C.X <= v.C.X+v.Width && b.C.X >= v.C.X {
				index = append(index, i)
				fmt.Println("Buzz")
				continue
			}
		}
	}
	return
}

func (b *Bullet) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.SetMoving()
	} else if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		b.SetStop()
	}

	if b.Moving {
		b.C.BounceUpward(BOUNCE_UPWARD_ANGLE, BOUNCE_DISTANCE)
	}
	return nil
}
