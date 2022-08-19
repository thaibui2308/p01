package ps

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/thaibui2308/p01/breakout"
)

type Vector = breakout.Coordinate

type Particle struct {
	Location     Vector
	Velocity     Vector
	Acceleration Vector
	Lifespan     float64
}

func NewParticle(emitter breakout.Collidable) *Particle {
	return &Particle{
		Location: breakout.Coordinate{
			X: emitter.Location().X,
			Y: emitter.Location().Y,
		},
		Velocity:     *RandVelocity(),
		Acceleration: breakout.Coordinate{X: 0, Y: -.05},
		Lifespan:     255,
	}
}

func (p *Particle) Update() error {
	p.Velocity = *p.Velocity.Add(&p.Acceleration)
	p.Location = *p.Location.Add(&p.Velocity)
	p.Lifespan -= 2.0
	return nil
}

func (p *Particle) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, p.Location.X, p.Location.Y, breakout.PARTICLE_WIDTH, breakout.BULLET_HEIGHT, color.RGBA{255, 255, 255, uint8(p.Lifespan)})
}

func (s *Particle) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return breakout.SCREEN_WIDTH, breakout.SCREEN_HEIGHT
}

func (p *Particle) IsDead() bool {
	return p.Lifespan < 0
}
