package breakout

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleSystem struct {
	Emitter  Collidable
	Paricles []*Particle
}

func NewParticleSystem(emitter Collidable) *ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < 9; i++ {
		particles = append(particles, NewParticle(emitter))
	}
	spacingX := emitter.W() / 9
	spacingY := emitter.H() / 3
	for i := 0; i < len(particles); i++ {
		tmp := 1.0
		if i%2 == 0 {
			tmp = 1.0
		} else {
			tmp = -1.0
		}
		new_X := particles[i].Location.GetX() + float64(i%10)*spacingX
		new_Y := particles[i].Location.GetY() + float64(i%3)*spacingY
		particles[i].Location.SetX(new_X)
		particles[i].Location.SetY(new_Y)
		dX := particles[i].Velocity.GetX() + float64(i+1)*tmp*.5
		dY := particles[i].Velocity.GetY() + float64(i+1)*tmp*.5
		particles[i].Velocity.SetX(dX)
		particles[i].Velocity.SetY(dY)
	}
	return &ParticleSystem{
		Emitter:  emitter,
		Paricles: particles,
	}
}

func (s *ParticleSystem) Update() {
	for _, v := range s.Paricles {
		_ = v.Update()

	}

}

func (s *ParticleSystem) Draw(dst *ebiten.Image) {
	for _, v := range s.Paricles {
		v.Draw(dst)
	}
}

func (s *ParticleSystem) Remove(index int) {
	s.Paricles[index] = s.Paricles[len(s.Paricles)-1]
	s.Paricles = s.Paricles[:len(s.Paricles)-1]
}
