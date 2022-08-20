package breakout

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Vector = Coordinate

type Particle struct {
	Location     Vector
	Velocity     Vector
	Acceleration Vector
	Lifespan     float64
}

func NewParticle(emitter Collidable) *Particle {
	return &Particle{
		Location: Coordinate{
			X: emitter.Location().X,
			Y: emitter.Location().Y,
		},
		Velocity:     *RandVelocity(),
		Acceleration: Coordinate{X: 0, Y: -.05},
		Lifespan:     255,
	}
}

func (p *Particle) Update() error {
	p.Velocity = *p.Velocity.Add(&p.Acceleration)
	p.Location = *p.Location.Add(&p.Velocity)
	p.Lifespan -= 5

	if int(p.Lifespan)%255 == 0 {
		p.Velocity.ReflectOx()
	}

	return nil
}

func (p *Particle) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, p.Location.X, p.Location.Y, PARTICLE_WIDTH, BULLET_HEIGHT, color.RGBA{255, 255, 255, uint8(math.Abs(p.Lifespan))})
}
