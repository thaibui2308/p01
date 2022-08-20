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
	ebitenutil.DrawRect(dst, float64(b.C.X), float64(b.C.Y), b.Width, b.Height, color.RGBA{0xff, 0xeb, 0xe7, 0xff})

}

func (b *Bullet) Update(game *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.SetMoving()
		game.Start = true
	}

	if game.Start {
		if b.Moving {
			// Check for collision with Player's block
			_ = b.DetectCollision(game.Player)

			// Check for collision with screen
			if b.C.X <= 0 || b.C.X+b.Width >= SCREEN_WIDTH {
				b.Translation.ReflectOy()
			} else if b.C.Y <= 0 {
				b.Translation.ReflectOx()
			} else if b.C.Y+b.Height > SCREEN_HEIGHT {
				// Check to see how many life Player still has left
				if game.LifeLeft > 1 {
					game.LifeLeft--
					game.Start = false
					b.ResetTo(&Coordinate{
						X: game.Player.C.X + (PLAYER_WIDTH-BULLET_WIDTH)/2,
						Y: PLAYER_Y - BULLET_HEIGHT,
					})
				} else {
					game.GameOver = true
				}
			}

			// Check for collision to adjust translation vector
			for i, v := range game.TargetSprite.Sprite {
				// Collision at the top and bottom of the Target's block
				hit := b.DetectCollision(&v)
				if hit {
					game.AppendSystem(NewParticleSystem(&v))
					game.TargetSprite.Remove(i)
					game.TargetSprite.HitSprite++
				}

			}

			b.C.Translate(&b.Translation)
		}
	} else {
		b.C.X = game.Player.C.X + BULLET_WIDTH
	}
	return nil
}

func (b *Bullet) DetectCollision(collidable Collidable) bool {
	c := collidable.Location()
	Height := collidable.H()
	Width := collidable.W()
	// Pre-Collision enhancement
	pre := b.Translation.Add(&b.Translation)
	if b.C.Add(pre).X+b.Width >= c.X && b.C.Add(pre).X <= c.X+Width {
		if b.C.Add(pre).Y <= c.Y+Height && b.C.Add(pre).Y >= c.Y {
			b.Translation.ReflectOx()
			return true
		}
	} else if b.C.Add(pre).Y+b.Height >= c.Y && b.C.Add(pre).Y <= c.Y+Height {
		if b.C.Add(pre).X <= c.X+Width && b.C.Add(pre).X >= c.X {
			b.Translation.ReflectOy()
			return true
		}
	}
	return false
}

func (b *Bullet) ResetTo(c *Coordinate) {
	b.C.SetX(c.GetX())
	b.C.SetY(c.GetY())
	b.Translation = *INITIAL_BULLET_DIRECTION
	b.C.Translate(&b.Translation)
}
