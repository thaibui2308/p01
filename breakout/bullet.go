package breakout

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Bullet struct {
	C           Coordinate
	Height      float64
	Width       float64
	Moving      bool
	Translation Coordinate
}

func NewBullet() *Bullet {
	bullet_coordinate := &Coordinate{
		X: BULLET_X,
		Y: BULLET_Y,
	}
	return &Bullet{
		C:           *bullet_coordinate,
		Height:      BULLET_HEIGHT,
		Width:       BULLET_WIDTH,
		Moving:      false,
		Translation: *INITIAL_BULLET_DIRECTION,
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

func (b *Bullet) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.SetMoving()
		game.Start = true
	}

	if game.Start {
		if b.Moving {
			// Check for collision with Player's block
			if b.C.Add(&b.Translation).X+b.Width >= game.Player.C.X && b.C.Add(&b.Translation).X <= game.Player.C.X+game.Player.Width {
				if b.C.Add(&b.Translation).Y <= game.Player.C.Y+game.Player.Height && b.C.Add(&b.Translation).Y >= game.Player.C.Y {
					// adjusting translation vector based on the collision type
					b.Translation.ReflectOx()
				}
			} else if b.C.Add(&b.Translation).Y+b.Height >= game.Player.C.Y && b.C.Add(&b.Translation).Y <= game.Player.C.Y+game.Player.Height {
				if b.C.Add(&b.Translation).X <= game.Player.C.X+game.Player.Width && b.C.Add(&b.Translation).X >= game.Player.C.X {
					b.Translation.ReflectOy()
				}
			}

			// Check for collision with screen
			if b.C.X <= 0 || b.C.X+b.Width >= SCREEN_WIDTH {
				b.Translation.ReflectOy()
			} else if b.C.Y <= 0 {
				b.Translation.ReflectOx()
			} else if b.C.Y+b.Height > SCREEN_HEIGHT {
				game.GameOver = true
			}

			// Check for collision to adjust translation vector
			for i, v := range game.TargetSprite.Sprite {
				// Collision at the top and bottom of the Target's block
				if b.C.Add(&b.Translation).X+b.Width >= v.C.X && b.C.Add(&b.Translation).X <= v.C.X+v.Width {
					if b.C.Add(&b.Translation).Y <= v.C.Y+v.Height && b.C.Add(&b.Translation).Y >= v.C.Y {
						game.TargetSprite.Remove(i)
						b.Translation.ReflectOx()
					}
					// Collision at the left and right edge of Target's block
				} else if b.C.Add(&b.Translation).Y+b.Height >= v.C.Y && b.C.Add(&b.Translation).Y <= v.C.Y+v.Height {
					if b.C.Add(&b.Translation).X <= v.C.X+v.Width && b.C.Add(&b.Translation).X >= v.C.X {
						game.TargetSprite.Remove(i)
						b.Translation.ReflectOy()
					}
				}
			}

			b.C.Translate(&b.Translation)
		}
	} else {
		b.C.X = game.Player.C.X + BULLET_WIDTH
	}
	return nil
}
